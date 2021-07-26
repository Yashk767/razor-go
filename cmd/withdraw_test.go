package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"razor/utils"
	"testing"
)

func Test_withdraw(t *testing.T) {
	t.Run("Staker is able to withdraw", func(t *testing.T) {
		txnArgs := getTransactionArgs()
		opts:= utils.GetOptions(false, txnArgs.AccountAddress, "")
		approve(txnArgs)
		stakeCoins(txnArgs)

		stakeManager := utils.GetStakeManager(txnArgs.Client)
		stakerId,_ := stakeManager.GetStakerId(&opts, common.HexToAddress(txnArgs.AccountAddress))
		unstake(txnArgs, stakerId)

		lock,_ := utils.GetLock(txnArgs.Client, txnArgs.AccountAddress, stakerId)
		withdrawReleasePeriod, _ := utils.GetWithdrawReleasePeriod(txnArgs.Client, txnArgs.AccountAddress)
		withdrawBefore := big.NewInt(0).Add(lock.WithdrawAfter, withdrawReleasePeriod)
		epoch, _ := WaitForCommitState(txnArgs.Client, txnArgs.AccountAddress, "withdraw")
		withdrawStatus := 0
		wantStatus := 1
		for i := epoch; i.Cmp(withdrawBefore) < 0; {
			if epoch.Cmp(lock.WithdrawAfter) >= 0 && epoch.Cmp(withdrawBefore) <= 0 {
				withdrawStatus = withdraw(txnArgs.Client, txnArgs, epoch, stakerId)
				break
			}
		}
		if withdrawStatus != wantStatus {
			t.Errorf("Error in withdraw")
		}
	})

		t.Run("Delegator is able to withdraw", func(t *testing.T) {
			txnArgs := getTransactionArgs()
			txnDelegatorArgs := getTransactionArgsDelegator1()
			opts:= utils.GetOptions(false, txnArgs.AccountAddress, "")
			approve(txnArgs)
			stakeCoins(txnArgs)

			stakeManager := utils.GetStakeManager(txnArgs.Client)
			stakerId,_ := stakeManager.GetStakerId(&opts, common.HexToAddress(txnArgs.AccountAddress))

			delegate(txnDelegatorArgs, stakerId)
			unstake(txnDelegatorArgs, stakerId)

			lock,_ := utils.GetLock(txnDelegatorArgs.Client, txnDelegatorArgs.AccountAddress, stakerId)
			withdrawReleasePeriod, _ := utils.GetWithdrawReleasePeriod(txnDelegatorArgs.Client, txnDelegatorArgs.AccountAddress)
			withdrawBefore := big.NewInt(0).Add(lock.WithdrawAfter, withdrawReleasePeriod)
			epoch, _ := WaitForCommitState(txnDelegatorArgs.Client, txnDelegatorArgs.AccountAddress, "withdraw")
			withdrawStatus := 0
			wantStatus := 1
			for i := epoch; i.Cmp(withdrawBefore) < 0; {
				if epoch.Cmp(lock.WithdrawAfter) >= 0 && epoch.Cmp(withdrawBefore) <= 0 {
					withdrawStatus = withdraw(txnDelegatorArgs.Client, txnDelegatorArgs, epoch, stakerId)
					break
				}
			}
			if withdrawStatus != wantStatus {
				t.Errorf("Error in withdraw")
			}
		})

}

func Test_checkForCommitStateAndWithdraw(t *testing.T) {

}

