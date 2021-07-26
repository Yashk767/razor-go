package cmd

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/pkg/bindings"
	"razor/utils"
	"testing"
)

func TestSetCommission(t *testing.T) {
	txnArgs := getTransactionArgs()
	approve(txnArgs)
	stakeCoins(txnArgs)
	stakeManager := utils.GetStakeManager(txnArgs.Client)
	txnOpts := utils.GetTxnOpts(txnArgs)
	SetAcceptance := true
	hash, err := stakeManager.SetDelegationAcceptance(txnOpts, SetAcceptance)
	fmt.Println(hash)
	if err != nil {
		fmt.Println("Error in executing SetDelegationAcceptance from contracts")
	}
	if hash == nil {
		fmt.Println("Staker is not able to set delegationAcceptance")
	}

	t.Run("Staker is able to initialize the commission", func(t *testing.T) {
		//stakerID,_ := utils.GetStakerId(txnArgs.Client, txnArgs.AccountAddress)
		//staker,_ := utils.GetStaker(txnArgs.Client, txnArgs.AccountAddress, stakerID)
		//fmt.Println(staker.Commission)
		txnOpts := utils.GetTxnOpts(txnArgs)
		NewCommission := big.NewInt(10)
		hash, err := stakeManager.SetCommission(txnOpts, NewCommission)
		if err != nil {
			t.Errorf("Error in executing SetCommission from contracts : %s", err)
		}
		if hash == nil {
			t.Errorf("Staker is not able to set commission")
		}
	})
}

func TestDecreaseCommission(t *testing.T) {
	txnArgs := getTransactionArgs()
	_txnOpts := utils.GetTxnOpts(txnArgs)
	_stakeManager := utils.GetStakeManager(txnArgs.Client)
	_stakerId,_ := utils.GetStakerId(txnArgs.Client, txnArgs.AccountAddress)
	staker,_ := utils.GetStaker(txnArgs.Client, txnArgs.AccountAddress, _stakerId )
	oldCommission := staker.Commission
	NewCommission1 := big.NewInt(0)
	NewCommission2 := big.NewInt(0)

	type args struct {
		client       *ethclient.Client
		stakeManager *bindings.StakeManager
		stakerId     *big.Int
		txnOpts      *bind.TransactOpts
		commission   *big.Int
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "Test1: Staker is able to decrease the commission",
			args: args{
				client:       txnArgs.Client,
				stakeManager: _stakeManager,
				stakerId:     _stakerId,
				txnOpts:      _txnOpts,
				commission:   NewCommission1.Sub(oldCommission, big.NewInt(1)),
			},
			wantErr: false,
		},
		{
			name: "Test2: Staker cannot increase its commission",
			args: args{
				client:       txnArgs.Client,
				stakeManager: _stakeManager,
				stakerId:     _stakerId,
				txnOpts:      _txnOpts,
				commission:   NewCommission2.Add(oldCommission, big.NewInt(1)),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decreaseCommissionStatus := DecreaseCommission(tt.args.client, tt.args.stakeManager, tt.args.stakerId, tt.args.txnOpts, tt.args.commission)
			if (tt.wantErr) && (decreaseCommissionStatus == 1) {
				t.Errorf("Staker is not able to decrease its commission")
			}
		})
	}

}

