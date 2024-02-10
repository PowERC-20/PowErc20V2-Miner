package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"peth-v2-miner/abi"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
	"github.com/gosuri/uilive"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var (
	logger = logrus.New()
)

// Define a struct for the configuration
type Config struct {
	InfuraURL       string   `yaml:"infuraURL"`
	ContractAddress string   `yaml:"contractAddress"`
	WorkerCount     int      `yaml:"workerCount"`
	MaxMiner        int      `yaml:"maxMiner"`
	PrivateKeys     []string `yaml:"privateKeys"`
	ReferrerAddress string   `yaml:"referrerAddress"`
}

// ReadConfig reads a config file and unmarshals it into a Config struct
func ReadConfig(configPath string) (*Config, error) {
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// WriteDefaultConfig writes a default configuration file
func WriteDefaultConfig(configPath string) error {
	defaultConfig := Config{
		InfuraURL:       "https://rpc.ankr.com/base",
		ContractAddress: "0x2Bd1D2Ae9C41Ae12082B20F00E36Eb6466703342",
		WorkerCount:     10,
		MaxMiner:        5,
		PrivateKeys:     []string{"00000"},
		ReferrerAddress: "",
	}
	data, err := yaml.Marshal(&defaultConfig)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

func mineWorker(ctx context.Context, wg *sync.WaitGroup, fromAddress common.Address, resultChan chan<- *big.Int, errorChan chan<- error, challenge *big.Int, target *big.Int, hashCount *int64) {
	defer wg.Done()

	var nonce *big.Int
	var err error

	for {
		select {
		case <-ctx.Done():
			return
		default:
			nonce, err = rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 256))
			if err != nil {
				errorChan <- fmt.Errorf("failed to generate random nonce: %v", err)
				return
			}

			noncePadded := common.LeftPadBytes(nonce.Bytes(), 32)
			challengePadded := common.LeftPadBytes(challenge.Bytes(), 32)
			addressBytes := fromAddress.Bytes()
			data := append(noncePadded, append(challengePadded, addressBytes...)...)
			// hash := sha256.Sum256(data)
			hash := crypto.Keccak256Hash(data)
			hashInt := new(big.Int).SetBytes(hash[:])
			if hashInt.Cmp(target) == -1 {
				resultChan <- nonce
				return
			}
			atomic.AddInt64(hashCount, 1)
		}
	}
}

