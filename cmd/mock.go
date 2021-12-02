package cmd

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"
)

type ProposeUtilsMock struct{}

type UtilsCmdMock struct{}

var WaitForAppropriateStateMock func(*ethclient.Client, string, string, UtilsStruct, ...int) (uint32, error)

var WaitIfCommitStateMock func(*ethclient.Client, string, string, UtilsStruct) (uint32, error)

var AssignAmountInWeiMock func(flagSet *pflag.FlagSet, utilsStruct UtilsStruct) (*big.Int, error)

var getBiggestInfluenceAndIdMock func(*ethclient.Client, string, uint32, UtilsStruct) (*big.Int, uint32, error)

var getIterationMock func(*ethclient.Client, string, types.ElectedProposer, UtilsStruct) int

var isElectedProposerMock func(*ethclient.Client, string, types.ElectedProposer, UtilsStruct) bool

var pseudoRandomNumberGeneratorMock func([]byte, uint32, []byte) *big.Int

var MakeBlockMock func(*ethclient.Client, string, bool, UtilsStruct) ([]uint32, error)

var getSortedVotesMock func(*ethclient.Client, string, uint8, uint32, UtilsStruct) ([]*big.Int, error)

var influencedMedianMock func([]*big.Int, *big.Int) *big.Int

var CheckCurrentStatusMock func(*ethclient.Client, string, uint8, UtilsStruct) (bool, error)

var DisputeMock func(*ethclient.Client, types.Configurations, types.Account, uint32, uint8, int, UtilsStruct) error

var GetEpochAndStateMock func(*ethclient.Client, string, UtilsStruct) (uint32, int64, error)

var GiveSortedMock func(*ethclient.Client, *bindings.BlockManager, *bind.TransactOpts, uint32, uint8, []uint32)

var UnstakeMock func(types.Configurations, *ethclient.Client, types.UnstakeInput, UtilsStruct) (types.TransactionOptions, error)

var AutoWithdrawMock func(types.TransactionOptions, uint32, UtilsStruct) error

var withdrawFundsMock func(*ethclient.Client, types.Account, types.Configurations, uint32, UtilsStruct) (common.Hash, error)

var claimBountyMock func(types.Configurations, *ethclient.Client, types.RedeemBountyInput, UtilsStruct) (common.Hash, error)

var SetCommissionMock func(*ethclient.Client, uint32, *bind.TransactOpts, uint8, UtilsStruct) error

var DecreaseCommissionMock func(*ethclient.Client, uint32, *bind.TransactOpts, uint8, UtilsStruct) error

var DecreaseCommissionPromptMock func() bool

var WithdrawMock func(*ethclient.Client, *bind.TransactOpts, uint32, uint32, UtilsStruct) (common.Hash, error)

var GetConfigDataMock func() (types.Configurations, error)

var getBufferPercentMock func() (int32, error)

func (proposeUtilsMock ProposeUtilsMock) getBiggestInfluenceAndId(client *ethclient.Client, address string, epoch uint32, utilsStruct UtilsStruct) (*big.Int, uint32, error) {
	return getBiggestInfluenceAndIdMock(client, address, epoch, utilsStruct)
}

func (proposeUtilsMock ProposeUtilsMock) getIteration(client *ethclient.Client, address string, proposer types.ElectedProposer, utilsStruct UtilsStruct) int {
	return getIterationMock(client, address, proposer, utilsStruct)
}

func (proposeUtilsMock ProposeUtilsMock) isElectedProposer(client *ethclient.Client, address string, proposer types.ElectedProposer, utilsStruct UtilsStruct) bool {
	return isElectedProposerMock(client, address, proposer, utilsStruct)
}

func (proposeUtilsMock ProposeUtilsMock) pseudoRandomNumberGenerator(seed []byte, max uint32, blockHashes []byte) *big.Int {
	return pseudoRandomNumberGeneratorMock(seed, max, blockHashes)
}

func (proposeUtilsMock ProposeUtilsMock) MakeBlock(client *ethclient.Client, address string, rogueMode bool, utilsStruct UtilsStruct) ([]uint32, error) {
	return MakeBlockMock(client, address, rogueMode, utilsStruct)
}

func (proposeUtilsMock ProposeUtilsMock) getSortedVotes(client *ethclient.Client, address string, assetId uint8, epoch uint32, utilsStruct UtilsStruct) ([]*big.Int, error) {
	return getSortedVotesMock(client, address, assetId, epoch, utilsStruct)
}

