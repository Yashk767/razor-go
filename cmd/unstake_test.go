package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"razor/utils"
	"testing"
)

func Test_unstake(t *testing.T) {
	txnArgs := getTransactionArgs()
	opts := utils.GetOptions(false, txnArgs.AccountAddress, "")
	stakeManager := utils.GetStakeManager(txnArgs.Client)

	t.Run("Test1: Staker is able to unstake", func(t *testing.T) {
		approve(txnArgs)
		stakeCoins(txnArgs)
		stakerId,_ := stakeManager.GetStakerId(&opts, common.HexToAddress(txnArgs.AccountAddress))
		staker,_ := stakeManager.GetStaker(&opts, stakerId)

		unstake(txnArgs, stakerId)
		newLock,_ := stakeManager.Locks(&opts,common.HexToAddress(txnArgs.AccountAddress), staker.TokenAddress)
		newLockedAmount := newLock.Amount

		if got := newLockedAmount; got.Cmp(txnArgs.Amount) !=0 {
			t.Error("Locked amount is not equal to requested lock amount")
		}
	})

	t.Run("Test2: Staker cannot unstake when there is existing lock", func(t *testing.T) {
		wantStatus := 0

		//defer func() { log.StandardLogger().ExitFunc = nil }()
		//var fatal bool
		//log.StandardLogger().ExitFunc = func(int){ fatal = true }

		stakerId,_ := stakeManager.GetStakerId(&opts, common.HexToAddress(txnArgs.AccountAddress))
		//fatal = false
		unstakeStatus := unstake(txnArgs, stakerId)
		if wantStatus != unstakeStatus {
			t.Error("There is no existing lock.")

		}
	})

}
