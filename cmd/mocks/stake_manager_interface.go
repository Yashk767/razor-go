// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	big "math/big"
	bindings "razor/pkg/bindings"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	coretypes "razor/core/types"

	ethclient "github.com/ethereum/go-ethereum/ethclient"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// StakeManagerInterface is an autogenerated mock type for the StakeManagerInterface type
type StakeManagerInterface struct {
	mock.Mock
}

// ApproveUnstake provides a mock function with given fields: client, opts, staker, amount
func (_m *StakeManagerInterface) ApproveUnstake(client *ethclient.Client, opts *bind.TransactOpts, staker bindings.StructsStaker, amount *big.Int) (*types.Transaction, error) {
	ret := _m.Called(client, opts, staker, amount)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, bindings.StructsStaker, *big.Int) *types.Transaction); ok {
		r0 = rf(client, opts, staker, amount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, bindings.StructsStaker, *big.Int) error); ok {
		r1 = rf(client, opts, staker, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClaimStakeReward provides a mock function with given fields: client, opts
func (_m *StakeManagerInterface) ClaimStakeReward(client *ethclient.Client, opts *bind.TransactOpts) (*types.Transaction, error) {
	ret := _m.Called(client, opts)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts) *types.Transaction); ok {
		r0 = rf(client, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts) error); ok {
		r1 = rf(client, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delegate provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *StakeManagerInterface) Delegate(_a0 *ethclient.Client, _a1 *bind.TransactOpts, _a2 uint32, _a3 *big.Int) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) *types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBountyLock provides a mock function with given fields: _a0, _a1, _a2
func (_m *StakeManagerInterface) GetBountyLock(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint32) (coretypes.BountyLock, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 coretypes.BountyLock
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32) coretypes.BountyLock); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(coretypes.BountyLock)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMaturity provides a mock function with given fields: _a0, _a1, _a2
func (_m *StakeManagerInterface) GetMaturity(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint32) (uint16, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32) uint16); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitiateWithdraw provides a mock function with given fields: _a0, _a1, _a2
func (_m *StakeManagerInterface) InitiateWithdraw(_a0 *ethclient.Client, _a1 *bind.TransactOpts, _a2 uint32) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32) *types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RedeemBounty provides a mock function with given fields: _a0, _a1, _a2
func (_m *StakeManagerInterface) RedeemBounty(_a0 *ethclient.Client, _a1 *bind.TransactOpts, _a2 uint32) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32) *types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetUnstakeLock provides a mock function with given fields: _a0, _a1, _a2
func (_m *StakeManagerInterface) ResetUnstakeLock(_a0 *ethclient.Client, _a1 *bind.TransactOpts, _a2 uint32) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32) *types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetDelegationAcceptance provides a mock function with given fields: _a0, _a1, _a2
func (_m *StakeManagerInterface) SetDelegationAcceptance(_a0 *ethclient.Client, _a1 *bind.TransactOpts, _a2 bool) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, bool) *types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, bool) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stake provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *StakeManagerInterface) Stake(_a0 *ethclient.Client, _a1 *bind.TransactOpts, _a2 uint32, _a3 *big.Int) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) *types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StakerInfo provides a mock function with given fields: _a0, _a1, _a2
func (_m *StakeManagerInterface) StakerInfo(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint32) (coretypes.Staker, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 coretypes.Staker
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32) coretypes.Staker); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(coretypes.Staker)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnlockWithdraw provides a mock function with given fields: client, opts, stakerId
func (_m *StakeManagerInterface) UnlockWithdraw(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32) (*types.Transaction, error) {
	ret := _m.Called(client, opts, stakerId)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32) *types.Transaction); ok {
		r0 = rf(client, opts, stakerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32) error); ok {
		r1 = rf(client, opts, stakerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unstake provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *StakeManagerInterface) Unstake(_a0 *ethclient.Client, _a1 *bind.TransactOpts, _a2 uint32, _a3 *big.Int) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) *types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCommission provides a mock function with given fields: _a0, _a1, _a2
func (_m *StakeManagerInterface) UpdateCommission(_a0 *ethclient.Client, _a1 *bind.TransactOpts, _a2 uint8) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint8) *types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint8) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
