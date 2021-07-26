package cmd

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"razor/core/types"
	"razor/utils"
	"testing"
)

func getTransactionArgsDelegator1() types.TransactionOptions {
	password := "12345"
	delegatorAddress := "0x5f0e92b8EC08f1c4C589c9dE5d34b4a54e58652A"
	provider := "http://127.0.0.1:8545/"
	client, err := ethclient.Dial(provider)
	if err != nil {
		log.Fatal("Error in connecting...\n", err)
	}
	log.Info("Connected to: ", provider)
	if err != nil {
		log.Fatal(err)
	}
	balance, err := utils.FetchBalance(client, delegatorAddress)
	amount := "1000"
	amountInWei := utils.GetAmountWithChecks(amount, balance)

	txnArgs := types.TransactionOptions{
		Client:         client,
		AccountAddress: delegatorAddress,
		Amount:         amountInWei,
		Password:       password,
		ChainId:        big.NewInt(1337),
		GasMultiplier:  10,
	}
	return txnArgs
}

func Test_delegate(t *testing.T) {
	txnArgs := getTransactionArgs()
	approve(txnArgs)
	stakeCoins(txnArgs)
	opts := utils.GetOptions(false, txnArgs.AccountAddress, "")
	txnOpts := utils.GetTxnOpts(txnArgs)
	stakeManager := utils.GetStakeManager(txnArgs.Client)
	stakerId, _ := stakeManager.GetStakerId(&opts, common.HexToAddress(txnArgs.AccountAddress))

	t.Run("Staker is able to set delegationAcceptance to true ", func(t *testing.T) {
		SetAcceptance := true
		hash, err := stakeManager.SetDelegationAcceptance(txnOpts, SetAcceptance)
		fmt.Println(hash)
		if err != nil {
			t.Errorf("Error in executing SetDelegationAccptance from contracts")
		}
		staker, _ := stakeManager.GetStaker(&opts, stakerId)
		fmt.Println(staker.AcceptDelegation)
		if hash == nil {
			t.Errorf("Staker is not able to set delegationAcceptance")
		}
	})

	t.Run("Test1: Delegator is able to delegate", func(t *testing.T) {
		staker, _ := stakeManager.GetStaker(&opts, stakerId)
		fmt.Println(staker.AcceptDelegation)
		wantStatus := 1
		txnArgsDelegator := getTransactionArgsDelegator1()
		approve(txnArgsDelegator)
		delegationStatus := delegate(txnArgsDelegator, stakerId)
		fmt.Println(delegationStatus)
		if delegationStatus != wantStatus {
			t.Errorf("Delegate is not able to delegate")
		}
	})

	t.Run("Test2: Delegator is able to delegate again to the same staker", func(t *testing.T) {
		wantStatus := 1
		txnArgsDelegator := getTransactionArgsDelegator1()
		approve(txnArgsDelegator)
		delegationStatus := delegate(txnArgsDelegator, stakerId)
		if delegationStatus != wantStatus {
			t.Errorf("Delegate is not able to delegate")
		}
	})

	t.Run("Test3: Delegated RZR's are updated correctly after delegation", func(t *testing.T) {
		txnArgsDelegator := getTransactionArgsDelegator1()
		staker, _ := stakeManager.GetStaker(&opts, stakerId)
		previousStake := staker.Stake
		approve(txnArgsDelegator)
		delegate(txnArgsDelegator, stakerId)

		staker1, _ := stakeManager.GetStaker(&opts, stakerId)
		newStake := staker1.Stake
		expectedStake := big.NewInt(0)
		expectedStake.Add(previousStake, txnArgs.Amount)
		if got := newStake; got.Cmp(expectedStake) != 0 {
			t.Errorf("Error in delegating the stake amount to staker")
		}
	})
}
	//opts := utils.GetOptions(false, txnArgs.AccountAddress, "")
	//txnOpts := utils.GetTxnOpts(txnArgs)
	//stakeManager := utils.GetStakeManager(txnArgs.Client)
	//stakeManager.SetDelegationAcceptance(txnOpts, true)
	//stakerId,_ := stakeManager.GetStakerId(&opts, common.HexToAddress(txnArgs.AccountAddress))
	//staker,_ := stakeManager.GetStaker(&opts, stakerId)
	//fmt.Println(staker.AcceptDelegation)

	//type args struct {
	//	txnArgs   types.TransactionOptions
	//	_stakerId *big.Int
	//}
	//t.Run("Test1: Delegator is not able to delegate to the staker who hasn't staked yet", func(t *testing.T) {
	//	wantErr := true
	//	expectedFatal:= true
	//	txnArgsDelegator := getTransactionArgsDelegator1()
	//
	//	defer func() { log.StandardLogger().ExitFunc = nil }()
	//	var fatal bool
	//	log.StandardLogger().ExitFunc = func(int){ fatal = true }
	//
	//	delegationStatus := delegate(txnArgsDelegator, stakerId)
	//	if (delegationStatus != 1) != wantErr {
	//		t.Errorf("Chosen Staker has staked as stakerId : %s", stakerId)
	//	}
	//	if expectedFatal {
	//		assert.Equal(t, expectedFatal, fatal)
	//	}
	//})


	//tests := []struct {
	//	name string
	//	args args
	//	wantErr bool
	//}{
	//	{
	//		name: "Test1: Delegator is not able to delegate to the staker who hasn't staked yet",
	//		args: args{
	//			txnArgs: getTransactionArgsDelegator1(),
	//			_stakerId: stakerId,
	//		},
	//		wantErr: false,
	//	},
	//	{
	//		name: "Test2: Delegator is able to delegate twice to the same staker",
	//		args: args{
	//			txnArgs: getTransactionArgsDelegator1(),
	//			_stakerId: stakerId,
	//		},
	//		wantErr: false,
	//	},
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//	})

