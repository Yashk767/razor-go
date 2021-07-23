package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"razor/pkg/bindings"

	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"razor/core/types"
	"razor/utils"
	"testing"
)

func getTransactionArgs() types.TransactionOptions{
	password := "12345"
	address := "0x5f0e92b8EC08f1c4C589c9dE5d34b4a54e58652A"
	provider := "http://127.0.0.1:8545/"
	client, err := ethclient.Dial(provider)
	if err != nil {
		log.Fatal("Error in connecting...\n", err)
	}
	log.Info("Connected to: ", provider)
	if err != nil {
		log.Fatal(err)
	}
	balance, err := utils.FetchBalance(client, address)
	amount := "1000"
	amountInWei := utils.GetAmountWithChecks(amount, balance)

	txnArgs := types.TransactionOptions{
		Client:         client,
		AccountAddress: address,
		Amount:         amountInWei,
		Password:       password,
		ChainId:        big.NewInt(1337),
		GasMultiplier:  10,
	}
	return txnArgs
}

func Test_approve(t *testing.T) {
	txnArgs := getTransactionArgs()
	approvedStatus := approve(txnArgs)
	if got:= approvedStatus; got != 1 {
		t.Errorf("Transaction is not Approved")
	}

	schellingCoinAddress := "0x1Ca31A8832c532DD22AcFB0a3cf7b9F87F4574e6"
	stakeManagerAddress := "0xBfBA8b5F2CaD3Fdd82893BECDBfb0193C878CED6"
	opts := utils.GetOptions(false, txnArgs.AccountAddress, "")

	tokenManager, _ := bindings.NewSchellingCoin(common.HexToAddress(schellingCoinAddress), txnArgs.Client)
	allowance, _ := tokenManager.Allowance(&opts, common.HexToAddress(txnArgs.AccountAddress), common.HexToAddress(stakeManagerAddress))

	if got := allowance; got.Cmp(txnArgs.Amount) != 1 && got.Cmp(txnArgs.Amount) != 0 {
		t.Errorf("Amount of SC's approved are less than the required amount of SC's to be sent")
	}
}

func Test_stakeCoins(t *testing.T) {
	stakeManagerAddress := "0xBfBA8b5F2CaD3Fdd82893BECDBfb0193C878CED6"
	txnArgs := getTransactionArgs()
	opts := utils.GetOptions(false, txnArgs.AccountAddress, "")
	stakeManager, _ := bindings.NewStakeManager(common.HexToAddress(stakeManagerAddress), txnArgs.Client)
	stakerId,_ := stakeManager.GetStakerId(&opts, common.HexToAddress(txnArgs.AccountAddress))

	if stakerId.Cmp(big.NewInt(0)) == 0 {
		approve(txnArgs)
		stakeCoins(txnArgs)

		newStakerID,_ := stakeManager.GetStakerId(&opts, common.HexToAddress(txnArgs.AccountAddress))
		staker,_ := stakeManager.GetStaker(&opts, newStakerID)
		expectedStake := big.NewInt(0)
		expectedStake.Add(expectedStake, txnArgs.Amount)
		newStake := staker.Stake

		if got := newStake; got.Cmp(expectedStake) != 0 {
			t.Errorf("Error in Staking first time")
		}
	}
	if stakerId.Cmp(big.NewInt(0)) != 0 {
		staker,_ := stakeManager.GetStaker(&opts, stakerId)
		previousStake := staker.Stake

		approve(txnArgs)
		stakeCoins(txnArgs)
		staker1,_ := stakeManager.GetStaker(&opts, stakerId)
		newStake := staker1.Stake

		expectedStake := big.NewInt(0)
		expectedStake.Add(previousStake, txnArgs.Amount)

		if got := newStake; got.Cmp(expectedStake) != 0 {
			t.Errorf("Error in Staking the stake amount")
		}
	}

}