func main() {
	banner := `
//  ____    __        _______ ____   ____ ____   ___    __  __ _                 
// |  _ \ __\ \      / / ____|  _ \ / ___|___ \ / _ \  |  \/  (_)_ __   ___ _ __ 
// | |_) / _ \ \ /\ / /|  _| | |_) | |     __) | | | | | |\/| | | '_ \ / _ \ '__|
// |  __/ (_) \ V  V / | |___|  _ <| |___ / __/| |_| | | |  | | | | | |  __/ |   
// |_|   \___/ \_/\_/  |_____|_| \_\\____|_____|\___/  |_|  |_|_|_| |_|\___|_|   
	`
	fmt.Println(banner)

	// Check if the config file exists, if not, create a default one
	configPath := "config.yaml"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err := WriteDefaultConfig(configPath)
		if err != nil {
			logger.Fatalf("Failed to create default config file: %v", err)
		}
		logger.Fatalf("Please modify the content in the config.yaml.")
	}

	// Read the configuration from the file
	config, err := ReadConfig(configPath)
	if err != nil {
		logger.Fatalf("Failed to read config file: %v", err)
	}

	writer := uilive.New()

	writer.Start()
	defer writer.Stop()

	logger.Info(color.GreenString("Establishing connection with Ethereum client..."))
	client, err := ethclient.Dial(config.InfuraURL)
	if err != nil {
		logger.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	logger.Info(color.GreenString("Successfully connected to Ethereum client."))

	privateKeys := config.PrivateKeys

	totalAttempts := 0

	for totalAttempts < config.MaxMiner {
		randIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(privateKeys))))
		if err != nil {
			logger.Fatalf("Failed to generate random index: %v", err)
		}
		wallet := privateKeys[randIndex.Int64()]

		privateKeyECDSA, err := crypto.HexToECDSA(wallet)
		if err != nil {
			logger.Fatalf("Error in parsing private key: %v", err)
		}

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			logger.Fatalf("Failed to get chainID: %v", err)
		}
		logger.Infof(color.GreenString("Successfully connected to Ethereum network with Chain ID: %v"), chainID)

		auth, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainID)
		if err != nil {
			logger.Fatalf("Failed to create transactor: %v", err)
		}

		contractAddr := common.HexToAddress(config.ContractAddress)
		contract, err := abi.NewFairToken(contractAddr, client)
		if err != nil {
			logger.Fatalf("Failed to instantiate a Token contract: %v", err)
		}
		logger.Info(color.GreenString("FairToken token contract successfully instantiated."))

		publicKey := privateKeyECDSA.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			logger.Fatal("Error casting public key to ECDSA")
		}
		accountAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		balance, err := contract.BalanceOf(nil, accountAddress)
		if err != nil {
			logger.Fatalf("Failed to get contract balance: %v", err)
		}
		logger.Infof(color.GreenString("Current contract balance: %d"), balance)

		referrerAddr := common.HexToAddress(config.ReferrerAddress)

		isValid, err := contract.IsValidReferrer(nil, referrerAddr)
		if err != nil {
			logger.Fatalf("Failed to verify referrer address: %v", err)
		}
		if !isValid {
			logger.Fatalf("Referrer address is not valid")
		}

		contractName, err := contract.Name(nil)
		if err != nil {
			logger.Fatalf("Failed to get contract name: %v", err)
		}
		logger.Infof(color.GreenString("Contract Name: %s"), color.RedString(contractName))

		challenge, err := contract.CurrentChallenge(nil)
		if err != nil {
			logger.Fatalf("Failed to get challenge: %v", err)
		}
		logger.Infof(color.GreenString("Current mining challenge number: %d"), challenge)

		difficulty, err := contract.MiningDifficulty(nil)
		if err != nil {
			logger.Fatalf("Failed to get difficulty: %v", err)
		}
		logger.Infof(color.GreenString("Current mining difficulty level: %d"), difficulty)

		difficultyUint := uint(difficulty.Uint64())
		target := new(big.Int).Lsh(big.NewInt(1), 256-difficultyUint)
		logger.Infof(color.GreenString("Target number is: %d"), target)

		resultChan := make(chan *big.Int)
		errorChan := make(chan error)
		refreshChan := make(chan bool)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		logger.Info(color.YellowString("Mining workers started..."))

		var hashCount int64

		var wg sync.WaitGroup
		for i := 0; i < config.WorkerCount; i++ {
			wg.Add(1)
			go mineWorker(ctx, &wg, auth.From, resultChan, errorChan, challenge, target, &hashCount)
		}

		ticker := time.NewTicker(1 * time.Second)

		go func() {
			for {
				select {
				case <-ticker.C:
					timestamp := time.Now().Format("2006-01-02 15:04:05")
					currentHashCount := atomic.SwapInt64(&hashCount, 0)
					hashesPerSecond := float64(currentHashCount)
					fmt.Fprintf(writer, "%s[%s] %s\n", color.BlueString("Mining"), timestamp, color.GreenString("Total hashes per second: %8.2f", hashesPerSecond))
				}
			}
		}()

		refreshTicker := time.NewTicker(10 * time.Second)
		go func() {
			for {
				select {
				case <-refreshTicker.C:
					newDifficulty, err := contract.MiningDifficulty(nil)
					if err != nil {
						logger.Errorf("Failed to get difficulty: %v\n", err)
						continue
					}
					newChallenge, err := contract.CurrentChallenge(nil)
					if err != nil {
						logger.Errorf("Failed to get challenge: %v\n", err)
						continue
					}

					timestamp := time.Now().Format("2006-01-02 15:04:05")
					fmt.Fprintf(writer, "%s[%s] %s\n", color.BlueString("Difficulty"), timestamp, color.GreenString("Current mining difficulty level: %d", newDifficulty))
					if difficulty.Cmp(newDifficulty) != 0 || challenge.Cmp(newChallenge) != 0 {
						difficulty = newDifficulty
						challenge = newChallenge

						refreshChan <- true
					}
				}
			}
		}()

		select {
		case <-refreshChan:
			ticker.Stop()
			refreshTicker.Stop()
			cancel()
			wg.Wait()
			logger.Infof(color.GreenString("Refresh mining difficulty level: %d"), difficulty)

		case nonce := <-resultChan:
			ticker.Stop()
			refreshTicker.Stop()
			cancel()
			wg.Wait()
			logger.Infof(color.GreenString("Successfully discovered a valid nonce: %d"), nonce)
			logger.Info(color.YellowString("Submitting mining transaction with nonce..."))
			auth.Value = big.NewInt(1000000000000000)
			tx, err := contract.Mine(auth, nonce, referrerAddr)
			if err != nil {
				logger.Errorf("Failed to submit mine transaction: %v", err)
				break
			}
			receipt, err := bind.WaitMined(context.Background(), client, tx)
			if err != nil {
				logger.Errorf("Failed to mine the transaction: %v", err)
				break
			}
			logger.Infof(color.GreenString("Mining transaction successfully confirmed, Transaction Hash: %s"), color.CyanString(receipt.TxHash.Hex()))

			totalAttempts++
		case err := <-errorChan:
			cancel()
			wg.Wait()
			logger.Errorf("Mining operation failed due to an error: %v", err)
			break
		}

	}

	logger.Info(color.GreenString("Mining process successfully completed"))
}
