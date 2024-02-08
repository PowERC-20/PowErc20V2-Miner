// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// FairTokenStake is an auto generated low-level Go binding around an user-defined struct.
type FairTokenStake struct {
	Amount           *big.Int
	StartTime        *big.Int
	RewardPaid       *big.Int
	RewardMultiplier *big.Int
}

// FairTokenMetaData contains all meta data concerning the FairToken contract.
var FairTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wethAmount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wethAmount\",\"type\":\"uint256\"}],\"name\":\"buybackFairTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lzEndpoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"referrerAddresses\",\"type\":\"address[]\"}],\"name\":\"importPethHolder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethSpent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fairTokensBought\",\"type\":\"uint256\"}],\"name\":\"Buyback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newChallenge\",\"type\":\"uint256\"}],\"name\":\"ChallengeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ClaimReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDifficulty\",\"type\":\"uint256\"}],\"name\":\"DifficultyAdjusted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"FeePaid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"}],\"name\":\"forceResumeReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_reason\",\"type\":\"bytes\"}],\"name\":\"MessageFailed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"mine\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challenge\",\"type\":\"uint256\"}],\"name\":\"MiningSuccess\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"nonblockingLzReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"ReferrerSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"retryMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_payloadHash\",\"type\":\"bytes32\"}],\"name\":\"RetryMessageSuccess\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_configType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_config\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_allowed\",\"type\":\"bool\"}],\"name\":\"setEarlyUnstakeAllowed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_packetType\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_minGas\",\"type\":\"uint256\"}],\"name\":\"setMinDstGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_type\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_minDstGas\",\"type\":\"uint256\"}],\"name\":\"SetMinDstGas\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"}],\"name\":\"setPayloadSizeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_precrime\",\"type\":\"address\"}],\"name\":\"setPrecrime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"precrime\",\"type\":\"address\"}],\"name\":\"SetPrecrime\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"}],\"name\":\"setReceiveVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"}],\"name\":\"setSendVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasuryAddress\",\"type\":\"address\"}],\"name\":\"setTreasuryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_path\",\"type\":\"bytes\"}],\"name\":\"setTrustedRemote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_path\",\"type\":\"bytes\"}],\"name\":\"SetTrustedRemote\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_remoteAddress\",\"type\":\"bytes\"}],\"name\":\"setTrustedRemoteAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_remoteAddress\",\"type\":\"bytes\"}],\"name\":\"SetTrustedRemoteAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"StakeToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakeAll\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeIndex\",\"type\":\"uint256\"}],\"name\":\"unstakeOne\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"UnStakeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"WithdrawToken\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"calculateReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentChallenge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentStage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_PAYLOAD_SIZE_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dstChainId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"earlyUnstakeAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"earlyUnstakeAllowedTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"estimateStakeFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"failedMessages\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getAllStakes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardMultiplier\",\"type\":\"uint256\"}],\"internalType\":\"structFairToken.Stake[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_configType\",\"type\":\"uint256\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"}],\"name\":\"getTrustedRemoteAddress\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"}],\"name\":\"isTrustedRemote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"isValidReferrer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAdjustmentBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lzEndpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"minDstGasLookup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"miningAttempts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"miningDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"miningReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nonceUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"payloadSizeLimitLookup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingTax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pethHolder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"posPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"posPoolStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"powPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"precrime\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"referrerCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"referrerRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"referrers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardMultiplier\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMined\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRewardsDistributed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"trustedRemoteLookup\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Pair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Router\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FairTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use FairTokenMetaData.ABI instead.
var FairTokenABI = FairTokenMetaData.ABI

// FairToken is an auto generated Go binding around an Ethereum contract.
type FairToken struct {
	FairTokenCaller     // Read-only binding to the contract
	FairTokenTransactor // Write-only binding to the contract
	FairTokenFilterer   // Log filterer for contract events
}

// FairTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type FairTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FairTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FairTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FairTokenSession struct {
	Contract     *FairToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FairTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FairTokenCallerSession struct {
	Contract *FairTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FairTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FairTokenTransactorSession struct {
	Contract     *FairTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FairTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type FairTokenRaw struct {
	Contract *FairToken // Generic contract binding to access the raw methods on
}

// FairTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FairTokenCallerRaw struct {
	Contract *FairTokenCaller // Generic read-only contract binding to access the raw methods on
}

// FairTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FairTokenTransactorRaw struct {
	Contract *FairTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFairToken creates a new instance of FairToken, bound to a specific deployed contract.
func NewFairToken(address common.Address, backend bind.ContractBackend) (*FairToken, error) {
	contract, err := bindFairToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FairToken{FairTokenCaller: FairTokenCaller{contract: contract}, FairTokenTransactor: FairTokenTransactor{contract: contract}, FairTokenFilterer: FairTokenFilterer{contract: contract}}, nil
}

// NewFairTokenCaller creates a new read-only instance of FairToken, bound to a specific deployed contract.
func NewFairTokenCaller(address common.Address, caller bind.ContractCaller) (*FairTokenCaller, error) {
	contract, err := bindFairToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FairTokenCaller{contract: contract}, nil
}

// NewFairTokenTransactor creates a new write-only instance of FairToken, bound to a specific deployed contract.
func NewFairTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*FairTokenTransactor, error) {
	contract, err := bindFairToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FairTokenTransactor{contract: contract}, nil
}

// NewFairTokenFilterer creates a new log filterer instance of FairToken, bound to a specific deployed contract.
func NewFairTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*FairTokenFilterer, error) {
	contract, err := bindFairToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FairTokenFilterer{contract: contract}, nil
}

// bindFairToken binds a generic wrapper to an already deployed contract.
func bindFairToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FairTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FairToken *FairTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FairToken.Contract.FairTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FairToken *FairTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FairToken.Contract.FairTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FairToken *FairTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FairToken.Contract.FairTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FairToken *FairTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FairToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FairToken *FairTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FairToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FairToken *FairTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FairToken.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTPAYLOADSIZELIMIT is a free data retrieval call binding the contract method 0xc4461834.
//
// Solidity: function DEFAULT_PAYLOAD_SIZE_LIMIT() view returns(uint256)
func (_FairToken *FairTokenCaller) DEFAULTPAYLOADSIZELIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "DEFAULT_PAYLOAD_SIZE_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTPAYLOADSIZELIMIT is a free data retrieval call binding the contract method 0xc4461834.
//
// Solidity: function DEFAULT_PAYLOAD_SIZE_LIMIT() view returns(uint256)
func (_FairToken *FairTokenSession) DEFAULTPAYLOADSIZELIMIT() (*big.Int, error) {
	return _FairToken.Contract.DEFAULTPAYLOADSIZELIMIT(&_FairToken.CallOpts)
}

