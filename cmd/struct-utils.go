package cmd

import (
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/razorInterface"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
)

type ProposeUtils struct{}
type UtilsCmd struct{}
type UtilsStruct struct {
	razorUtils        razorInterface.UtilsInterface
	cmdUtils          utilsCmdInterface
	proposeUtils      proposeUtilsInterface
	blockManagerUtils razorInterface.BlockManagerInterface
	stakeManagerUtils razorInterface.StakeManagerInterface
	transactionUtils  razorInterface.TransactionInterface
	assetManagerUtils razorInterface.AssetManagerInterface
	voteManagerUtils  razorInterface.VoteManagerInterface
	tokenManagerUtils razorInterface.TokenManagerInterface
	keystoreUtils     razorInterface.KeystoreInterface
	accountUtils      razorInterface.AccountInterface
	flagSetUtils      razorInterface.FlagSetInterface
	cryptoUtils       razorInterface.CryptoInterface
}

func (proposeUtils ProposeUtils) getBiggestInfluenceAndId(client *ethclient.Client, address string, epoch uint32, utilsStruct UtilsStruct) (*big.Int, uint32, error) {
	return getBiggestInfluenceAndId(client, address, epoch, utilsStruct)
}

func (proposeUtils ProposeUtils) getIteration(client *ethclient.Client, address string, proposer types.ElectedProposer, utilsStruct UtilsStruct) int {
	return getIteration(client, address, proposer, utilsStruct)
}

func (proposeUtils ProposeUtils) isElectedProposer(client *ethclient.Client, address string, proposer types.ElectedProposer, utilsStruct UtilsStruct) bool {
	return isElectedProposer(client, address, proposer, utilsStruct)
}

func (proposeUtils ProposeUtils) pseudoRandomNumberGenerator(seed []byte, max uint32, blockHashes []byte) *big.Int {
	return pseudoRandomNumberGenerator(seed, max, blockHashes)
}

func (proposeUtils ProposeUtils) MakeBlock(client *ethclient.Client, address string, rogueMode bool, utilsStruct UtilsStruct) ([]uint32, error) {
	return MakeBlock(client, address, rogueMode, utilsStruct)
}

func (proposeUtils ProposeUtils) getSortedVotes(client *ethclient.Client, address string, assetId uint8, epoch uint32, utilsStruct UtilsStruct) ([]*big.Int, error) {
	return getSortedVotes(client, address, assetId, epoch, utilsStruct)
}

func (proposeUtils ProposeUtils) influencedMedian(sortedVotes []*big.Int, totalInfluenceRevealed *big.Int) *big.Int {
	return influencedMedian(sortedVotes, totalInfluenceRevealed)
}

func (cmdUtils UtilsCmd) SetCommission(client *ethclient.Client, stakerId uint32, txnOpts *bind.TransactOpts, commission uint8, utilsStruct UtilsStruct) error {
	return SetCommission(client, stakerId, txnOpts, commission, utilsStruct)
}

func (cmdUtils UtilsCmd) DecreaseCommission(client *ethclient.Client, stakerId uint32, txnOpts *bind.TransactOpts, commission uint8, utilsStruct UtilsStruct) error {
	return DecreaseCommission(client, stakerId, txnOpts, commission, utilsStruct)
}

func (cmdUtils UtilsCmd) DecreaseCommissionPrompt() bool {
	return DecreaseCommissionPrompt()
}

func (cmdUtils UtilsCmd) Withdraw(client *ethclient.Client, txnOpts *bind.TransactOpts, epoch uint32, stakerId uint32, utilsStruct UtilsStruct) (common.Hash, error) {
	return withdraw(client, txnOpts, epoch, stakerId, utilsStruct)
}

func (cmdUtils UtilsCmd) CheckCurrentStatus(client *ethclient.Client, address string, assetId uint8, utilsStruct UtilsStruct) (bool, error) {
	return CheckCurrentStatus(client, address, assetId, utilsStruct)
}

func (cmdUtils UtilsCmd) Dispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, blockId uint8, assetId int, utilsStruct UtilsStruct) error {
	return Dispute(client, config, account, epoch, blockId, assetId, utilsStruct)
}

func (cmdUtils UtilsCmd) GiveSorted(client *ethclient.Client, blockManager *bindings.BlockManager, txnOpts *bind.TransactOpts, epoch uint32, assetId uint8, sortedStakers []uint32) {
	GiveSorted(client, blockManager, txnOpts, epoch, assetId, sortedStakers)
}

func (cmdUtils UtilsCmd) GetEpochAndState(client *ethclient.Client, accountAddress string, utilsStruct UtilsStruct) (uint32, int64, error) {
	return GetEpochAndState(client, accountAddress, utilsStruct)
}

func (cmdUtils UtilsCmd) WaitForAppropriateState(client *ethclient.Client, accountAddress string, action string, utilsStruct UtilsStruct, states ...int) (uint32, error) {
	return WaitForAppropriateState(client, accountAddress, action, utilsStruct, states...)
}

func (cmdUtils UtilsCmd) WaitIfCommitState(client *ethclient.Client, accountAddress string, action string, utilsStruct UtilsStruct) (uint32, error) {
	return WaitIfCommitState(client, accountAddress, action, utilsStruct)
}

func (cmdUtils UtilsCmd) AssignAmountInWei(flagSet *pflag.FlagSet, utilsStruct UtilsStruct) (*big.Int, error) {
	return AssignAmountInWei(flagSet, utilsStruct)
}

func (cmdUtils UtilsCmd) Unstake(config types.Configurations, client *ethclient.Client, unstakeInput types.UnstakeInput, utilsStruct UtilsStruct) (types.TransactionOptions, error) {
	return Unstake(config, client, unstakeInput, utilsStruct)
}

func (cmdUtils UtilsCmd) AutoWithdraw(txnArgs types.TransactionOptions, stakerId uint32, utilsStruct UtilsStruct) error {
	return AutoWithdraw(txnArgs, stakerId, utilsStruct)
}

func (cmdUtils UtilsCmd) withdrawFunds(client *ethclient.Client, account types.Account, configurations types.Configurations, stakerId uint32, utilsStruct UtilsStruct) (common.Hash, error) {
	return withdrawFunds(client, account, configurations, stakerId, utilsStruct)
}

func (cmdUtils UtilsCmd) claimBounty(config types.Configurations, client *ethclient.Client, redeemBountyInput types.RedeemBountyInput, utilsStruct UtilsStruct) (common.Hash, error) {
	return claimBounty(config, client, redeemBountyInput, utilsStruct)
}

func (cmdUtils UtilsCmd) GetConfigData() (types.Configurations, error) {
	return GetConfigData()
}

func (cmdUtils UtilsCmd) getBufferPercent() (int32, error) {
	return getBufferPercent()
}
