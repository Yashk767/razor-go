package cmd

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"
)

type utilsCmdInterface interface {
	SetCommission(*ethclient.Client, uint32, *bind.TransactOpts, uint8, UtilsStruct) error
	DecreaseCommission(*ethclient.Client, uint32, *bind.TransactOpts, uint8, UtilsStruct) error
	DecreaseCommissionPrompt() bool
	Withdraw(*ethclient.Client, *bind.TransactOpts, uint32, uint32, UtilsStruct) (common.Hash, error)
	CheckCurrentStatus(*ethclient.Client, string, uint8, UtilsStruct) (bool, error)
	Dispute(*ethclient.Client, types.Configurations, types.Account, uint32, uint8, int, UtilsStruct) error
	GiveSorted(*ethclient.Client, *bindings.BlockManager, *bind.TransactOpts, uint32, uint8, []uint32)
	GetEpochAndState(*ethclient.Client, string, UtilsStruct) (uint32, int64, error)
	WaitForAppropriateState(*ethclient.Client, string, string, UtilsStruct, ...int) (uint32, error)
	WaitIfCommitState(*ethclient.Client, string, string, UtilsStruct) (uint32, error)
	AssignAmountInWei(*pflag.FlagSet, UtilsStruct) (*big.Int, error)
	Unstake(types.Configurations, *ethclient.Client, types.UnstakeInput, UtilsStruct) (types.TransactionOptions, error)
	AutoWithdraw(types.TransactionOptions, uint32, UtilsStruct) error
	withdrawFunds(*ethclient.Client, types.Account, types.Configurations, uint32, UtilsStruct) (common.Hash, error)
	Create(string, UtilsStruct) (accounts.Account, error)
	claimBounty(types.Configurations, *ethclient.Client, types.RedeemBountyInput, UtilsStruct) (common.Hash, error)
	GetConfigData() (types.Configurations, error)
	getBufferPercent() (int32, error)
}

type proposeUtilsInterface interface {
	getBiggestInfluenceAndId(*ethclient.Client, string, uint32, UtilsStruct) (*big.Int, uint32, error)
	getIteration(*ethclient.Client, string, types.ElectedProposer, UtilsStruct) int
	isElectedProposer(*ethclient.Client, string, types.ElectedProposer, UtilsStruct) bool
	pseudoRandomNumberGenerator([]byte, uint32, []byte) *big.Int
	MakeBlock(*ethclient.Client, string, bool, UtilsStruct) ([]uint32, error)
	getSortedVotes(*ethclient.Client, string, uint8, uint32, UtilsStruct) ([]*big.Int, error)
	influencedMedian([]*big.Int, *big.Int) *big.Int
}