// DEFAULTPAYLOADSIZELIMIT is a free data retrieval call binding the contract method 0xc4461834.
//
// Solidity: function DEFAULT_PAYLOAD_SIZE_LIMIT() view returns(uint256)
func (_FairToken *FairTokenCallerSession) DEFAULTPAYLOADSIZELIMIT() (*big.Int, error) {
	return _FairToken.Contract.DEFAULTPAYLOADSIZELIMIT(&_FairToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FairToken *FairTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FairToken *FairTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FairToken.Contract.Allowance(&_FairToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FairToken *FairTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FairToken.Contract.Allowance(&_FairToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FairToken *FairTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FairToken *FairTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _FairToken.Contract.BalanceOf(&_FairToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FairToken *FairTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _FairToken.Contract.BalanceOf(&_FairToken.CallOpts, account)
}

// CalculateReward is a free data retrieval call binding the contract method 0xd82e3962.
//
// Solidity: function calculateReward(address user) view returns(uint256 totalReward)
func (_FairToken *FairTokenCaller) CalculateReward(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "calculateReward", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateReward is a free data retrieval call binding the contract method 0xd82e3962.
//
// Solidity: function calculateReward(address user) view returns(uint256 totalReward)
func (_FairToken *FairTokenSession) CalculateReward(user common.Address) (*big.Int, error) {
	return _FairToken.Contract.CalculateReward(&_FairToken.CallOpts, user)
}

// CalculateReward is a free data retrieval call binding the contract method 0xd82e3962.
//
// Solidity: function calculateReward(address user) view returns(uint256 totalReward)
func (_FairToken *FairTokenCallerSession) CalculateReward(user common.Address) (*big.Int, error) {
	return _FairToken.Contract.CalculateReward(&_FairToken.CallOpts, user)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x51bdd585.
//
// Solidity: function currentChallenge() view returns(uint256)
func (_FairToken *FairTokenCaller) CurrentChallenge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "currentChallenge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentChallenge is a free data retrieval call binding the contract method 0x51bdd585.
//
// Solidity: function currentChallenge() view returns(uint256)
func (_FairToken *FairTokenSession) CurrentChallenge() (*big.Int, error) {
	return _FairToken.Contract.CurrentChallenge(&_FairToken.CallOpts)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x51bdd585.
//
// Solidity: function currentChallenge() view returns(uint256)
func (_FairToken *FairTokenCallerSession) CurrentChallenge() (*big.Int, error) {
	return _FairToken.Contract.CurrentChallenge(&_FairToken.CallOpts)
}

// CurrentStage is a free data retrieval call binding the contract method 0x5bf5d54c.
//
// Solidity: function currentStage() view returns(uint256)
func (_FairToken *FairTokenCaller) CurrentStage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "currentStage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentStage is a free data retrieval call binding the contract method 0x5bf5d54c.
//
// Solidity: function currentStage() view returns(uint256)
func (_FairToken *FairTokenSession) CurrentStage() (*big.Int, error) {
	return _FairToken.Contract.CurrentStage(&_FairToken.CallOpts)
}

// CurrentStage is a free data retrieval call binding the contract method 0x5bf5d54c.
//
// Solidity: function currentStage() view returns(uint256)
func (_FairToken *FairTokenCallerSession) CurrentStage() (*big.Int, error) {
	return _FairToken.Contract.CurrentStage(&_FairToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FairToken *FairTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FairToken *FairTokenSession) Decimals() (uint8, error) {
	return _FairToken.Contract.Decimals(&_FairToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FairToken *FairTokenCallerSession) Decimals() (uint8, error) {
	return _FairToken.Contract.Decimals(&_FairToken.CallOpts)
}

// DstChainId is a free data retrieval call binding the contract method 0x30c593f7.
//
// Solidity: function dstChainId() view returns(uint16)
func (_FairToken *FairTokenCaller) DstChainId(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "dstChainId")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DstChainId is a free data retrieval call binding the contract method 0x30c593f7.
//
// Solidity: function dstChainId() view returns(uint16)
func (_FairToken *FairTokenSession) DstChainId() (uint16, error) {
	return _FairToken.Contract.DstChainId(&_FairToken.CallOpts)
}

// DstChainId is a free data retrieval call binding the contract method 0x30c593f7.
//
// Solidity: function dstChainId() view returns(uint16)
func (_FairToken *FairTokenCallerSession) DstChainId() (uint16, error) {
	return _FairToken.Contract.DstChainId(&_FairToken.CallOpts)
}

// EarlyUnstakeAllowed is a free data retrieval call binding the contract method 0x8ac1bd0a.
//
// Solidity: function earlyUnstakeAllowed() view returns(bool)
func (_FairToken *FairTokenCaller) EarlyUnstakeAllowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "earlyUnstakeAllowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EarlyUnstakeAllowed is a free data retrieval call binding the contract method 0x8ac1bd0a.
//
// Solidity: function earlyUnstakeAllowed() view returns(bool)
func (_FairToken *FairTokenSession) EarlyUnstakeAllowed() (bool, error) {
	return _FairToken.Contract.EarlyUnstakeAllowed(&_FairToken.CallOpts)
}

// EarlyUnstakeAllowed is a free data retrieval call binding the contract method 0x8ac1bd0a.
//
// Solidity: function earlyUnstakeAllowed() view returns(bool)
func (_FairToken *FairTokenCallerSession) EarlyUnstakeAllowed() (bool, error) {
	return _FairToken.Contract.EarlyUnstakeAllowed(&_FairToken.CallOpts)
}

// EarlyUnstakeAllowedTimestamp is a free data retrieval call binding the contract method 0xc30a50b8.
//
// Solidity: function earlyUnstakeAllowedTimestamp() view returns(uint256)
func (_FairToken *FairTokenCaller) EarlyUnstakeAllowedTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "earlyUnstakeAllowedTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EarlyUnstakeAllowedTimestamp is a free data retrieval call binding the contract method 0xc30a50b8.
//
// Solidity: function earlyUnstakeAllowedTimestamp() view returns(uint256)
func (_FairToken *FairTokenSession) EarlyUnstakeAllowedTimestamp() (*big.Int, error) {
	return _FairToken.Contract.EarlyUnstakeAllowedTimestamp(&_FairToken.CallOpts)
}

// EarlyUnstakeAllowedTimestamp is a free data retrieval call binding the contract method 0xc30a50b8.
//
// Solidity: function earlyUnstakeAllowedTimestamp() view returns(uint256)
func (_FairToken *FairTokenCallerSession) EarlyUnstakeAllowedTimestamp() (*big.Int, error) {
	return _FairToken.Contract.EarlyUnstakeAllowedTimestamp(&_FairToken.CallOpts)
}

// EstimateStakeFees is a free data retrieval call binding the contract method 0xb48aa744.
//
// Solidity: function estimateStakeFees(uint256 _amount) view returns(uint256 nativeFee)
func (_FairToken *FairTokenCaller) EstimateStakeFees(opts *bind.CallOpts, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "estimateStakeFees", _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateStakeFees is a free data retrieval call binding the contract method 0xb48aa744.
//
// Solidity: function estimateStakeFees(uint256 _amount) view returns(uint256 nativeFee)
func (_FairToken *FairTokenSession) EstimateStakeFees(_amount *big.Int) (*big.Int, error) {
	return _FairToken.Contract.EstimateStakeFees(&_FairToken.CallOpts, _amount)
}

// EstimateStakeFees is a free data retrieval call binding the contract method 0xb48aa744.
//
// Solidity: function estimateStakeFees(uint256 _amount) view returns(uint256 nativeFee)
func (_FairToken *FairTokenCallerSession) EstimateStakeFees(_amount *big.Int) (*big.Int, error) {
	return _FairToken.Contract.EstimateStakeFees(&_FairToken.CallOpts, _amount)
}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_FairToken *FairTokenCaller) FailedMessages(opts *bind.CallOpts, arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "failedMessages", arg0, arg1, arg2)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_FairToken *FairTokenSession) FailedMessages(arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	return _FairToken.Contract.FailedMessages(&_FairToken.CallOpts, arg0, arg1, arg2)
}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_FairToken *FairTokenCallerSession) FailedMessages(arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	return _FairToken.Contract.FailedMessages(&_FairToken.CallOpts, arg0, arg1, arg2)
}

// FeeDenominator is a free data retrieval call binding the contract method 0x180b0d7e.
//
// Solidity: function feeDenominator() view returns(uint256)
func (_FairToken *FairTokenCaller) FeeDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "feeDenominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeDenominator is a free data retrieval call binding the contract method 0x180b0d7e.
//
// Solidity: function feeDenominator() view returns(uint256)
func (_FairToken *FairTokenSession) FeeDenominator() (*big.Int, error) {
	return _FairToken.Contract.FeeDenominator(&_FairToken.CallOpts)
}

// FeeDenominator is a free data retrieval call binding the contract method 0x180b0d7e.
//
// Solidity: function feeDenominator() view returns(uint256)
func (_FairToken *FairTokenCallerSession) FeeDenominator() (*big.Int, error) {
	return _FairToken.Contract.FeeDenominator(&_FairToken.CallOpts)
}

// GetAllStakes is a free data retrieval call binding the contract method 0x04238994.
//
// Solidity: function getAllStakes(address user) view returns((uint256,uint256,uint256,uint256)[])
func (_FairToken *FairTokenCaller) GetAllStakes(opts *bind.CallOpts, user common.Address) ([]FairTokenStake, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "getAllStakes", user)

	if err != nil {
		return *new([]FairTokenStake), err
	}

	out0 := *abi.ConvertType(out[0], new([]FairTokenStake)).(*[]FairTokenStake)

	return out0, err

}

// GetAllStakes is a free data retrieval call binding the contract method 0x04238994.
//
// Solidity: function getAllStakes(address user) view returns((uint256,uint256,uint256,uint256)[])
func (_FairToken *FairTokenSession) GetAllStakes(user common.Address) ([]FairTokenStake, error) {
	return _FairToken.Contract.GetAllStakes(&_FairToken.CallOpts, user)
}

// GetAllStakes is a free data retrieval call binding the contract method 0x04238994.
//
// Solidity: function getAllStakes(address user) view returns((uint256,uint256,uint256,uint256)[])
func (_FairToken *FairTokenCallerSession) GetAllStakes(user common.Address) ([]FairTokenStake, error) {
	return _FairToken.Contract.GetAllStakes(&_FairToken.CallOpts, user)
}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_FairToken *FairTokenCaller) GetConfig(opts *bind.CallOpts, _version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "getConfig", _version, _chainId, arg2, _configType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_FairToken *FairTokenSession) GetConfig(_version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	return _FairToken.Contract.GetConfig(&_FairToken.CallOpts, _version, _chainId, arg2, _configType)
}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_FairToken *FairTokenCallerSession) GetConfig(_version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	return _FairToken.Contract.GetConfig(&_FairToken.CallOpts, _version, _chainId, arg2, _configType)
}

// GetTrustedRemoteAddress is a free data retrieval call binding the contract method 0x9f38369a.
//
// Solidity: function getTrustedRemoteAddress(uint16 _remoteChainId) view returns(bytes)
func (_FairToken *FairTokenCaller) GetTrustedRemoteAddress(opts *bind.CallOpts, _remoteChainId uint16) ([]byte, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "getTrustedRemoteAddress", _remoteChainId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTrustedRemoteAddress is a free data retrieval call binding the contract method 0x9f38369a.
//
// Solidity: function getTrustedRemoteAddress(uint16 _remoteChainId) view returns(bytes)
func (_FairToken *FairTokenSession) GetTrustedRemoteAddress(_remoteChainId uint16) ([]byte, error) {
	return _FairToken.Contract.GetTrustedRemoteAddress(&_FairToken.CallOpts, _remoteChainId)
}

// GetTrustedRemoteAddress is a free data retrieval call binding the contract method 0x9f38369a.
//
// Solidity: function getTrustedRemoteAddress(uint16 _remoteChainId) view returns(bytes)
func (_FairToken *FairTokenCallerSession) GetTrustedRemoteAddress(_remoteChainId uint16) ([]byte, error) {
	return _FairToken.Contract.GetTrustedRemoteAddress(&_FairToken.CallOpts, _remoteChainId)
}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_FairToken *FairTokenCaller) IsTrustedRemote(opts *bind.CallOpts, _srcChainId uint16, _srcAddress []byte) (bool, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "isTrustedRemote", _srcChainId, _srcAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_FairToken *FairTokenSession) IsTrustedRemote(_srcChainId uint16, _srcAddress []byte) (bool, error) {
	return _FairToken.Contract.IsTrustedRemote(&_FairToken.CallOpts, _srcChainId, _srcAddress)
}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_FairToken *FairTokenCallerSession) IsTrustedRemote(_srcChainId uint16, _srcAddress []byte) (bool, error) {
	return _FairToken.Contract.IsTrustedRemote(&_FairToken.CallOpts, _srcChainId, _srcAddress)
}

// IsValidReferrer is a free data retrieval call binding the contract method 0xd52f199f.
//
// Solidity: function isValidReferrer(address referrer) view returns(bool)
func (_FairToken *FairTokenCaller) IsValidReferrer(opts *bind.CallOpts, referrer common.Address) (bool, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "isValidReferrer", referrer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidReferrer is a free data retrieval call binding the contract method 0xd52f199f.
//
// Solidity: function isValidReferrer(address referrer) view returns(bool)
func (_FairToken *FairTokenSession) IsValidReferrer(referrer common.Address) (bool, error) {
	return _FairToken.Contract.IsValidReferrer(&_FairToken.CallOpts, referrer)
}

// IsValidReferrer is a free data retrieval call binding the contract method 0xd52f199f.
//
// Solidity: function isValidReferrer(address referrer) view returns(bool)
func (_FairToken *FairTokenCallerSession) IsValidReferrer(referrer common.Address) (bool, error) {
	return _FairToken.Contract.IsValidReferrer(&_FairToken.CallOpts, referrer)
}

// LastAdjustmentBlockNumber is a free data retrieval call binding the contract method 0xb1cc23ac.
//
// Solidity: function lastAdjustmentBlockNumber() view returns(uint256)
func (_FairToken *FairTokenCaller) LastAdjustmentBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "lastAdjustmentBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastAdjustmentBlockNumber is a free data retrieval call binding the contract method 0xb1cc23ac.
//
// Solidity: function lastAdjustmentBlockNumber() view returns(uint256)
func (_FairToken *FairTokenSession) LastAdjustmentBlockNumber() (*big.Int, error) {
	return _FairToken.Contract.LastAdjustmentBlockNumber(&_FairToken.CallOpts)
}

// LastAdjustmentBlockNumber is a free data retrieval call binding the contract method 0xb1cc23ac.
//
// Solidity: function lastAdjustmentBlockNumber() view returns(uint256)
func (_FairToken *FairTokenCallerSession) LastAdjustmentBlockNumber() (*big.Int, error) {
	return _FairToken.Contract.LastAdjustmentBlockNumber(&_FairToken.CallOpts)
}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_FairToken *FairTokenCaller) LzEndpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "lzEndpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_FairToken *FairTokenSession) LzEndpoint() (common.Address, error) {
	return _FairToken.Contract.LzEndpoint(&_FairToken.CallOpts)
}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_FairToken *FairTokenCallerSession) LzEndpoint() (common.Address, error) {
	return _FairToken.Contract.LzEndpoint(&_FairToken.CallOpts)
}

// MinDstGasLookup is a free data retrieval call binding the contract method 0x8cfd8f5c.
//
// Solidity: function minDstGasLookup(uint16 , uint16 ) view returns(uint256)
func (_FairToken *FairTokenCaller) MinDstGasLookup(opts *bind.CallOpts, arg0 uint16, arg1 uint16) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "minDstGasLookup", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDstGasLookup is a free data retrieval call binding the contract method 0x8cfd8f5c.
//
// Solidity: function minDstGasLookup(uint16 , uint16 ) view returns(uint256)
func (_FairToken *FairTokenSession) MinDstGasLookup(arg0 uint16, arg1 uint16) (*big.Int, error) {
	return _FairToken.Contract.MinDstGasLookup(&_FairToken.CallOpts, arg0, arg1)
}

// MinDstGasLookup is a free data retrieval call binding the contract method 0x8cfd8f5c.
//
// Solidity: function minDstGasLookup(uint16 , uint16 ) view returns(uint256)
func (_FairToken *FairTokenCallerSession) MinDstGasLookup(arg0 uint16, arg1 uint16) (*big.Int, error) {
	return _FairToken.Contract.MinDstGasLookup(&_FairToken.CallOpts, arg0, arg1)
}

// MiningAttempts is a free data retrieval call binding the contract method 0xcd5d7cfb.
//
// Solidity: function miningAttempts() view returns(uint256)
func (_FairToken *FairTokenCaller) MiningAttempts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "miningAttempts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiningAttempts is a free data retrieval call binding the contract method 0xcd5d7cfb.
//
// Solidity: function miningAttempts() view returns(uint256)
func (_FairToken *FairTokenSession) MiningAttempts() (*big.Int, error) {
	return _FairToken.Contract.MiningAttempts(&_FairToken.CallOpts)
}

// MiningAttempts is a free data retrieval call binding the contract method 0xcd5d7cfb.
//
// Solidity: function miningAttempts() view returns(uint256)
func (_FairToken *FairTokenCallerSession) MiningAttempts() (*big.Int, error) {
	return _FairToken.Contract.MiningAttempts(&_FairToken.CallOpts)
}

// MiningDifficulty is a free data retrieval call binding the contract method 0xb40b9bb6.
//
// Solidity: function miningDifficulty() view returns(uint256)
func (_FairToken *FairTokenCaller) MiningDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "miningDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiningDifficulty is a free data retrieval call binding the contract method 0xb40b9bb6.
//
// Solidity: function miningDifficulty() view returns(uint256)
func (_FairToken *FairTokenSession) MiningDifficulty() (*big.Int, error) {
	return _FairToken.Contract.MiningDifficulty(&_FairToken.CallOpts)
}

// MiningDifficulty is a free data retrieval call binding the contract method 0xb40b9bb6.
//
// Solidity: function miningDifficulty() view returns(uint256)
func (_FairToken *FairTokenCallerSession) MiningDifficulty() (*big.Int, error) {
	return _FairToken.Contract.MiningDifficulty(&_FairToken.CallOpts)
}

// MiningReward is a free data retrieval call binding the contract method 0x4ac2d103.
//
// Solidity: function miningReward() view returns(uint256)
func (_FairToken *FairTokenCaller) MiningReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "miningReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiningReward is a free data retrieval call binding the contract method 0x4ac2d103.
//
// Solidity: function miningReward() view returns(uint256)
func (_FairToken *FairTokenSession) MiningReward() (*big.Int, error) {
	return _FairToken.Contract.MiningReward(&_FairToken.CallOpts)
}

// MiningReward is a free data retrieval call binding the contract method 0x4ac2d103.
//
// Solidity: function miningReward() view returns(uint256)
func (_FairToken *FairTokenCallerSession) MiningReward() (*big.Int, error) {
	return _FairToken.Contract.MiningReward(&_FairToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FairToken *FairTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FairToken *FairTokenSession) Name() (string, error) {
	return _FairToken.Contract.Name(&_FairToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FairToken *FairTokenCallerSession) Name() (string, error) {
	return _FairToken.Contract.Name(&_FairToken.CallOpts)
}

// NonceUsed is a free data retrieval call binding the contract method 0x1647795e.
//
// Solidity: function nonceUsed(address , uint256 ) view returns(bool)
func (_FairToken *FairTokenCaller) NonceUsed(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "nonceUsed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NonceUsed is a free data retrieval call binding the contract method 0x1647795e.
//
// Solidity: function nonceUsed(address , uint256 ) view returns(bool)
func (_FairToken *FairTokenSession) NonceUsed(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _FairToken.Contract.NonceUsed(&_FairToken.CallOpts, arg0, arg1)
}

// NonceUsed is a free data retrieval call binding the contract method 0x1647795e.
//
// Solidity: function nonceUsed(address , uint256 ) view returns(bool)
func (_FairToken *FairTokenCallerSession) NonceUsed(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _FairToken.Contract.NonceUsed(&_FairToken.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FairToken *FairTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FairToken *FairTokenSession) Owner() (common.Address, error) {
	return _FairToken.Contract.Owner(&_FairToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FairToken *FairTokenCallerSession) Owner() (common.Address, error) {
	return _FairToken.Contract.Owner(&_FairToken.CallOpts)
}

// PayloadSizeLimitLookup is a free data retrieval call binding the contract method 0x3f1f4fa4.
//
// Solidity: function payloadSizeLimitLookup(uint16 ) view returns(uint256)
func (_FairToken *FairTokenCaller) PayloadSizeLimitLookup(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "payloadSizeLimitLookup", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PayloadSizeLimitLookup is a free data retrieval call binding the contract method 0x3f1f4fa4.
//
// Solidity: function payloadSizeLimitLookup(uint16 ) view returns(uint256)
func (_FairToken *FairTokenSession) PayloadSizeLimitLookup(arg0 uint16) (*big.Int, error) {
	return _FairToken.Contract.PayloadSizeLimitLookup(&_FairToken.CallOpts, arg0)
}

// PayloadSizeLimitLookup is a free data retrieval call binding the contract method 0x3f1f4fa4.
//
// Solidity: function payloadSizeLimitLookup(uint16 ) view returns(uint256)
func (_FairToken *FairTokenCallerSession) PayloadSizeLimitLookup(arg0 uint16) (*big.Int, error) {
	return _FairToken.Contract.PayloadSizeLimitLookup(&_FairToken.CallOpts, arg0)
}

// PendingTax is a free data retrieval call binding the contract method 0xf05ffa26.
//
// Solidity: function pendingTax() view returns(uint256)
func (_FairToken *FairTokenCaller) PendingTax(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "pendingTax")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingTax is a free data retrieval call binding the contract method 0xf05ffa26.
//
// Solidity: function pendingTax() view returns(uint256)
func (_FairToken *FairTokenSession) PendingTax() (*big.Int, error) {
	return _FairToken.Contract.PendingTax(&_FairToken.CallOpts)
}

// PendingTax is a free data retrieval call binding the contract method 0xf05ffa26.
//
// Solidity: function pendingTax() view returns(uint256)
func (_FairToken *FairTokenCallerSession) PendingTax() (*big.Int, error) {
	return _FairToken.Contract.PendingTax(&_FairToken.CallOpts)
}

// PethHolder is a free data retrieval call binding the contract method 0x6341eb29.
//
// Solidity: function pethHolder(address ) view returns(bool)
func (_FairToken *FairTokenCaller) PethHolder(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "pethHolder", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PethHolder is a free data retrieval call binding the contract method 0x6341eb29.
//
// Solidity: function pethHolder(address ) view returns(bool)
func (_FairToken *FairTokenSession) PethHolder(arg0 common.Address) (bool, error) {
	return _FairToken.Contract.PethHolder(&_FairToken.CallOpts, arg0)
}

// PethHolder is a free data retrieval call binding the contract method 0x6341eb29.
//
// Solidity: function pethHolder(address ) view returns(bool)
func (_FairToken *FairTokenCallerSession) PethHolder(arg0 common.Address) (bool, error) {
	return _FairToken.Contract.PethHolder(&_FairToken.CallOpts, arg0)
}

// PosPool is a free data retrieval call binding the contract method 0x92596d54.
//
// Solidity: function posPool() view returns(uint256)
func (_FairToken *FairTokenCaller) PosPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "posPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PosPool is a free data retrieval call binding the contract method 0x92596d54.
//
// Solidity: function posPool() view returns(uint256)
func (_FairToken *FairTokenSession) PosPool() (*big.Int, error) {
	return _FairToken.Contract.PosPool(&_FairToken.CallOpts)
}

// PosPool is a free data retrieval call binding the contract method 0x92596d54.
//
// Solidity: function posPool() view returns(uint256)
func (_FairToken *FairTokenCallerSession) PosPool() (*big.Int, error) {
	return _FairToken.Contract.PosPool(&_FairToken.CallOpts)
}

// PosPoolStartTime is a free data retrieval call binding the contract method 0x51009286.
//
// Solidity: function posPoolStartTime() view returns(uint256)
func (_FairToken *FairTokenCaller) PosPoolStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "posPoolStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PosPoolStartTime is a free data retrieval call binding the contract method 0x51009286.
//
// Solidity: function posPoolStartTime() view returns(uint256)
func (_FairToken *FairTokenSession) PosPoolStartTime() (*big.Int, error) {
	return _FairToken.Contract.PosPoolStartTime(&_FairToken.CallOpts)
}

// PosPoolStartTime is a free data retrieval call binding the contract method 0x51009286.
//
// Solidity: function posPoolStartTime() view returns(uint256)
func (_FairToken *FairTokenCallerSession) PosPoolStartTime() (*big.Int, error) {
	return _FairToken.Contract.PosPoolStartTime(&_FairToken.CallOpts)
}

// PowPool is a free data retrieval call binding the contract method 0xacd40b93.
//
// Solidity: function powPool() view returns(uint256)
func (_FairToken *FairTokenCaller) PowPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "powPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PowPool is a free data retrieval call binding the contract method 0xacd40b93.
//
// Solidity: function powPool() view returns(uint256)
func (_FairToken *FairTokenSession) PowPool() (*big.Int, error) {
	return _FairToken.Contract.PowPool(&_FairToken.CallOpts)
}

// PowPool is a free data retrieval call binding the contract method 0xacd40b93.
//
// Solidity: function powPool() view returns(uint256)
func (_FairToken *FairTokenCallerSession) PowPool() (*big.Int, error) {
	return _FairToken.Contract.PowPool(&_FairToken.CallOpts)
}

// Precrime is a free data retrieval call binding the contract method 0x950c8a74.
//
// Solidity: function precrime() view returns(address)
func (_FairToken *FairTokenCaller) Precrime(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "precrime")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Precrime is a free data retrieval call binding the contract method 0x950c8a74.
//
// Solidity: function precrime() view returns(address)
func (_FairToken *FairTokenSession) Precrime() (common.Address, error) {
	return _FairToken.Contract.Precrime(&_FairToken.CallOpts)
}

// Precrime is a free data retrieval call binding the contract method 0x950c8a74.
//
// Solidity: function precrime() view returns(address)
func (_FairToken *FairTokenCallerSession) Precrime() (common.Address, error) {
	return _FairToken.Contract.Precrime(&_FairToken.CallOpts)
}

// ReferrerCounts is a free data retrieval call binding the contract method 0xacbfeb28.
//
// Solidity: function referrerCounts(address ) view returns(uint256)
func (_FairToken *FairTokenCaller) ReferrerCounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "referrerCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReferrerCounts is a free data retrieval call binding the contract method 0xacbfeb28.
//
// Solidity: function referrerCounts(address ) view returns(uint256)
func (_FairToken *FairTokenSession) ReferrerCounts(arg0 common.Address) (*big.Int, error) {
	return _FairToken.Contract.ReferrerCounts(&_FairToken.CallOpts, arg0)
}

// ReferrerCounts is a free data retrieval call binding the contract method 0xacbfeb28.
//
// Solidity: function referrerCounts(address ) view returns(uint256)
func (_FairToken *FairTokenCallerSession) ReferrerCounts(arg0 common.Address) (*big.Int, error) {
	return _FairToken.Contract.ReferrerCounts(&_FairToken.CallOpts, arg0)
}

// ReferrerRewards is a free data retrieval call binding the contract method 0x84e311a4.
//
// Solidity: function referrerRewards(address ) view returns(uint256)
func (_FairToken *FairTokenCaller) ReferrerRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "referrerRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReferrerRewards is a free data retrieval call binding the contract method 0x84e311a4.
//
// Solidity: function referrerRewards(address ) view returns(uint256)
func (_FairToken *FairTokenSession) ReferrerRewards(arg0 common.Address) (*big.Int, error) {
	return _FairToken.Contract.ReferrerRewards(&_FairToken.CallOpts, arg0)
}

// ReferrerRewards is a free data retrieval call binding the contract method 0x84e311a4.
//
// Solidity: function referrerRewards(address ) view returns(uint256)
func (_FairToken *FairTokenCallerSession) ReferrerRewards(arg0 common.Address) (*big.Int, error) {
	return _FairToken.Contract.ReferrerRewards(&_FairToken.CallOpts, arg0)
}

// Referrers is a free data retrieval call binding the contract method 0x4a3b68cc.
//
// Solidity: function referrers(address ) view returns(address)
func (_FairToken *FairTokenCaller) Referrers(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "referrers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Referrers is a free data retrieval call binding the contract method 0x4a3b68cc.
//
// Solidity: function referrers(address ) view returns(address)
func (_FairToken *FairTokenSession) Referrers(arg0 common.Address) (common.Address, error) {
	return _FairToken.Contract.Referrers(&_FairToken.CallOpts, arg0)
}

// Referrers is a free data retrieval call binding the contract method 0x4a3b68cc.
//
// Solidity: function referrers(address ) view returns(address)
func (_FairToken *FairTokenCallerSession) Referrers(arg0 common.Address) (common.Address, error) {
	return _FairToken.Contract.Referrers(&_FairToken.CallOpts, arg0)
}

// SellFee is a free data retrieval call binding the contract method 0x2b14ca56.
//
// Solidity: function sellFee() view returns(uint256)
func (_FairToken *FairTokenCaller) SellFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "sellFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellFee is a free data retrieval call binding the contract method 0x2b14ca56.
//
// Solidity: function sellFee() view returns(uint256)
func (_FairToken *FairTokenSession) SellFee() (*big.Int, error) {
	return _FairToken.Contract.SellFee(&_FairToken.CallOpts)
}

// SellFee is a free data retrieval call binding the contract method 0x2b14ca56.
//
// Solidity: function sellFee() view returns(uint256)
func (_FairToken *FairTokenCallerSession) SellFee() (*big.Int, error) {
	return _FairToken.Contract.SellFee(&_FairToken.CallOpts)
}

// Stakes is a free data retrieval call binding the contract method 0x584b62a1.
//
// Solidity: function stakes(address , uint256 ) view returns(uint256 amount, uint256 startTime, uint256 rewardPaid, uint256 rewardMultiplier)
func (_FairToken *FairTokenCaller) Stakes(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Amount           *big.Int
	StartTime        *big.Int
	RewardPaid       *big.Int
	RewardMultiplier *big.Int
}, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "stakes", arg0, arg1)

	outstruct := new(struct {
		Amount           *big.Int
		StartTime        *big.Int
		RewardPaid       *big.Int
		RewardMultiplier *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RewardPaid = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RewardMultiplier = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Stakes is a free data retrieval call binding the contract method 0x584b62a1.
//
// Solidity: function stakes(address , uint256 ) view returns(uint256 amount, uint256 startTime, uint256 rewardPaid, uint256 rewardMultiplier)
func (_FairToken *FairTokenSession) Stakes(arg0 common.Address, arg1 *big.Int) (struct {
	Amount           *big.Int
	StartTime        *big.Int
	RewardPaid       *big.Int
	RewardMultiplier *big.Int
}, error) {
	return _FairToken.Contract.Stakes(&_FairToken.CallOpts, arg0, arg1)
}

// Stakes is a free data retrieval call binding the contract method 0x584b62a1.
//
// Solidity: function stakes(address , uint256 ) view returns(uint256 amount, uint256 startTime, uint256 rewardPaid, uint256 rewardMultiplier)
func (_FairToken *FairTokenCallerSession) Stakes(arg0 common.Address, arg1 *big.Int) (struct {
	Amount           *big.Int
	StartTime        *big.Int
	RewardPaid       *big.Int
	RewardMultiplier *big.Int
}, error) {
	return _FairToken.Contract.Stakes(&_FairToken.CallOpts, arg0, arg1)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FairToken *FairTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FairToken *FairTokenSession) Symbol() (string, error) {
	return _FairToken.Contract.Symbol(&_FairToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FairToken *FairTokenCallerSession) Symbol() (string, error) {
	return _FairToken.Contract.Symbol(&_FairToken.CallOpts)
}

// TotalMined is a free data retrieval call binding the contract method 0x5556db65.
//
// Solidity: function totalMined() view returns(uint256)
func (_FairToken *FairTokenCaller) TotalMined(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "totalMined")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMined is a free data retrieval call binding the contract method 0x5556db65.
//
// Solidity: function totalMined() view returns(uint256)
func (_FairToken *FairTokenSession) TotalMined() (*big.Int, error) {
	return _FairToken.Contract.TotalMined(&_FairToken.CallOpts)
}

// TotalMined is a free data retrieval call binding the contract method 0x5556db65.
//
// Solidity: function totalMined() view returns(uint256)
func (_FairToken *FairTokenCallerSession) TotalMined() (*big.Int, error) {
	return _FairToken.Contract.TotalMined(&_FairToken.CallOpts)
}

// TotalRewardsDistributed is a free data retrieval call binding the contract method 0xee172546.
//
// Solidity: function totalRewardsDistributed() view returns(uint256)
func (_FairToken *FairTokenCaller) TotalRewardsDistributed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "totalRewardsDistributed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRewardsDistributed is a free data retrieval call binding the contract method 0xee172546.
//
// Solidity: function totalRewardsDistributed() view returns(uint256)
func (_FairToken *FairTokenSession) TotalRewardsDistributed() (*big.Int, error) {
	return _FairToken.Contract.TotalRewardsDistributed(&_FairToken.CallOpts)
}

// TotalRewardsDistributed is a free data retrieval call binding the contract method 0xee172546.
//
// Solidity: function totalRewardsDistributed() view returns(uint256)
func (_FairToken *FairTokenCallerSession) TotalRewardsDistributed() (*big.Int, error) {
	return _FairToken.Contract.TotalRewardsDistributed(&_FairToken.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_FairToken *FairTokenCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_FairToken *FairTokenSession) TotalStaked() (*big.Int, error) {
	return _FairToken.Contract.TotalStaked(&_FairToken.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_FairToken *FairTokenCallerSession) TotalStaked() (*big.Int, error) {
	return _FairToken.Contract.TotalStaked(&_FairToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FairToken *FairTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FairToken *FairTokenSession) TotalSupply() (*big.Int, error) {
	return _FairToken.Contract.TotalSupply(&_FairToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FairToken *FairTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _FairToken.Contract.TotalSupply(&_FairToken.CallOpts)
}

// TransferEnabled is a free data retrieval call binding the contract method 0x4cd412d5.
//
// Solidity: function transferEnabled() view returns(bool)
func (_FairToken *FairTokenCaller) TransferEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "transferEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransferEnabled is a free data retrieval call binding the contract method 0x4cd412d5.
//
// Solidity: function transferEnabled() view returns(bool)
func (_FairToken *FairTokenSession) TransferEnabled() (bool, error) {
	return _FairToken.Contract.TransferEnabled(&_FairToken.CallOpts)
}

// TransferEnabled is a free data retrieval call binding the contract method 0x4cd412d5.
//
// Solidity: function transferEnabled() view returns(bool)
func (_FairToken *FairTokenCallerSession) TransferEnabled() (bool, error) {
	return _FairToken.Contract.TransferEnabled(&_FairToken.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_FairToken *FairTokenCaller) TreasuryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "treasuryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_FairToken *FairTokenSession) TreasuryAddress() (common.Address, error) {
	return _FairToken.Contract.TreasuryAddress(&_FairToken.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_FairToken *FairTokenCallerSession) TreasuryAddress() (common.Address, error) {
	return _FairToken.Contract.TreasuryAddress(&_FairToken.CallOpts)
}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_FairToken *FairTokenCaller) TrustedRemoteLookup(opts *bind.CallOpts, arg0 uint16) ([]byte, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "trustedRemoteLookup", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_FairToken *FairTokenSession) TrustedRemoteLookup(arg0 uint16) ([]byte, error) {
	return _FairToken.Contract.TrustedRemoteLookup(&_FairToken.CallOpts, arg0)
}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_FairToken *FairTokenCallerSession) TrustedRemoteLookup(arg0 uint16) ([]byte, error) {
	return _FairToken.Contract.TrustedRemoteLookup(&_FairToken.CallOpts, arg0)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_FairToken *FairTokenCaller) UniswapV2Pair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "uniswapV2Pair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_FairToken *FairTokenSession) UniswapV2Pair() (common.Address, error) {
	return _FairToken.Contract.UniswapV2Pair(&_FairToken.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_FairToken *FairTokenCallerSession) UniswapV2Pair() (common.Address, error) {
	return _FairToken.Contract.UniswapV2Pair(&_FairToken.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_FairToken *FairTokenCaller) UniswapV2Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "uniswapV2Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_FairToken *FairTokenSession) UniswapV2Router() (common.Address, error) {
	return _FairToken.Contract.UniswapV2Router(&_FairToken.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_FairToken *FairTokenCallerSession) UniswapV2Router() (common.Address, error) {
	return _FairToken.Contract.UniswapV2Router(&_FairToken.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_FairToken *FairTokenCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FairToken.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_FairToken *FairTokenSession) Weth() (common.Address, error) {
	return _FairToken.Contract.Weth(&_FairToken.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_FairToken *FairTokenCallerSession) Weth() (common.Address, error) {
	return _FairToken.Contract.Weth(&_FairToken.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9cd441da.
//
// Solidity: function addLiquidity(uint256 tokenAmount, uint256 wethAmount) returns()
func (_FairToken *FairTokenTransactor) AddLiquidity(opts *bind.TransactOpts, tokenAmount *big.Int, wethAmount *big.Int) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "addLiquidity", tokenAmount, wethAmount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9cd441da.
//
// Solidity: function addLiquidity(uint256 tokenAmount, uint256 wethAmount) returns()
func (_FairToken *FairTokenSession) AddLiquidity(tokenAmount *big.Int, wethAmount *big.Int) (*types.Transaction, error) {
	return _FairToken.Contract.AddLiquidity(&_FairToken.TransactOpts, tokenAmount, wethAmount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9cd441da.
//
// Solidity: function addLiquidity(uint256 tokenAmount, uint256 wethAmount) returns()
func (_FairToken *FairTokenTransactorSession) AddLiquidity(tokenAmount *big.Int, wethAmount *big.Int) (*types.Transaction, error) {
	return _FairToken.Contract.AddLiquidity(&_FairToken.TransactOpts, tokenAmount, wethAmount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_FairToken *FairTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_FairToken *FairTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _FairToken.Contract.Approve(&_FairToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_FairToken *FairTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _FairToken.Contract.Approve(&_FairToken.TransactOpts, spender, value)
}

// BuybackFairTokens is a paid mutator transaction binding the contract method 0x6bffccb1.
//
// Solidity: function buybackFairTokens(uint256 wethAmount) returns()
func (_FairToken *FairTokenTransactor) BuybackFairTokens(opts *bind.TransactOpts, wethAmount *big.Int) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "buybackFairTokens", wethAmount)
}

// BuybackFairTokens is a paid mutator transaction binding the contract method 0x6bffccb1.
//
// Solidity: function buybackFairTokens(uint256 wethAmount) returns()
func (_FairToken *FairTokenSession) BuybackFairTokens(wethAmount *big.Int) (*types.Transaction, error) {
	return _FairToken.Contract.BuybackFairTokens(&_FairToken.TransactOpts, wethAmount)
}

// BuybackFairTokens is a paid mutator transaction binding the contract method 0x6bffccb1.
//
// Solidity: function buybackFairTokens(uint256 wethAmount) returns()
func (_FairToken *FairTokenTransactorSession) BuybackFairTokens(wethAmount *big.Int) (*types.Transaction, error) {
	return _FairToken.Contract.BuybackFairTokens(&_FairToken.TransactOpts, wethAmount)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_FairToken *FairTokenTransactor) ClaimReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "claimReward")
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_FairToken *FairTokenSession) ClaimReward() (*types.Transaction, error) {
	return _FairToken.Contract.ClaimReward(&_FairToken.TransactOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_FairToken *FairTokenTransactorSession) ClaimReward() (*types.Transaction, error) {
	return _FairToken.Contract.ClaimReward(&_FairToken.TransactOpts)
}

// EnableTransfers is a paid mutator transaction binding the contract method 0xaf35c6c7.
//
// Solidity: function enableTransfers() returns()
func (_FairToken *FairTokenTransactor) EnableTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "enableTransfers")
}

// EnableTransfers is a paid mutator transaction binding the contract method 0xaf35c6c7.
//
// Solidity: function enableTransfers() returns()
func (_FairToken *FairTokenSession) EnableTransfers() (*types.Transaction, error) {
	return _FairToken.Contract.EnableTransfers(&_FairToken.TransactOpts)
}

// EnableTransfers is a paid mutator transaction binding the contract method 0xaf35c6c7.
//
// Solidity: function enableTransfers() returns()
func (_FairToken *FairTokenTransactorSession) EnableTransfers() (*types.Transaction, error) {
	return _FairToken.Contract.EnableTransfers(&_FairToken.TransactOpts)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_FairToken *FairTokenTransactor) ForceResumeReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "forceResumeReceive", _srcChainId, _srcAddress)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_FairToken *FairTokenSession) ForceResumeReceive(_srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _FairToken.Contract.ForceResumeReceive(&_FairToken.TransactOpts, _srcChainId, _srcAddress)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_FairToken *FairTokenTransactorSession) ForceResumeReceive(_srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _FairToken.Contract.ForceResumeReceive(&_FairToken.TransactOpts, _srcChainId, _srcAddress)
}

// ImportPethHolder is a paid mutator transaction binding the contract method 0x28915d89.
//
// Solidity: function importPethHolder(address[] referrerAddresses) returns()
func (_FairToken *FairTokenTransactor) ImportPethHolder(opts *bind.TransactOpts, referrerAddresses []common.Address) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "importPethHolder", referrerAddresses)
}

// ImportPethHolder is a paid mutator transaction binding the contract method 0x28915d89.
//
// Solidity: function importPethHolder(address[] referrerAddresses) returns()
func (_FairToken *FairTokenSession) ImportPethHolder(referrerAddresses []common.Address) (*types.Transaction, error) {
	return _FairToken.Contract.ImportPethHolder(&_FairToken.TransactOpts, referrerAddresses)
}

// ImportPethHolder is a paid mutator transaction binding the contract method 0x28915d89.
//
// Solidity: function importPethHolder(address[] referrerAddresses) returns()
func (_FairToken *FairTokenTransactorSession) ImportPethHolder(referrerAddresses []common.Address) (*types.Transaction, error) {
	return _FairToken.Contract.ImportPethHolder(&_FairToken.TransactOpts, referrerAddresses)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_FairToken *FairTokenTransactor) LzReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "lzReceive", _srcChainId, _srcAddress, _nonce, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_FairToken *FairTokenSession) LzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _FairToken.Contract.LzReceive(&_FairToken.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_FairToken *FairTokenTransactorSession) LzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _FairToken.Contract.LzReceive(&_FairToken.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// Mine is a paid mutator transaction binding the contract method 0x414da005.
//
// Solidity: function mine(uint256 nonce, address referrer) payable returns()
func (_FairToken *FairTokenTransactor) Mine(opts *bind.TransactOpts, nonce *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "mine", nonce, referrer)
}

// Mine is a paid mutator transaction binding the contract method 0x414da005.
//
// Solidity: function mine(uint256 nonce, address referrer) payable returns()
func (_FairToken *FairTokenSession) Mine(nonce *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _FairToken.Contract.Mine(&_FairToken.TransactOpts, nonce, referrer)
}

// Mine is a paid mutator transaction binding the contract method 0x414da005.
//
// Solidity: function mine(uint256 nonce, address referrer) payable returns()
func (_FairToken *FairTokenTransactorSession) Mine(nonce *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _FairToken.Contract.Mine(&_FairToken.TransactOpts, nonce, referrer)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_FairToken *FairTokenTransactor) NonblockingLzReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "nonblockingLzReceive", _srcChainId, _srcAddress, _nonce, _payload)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_FairToken *FairTokenSession) NonblockingLzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _FairToken.Contract.NonblockingLzReceive(&_FairToken.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_FairToken *FairTokenTransactorSession) NonblockingLzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _FairToken.Contract.NonblockingLzReceive(&_FairToken.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FairToken *FairTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FairToken *FairTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _FairToken.Contract.RenounceOwnership(&_FairToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FairToken *FairTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FairToken.Contract.RenounceOwnership(&_FairToken.TransactOpts)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_FairToken *FairTokenTransactor) RetryMessage(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "retryMessage", _srcChainId, _srcAddress, _nonce, _payload)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_FairToken *FairTokenSession) RetryMessage(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _FairToken.Contract.RetryMessage(&_FairToken.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_FairToken *FairTokenTransactorSession) RetryMessage(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _FairToken.Contract.RetryMessage(&_FairToken.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_FairToken *FairTokenTransactor) SetConfig(opts *bind.TransactOpts, _version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "setConfig", _version, _chainId, _configType, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_FairToken *FairTokenSession) SetConfig(_version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _FairToken.Contract.SetConfig(&_FairToken.TransactOpts, _version, _chainId, _configType, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_FairToken *FairTokenTransactorSession) SetConfig(_version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _FairToken.Contract.SetConfig(&_FairToken.TransactOpts, _version, _chainId, _configType, _config)
}

// SetEarlyUnstakeAllowed is a paid mutator transaction binding the contract method 0x31cfc3e1.
//
// Solidity: function setEarlyUnstakeAllowed(bool _allowed) returns()
func (_FairToken *FairTokenTransactor) SetEarlyUnstakeAllowed(opts *bind.TransactOpts, _allowed bool) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "setEarlyUnstakeAllowed", _allowed)
}

// SetEarlyUnstakeAllowed is a paid mutator transaction binding the contract method 0x31cfc3e1.
//
// Solidity: function setEarlyUnstakeAllowed(bool _allowed) returns()
func (_FairToken *FairTokenSession) SetEarlyUnstakeAllowed(_allowed bool) (*types.Transaction, error) {
	return _FairToken.Contract.SetEarlyUnstakeAllowed(&_FairToken.TransactOpts, _allowed)
}

// SetEarlyUnstakeAllowed is a paid mutator transaction binding the contract method 0x31cfc3e1.
//
// Solidity: function setEarlyUnstakeAllowed(bool _allowed) returns()
func (_FairToken *FairTokenTransactorSession) SetEarlyUnstakeAllowed(_allowed bool) (*types.Transaction, error) {
	return _FairToken.Contract.SetEarlyUnstakeAllowed(&_FairToken.TransactOpts, _allowed)
}

// SetMinDstGas is a paid mutator transaction binding the contract method 0xdf2a5b3b.
//
// Solidity: function setMinDstGas(uint16 _dstChainId, uint16 _packetType, uint256 _minGas) returns()
func (_FairToken *FairTokenTransactor) SetMinDstGas(opts *bind.TransactOpts, _dstChainId uint16, _packetType uint16, _minGas *big.Int) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "setMinDstGas", _dstChainId, _packetType, _minGas)
}

// SetMinDstGas is a paid mutator transaction binding the contract method 0xdf2a5b3b.
//
// Solidity: function setMinDstGas(uint16 _dstChainId, uint16 _packetType, uint256 _minGas) returns()
func (_FairToken *FairTokenSession) SetMinDstGas(_dstChainId uint16, _packetType uint16, _minGas *big.Int) (*types.Transaction, error) {
	return _FairToken.Contract.SetMinDstGas(&_FairToken.TransactOpts, _dstChainId, _packetType, _minGas)
}

// SetMinDstGas is a paid mutator transaction binding the contract method 0xdf2a5b3b.
//
// Solidity: function setMinDstGas(uint16 _dstChainId, uint16 _packetType, uint256 _minGas) returns()
func (_FairToken *FairTokenTransactorSession) SetMinDstGas(_dstChainId uint16, _packetType uint16, _minGas *big.Int) (*types.Transaction, error) {
	return _FairToken.Contract.SetMinDstGas(&_FairToken.TransactOpts, _dstChainId, _packetType, _minGas)
}

// SetPayloadSizeLimit is a paid mutator transaction binding the contract method 0x0df37483.
//
// Solidity: function setPayloadSizeLimit(uint16 _dstChainId, uint256 _size) returns()
func (_FairToken *FairTokenTransactor) SetPayloadSizeLimit(opts *bind.TransactOpts, _dstChainId uint16, _size *big.Int) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "setPayloadSizeLimit", _dstChainId, _size)
}

// SetPayloadSizeLimit is a paid mutator transaction binding the contract method 0x0df37483.
//
// Solidity: function setPayloadSizeLimit(uint16 _dstChainId, uint256 _size) returns()
func (_FairToken *FairTokenSession) SetPayloadSizeLimit(_dstChainId uint16, _size *big.Int) (*types.Transaction, error) {
	return _FairToken.Contract.SetPayloadSizeLimit(&_FairToken.TransactOpts, _dstChainId, _size)
}

// SetPayloadSizeLimit is a paid mutator transaction binding the contract method 0x0df37483.
//
// Solidity: function setPayloadSizeLimit(uint16 _dstChainId, uint256 _size) returns()
func (_FairToken *FairTokenTransactorSession) SetPayloadSizeLimit(_dstChainId uint16, _size *big.Int) (*types.Transaction, error) {
	return _FairToken.Contract.SetPayloadSizeLimit(&_FairToken.TransactOpts, _dstChainId, _size)
}

// SetPrecrime is a paid mutator transaction binding the contract method 0xbaf3292d.
//
// Solidity: function setPrecrime(address _precrime) returns()
func (_FairToken *FairTokenTransactor) SetPrecrime(opts *bind.TransactOpts, _precrime common.Address) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "setPrecrime", _precrime)
}

// SetPrecrime is a paid mutator transaction binding the contract method 0xbaf3292d.
//
// Solidity: function setPrecrime(address _precrime) returns()
func (_FairToken *FairTokenSession) SetPrecrime(_precrime common.Address) (*types.Transaction, error) {
	return _FairToken.Contract.SetPrecrime(&_FairToken.TransactOpts, _precrime)
}

// SetPrecrime is a paid mutator transaction binding the contract method 0xbaf3292d.
//
// Solidity: function setPrecrime(address _precrime) returns()
func (_FairToken *FairTokenTransactorSession) SetPrecrime(_precrime common.Address) (*types.Transaction, error) {
	return _FairToken.Contract.SetPrecrime(&_FairToken.TransactOpts, _precrime)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_FairToken *FairTokenTransactor) SetReceiveVersion(opts *bind.TransactOpts, _version uint16) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "setReceiveVersion", _version)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_FairToken *FairTokenSession) SetReceiveVersion(_version uint16) (*types.Transaction, error) {
	return _FairToken.Contract.SetReceiveVersion(&_FairToken.TransactOpts, _version)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_FairToken *FairTokenTransactorSession) SetReceiveVersion(_version uint16) (*types.Transaction, error) {
	return _FairToken.Contract.SetReceiveVersion(&_FairToken.TransactOpts, _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_FairToken *FairTokenTransactor) SetSendVersion(opts *bind.TransactOpts, _version uint16) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "setSendVersion", _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_FairToken *FairTokenSession) SetSendVersion(_version uint16) (*types.Transaction, error) {
	return _FairToken.Contract.SetSendVersion(&_FairToken.TransactOpts, _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_FairToken *FairTokenTransactorSession) SetSendVersion(_version uint16) (*types.Transaction, error) {
	return _FairToken.Contract.SetSendVersion(&_FairToken.TransactOpts, _version)
}

// SetTreasuryAddress is a paid mutator transaction binding the contract method 0x6605bfda.
//
// Solidity: function setTreasuryAddress(address _treasuryAddress) returns()
func (_FairToken *FairTokenTransactor) SetTreasuryAddress(opts *bind.TransactOpts, _treasuryAddress common.Address) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "setTreasuryAddress", _treasuryAddress)
}

// SetTreasuryAddress is a paid mutator transaction binding the contract method 0x6605bfda.
//
// Solidity: function setTreasuryAddress(address _treasuryAddress) returns()
func (_FairToken *FairTokenSession) SetTreasuryAddress(_treasuryAddress common.Address) (*types.Transaction, error) {
	return _FairToken.Contract.SetTreasuryAddress(&_FairToken.TransactOpts, _treasuryAddress)
}

// SetTreasuryAddress is a paid mutator transaction binding the contract method 0x6605bfda.
//
// Solidity: function setTreasuryAddress(address _treasuryAddress) returns()
func (_FairToken *FairTokenTransactorSession) SetTreasuryAddress(_treasuryAddress common.Address) (*types.Transaction, error) {
	return _FairToken.Contract.SetTreasuryAddress(&_FairToken.TransactOpts, _treasuryAddress)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _remoteChainId, bytes _path) returns()
func (_FairToken *FairTokenTransactor) SetTrustedRemote(opts *bind.TransactOpts, _remoteChainId uint16, _path []byte) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "setTrustedRemote", _remoteChainId, _path)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _remoteChainId, bytes _path) returns()
func (_FairToken *FairTokenSession) SetTrustedRemote(_remoteChainId uint16, _path []byte) (*types.Transaction, error) {
	return _FairToken.Contract.SetTrustedRemote(&_FairToken.TransactOpts, _remoteChainId, _path)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _remoteChainId, bytes _path) returns()
func (_FairToken *FairTokenTransactorSession) SetTrustedRemote(_remoteChainId uint16, _path []byte) (*types.Transaction, error) {
	return _FairToken.Contract.SetTrustedRemote(&_FairToken.TransactOpts, _remoteChainId, _path)
}

// SetTrustedRemoteAddress is a paid mutator transaction binding the contract method 0xa6c3d165.
//
// Solidity: function setTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress) returns()
func (_FairToken *FairTokenTransactor) SetTrustedRemoteAddress(opts *bind.TransactOpts, _remoteChainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "setTrustedRemoteAddress", _remoteChainId, _remoteAddress)
}

