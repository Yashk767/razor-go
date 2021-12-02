package cmd

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"razor/razorInterface"
	"razor/utils"
)

var accountUtils razorInterface.AccountInterface

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create command can be used to create new accounts",
	Long: `For a new user to start doing anything, an account is required. This command helps the user to create a new account secured by a password so that only that user would be able to use the account

Example: 
  ./razor create`,
	Run: func(cmd *cobra.Command, args []string) {
		utilsStruct := UtilsStruct{
			razorUtils:   razorUtils,
			accountUtils: accountUtils,
		}
		account, err := utilsStruct.Create(cmd.Flags())
		utils.CheckError("Create error: ", err)
		log.Info("Account address: ", account.Address)
		log.Info("Keystore Path: ", account.URL)
	},
}

func (utilsStruct UtilsStruct) Create(flagSet *pflag.FlagSet) (accounts.Account, error) {
	password := utilsStruct.razorUtils.AssignPassword(flagSet)
	path, err := utilsStruct.razorUtils.GetDefaultPath()
	if err != nil {
		log.Error("Error in fetching .razor directory")
		return accounts.Account{Address: common.Address{0x00}}, err
	}
	account := utilsStruct.accountUtils.CreateAccount(path, password)
	return account, nil
}

func init() {
	razorUtils = razorInterface.Utils{}
	accountUtils = razorInterface.AccountUtils{}

	rootCmd.AddCommand(createCmd)

	var (
		Password string
	)

	createCmd.Flags().StringVarP(&Password, "password", "", "", "password file path to protect the keystore")
}
