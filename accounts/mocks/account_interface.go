// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	ecdsa "crypto/ecdsa"

	accounts "github.com/ethereum/go-ethereum/accounts"

	keystore "github.com/ethereum/go-ethereum/accounts/keystore"

	mock "github.com/stretchr/testify/mock"

	types "razor/core/types"
)

// AccountInterface is an autogenerated mock type for the AccountInterface type
type AccountInterface struct {
	mock.Mock
}

// Accounts provides a mock function with given fields: _a0
func (_m *AccountInterface) Accounts(_a0 string) []accounts.Account {
	ret := _m.Called(_a0)

	var r0 []accounts.Account
	if rf, ok := ret.Get(0).(func(string) []accounts.Account); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]accounts.Account)
		}
	}

	return r0
}

// CreateAccount provides a mock function with given fields: _a0, _a1
func (_m *AccountInterface) CreateAccount(_a0 string, _a1 string) accounts.Account {
	ret := _m.Called(_a0, _a1)

	var r0 accounts.Account
	if rf, ok := ret.Get(0).(func(string, string) accounts.Account); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(accounts.Account)
	}

	return r0
}

// DecryptKey provides a mock function with given fields: _a0, _a1
func (_m *AccountInterface) DecryptKey(_a0 []byte, _a1 string) (*keystore.Key, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *keystore.Key
	if rf, ok := ret.Get(0).(func([]byte, string) *keystore.Key); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keystore.Key)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPrivateKey provides a mock function with given fields: _a0, _a1, _a2
func (_m *AccountInterface) GetPrivateKey(_a0 string, _a1 string, _a2 string) *ecdsa.PrivateKey {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *ecdsa.PrivateKey
	if rf, ok := ret.Get(0).(func(string, string, string) *ecdsa.PrivateKey); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecdsa.PrivateKey)
		}
	}

	return r0
}

// GetPrivateKeyFromKeystore provides a mock function with given fields: _a0, _a1
func (_m *AccountInterface) GetPrivateKeyFromKeystore(_a0 string, _a1 string) *ecdsa.PrivateKey {
	ret := _m.Called(_a0, _a1)

	var r0 *ecdsa.PrivateKey
	if rf, ok := ret.Get(0).(func(string, string) *ecdsa.PrivateKey); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecdsa.PrivateKey)
		}
	}

	return r0
}

// NewAccount provides a mock function with given fields: _a0, _a1
func (_m *AccountInterface) NewAccount(_a0 string, _a1 string) (accounts.Account, error) {
	ret := _m.Called(_a0, _a1)

	var r0 accounts.Account
	if rf, ok := ret.Get(0).(func(string, string) accounts.Account); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(accounts.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadFile provides a mock function with given fields: _a0
func (_m *AccountInterface) ReadFile(_a0 string) ([]byte, error) {
	ret := _m.Called(_a0)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Sign provides a mock function with given fields: _a0, _a1
func (_m *AccountInterface) Sign(_a0 []byte, _a1 *ecdsa.PrivateKey) ([]byte, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte, *ecdsa.PrivateKey) []byte); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, *ecdsa.PrivateKey) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignData provides a mock function with given fields: _a0, _a1, _a2
func (_m *AccountInterface) SignData(_a0 []byte, _a1 types.Account, _a2 string) ([]byte, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte, types.Account, string) []byte); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, types.Account, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
