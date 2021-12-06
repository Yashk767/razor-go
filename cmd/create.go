package cmd

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	razorAccounts "razor/accounts"
	"razor/razorInterface"
	"razor/utils"
)

var accountUtils razorAccounts.AccountInterface

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create command can be used to create new accounts",
	Long: `For a new user to start doing anything, an account is required. This command helps the user to create a new account secured by a password so that only that user would be able to use the account

Example: 
  ./razor create`,
	Run: initialiseCreate,
}

func initialiseCreate(cmd *cobra.Command, args []string) {
	utilsStruct := UtilsStruct{
		razorUtils: razorUtils,
		cmdUtils:   cmdUtils,
		accountUtils: razorAccounts.AccountUtilsStruct{
			KeystoreUtils: keystoreUtils,
			RazorUtils:    razorUtils,
			AccountUtils:  accountUtils,
			CryptoUtils:   cryptoUtils,
		},
	}
	utilsStruct.executeCreate(cmd.Flags())
}

func (utilsStruct UtilsStruct) executeCreate(flagSet *pflag.FlagSet) {
	password := utilsStruct.razorUtils.AssignPassword(flagSet)
	account, err := utilsStruct.cmdUtils.Create(password, utilsStruct)
	utils.CheckError("Create error: ", err)
	log.Info("Account address: ", account.Address)
	log.Info("Keystore Path: ", account.URL)
}

func Create(password string, utilsStruct UtilsStruct) (accounts.Account, error) {
	path, err := utilsStruct.razorUtils.GetDefaultPath()
	if err != nil {
		log.Error("Error in fetching .razor directory")
		return accounts.Account{Address: common.Address{0x00}}, err
	}
	account := utilsStruct.accountUtils.AccountUtils.CreateAccount(path, password, utilsStruct.accountUtils)
	return account, nil
}

func init() {
	razorUtils = razorInterface.Utils{}
	accountUtils = razorAccounts.AccountUtils{}
	cmdUtils = UtilsCmd{}
	keystoreUtils = razorInterface.KeystoreUtils{}
	cryptoUtils = razorInterface.CryptoUtils{}

	rootCmd.AddCommand(createCmd)

	var (
		Password string
	)

	createCmd.Flags().StringVarP(&Password, "password", "", "", "password file path to protect the keystore")
}
