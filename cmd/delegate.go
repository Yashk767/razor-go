package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"
)

// delegateCmd represents the delegate command
var delegateCmd = &cobra.Command{
	Use:   "delegate",
	Short: "delegate is used by delegator to stake coins on the network without setting up a node",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in getting config: ", err)

		password := utils.PasswordPrompt()
		address, _ := cmd.Flags().GetString("address")
		stakerId, _ := cmd.Flags().GetString("stakerId")
		amount, _ := cmd.Flags().GetString("amount")

		client := utils.ConnectToClient(config.Provider)

		balance, err := utils.FetchBalance(client, address)
		utils.CheckError("Error in fetching balance for account " + address + ": ", err)

		amountInWei := utils.GetAmountWithChecks(amount, balance)


		_stakerId, ok := new(big.Int).SetString(stakerId, 10)
		if !ok {
			log.Fatal("SetString error while converting stakerId")
		}

		txnOpts := types.TransactionOptions{
			Client:         client,
			Password:       password,
			Amount:         amountInWei,
			AccountAddress: address,
			ChainId:        core.ChainId,
			GasMultiplier:  config.GasMultiplier,
		}

		approve(txnOpts)
		delegate(txnOpts, _stakerId)


	},
}

func delegate(txnArgs types.TransactionOptions, _stakerId *big.Int) int {
	stakeManager := utils.GetStakeManager(txnArgs.Client)
	epoch, err := WaitForCommitState(txnArgs.Client, txnArgs.AccountAddress, "delegate")
	utils.CheckError("Error in fetching epoch: ", err)
	txnOpts := utils.GetOptions(false, txnArgs.AccountAddress, "")
	staker,_ := stakeManager.GetStaker(&txnOpts, _stakerId)

	if !staker.AcceptDelegation {
		log.Info("Chosen staker does not accept delegation")
		return 0
	}
	//log.Infof("Delegating %s razors to Staker %s", txnArgs.Amount, _stakerId)
	log.Infof("Delegating razors to Staker %s", _stakerId)
	txn, err := stakeManager.Delegate(utils.GetTxnOpts(txnArgs), epoch, txnArgs.Amount, _stakerId)
	if err != nil {
		utils.CheckError("Error in delegating: ", err)
	}
	log.Infof("Sending Delegate transaction...")
	log.Infof("Transaction hash: %s", txn.Hash())
	delegateStatus := utils.WaitForBlockCompletion(txnArgs.Client, txn.Hash().String())
	return delegateStatus
}

func init() {
	rootCmd.AddCommand(delegateCmd)
	var (
		Amount   string
		Address  string
		StakerId string
	)

	delegateCmd.Flags().StringVarP(&Amount, "amount", "a", "0", "amount to stake (in Wei)")
	delegateCmd.Flags().StringVarP(&Address, "address", "", "", "your account address")
	delegateCmd.Flags().StringVarP(&StakerId, "stakerId", "", "", "staker id")

	delegateCmd.MarkFlagRequired("amount")
	delegateCmd.MarkFlagRequired("address")
	delegateCmd.MarkFlagRequired("stakerId")

}
