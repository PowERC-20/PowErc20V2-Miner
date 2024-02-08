# PoWERC20 V2 Miner

This Go program is a specialized miner for the PETHv2 token on the Ethereum Base Chain. It connects to an Ethereum client, reads a configuration file, and interacts with a specified smart contract to execute mining operations. Utilizing concurrent goroutines, the program efficiently parallelizes mining efforts to optimize your chances of success while providing real-time mining statistics on the console.

## Features

- **Base Chain Connectivity**: Connects to the Ethereum Base Chain using a provided RPC URL.
- **Smart Contract Mining**: Mines PETHv2 tokens by interacting with the designated smart contract.
- **Concurrent Mining**: Employs multiple goroutines for simultaneous mining operations.
- **Configurable Settings**: Customizable through a YAML configuration file for ease of use.
- **Live Statistics**: Offers a live update of mining statistics using `uilive`.

## Getting Started

### Prerequisites

- Go (version 1.21.x or later recommended)
- An Ethereum wallet with a balance of ETH to cover gas fees
- Access to an Ethereum Base Chain client endpoint (e.g., Infura)

### Installation


1. Clone the repository:

   ```sh
   git clone https://github.com/PowERC-20/PowErc20V2-Miner
   ```

2. Change directory to the cloned repository:

   ```sh
   cd PowErc20V2-Miner
   ```

3. Build the miner:

   ```sh
   go build .
   ```

### Configuration

Upon initial run of the tool, a config.yaml configuration file will be generated. Modify this configuration file with appropriate settings:

  - `infuraURL`: Your Ethereum Base Chain client RPC URL.
  - `contractAddress`: The smart contract address for PETHv2 token mining.
  - `workerCount`: Number of goroutines to deploy for mining.
  - `maxMiner`: Maximum number of attempts before the miner stops.
  - `privateKeys`: List of private keys for the wallets you intend to use.
  - `referrerAddress`: A referrer address if one is required.



### Usage

Execute the mining program:

```sh
./peth-v2-miner
```

Monitor your terminal for live updates on the mining process. The miner will continue to work according to your `config.yaml` settings.

## Contact

- Twitter - [@Powerc_20](https://twitter.com/powerc_20)
- Project Link: [https://v2.powerc20.com/](https://v2.powerc20.com/)

## Contributing

Your contributions are welcome. Please adhere to the project's coding standards and include tests for any new features or fixes.

## Declare

The project code is completely open source. The released version is compiled using Github Actions. If you have any questions about the code, please feel free to raise them or submit the code for security auditing anywhere.

## Acknowledgements

- [go-ethereum](https://github.com/ethereum/go-ethereum)
- [uilive](https://github.com/gosuri/uilive)
- [logrus](https://github.com/sirupsen/logrus)
- [color](https://github.com/fatih/color)
- [yaml.v2](https://gopkg.in/yaml.v2)
