// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	pflag "github.com/spf13/pflag"
	mock "github.com/stretchr/testify/mock"
)

// FlagSetInterface is an autogenerated mock type for the FlagSetInterface type
type FlagSetInterface struct {
	mock.Mock
}

// GetBoolAutoWithdraw provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetBoolAutoWithdraw(_a0 *pflag.FlagSet) (bool, error) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBoolRogue provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetBoolRogue(_a0 *pflag.FlagSet) (bool, error) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFloat32GasLimit provides a mock function with given fields: set
func (_m *FlagSetInterface) GetFloat32GasLimit(set *pflag.FlagSet) (float32, error) {
	ret := _m.Called(set)

	var r0 float32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) float32); ok {
		r0 = rf(set)
	} else {
		r0 = ret.Get(0).(float32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(set)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFloat32GasMultiplier provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetFloat32GasMultiplier(_a0 *pflag.FlagSet) (float32, error) {
	ret := _m.Called(_a0)

	var r0 float32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) float32); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(float32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInt32Buffer provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetInt32Buffer(_a0 *pflag.FlagSet) (int32, error) {
	ret := _m.Called(_a0)

	var r0 int32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) int32); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInt32GasPrice provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetInt32GasPrice(_a0 *pflag.FlagSet) (int32, error) {
	ret := _m.Called(_a0)

	var r0 int32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) int32); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInt32Wait provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetInt32Wait(_a0 *pflag.FlagSet) (int32, error) {
	ret := _m.Called(_a0)

	var r0 int32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) int32); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInt8Power provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetInt8Power(_a0 *pflag.FlagSet) (int8, error) {
	ret := _m.Called(_a0)

	var r0 int8
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) int8); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRootFloat32GasLimit provides a mock function with given fields:
func (_m *FlagSetInterface) GetRootFloat32GasLimit() (float32, error) {
	ret := _m.Called()

	var r0 float32
	if rf, ok := ret.Get(0).(func() float32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRootFloat32GasMultiplier provides a mock function with given fields:
func (_m *FlagSetInterface) GetRootFloat32GasMultiplier() (float32, error) {
	ret := _m.Called()

	var r0 float32
	if rf, ok := ret.Get(0).(func() float32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRootInt32Buffer provides a mock function with given fields:
func (_m *FlagSetInterface) GetRootInt32Buffer() (int32, error) {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRootInt32GasPrice provides a mock function with given fields:
func (_m *FlagSetInterface) GetRootInt32GasPrice() (int32, error) {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRootInt32Wait provides a mock function with given fields:
func (_m *FlagSetInterface) GetRootInt32Wait() (int32, error) {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRootStringLogLevel provides a mock function with given fields:
func (_m *FlagSetInterface) GetRootStringLogLevel() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRootStringProvider provides a mock function with given fields:
func (_m *FlagSetInterface) GetRootStringProvider() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringAddress provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetStringAddress(_a0 *pflag.FlagSet) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringFrom provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetStringFrom(_a0 *pflag.FlagSet) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringLogLevel provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetStringLogLevel(_a0 *pflag.FlagSet) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringName provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetStringName(_a0 *pflag.FlagSet) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringPow provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetStringPow(_a0 *pflag.FlagSet) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringProvider provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetStringProvider(_a0 *pflag.FlagSet) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringSelector provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetStringSelector(_a0 *pflag.FlagSet) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringSliceRogueMode provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetStringSliceRogueMode(_a0 *pflag.FlagSet) ([]string, error) {
	ret := _m.Called(_a0)

	var r0 []string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringStatus provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetStringStatus(_a0 *pflag.FlagSet) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringTo provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetStringTo(_a0 *pflag.FlagSet) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringUrl provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetStringUrl(_a0 *pflag.FlagSet) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringValue provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetStringValue(_a0 *pflag.FlagSet) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint16AssetId provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetUint16AssetId(_a0 *pflag.FlagSet) (uint16, error) {
	ret := _m.Called(_a0)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint16); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint16CollectionId provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetUint16CollectionId(_a0 *pflag.FlagSet) (uint16, error) {
	ret := _m.Called(_a0)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint16); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint16JobId provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetUint16JobId(_a0 *pflag.FlagSet) (uint16, error) {
	ret := _m.Called(_a0)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint16); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint16Tolerance provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetUint16Tolerance(_a0 *pflag.FlagSet) (uint16, error) {
	ret := _m.Called(_a0)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint16); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint32Aggregation provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetUint32Aggregation(_a0 *pflag.FlagSet) (uint32, error) {
	ret := _m.Called(_a0)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint32); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint32BountyId provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetUint32BountyId(_a0 *pflag.FlagSet) (uint32, error) {
	ret := _m.Called(_a0)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint32); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint32StakerId provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetUint32StakerId(_a0 *pflag.FlagSet) (uint32, error) {
	ret := _m.Called(_a0)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint32); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint8Commission provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetUint8Commission(_a0 *pflag.FlagSet) (uint8, error) {
	ret := _m.Called(_a0)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint8); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint8SelectorType provides a mock function with given fields: set
func (_m *FlagSetInterface) GetUint8SelectorType(set *pflag.FlagSet) (uint8, error) {
	ret := _m.Called(set)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint8); ok {
		r0 = rf(set)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(set)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint8Weight provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetUint8Weight(_a0 *pflag.FlagSet) (uint8, error) {
	ret := _m.Called(_a0)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint8); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUintSliceJobIds provides a mock function with given fields: _a0
func (_m *FlagSetInterface) GetUintSliceJobIds(_a0 *pflag.FlagSet) ([]uint, error) {
	ret := _m.Called(_a0)

	var r0 []uint
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) []uint); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
