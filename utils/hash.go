package utils

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"runtime"
)

func EcRecover(data, sig hexutil.Bytes) (common.Address, error) {
	if len(sig) != 65 {
		return common.Address{}, fmt.Errorf("signature must be 65 bytes long")
	}

	rpk, err := crypto.SigToPub(SignHash(data), sig)
	if err != nil {
		return common.Address{}, err
	}
	log.Info("AAAA Commit Num of go routines 407: ", runtime.NumGoroutine())
	return crypto.PubkeyToAddress(*rpk), nil
}

func SignHash(data []byte) []byte {
	log.Info("AAAA Commit Num of go routines 401: ", runtime.NumGoroutine())
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	log.Info("AAAA Commit Num of go routines 402: ", runtime.NumGoroutine())
	return crypto.Keccak256([]byte(msg))
}

func generateCacheKey(url string, body map[string]interface{}) (string, error) {
	var bodyString string
	if body != nil {
		// Convert the body to a JSON string
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			log.Error("Error in marshalling body: ", err)
			return "", err
		}
		bodyString = string(bodyBytes)
	}

	hash := solsha3.SoliditySHA3([]string{"string", "string"}, []interface{}{url, bodyString})
	return common.BytesToHash(hash).Hex(), nil
}
