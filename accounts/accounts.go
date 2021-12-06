package accounts

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts"
	"razor/core/types"
	"razor/logger"
	"strings"
)

var log = logger.NewLogger()

func CreateAccount(path string, password string, accountUtils AccountUtilsStruct) accounts.Account {
	newAcc, err := accountUtils.KeystoreUtils.NewAccount(path, password)
	if err != nil {
		log.Fatal("Error in creating account: ", err)
	}
	return newAcc
}

func getPrivateKeyFromKeystore(keystorePath string, password string, accountUtils AccountUtilsStruct) *ecdsa.PrivateKey {
	jsonBytes, err := accountUtils.RazorUtils.ReadFile(keystorePath)
	if err != nil {
		log.Fatal("Error in reading keystore: ", err)
	}
	key, err := accountUtils.KeystoreUtils.DecryptKey(jsonBytes, password)
	if err != nil {
		log.Fatal("Error in fetching private key: ", err)
	}
	return key.PrivateKey
}

func GetPrivateKey(address string, password string, keystorePath string, accountUtilsStruct AccountUtilsStruct) *ecdsa.PrivateKey {
	allAccounts := accountUtilsStruct.KeystoreUtils.Accounts(keystorePath)
	for _, account := range allAccounts {
		if strings.EqualFold(account.Address.Hex(), address) {
			return accountUtilsStruct.AccountUtils.getPrivateKeyFromKeystore(account.URL.Path, password, accountUtilsStruct)
		}
	}
	return nil
}

func Sign(hash []byte, account types.Account, defaultPath string, accountUtilsStruct AccountUtilsStruct) ([]byte, error) {
	privateKey := accountUtilsStruct.AccountUtils.GetPrivateKey(account.Address, account.Password, defaultPath, accountUtilsStruct)
	return accountUtilsStruct.CryptoUtils.Sign(hash, privateKey)
}