func (proposeUtilsMock ProposeUtilsMock) influencedMedian(sortedVotes []*big.Int, totalInfluenceRevealed *big.Int) *big.Int {
	return influencedMedianMock(sortedVotes, totalInfluenceRevealed)
}

func (utilsCmdMock UtilsCmdMock) SetCommission(client *ethclient.Client, stakerId uint32, opts *bind.TransactOpts, commission uint8, utilsStruct UtilsStruct) error {
	return SetCommissionMock(client, stakerId, opts, commission, utilsStruct)
}

func (utilsCmdMock UtilsCmdMock) DecreaseCommission(client *ethclient.Client, stakerId uint32, opts *bind.TransactOpts, commission uint8, utilsStruct UtilsStruct) error {
	return DecreaseCommissionMock(client, stakerId, opts, commission, utilsStruct)
}

func (utilsCmdMock UtilsCmdMock) DecreaseCommissionPrompt() bool {
	return DecreaseCommissionPromptMock()
}

func (utilsCmdMock UtilsCmdMock) Withdraw(client *ethclient.Client, txnOpts *bind.TransactOpts, epoch uint32, stakerId uint32, utilsStruct UtilsStruct) (common.Hash, error) {
	return WithdrawMock(client, txnOpts, epoch, stakerId, utilsStruct)
}

func (utilsCmdMock UtilsCmdMock) CheckCurrentStatus(client *ethclient.Client, address string, assetId uint8, utilsStruct UtilsStruct) (bool, error) {
	return CheckCurrentStatusMock(client, address, assetId, utilsStruct)
}

func (utilsCmdMock UtilsCmdMock) Dispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, blockId uint8, assetId int, utilsStruct UtilsStruct) error {
	return DisputeMock(client, config, account, epoch, blockId, assetId, utilsStruct)
}

func (utilsCmdMock UtilsCmdMock) GiveSorted(client *ethclient.Client, blockManager *bindings.BlockManager, txnOpts *bind.TransactOpts, epoch uint32, assetId uint8, sortedStakers []uint32) {
	GiveSortedMock(client, blockManager, txnOpts, epoch, assetId, sortedStakers)
}

func (utilsCmdMock UtilsCmdMock) GetEpochAndState(client *ethclient.Client, accountAddress string, utilsStruct UtilsStruct) (uint32, int64, error) {
	return GetEpochAndStateMock(client, accountAddress, utilsStruct)
}

func (utilsCmdMock UtilsCmdMock) WaitForAppropriateState(client *ethclient.Client, accountAddress string, action string, utilsStruct UtilsStruct, states ...int) (uint32, error) {
	return WaitForAppropriateStateMock(client, accountAddress, action, utilsStruct, states...)
}

func (utilsCmdMock UtilsCmdMock) WaitIfCommitState(client *ethclient.Client, accountAddress string, action string, utilsStruct UtilsStruct) (uint32, error) {
	return WaitIfCommitStateMock(client, accountAddress, action, utilsStruct)
}

func (utilsCmdMock UtilsCmdMock) AssignAmountInWei(flagSet *pflag.FlagSet, utilsStruct UtilsStruct) (*big.Int, error) {
	return AssignAmountInWeiMock(flagSet, utilsStruct)
}

func (utilsCmdMock UtilsCmdMock) Unstake(config types.Configurations, client *ethclient.Client, inputUnstake types.UnstakeInput, utilsStruct UtilsStruct) (types.TransactionOptions, error) {
	return UnstakeMock(config, client, inputUnstake, utilsStruct)
}

func (utilsCmdMock UtilsCmdMock) AutoWithdraw(txnArgs types.TransactionOptions, stakerId uint32, utilsStruct UtilsStruct) error {
	return AutoWithdrawMock(txnArgs, stakerId, utilsStruct)
}

func (utilsCmdMock UtilsCmdMock) withdrawFunds(client *ethclient.Client, account types.Account, configurations types.Configurations, stakerId uint32, utilsStruct UtilsStruct) (common.Hash, error) {
	return withdrawFundsMock(client, account, configurations, stakerId, utilsStruct)
}

func (utilsCmdMock UtilsCmdMock) claimBounty(config types.Configurations, client *ethclient.Client, redeemBountyInput types.RedeemBountyInput, utilsStruct UtilsStruct) (common.Hash, error) {
	return claimBountyMock(config, client, redeemBountyInput, utilsStruct)
}

func (utilsCmdMock UtilsCmdMock) GetConfigData() (types.Configurations, error) {
	return GetConfigDataMock()
}

func (utilsCmdMock UtilsCmdMock) getBufferPercent() (int32, error) {
	return getBufferPercentMock()
}
