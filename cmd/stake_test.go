package cmd

import (
	"razor/core/types"
	"testing"
)

func Test_stakeCoins(t *testing.T) {
	type args struct {
		txnArgs types.TransactionOptions
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