// SetTrustedRemoteAddress is a paid mutator transaction binding the contract method 0xa6c3d165.
//
// Solidity: function setTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress) returns()
func (_FairToken *FairTokenSession) SetTrustedRemoteAddress(_remoteChainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _FairToken.Contract.SetTrustedRemoteAddress(&_FairToken.TransactOpts, _remoteChainId, _remoteAddress)
}

// SetTrustedRemoteAddress is a paid mutator transaction binding the contract method 0xa6c3d165.
//
// Solidity: function setTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress) returns()
func (_FairToken *FairTokenTransactorSession) SetTrustedRemoteAddress(_remoteChainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _FairToken.Contract.SetTrustedRemoteAddress(&_FairToken.TransactOpts, _remoteChainId, _remoteAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_FairToken *FairTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_FairToken *FairTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FairToken.Contract.Transfer(&_FairToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_FairToken *FairTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FairToken.Contract.Transfer(&_FairToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_FairToken *FairTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_FairToken *FairTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FairToken.Contract.TransferFrom(&_FairToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_FairToken *FairTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FairToken.Contract.TransferFrom(&_FairToken.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FairToken *FairTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FairToken *FairTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FairToken.Contract.TransferOwnership(&_FairToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FairToken *FairTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FairToken.Contract.TransferOwnership(&_FairToken.TransactOpts, newOwner)
}

// UnstakeAll is a paid mutator transaction binding the contract method 0x35322f37.
//
// Solidity: function unstakeAll() payable returns()
func (_FairToken *FairTokenTransactor) UnstakeAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "unstakeAll")
}

// UnstakeAll is a paid mutator transaction binding the contract method 0x35322f37.
//
// Solidity: function unstakeAll() payable returns()
func (_FairToken *FairTokenSession) UnstakeAll() (*types.Transaction, error) {
	return _FairToken.Contract.UnstakeAll(&_FairToken.TransactOpts)
}

// UnstakeAll is a paid mutator transaction binding the contract method 0x35322f37.
//
// Solidity: function unstakeAll() payable returns()
func (_FairToken *FairTokenTransactorSession) UnstakeAll() (*types.Transaction, error) {
	return _FairToken.Contract.UnstakeAll(&_FairToken.TransactOpts)
}

// UnstakeOne is a paid mutator transaction binding the contract method 0x2c6cff06.
//
// Solidity: function unstakeOne(uint256 stakeIndex) payable returns()
func (_FairToken *FairTokenTransactor) UnstakeOne(opts *bind.TransactOpts, stakeIndex *big.Int) (*types.Transaction, error) {
	return _FairToken.contract.Transact(opts, "unstakeOne", stakeIndex)
}

// UnstakeOne is a paid mutator transaction binding the contract method 0x2c6cff06.
//
// Solidity: function unstakeOne(uint256 stakeIndex) payable returns()
func (_FairToken *FairTokenSession) UnstakeOne(stakeIndex *big.Int) (*types.Transaction, error) {
	return _FairToken.Contract.UnstakeOne(&_FairToken.TransactOpts, stakeIndex)
}

// UnstakeOne is a paid mutator transaction binding the contract method 0x2c6cff06.
//
// Solidity: function unstakeOne(uint256 stakeIndex) payable returns()
func (_FairToken *FairTokenTransactorSession) UnstakeOne(stakeIndex *big.Int) (*types.Transaction, error) {
	return _FairToken.Contract.UnstakeOne(&_FairToken.TransactOpts, stakeIndex)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FairToken *FairTokenTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FairToken.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FairToken *FairTokenSession) Receive() (*types.Transaction, error) {
	return _FairToken.Contract.Receive(&_FairToken.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FairToken *FairTokenTransactorSession) Receive() (*types.Transaction, error) {
	return _FairToken.Contract.Receive(&_FairToken.TransactOpts)
}

// FairTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FairToken contract.
type FairTokenApprovalIterator struct {
	Event *FairTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FairTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FairTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FairTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTokenApproval represents a Approval event raised by the FairToken contract.
type FairTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FairToken *FairTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*FairTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FairToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &FairTokenApprovalIterator{contract: _FairToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FairToken *FairTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FairTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FairToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTokenApproval)
				if err := _FairToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FairToken *FairTokenFilterer) ParseApproval(log types.Log) (*FairTokenApproval, error) {
	event := new(FairTokenApproval)
	if err := _FairToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTokenBuybackIterator is returned from FilterBuyback and is used to iterate over the raw logs and unpacked data for Buyback events raised by the FairToken contract.
type FairTokenBuybackIterator struct {
	Event *FairTokenBuyback // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FairTokenBuybackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTokenBuyback)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FairTokenBuyback)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FairTokenBuybackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTokenBuybackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTokenBuyback represents a Buyback event raised by the FairToken contract.
type FairTokenBuyback struct {
	Pair             common.Address
	EthSpent         *big.Int
	FairTokensBought *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBuyback is a free log retrieval operation binding the contract event 0x2dcc2439519c7d06fca9f8ae01e07f4f3c6ca21b5cdf8eff42cb75cf34d223c9.
//
// Solidity: event Buyback(address indexed pair, uint256 ethSpent, uint256 fairTokensBought)
func (_FairToken *FairTokenFilterer) FilterBuyback(opts *bind.FilterOpts, pair []common.Address) (*FairTokenBuybackIterator, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}

	logs, sub, err := _FairToken.contract.FilterLogs(opts, "Buyback", pairRule)
	if err != nil {
		return nil, err
	}
	return &FairTokenBuybackIterator{contract: _FairToken.contract, event: "Buyback", logs: logs, sub: sub}, nil
}

// WatchBuyback is a free log subscription operation binding the contract event 0x2dcc2439519c7d06fca9f8ae01e07f4f3c6ca21b5cdf8eff42cb75cf34d223c9.
//
// Solidity: event Buyback(address indexed pair, uint256 ethSpent, uint256 fairTokensBought)
func (_FairToken *FairTokenFilterer) WatchBuyback(opts *bind.WatchOpts, sink chan<- *FairTokenBuyback, pair []common.Address) (event.Subscription, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}

	logs, sub, err := _FairToken.contract.WatchLogs(opts, "Buyback", pairRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTokenBuyback)
				if err := _FairToken.contract.UnpackLog(event, "Buyback", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBuyback is a log parse operation binding the contract event 0x2dcc2439519c7d06fca9f8ae01e07f4f3c6ca21b5cdf8eff42cb75cf34d223c9.
//
// Solidity: event Buyback(address indexed pair, uint256 ethSpent, uint256 fairTokensBought)
func (_FairToken *FairTokenFilterer) ParseBuyback(log types.Log) (*FairTokenBuyback, error) {
	event := new(FairTokenBuyback)
	if err := _FairToken.contract.UnpackLog(event, "Buyback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTokenChallengeUpdatedIterator is returned from FilterChallengeUpdated and is used to iterate over the raw logs and unpacked data for ChallengeUpdated events raised by the FairToken contract.
type FairTokenChallengeUpdatedIterator struct {
	Event *FairTokenChallengeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FairTokenChallengeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTokenChallengeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FairTokenChallengeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FairTokenChallengeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTokenChallengeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTokenChallengeUpdated represents a ChallengeUpdated event raised by the FairToken contract.
type FairTokenChallengeUpdated struct {
	NewChallenge *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterChallengeUpdated is a free log retrieval operation binding the contract event 0xe2279210f8995d05b27b8bfa6b0780141d7eb986659e40f88b9d872a1a8a280a.
//
// Solidity: event ChallengeUpdated(uint256 newChallenge)
func (_FairToken *FairTokenFilterer) FilterChallengeUpdated(opts *bind.FilterOpts) (*FairTokenChallengeUpdatedIterator, error) {

	logs, sub, err := _FairToken.contract.FilterLogs(opts, "ChallengeUpdated")
	if err != nil {
		return nil, err
	}
	return &FairTokenChallengeUpdatedIterator{contract: _FairToken.contract, event: "ChallengeUpdated", logs: logs, sub: sub}, nil
}

// WatchChallengeUpdated is a free log subscription operation binding the contract event 0xe2279210f8995d05b27b8bfa6b0780141d7eb986659e40f88b9d872a1a8a280a.
//
// Solidity: event ChallengeUpdated(uint256 newChallenge)
func (_FairToken *FairTokenFilterer) WatchChallengeUpdated(opts *bind.WatchOpts, sink chan<- *FairTokenChallengeUpdated) (event.Subscription, error) {

	logs, sub, err := _FairToken.contract.WatchLogs(opts, "ChallengeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTokenChallengeUpdated)
				if err := _FairToken.contract.UnpackLog(event, "ChallengeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChallengeUpdated is a log parse operation binding the contract event 0xe2279210f8995d05b27b8bfa6b0780141d7eb986659e40f88b9d872a1a8a280a.
//
// Solidity: event ChallengeUpdated(uint256 newChallenge)
func (_FairToken *FairTokenFilterer) ParseChallengeUpdated(log types.Log) (*FairTokenChallengeUpdated, error) {
	event := new(FairTokenChallengeUpdated)
	if err := _FairToken.contract.UnpackLog(event, "ChallengeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTokenClaimRewardIterator is returned from FilterClaimReward and is used to iterate over the raw logs and unpacked data for ClaimReward events raised by the FairToken contract.
type FairTokenClaimRewardIterator struct {
	Event *FairTokenClaimReward // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FairTokenClaimRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTokenClaimReward)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FairTokenClaimReward)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FairTokenClaimRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTokenClaimRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTokenClaimReward represents a ClaimReward event raised by the FairToken contract.
type FairTokenClaimReward struct {
	User   common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimReward is a free log retrieval operation binding the contract event 0xe74e5c9d4ac1fc33412485f18c159a0a391efe287ab3fd271123f30e6bacf4e3.
//
// Solidity: event ClaimReward(address indexed user, uint256 amount, uint256 time)
func (_FairToken *FairTokenFilterer) FilterClaimReward(opts *bind.FilterOpts, user []common.Address) (*FairTokenClaimRewardIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FairToken.contract.FilterLogs(opts, "ClaimReward", userRule)
	if err != nil {
		return nil, err
	}
	return &FairTokenClaimRewardIterator{contract: _FairToken.contract, event: "ClaimReward", logs: logs, sub: sub}, nil
}

// WatchClaimReward is a free log subscription operation binding the contract event 0xe74e5c9d4ac1fc33412485f18c159a0a391efe287ab3fd271123f30e6bacf4e3.
//
// Solidity: event ClaimReward(address indexed user, uint256 amount, uint256 time)
func (_FairToken *FairTokenFilterer) WatchClaimReward(opts *bind.WatchOpts, sink chan<- *FairTokenClaimReward, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FairToken.contract.WatchLogs(opts, "ClaimReward", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTokenClaimReward)
				if err := _FairToken.contract.UnpackLog(event, "ClaimReward", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimReward is a log parse operation binding the contract event 0xe74e5c9d4ac1fc33412485f18c159a0a391efe287ab3fd271123f30e6bacf4e3.
//
// Solidity: event ClaimReward(address indexed user, uint256 amount, uint256 time)
func (_FairToken *FairTokenFilterer) ParseClaimReward(log types.Log) (*FairTokenClaimReward, error) {
	event := new(FairTokenClaimReward)
	if err := _FairToken.contract.UnpackLog(event, "ClaimReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTokenDifficultyAdjustedIterator is returned from FilterDifficultyAdjusted and is used to iterate over the raw logs and unpacked data for DifficultyAdjusted events raised by the FairToken contract.
type FairTokenDifficultyAdjustedIterator struct {
	Event *FairTokenDifficultyAdjusted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FairTokenDifficultyAdjustedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTokenDifficultyAdjusted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FairTokenDifficultyAdjusted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FairTokenDifficultyAdjustedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTokenDifficultyAdjustedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTokenDifficultyAdjusted represents a DifficultyAdjusted event raised by the FairToken contract.
type FairTokenDifficultyAdjusted struct {
	NewDifficulty *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDifficultyAdjusted is a free log retrieval operation binding the contract event 0x42ebc15644fe1ca39e719c989e249095a01e7bf6c4e2ea33889a7e33a7664caa.
//
// Solidity: event DifficultyAdjusted(uint256 newDifficulty)
func (_FairToken *FairTokenFilterer) FilterDifficultyAdjusted(opts *bind.FilterOpts) (*FairTokenDifficultyAdjustedIterator, error) {

	logs, sub, err := _FairToken.contract.FilterLogs(opts, "DifficultyAdjusted")
	if err != nil {
		return nil, err
	}
	return &FairTokenDifficultyAdjustedIterator{contract: _FairToken.contract, event: "DifficultyAdjusted", logs: logs, sub: sub}, nil
}

// WatchDifficultyAdjusted is a free log subscription operation binding the contract event 0x42ebc15644fe1ca39e719c989e249095a01e7bf6c4e2ea33889a7e33a7664caa.
//
// Solidity: event DifficultyAdjusted(uint256 newDifficulty)
func (_FairToken *FairTokenFilterer) WatchDifficultyAdjusted(opts *bind.WatchOpts, sink chan<- *FairTokenDifficultyAdjusted) (event.Subscription, error) {

	logs, sub, err := _FairToken.contract.WatchLogs(opts, "DifficultyAdjusted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTokenDifficultyAdjusted)
				if err := _FairToken.contract.UnpackLog(event, "DifficultyAdjusted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDifficultyAdjusted is a log parse operation binding the contract event 0x42ebc15644fe1ca39e719c989e249095a01e7bf6c4e2ea33889a7e33a7664caa.
//
// Solidity: event DifficultyAdjusted(uint256 newDifficulty)
func (_FairToken *FairTokenFilterer) ParseDifficultyAdjusted(log types.Log) (*FairTokenDifficultyAdjusted, error) {
	event := new(FairTokenDifficultyAdjusted)
	if err := _FairToken.contract.UnpackLog(event, "DifficultyAdjusted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTokenFeePaidIterator is returned from FilterFeePaid and is used to iterate over the raw logs and unpacked data for FeePaid events raised by the FairToken contract.
type FairTokenFeePaidIterator struct {
	Event *FairTokenFeePaid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FairTokenFeePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTokenFeePaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FairTokenFeePaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FairTokenFeePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTokenFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTokenFeePaid represents a FeePaid event raised by the FairToken contract.
type FairTokenFeePaid struct {
	User      common.Address
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeePaid is a free log retrieval operation binding the contract event 0x075a2720282fdf622141dae0b048ef90a21a7e57c134c76912d19d006b3b3f6f.
//
// Solidity: event FeePaid(address indexed user, uint256 feeAmount)
func (_FairToken *FairTokenFilterer) FilterFeePaid(opts *bind.FilterOpts, user []common.Address) (*FairTokenFeePaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FairToken.contract.FilterLogs(opts, "FeePaid", userRule)
	if err != nil {
		return nil, err
	}
	return &FairTokenFeePaidIterator{contract: _FairToken.contract, event: "FeePaid", logs: logs, sub: sub}, nil
}

// WatchFeePaid is a free log subscription operation binding the contract event 0x075a2720282fdf622141dae0b048ef90a21a7e57c134c76912d19d006b3b3f6f.
//
// Solidity: event FeePaid(address indexed user, uint256 feeAmount)
func (_FairToken *FairTokenFilterer) WatchFeePaid(opts *bind.WatchOpts, sink chan<- *FairTokenFeePaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FairToken.contract.WatchLogs(opts, "FeePaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTokenFeePaid)
				if err := _FairToken.contract.UnpackLog(event, "FeePaid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeePaid is a log parse operation binding the contract event 0x075a2720282fdf622141dae0b048ef90a21a7e57c134c76912d19d006b3b3f6f.
//
// Solidity: event FeePaid(address indexed user, uint256 feeAmount)
func (_FairToken *FairTokenFilterer) ParseFeePaid(log types.Log) (*FairTokenFeePaid, error) {
	event := new(FairTokenFeePaid)
	if err := _FairToken.contract.UnpackLog(event, "FeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTokenMessageFailedIterator is returned from FilterMessageFailed and is used to iterate over the raw logs and unpacked data for MessageFailed events raised by the FairToken contract.
type FairTokenMessageFailedIterator struct {
	Event *FairTokenMessageFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FairTokenMessageFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTokenMessageFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FairTokenMessageFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FairTokenMessageFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTokenMessageFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTokenMessageFailed represents a MessageFailed event raised by the FairToken contract.
type FairTokenMessageFailed struct {
	SrcChainId uint16
	SrcAddress []byte
	Nonce      uint64
	Payload    []byte
	Reason     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessageFailed is a free log retrieval operation binding the contract event 0xe183f33de2837795525b4792ca4cd60535bd77c53b7e7030060bfcf5734d6b0c.
//
// Solidity: event MessageFailed(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload, bytes _reason)
func (_FairToken *FairTokenFilterer) FilterMessageFailed(opts *bind.FilterOpts) (*FairTokenMessageFailedIterator, error) {

	logs, sub, err := _FairToken.contract.FilterLogs(opts, "MessageFailed")
	if err != nil {
		return nil, err
	}
	return &FairTokenMessageFailedIterator{contract: _FairToken.contract, event: "MessageFailed", logs: logs, sub: sub}, nil
}

// WatchMessageFailed is a free log subscription operation binding the contract event 0xe183f33de2837795525b4792ca4cd60535bd77c53b7e7030060bfcf5734d6b0c.
//
// Solidity: event MessageFailed(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload, bytes _reason)
func (_FairToken *FairTokenFilterer) WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *FairTokenMessageFailed) (event.Subscription, error) {

	logs, sub, err := _FairToken.contract.WatchLogs(opts, "MessageFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTokenMessageFailed)
				if err := _FairToken.contract.UnpackLog(event, "MessageFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageFailed is a log parse operation binding the contract event 0xe183f33de2837795525b4792ca4cd60535bd77c53b7e7030060bfcf5734d6b0c.
//
// Solidity: event MessageFailed(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload, bytes _reason)
func (_FairToken *FairTokenFilterer) ParseMessageFailed(log types.Log) (*FairTokenMessageFailed, error) {
	event := new(FairTokenMessageFailed)
	if err := _FairToken.contract.UnpackLog(event, "MessageFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTokenMiningSuccessIterator is returned from FilterMiningSuccess and is used to iterate over the raw logs and unpacked data for MiningSuccess events raised by the FairToken contract.
type FairTokenMiningSuccessIterator struct {
	Event *FairTokenMiningSuccess // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FairTokenMiningSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTokenMiningSuccess)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FairTokenMiningSuccess)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FairTokenMiningSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTokenMiningSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTokenMiningSuccess represents a MiningSuccess event raised by the FairToken contract.
type FairTokenMiningSuccess struct {
	From      common.Address
	Reward    *big.Int
	Challenge *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMiningSuccess is a free log retrieval operation binding the contract event 0xc16c7fcee27628d9bb8c56bb85faf8e8035e75caf7c467ffdfb993ebf821758c.
//
// Solidity: event MiningSuccess(address indexed from, uint256 reward, uint256 challenge)
func (_FairToken *FairTokenFilterer) FilterMiningSuccess(opts *bind.FilterOpts, from []common.Address) (*FairTokenMiningSuccessIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _FairToken.contract.FilterLogs(opts, "MiningSuccess", fromRule)
	if err != nil {
		return nil, err
	}
	return &FairTokenMiningSuccessIterator{contract: _FairToken.contract, event: "MiningSuccess", logs: logs, sub: sub}, nil
}

// WatchMiningSuccess is a free log subscription operation binding the contract event 0xc16c7fcee27628d9bb8c56bb85faf8e8035e75caf7c467ffdfb993ebf821758c.
//
// Solidity: event MiningSuccess(address indexed from, uint256 reward, uint256 challenge)
func (_FairToken *FairTokenFilterer) WatchMiningSuccess(opts *bind.WatchOpts, sink chan<- *FairTokenMiningSuccess, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _FairToken.contract.WatchLogs(opts, "MiningSuccess", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTokenMiningSuccess)
				if err := _FairToken.contract.UnpackLog(event, "MiningSuccess", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMiningSuccess is a log parse operation binding the contract event 0xc16c7fcee27628d9bb8c56bb85faf8e8035e75caf7c467ffdfb993ebf821758c.
//
// Solidity: event MiningSuccess(address indexed from, uint256 reward, uint256 challenge)
func (_FairToken *FairTokenFilterer) ParseMiningSuccess(log types.Log) (*FairTokenMiningSuccess, error) {
	event := new(FairTokenMiningSuccess)
	if err := _FairToken.contract.UnpackLog(event, "MiningSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FairToken contract.
type FairTokenOwnershipTransferredIterator struct {
	Event *FairTokenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FairTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTokenOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FairTokenOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FairTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTokenOwnershipTransferred represents a OwnershipTransferred event raised by the FairToken contract.
type FairTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FairToken *FairTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FairTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FairToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FairTokenOwnershipTransferredIterator{contract: _FairToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FairToken *FairTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FairTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FairToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTokenOwnershipTransferred)
				if err := _FairToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FairToken *FairTokenFilterer) ParseOwnershipTransferred(log types.Log) (*FairTokenOwnershipTransferred, error) {
	event := new(FairTokenOwnershipTransferred)
	if err := _FairToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTokenReferrerSetIterator is returned from FilterReferrerSet and is used to iterate over the raw logs and unpacked data for ReferrerSet events raised by the FairToken contract.
type FairTokenReferrerSetIterator struct {
	Event *FairTokenReferrerSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FairTokenReferrerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTokenReferrerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FairTokenReferrerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FairTokenReferrerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTokenReferrerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTokenReferrerSet represents a ReferrerSet event raised by the FairToken contract.
type FairTokenReferrerSet struct {
	User     common.Address
	Referrer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReferrerSet is a free log retrieval operation binding the contract event 0x5f7165288eef601591cf549e15ff19ef9060b7f71b9c115be946fa1fe7ebf68a.
//
// Solidity: event ReferrerSet(address indexed user, address indexed referrer)
func (_FairToken *FairTokenFilterer) FilterReferrerSet(opts *bind.FilterOpts, user []common.Address, referrer []common.Address) (*FairTokenReferrerSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var referrerRule []interface{}
	for _, referrerItem := range referrer {
		referrerRule = append(referrerRule, referrerItem)
	}

	logs, sub, err := _FairToken.contract.FilterLogs(opts, "ReferrerSet", userRule, referrerRule)
	if err != nil {
		return nil, err
	}
	return &FairTokenReferrerSetIterator{contract: _FairToken.contract, event: "ReferrerSet", logs: logs, sub: sub}, nil
}

// WatchReferrerSet is a free log subscription operation binding the contract event 0x5f7165288eef601591cf549e15ff19ef9060b7f71b9c115be946fa1fe7ebf68a.
//
// Solidity: event ReferrerSet(address indexed user, address indexed referrer)
func (_FairToken *FairTokenFilterer) WatchReferrerSet(opts *bind.WatchOpts, sink chan<- *FairTokenReferrerSet, user []common.Address, referrer []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var referrerRule []interface{}
	for _, referrerItem := range referrer {
		referrerRule = append(referrerRule, referrerItem)
	}

	logs, sub, err := _FairToken.contract.WatchLogs(opts, "ReferrerSet", userRule, referrerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTokenReferrerSet)
				if err := _FairToken.contract.UnpackLog(event, "ReferrerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReferrerSet is a log parse operation binding the contract event 0x5f7165288eef601591cf549e15ff19ef9060b7f71b9c115be946fa1fe7ebf68a.
//
// Solidity: event ReferrerSet(address indexed user, address indexed referrer)
func (_FairToken *FairTokenFilterer) ParseReferrerSet(log types.Log) (*FairTokenReferrerSet, error) {
	event := new(FairTokenReferrerSet)
	if err := _FairToken.contract.UnpackLog(event, "ReferrerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTokenRetryMessageSuccessIterator is returned from FilterRetryMessageSuccess and is used to iterate over the raw logs and unpacked data for RetryMessageSuccess events raised by the FairToken contract.
type FairTokenRetryMessageSuccessIterator struct {
	Event *FairTokenRetryMessageSuccess // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FairTokenRetryMessageSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTokenRetryMessageSuccess)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FairTokenRetryMessageSuccess)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FairTokenRetryMessageSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTokenRetryMessageSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTokenRetryMessageSuccess represents a RetryMessageSuccess event raised by the FairToken contract.
type FairTokenRetryMessageSuccess struct {
	SrcChainId  uint16
	SrcAddress  []byte
	Nonce       uint64
	PayloadHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRetryMessageSuccess is a free log retrieval operation binding the contract event 0xc264d91f3adc5588250e1551f547752ca0cfa8f6b530d243b9f9f4cab10ea8e5.
//
// Solidity: event RetryMessageSuccess(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes32 _payloadHash)
func (_FairToken *FairTokenFilterer) FilterRetryMessageSuccess(opts *bind.FilterOpts) (*FairTokenRetryMessageSuccessIterator, error) {

	logs, sub, err := _FairToken.contract.FilterLogs(opts, "RetryMessageSuccess")
	if err != nil {
		return nil, err
	}
	return &FairTokenRetryMessageSuccessIterator{contract: _FairToken.contract, event: "RetryMessageSuccess", logs: logs, sub: sub}, nil
}

// WatchRetryMessageSuccess is a free log subscription operation binding the contract event 0xc264d91f3adc5588250e1551f547752ca0cfa8f6b530d243b9f9f4cab10ea8e5.
//
// Solidity: event RetryMessageSuccess(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes32 _payloadHash)
func (_FairToken *FairTokenFilterer) WatchRetryMessageSuccess(opts *bind.WatchOpts, sink chan<- *FairTokenRetryMessageSuccess) (event.Subscription, error) {

	logs, sub, err := _FairToken.contract.WatchLogs(opts, "RetryMessageSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTokenRetryMessageSuccess)
				if err := _FairToken.contract.UnpackLog(event, "RetryMessageSuccess", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRetryMessageSuccess is a log parse operation binding the contract event 0xc264d91f3adc5588250e1551f547752ca0cfa8f6b530d243b9f9f4cab10ea8e5.
//
// Solidity: event RetryMessageSuccess(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes32 _payloadHash)
func (_FairToken *FairTokenFilterer) ParseRetryMessageSuccess(log types.Log) (*FairTokenRetryMessageSuccess, error) {
	event := new(FairTokenRetryMessageSuccess)
	if err := _FairToken.contract.UnpackLog(event, "RetryMessageSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTokenSetMinDstGasIterator is returned from FilterSetMinDstGas and is used to iterate over the raw logs and unpacked data for SetMinDstGas events raised by the FairToken contract.
type FairTokenSetMinDstGasIterator struct {
	Event *FairTokenSetMinDstGas // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FairTokenSetMinDstGasIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTokenSetMinDstGas)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FairTokenSetMinDstGas)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FairTokenSetMinDstGasIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTokenSetMinDstGasIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTokenSetMinDstGas represents a SetMinDstGas event raised by the FairToken contract.
type FairTokenSetMinDstGas struct {
	DstChainId uint16
	Type       uint16
	MinDstGas  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetMinDstGas is a free log retrieval operation binding the contract event 0x9d5c7c0b934da8fefa9c7760c98383778a12dfbfc0c3b3106518f43fb9508ac0.
//
// Solidity: event SetMinDstGas(uint16 _dstChainId, uint16 _type, uint256 _minDstGas)
func (_FairToken *FairTokenFilterer) FilterSetMinDstGas(opts *bind.FilterOpts) (*FairTokenSetMinDstGasIterator, error) {

	logs, sub, err := _FairToken.contract.FilterLogs(opts, "SetMinDstGas")
	if err != nil {
		return nil, err
	}
	return &FairTokenSetMinDstGasIterator{contract: _FairToken.contract, event: "SetMinDstGas", logs: logs, sub: sub}, nil
}

// WatchSetMinDstGas is a free log subscription operation binding the contract event 0x9d5c7c0b934da8fefa9c7760c98383778a12dfbfc0c3b3106518f43fb9508ac0.
//
// Solidity: event SetMinDstGas(uint16 _dstChainId, uint16 _type, uint256 _minDstGas)
func (_FairToken *FairTokenFilterer) WatchSetMinDstGas(opts *bind.WatchOpts, sink chan<- *FairTokenSetMinDstGas) (event.Subscription, error) {

	logs, sub, err := _FairToken.contract.WatchLogs(opts, "SetMinDstGas")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTokenSetMinDstGas)
				if err := _FairToken.contract.UnpackLog(event, "SetMinDstGas", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetMinDstGas is a log parse operation binding the contract event 0x9d5c7c0b934da8fefa9c7760c98383778a12dfbfc0c3b3106518f43fb9508ac0.
//
// Solidity: event SetMinDstGas(uint16 _dstChainId, uint16 _type, uint256 _minDstGas)
func (_FairToken *FairTokenFilterer) ParseSetMinDstGas(log types.Log) (*FairTokenSetMinDstGas, error) {
	event := new(FairTokenSetMinDstGas)
	if err := _FairToken.contract.UnpackLog(event, "SetMinDstGas", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTokenSetPrecrimeIterator is returned from FilterSetPrecrime and is used to iterate over the raw logs and unpacked data for SetPrecrime events raised by the FairToken contract.
type FairTokenSetPrecrimeIterator struct {
	Event *FairTokenSetPrecrime // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FairTokenSetPrecrimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTokenSetPrecrime)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FairTokenSetPrecrime)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FairTokenSetPrecrimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTokenSetPrecrimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTokenSetPrecrime represents a SetPrecrime event raised by the FairToken contract.
type FairTokenSetPrecrime struct {
	Precrime common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPrecrime is a free log retrieval operation binding the contract event 0x5db758e995a17ec1ad84bdef7e8c3293a0bd6179bcce400dff5d4c3d87db726b.
//
// Solidity: event SetPrecrime(address precrime)
func (_FairToken *FairTokenFilterer) FilterSetPrecrime(opts *bind.FilterOpts) (*FairTokenSetPrecrimeIterator, error) {

	logs, sub, err := _FairToken.contract.FilterLogs(opts, "SetPrecrime")
	if err != nil {
		return nil, err
	}
	return &FairTokenSetPrecrimeIterator{contract: _FairToken.contract, event: "SetPrecrime", logs: logs, sub: sub}, nil
}

// WatchSetPrecrime is a free log subscription operation binding the contract event 0x5db758e995a17ec1ad84bdef7e8c3293a0bd6179bcce400dff5d4c3d87db726b.
//
// Solidity: event SetPrecrime(address precrime)
func (_FairToken *FairTokenFilterer) WatchSetPrecrime(opts *bind.WatchOpts, sink chan<- *FairTokenSetPrecrime) (event.Subscription, error) {

	logs, sub, err := _FairToken.contract.WatchLogs(opts, "SetPrecrime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTokenSetPrecrime)
				if err := _FairToken.contract.UnpackLog(event, "SetPrecrime", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetPrecrime is a log parse operation binding the contract event 0x5db758e995a17ec1ad84bdef7e8c3293a0bd6179bcce400dff5d4c3d87db726b.
//
// Solidity: event SetPrecrime(address precrime)
func (_FairToken *FairTokenFilterer) ParseSetPrecrime(log types.Log) (*FairTokenSetPrecrime, error) {
	event := new(FairTokenSetPrecrime)
	if err := _FairToken.contract.UnpackLog(event, "SetPrecrime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTokenSetTrustedRemoteIterator is returned from FilterSetTrustedRemote and is used to iterate over the raw logs and unpacked data for SetTrustedRemote events raised by the FairToken contract.
type FairTokenSetTrustedRemoteIterator struct {
	Event *FairTokenSetTrustedRemote // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FairTokenSetTrustedRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTokenSetTrustedRemote)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FairTokenSetTrustedRemote)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FairTokenSetTrustedRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTokenSetTrustedRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTokenSetTrustedRemote represents a SetTrustedRemote event raised by the FairToken contract.
type FairTokenSetTrustedRemote struct {
	RemoteChainId uint16
	Path          []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedRemote is a free log retrieval operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _remoteChainId, bytes _path)
func (_FairToken *FairTokenFilterer) FilterSetTrustedRemote(opts *bind.FilterOpts) (*FairTokenSetTrustedRemoteIterator, error) {

	logs, sub, err := _FairToken.contract.FilterLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return &FairTokenSetTrustedRemoteIterator{contract: _FairToken.contract, event: "SetTrustedRemote", logs: logs, sub: sub}, nil
}

// WatchSetTrustedRemote is a free log subscription operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _remoteChainId, bytes _path)
func (_FairToken *FairTokenFilterer) WatchSetTrustedRemote(opts *bind.WatchOpts, sink chan<- *FairTokenSetTrustedRemote) (event.Subscription, error) {

	logs, sub, err := _FairToken.contract.WatchLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTokenSetTrustedRemote)
				if err := _FairToken.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTrustedRemote is a log parse operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _remoteChainId, bytes _path)
func (_FairToken *FairTokenFilterer) ParseSetTrustedRemote(log types.Log) (*FairTokenSetTrustedRemote, error) {
	event := new(FairTokenSetTrustedRemote)
	if err := _FairToken.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTokenSetTrustedRemoteAddressIterator is returned from FilterSetTrustedRemoteAddress and is used to iterate over the raw logs and unpacked data for SetTrustedRemoteAddress events raised by the FairToken contract.
type FairTokenSetTrustedRemoteAddressIterator struct {
	Event *FairTokenSetTrustedRemoteAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FairTokenSetTrustedRemoteAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTokenSetTrustedRemoteAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FairTokenSetTrustedRemoteAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FairTokenSetTrustedRemoteAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTokenSetTrustedRemoteAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTokenSetTrustedRemoteAddress represents a SetTrustedRemoteAddress event raised by the FairToken contract.
type FairTokenSetTrustedRemoteAddress struct {
	RemoteChainId uint16
	RemoteAddress []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedRemoteAddress is a free log retrieval operation binding the contract event 0x8c0400cfe2d1199b1a725c78960bcc2a344d869b80590d0f2bd005db15a572ce.
//
// Solidity: event SetTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress)
func (_FairToken *FairTokenFilterer) FilterSetTrustedRemoteAddress(opts *bind.FilterOpts) (*FairTokenSetTrustedRemoteAddressIterator, error) {

	logs, sub, err := _FairToken.contract.FilterLogs(opts, "SetTrustedRemoteAddress")
	if err != nil {
		return nil, err
	}
	return &FairTokenSetTrustedRemoteAddressIterator{contract: _FairToken.contract, event: "SetTrustedRemoteAddress", logs: logs, sub: sub}, nil
}

// WatchSetTrustedRemoteAddress is a free log subscription operation binding the contract event 0x8c0400cfe2d1199b1a725c78960bcc2a344d869b80590d0f2bd005db15a572ce.
//
// Solidity: event SetTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress)
func (_FairToken *FairTokenFilterer) WatchSetTrustedRemoteAddress(opts *bind.WatchOpts, sink chan<- *FairTokenSetTrustedRemoteAddress) (event.Subscription, error) {

	logs, sub, err := _FairToken.contract.WatchLogs(opts, "SetTrustedRemoteAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTokenSetTrustedRemoteAddress)
				if err := _FairToken.contract.UnpackLog(event, "SetTrustedRemoteAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTrustedRemoteAddress is a log parse operation binding the contract event 0x8c0400cfe2d1199b1a725c78960bcc2a344d869b80590d0f2bd005db15a572ce.
//
// Solidity: event SetTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress)
func (_FairToken *FairTokenFilterer) ParseSetTrustedRemoteAddress(log types.Log) (*FairTokenSetTrustedRemoteAddress, error) {
	event := new(FairTokenSetTrustedRemoteAddress)
	if err := _FairToken.contract.UnpackLog(event, "SetTrustedRemoteAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTokenStakeTokenIterator is returned from FilterStakeToken and is used to iterate over the raw logs and unpacked data for StakeToken events raised by the FairToken contract.
type FairTokenStakeTokenIterator struct {
	Event *FairTokenStakeToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FairTokenStakeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTokenStakeToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FairTokenStakeToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FairTokenStakeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTokenStakeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTokenStakeToken represents a StakeToken event raised by the FairToken contract.
type FairTokenStakeToken struct {
	User   common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeToken is a free log retrieval operation binding the contract event 0x1f42ead3b460988335c082e2eae9c3f3f9e8279142406f3dbd34af4e7c53a841.
//
// Solidity: event StakeToken(address indexed user, uint256 amount, uint256 time)
func (_FairToken *FairTokenFilterer) FilterStakeToken(opts *bind.FilterOpts, user []common.Address) (*FairTokenStakeTokenIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FairToken.contract.FilterLogs(opts, "StakeToken", userRule)
	if err != nil {
		return nil, err
	}
	return &FairTokenStakeTokenIterator{contract: _FairToken.contract, event: "StakeToken", logs: logs, sub: sub}, nil
}

// WatchStakeToken is a free log subscription operation binding the contract event 0x1f42ead3b460988335c082e2eae9c3f3f9e8279142406f3dbd34af4e7c53a841.
//
// Solidity: event StakeToken(address indexed user, uint256 amount, uint256 time)
func (_FairToken *FairTokenFilterer) WatchStakeToken(opts *bind.WatchOpts, sink chan<- *FairTokenStakeToken, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FairToken.contract.WatchLogs(opts, "StakeToken", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTokenStakeToken)
				if err := _FairToken.contract.UnpackLog(event, "StakeToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeToken is a log parse operation binding the contract event 0x1f42ead3b460988335c082e2eae9c3f3f9e8279142406f3dbd34af4e7c53a841.
//
// Solidity: event StakeToken(address indexed user, uint256 amount, uint256 time)
func (_FairToken *FairTokenFilterer) ParseStakeToken(log types.Log) (*FairTokenStakeToken, error) {
	event := new(FairTokenStakeToken)
	if err := _FairToken.contract.UnpackLog(event, "StakeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FairToken contract.
type FairTokenTransferIterator struct {
	Event *FairTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FairTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FairTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FairTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTokenTransfer represents a Transfer event raised by the FairToken contract.
type FairTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FairToken *FairTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FairTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FairToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FairTokenTransferIterator{contract: _FairToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FairToken *FairTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FairTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FairToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTokenTransfer)
				if err := _FairToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FairToken *FairTokenFilterer) ParseTransfer(log types.Log) (*FairTokenTransfer, error) {
	event := new(FairTokenTransfer)
	if err := _FairToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTokenUnStakeTokenIterator is returned from FilterUnStakeToken and is used to iterate over the raw logs and unpacked data for UnStakeToken events raised by the FairToken contract.
type FairTokenUnStakeTokenIterator struct {
	Event *FairTokenUnStakeToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FairTokenUnStakeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTokenUnStakeToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FairTokenUnStakeToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FairTokenUnStakeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTokenUnStakeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTokenUnStakeToken represents a UnStakeToken event raised by the FairToken contract.
type FairTokenUnStakeToken struct {
	User   common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnStakeToken is a free log retrieval operation binding the contract event 0x3cfaa7db2d8b7c8baa09aef6378ff7ca9bc7bcf74944c1d68561b923ba8ab8a9.
//
// Solidity: event UnStakeToken(address indexed user, uint256 amount, uint256 time)
func (_FairToken *FairTokenFilterer) FilterUnStakeToken(opts *bind.FilterOpts, user []common.Address) (*FairTokenUnStakeTokenIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FairToken.contract.FilterLogs(opts, "UnStakeToken", userRule)
	if err != nil {
		return nil, err
	}
	return &FairTokenUnStakeTokenIterator{contract: _FairToken.contract, event: "UnStakeToken", logs: logs, sub: sub}, nil
}

// WatchUnStakeToken is a free log subscription operation binding the contract event 0x3cfaa7db2d8b7c8baa09aef6378ff7ca9bc7bcf74944c1d68561b923ba8ab8a9.
//
// Solidity: event UnStakeToken(address indexed user, uint256 amount, uint256 time)
func (_FairToken *FairTokenFilterer) WatchUnStakeToken(opts *bind.WatchOpts, sink chan<- *FairTokenUnStakeToken, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FairToken.contract.WatchLogs(opts, "UnStakeToken", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTokenUnStakeToken)
				if err := _FairToken.contract.UnpackLog(event, "UnStakeToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnStakeToken is a log parse operation binding the contract event 0x3cfaa7db2d8b7c8baa09aef6378ff7ca9bc7bcf74944c1d68561b923ba8ab8a9.
//
// Solidity: event UnStakeToken(address indexed user, uint256 amount, uint256 time)
func (_FairToken *FairTokenFilterer) ParseUnStakeToken(log types.Log) (*FairTokenUnStakeToken, error) {
	event := new(FairTokenUnStakeToken)
	if err := _FairToken.contract.UnpackLog(event, "UnStakeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTokenWithdrawTokenIterator is returned from FilterWithdrawToken and is used to iterate over the raw logs and unpacked data for WithdrawToken events raised by the FairToken contract.
type FairTokenWithdrawTokenIterator struct {
	Event *FairTokenWithdrawToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FairTokenWithdrawTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTokenWithdrawToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FairTokenWithdrawToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FairTokenWithdrawTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTokenWithdrawTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTokenWithdrawToken represents a WithdrawToken event raised by the FairToken contract.
type FairTokenWithdrawToken struct {
	User   common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawToken is a free log retrieval operation binding the contract event 0x7575c9e41ba6d91ef21e39bd62b5b9df62fc9b0401379fdf9fe1fa372f41e7c1.
//
// Solidity: event WithdrawToken(address indexed user, uint256 amount, uint256 time)
func (_FairToken *FairTokenFilterer) FilterWithdrawToken(opts *bind.FilterOpts, user []common.Address) (*FairTokenWithdrawTokenIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FairToken.contract.FilterLogs(opts, "WithdrawToken", userRule)
	if err != nil {
		return nil, err
	}
	return &FairTokenWithdrawTokenIterator{contract: _FairToken.contract, event: "WithdrawToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawToken is a free log subscription operation binding the contract event 0x7575c9e41ba6d91ef21e39bd62b5b9df62fc9b0401379fdf9fe1fa372f41e7c1.
//
// Solidity: event WithdrawToken(address indexed user, uint256 amount, uint256 time)
func (_FairToken *FairTokenFilterer) WatchWithdrawToken(opts *bind.WatchOpts, sink chan<- *FairTokenWithdrawToken, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FairToken.contract.WatchLogs(opts, "WithdrawToken", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTokenWithdrawToken)
				if err := _FairToken.contract.UnpackLog(event, "WithdrawToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawToken is a log parse operation binding the contract event 0x7575c9e41ba6d91ef21e39bd62b5b9df62fc9b0401379fdf9fe1fa372f41e7c1.
//
// Solidity: event WithdrawToken(address indexed user, uint256 amount, uint256 time)
func (_FairToken *FairTokenFilterer) ParseWithdrawToken(log types.Log) (*FairTokenWithdrawToken, error) {
	event := new(FairTokenWithdrawToken)
	if err := _FairToken.contract.UnpackLog(event, "WithdrawToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
