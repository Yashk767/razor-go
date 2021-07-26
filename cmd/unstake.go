package cmd

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var unstakeCmd = &cobra.Command{
	Use:   "unstake",
	Short: "Unstake your razors",
	Long: `unstake allows user to unstake their sRzrs in the razor network
	For ex:
	unstake --address <address> --amount <amount_of_sRazors>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in getting config: ", err)

		address, _ := cmd.Flags().GetString("address")
		amount, _ := cmd.Flags().GetString("amount")
		stakerId, _ := cmd.Flags().GetString("stakerId")
		autoWithdraw, _ := cmd.Flags().GetBool("autoWithdraw")
		password := utils.PasswordPrompt()

		client := utils.ConnectToClient(config.Provider)

		balance, err := utils.FetchBalance(client, address)
		if err != nil {
			log.Fatal("Error in fetching balance: ", err)
		}

		amountInWei := utils.GetAmountWithChecks(amount, balance)

		utils.CheckError("Error in fetching epoch: ", err)
		_stakerId, ok := new(big.Int).SetString(stakerId, 10)
		if !ok {
			log.Fatal("Set string error in converting staker id")
		}

		lock, err := utils.GetLock(client, address, _stakerId)
		utils.CheckError("Error in getting lock: ", err)
		if lock.Amount.Cmp(big.NewInt(0)) != 0 {
			log.Fatal("Existing lock")
		}

		txnArgs := types.TransactionOptions{
			Client:         client,
			Password:       password,
			Amount:			amountInWei,
			AccountAddress: address,
			ChainId:        core.ChainId,
			GasMultiplier:  config.GasMultiplier,
		}

		unstake(txnArgs, _stakerId)

		if autoWithdraw {
			log.Info("Starting withdrawal now...")
			s := spinner.New(spinner.CharSets[9], 100 * time.Millisecond)
			s.Start()
			time.Sleep(time.Duration(core.EpochLength) * time.Second)
			s.Stop()
			checkForCommitStateAndWithdraw(client, types.Account{
				Address:  address,
				Password: password,
			}, config, _stakerId)
		}

	},
}

func unstake(txnArgs types.TransactionOptions, _stakerId *big.Int) int {
	stakeManager := utils.GetStakeManager(txnArgs.Client)
	opts := utils.GetOptions(false, txnArgs.AccountAddress, "")
	stakerId,_ := stakeManager.GetStakerId(&opts, common.HexToAddress(txnArgs.AccountAddress))
	staker,_ := stakeManager.GetStaker(&opts, stakerId)
	lock,_ := stakeManager.Locks(&opts,common.HexToAddress(txnArgs.AccountAddress), staker.TokenAddress)
	if lock.Amount.Cmp(big.NewInt(0)) != 0 {
		log.Info("Existing lock, cannot unstake")
		return 0
	}
	txnOpts := utils.GetTxnOpts(txnArgs)
	epoch, err := WaitForCommitState(txnArgs.Client, txnArgs.AccountAddress, "unstake")
	log.Info("Unstaking coins")
	txn, err := stakeManager.Unstake(txnOpts, epoch, _stakerId, txnArgs.Amount)
	utils.CheckError("Error in un-staking: ", err)
	log.Infof("Successfully unstaked %s sRazors", txnArgs.Amount)
	log.Info("Transaction hash: ", txn.Hash())
	unstakeStatus := utils.WaitForBlockCompletion(txnArgs.Client, fmt.Sprintf("%s", txn.Hash()))
	return unstakeStatus
}

func init() {
	rootCmd.AddCommand(unstakeCmd)

	var (
		Address               string
		StakerId              string
		AmountToUnStake       string
		WithdrawAutomatically bool
	)

	unstakeCmd.Flags().StringVarP(&Address, "address", "", "", "user's address")
	unstakeCmd.Flags().StringVarP(&StakerId, "stakerId", "", "", "staker id")
	unstakeCmd.Flags().StringVarP(&AmountToUnStake, "amount", "a", "0", "amount of sRazors to un-stake")
	unstakeCmd.Flags().BoolVarP(&WithdrawAutomatically, "autoWithdraw", "w", false, "withdraw after un-stake automatically")

	unstakeCmd.MarkFlagRequired("address")
	unstakeCmd.MarkFlagRequired("stakerId")
	unstakeCmd.MarkFlagRequired("amount")

}
