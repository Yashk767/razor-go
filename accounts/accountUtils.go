package accounts

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts"
	"razor/razorInterface"
)

type AccountUtilsStruct struct {
	keystoreUtils razorInterface.KeystoreInterface
	razorUtils    razorInterface.UtilsInterface
	cryptoUtils   razorInterface.CryptoInterface
	accountUtils  AccountInterface
}

type AccountInterface interface {
	CreateAccount(string, string, AccountUtilsStruct) accounts.Account
	getPrivateKeyFromKeystore(string, string, AccountUtilsStruct) *ecdsa.PrivateKey
	GetPrivateKey(string, string, string, AccountUtilsStruct) *ecdsa.PrivateKey
}

type AccountUtils struct{}

func (accountUtils AccountUtils) CreateAccount(path string, password string, accountUtilsStruct AccountUtilsStruct) accounts.Account {
	return CreateAccount(path, password, accountUtilsStruct)
}

func (accountUtils AccountUtils) getPrivateKeyFromKeystore(keystorePath string, password string, accountUtilsStruct AccountUtilsStruct) *ecdsa.PrivateKey {
	return getPrivateKeyFromKeystore(keystorePath, password, accountUtilsStruct)
}

func (accountUtils AccountUtils) GetPrivateKey(address string, password string, keystorePath string, accountUtilsStruct AccountUtilsStruct) *ecdsa.PrivateKey {
	return GetPrivateKey(address, password, keystorePath, accountUtilsStruct)
}
