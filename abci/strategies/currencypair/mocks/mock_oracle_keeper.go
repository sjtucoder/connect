// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	oracletypes "github.com/skip-mev/connect/v2/x/oracle/types"

	types "github.com/skip-mev/connect/v2/pkg/types"
)

// OracleKeeper is an autogenerated mock type for the OracleKeeper type
type OracleKeeper struct {
	mock.Mock
}

type OracleKeeper_Expecter struct {
	mock *mock.Mock
}

func (_m *OracleKeeper) EXPECT() *OracleKeeper_Expecter {
	return &OracleKeeper_Expecter{mock: &_m.Mock}
}

// GetAllCurrencyPairs provides a mock function with given fields: ctx
func (_m *OracleKeeper) GetAllCurrencyPairs(ctx context.Context) []types.CurrencyPair {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllCurrencyPairs")
	}

	var r0 []types.CurrencyPair
	if rf, ok := ret.Get(0).(func(context.Context) []types.CurrencyPair); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.CurrencyPair)
		}
	}

	return r0
}

// OracleKeeper_GetAllCurrencyPairs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllCurrencyPairs'
type OracleKeeper_GetAllCurrencyPairs_Call struct {
	*mock.Call
}

// GetAllCurrencyPairs is a helper method to define mock.On call
//   - ctx context.Context
func (_e *OracleKeeper_Expecter) GetAllCurrencyPairs(ctx interface{}) *OracleKeeper_GetAllCurrencyPairs_Call {
	return &OracleKeeper_GetAllCurrencyPairs_Call{Call: _e.mock.On("GetAllCurrencyPairs", ctx)}
}

func (_c *OracleKeeper_GetAllCurrencyPairs_Call) Run(run func(ctx context.Context)) *OracleKeeper_GetAllCurrencyPairs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *OracleKeeper_GetAllCurrencyPairs_Call) Return(_a0 []types.CurrencyPair) *OracleKeeper_GetAllCurrencyPairs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OracleKeeper_GetAllCurrencyPairs_Call) RunAndReturn(run func(context.Context) []types.CurrencyPair) *OracleKeeper_GetAllCurrencyPairs_Call {
	_c.Call.Return(run)
	return _c
}

// GetCurrencyPairFromID provides a mock function with given fields: ctx, id
func (_m *OracleKeeper) GetCurrencyPairFromID(ctx context.Context, id uint64) (types.CurrencyPair, bool) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetCurrencyPairFromID")
	}

	var r0 types.CurrencyPair
	var r1 bool
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (types.CurrencyPair, bool)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) types.CurrencyPair); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(types.CurrencyPair)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) bool); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// OracleKeeper_GetCurrencyPairFromID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCurrencyPairFromID'
type OracleKeeper_GetCurrencyPairFromID_Call struct {
	*mock.Call
}

// GetCurrencyPairFromID is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint64
func (_e *OracleKeeper_Expecter) GetCurrencyPairFromID(ctx interface{}, id interface{}) *OracleKeeper_GetCurrencyPairFromID_Call {
	return &OracleKeeper_GetCurrencyPairFromID_Call{Call: _e.mock.On("GetCurrencyPairFromID", ctx, id)}
}

func (_c *OracleKeeper_GetCurrencyPairFromID_Call) Run(run func(ctx context.Context, id uint64)) *OracleKeeper_GetCurrencyPairFromID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *OracleKeeper_GetCurrencyPairFromID_Call) Return(cp types.CurrencyPair, found bool) *OracleKeeper_GetCurrencyPairFromID_Call {
	_c.Call.Return(cp, found)
	return _c
}

func (_c *OracleKeeper_GetCurrencyPairFromID_Call) RunAndReturn(run func(context.Context, uint64) (types.CurrencyPair, bool)) *OracleKeeper_GetCurrencyPairFromID_Call {
	_c.Call.Return(run)
	return _c
}

// GetIDForCurrencyPair provides a mock function with given fields: ctx, cp
func (_m *OracleKeeper) GetIDForCurrencyPair(ctx context.Context, cp types.CurrencyPair) (uint64, bool) {
	ret := _m.Called(ctx, cp)

	if len(ret) == 0 {
		panic("no return value specified for GetIDForCurrencyPair")
	}

	var r0 uint64
	var r1 bool
	if rf, ok := ret.Get(0).(func(context.Context, types.CurrencyPair) (uint64, bool)); ok {
		return rf(ctx, cp)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.CurrencyPair) uint64); ok {
		r0 = rf(ctx, cp)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.CurrencyPair) bool); ok {
		r1 = rf(ctx, cp)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// OracleKeeper_GetIDForCurrencyPair_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIDForCurrencyPair'
type OracleKeeper_GetIDForCurrencyPair_Call struct {
	*mock.Call
}

// GetIDForCurrencyPair is a helper method to define mock.On call
//   - ctx context.Context
//   - cp types.CurrencyPair
func (_e *OracleKeeper_Expecter) GetIDForCurrencyPair(ctx interface{}, cp interface{}) *OracleKeeper_GetIDForCurrencyPair_Call {
	return &OracleKeeper_GetIDForCurrencyPair_Call{Call: _e.mock.On("GetIDForCurrencyPair", ctx, cp)}
}

func (_c *OracleKeeper_GetIDForCurrencyPair_Call) Run(run func(ctx context.Context, cp types.CurrencyPair)) *OracleKeeper_GetIDForCurrencyPair_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.CurrencyPair))
	})
	return _c
}

func (_c *OracleKeeper_GetIDForCurrencyPair_Call) Return(_a0 uint64, _a1 bool) *OracleKeeper_GetIDForCurrencyPair_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OracleKeeper_GetIDForCurrencyPair_Call) RunAndReturn(run func(context.Context, types.CurrencyPair) (uint64, bool)) *OracleKeeper_GetIDForCurrencyPair_Call {
	_c.Call.Return(run)
	return _c
}

// GetNumCurrencyPairs provides a mock function with given fields: ctx
func (_m *OracleKeeper) GetNumCurrencyPairs(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetNumCurrencyPairs")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OracleKeeper_GetNumCurrencyPairs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNumCurrencyPairs'
type OracleKeeper_GetNumCurrencyPairs_Call struct {
	*mock.Call
}

// GetNumCurrencyPairs is a helper method to define mock.On call
//   - ctx context.Context
func (_e *OracleKeeper_Expecter) GetNumCurrencyPairs(ctx interface{}) *OracleKeeper_GetNumCurrencyPairs_Call {
	return &OracleKeeper_GetNumCurrencyPairs_Call{Call: _e.mock.On("GetNumCurrencyPairs", ctx)}
}

func (_c *OracleKeeper_GetNumCurrencyPairs_Call) Run(run func(ctx context.Context)) *OracleKeeper_GetNumCurrencyPairs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *OracleKeeper_GetNumCurrencyPairs_Call) Return(_a0 uint64, _a1 error) *OracleKeeper_GetNumCurrencyPairs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OracleKeeper_GetNumCurrencyPairs_Call) RunAndReturn(run func(context.Context) (uint64, error)) *OracleKeeper_GetNumCurrencyPairs_Call {
	_c.Call.Return(run)
	return _c
}

// GetNumRemovedCurrencyPairs provides a mock function with given fields: ctx
func (_m *OracleKeeper) GetNumRemovedCurrencyPairs(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetNumRemovedCurrencyPairs")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OracleKeeper_GetNumRemovedCurrencyPairs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNumRemovedCurrencyPairs'
type OracleKeeper_GetNumRemovedCurrencyPairs_Call struct {
	*mock.Call
}

// GetNumRemovedCurrencyPairs is a helper method to define mock.On call
//   - ctx context.Context
func (_e *OracleKeeper_Expecter) GetNumRemovedCurrencyPairs(ctx interface{}) *OracleKeeper_GetNumRemovedCurrencyPairs_Call {
	return &OracleKeeper_GetNumRemovedCurrencyPairs_Call{Call: _e.mock.On("GetNumRemovedCurrencyPairs", ctx)}
}

func (_c *OracleKeeper_GetNumRemovedCurrencyPairs_Call) Run(run func(ctx context.Context)) *OracleKeeper_GetNumRemovedCurrencyPairs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *OracleKeeper_GetNumRemovedCurrencyPairs_Call) Return(_a0 uint64, _a1 error) *OracleKeeper_GetNumRemovedCurrencyPairs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OracleKeeper_GetNumRemovedCurrencyPairs_Call) RunAndReturn(run func(context.Context) (uint64, error)) *OracleKeeper_GetNumRemovedCurrencyPairs_Call {
	_c.Call.Return(run)
	return _c
}

// GetPriceForCurrencyPair provides a mock function with given fields: ctx, cp
func (_m *OracleKeeper) GetPriceForCurrencyPair(ctx context.Context, cp types.CurrencyPair) (oracletypes.QuotePrice, error) {
	ret := _m.Called(ctx, cp)

	if len(ret) == 0 {
		panic("no return value specified for GetPriceForCurrencyPair")
	}

	var r0 oracletypes.QuotePrice
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.CurrencyPair) (oracletypes.QuotePrice, error)); ok {
		return rf(ctx, cp)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.CurrencyPair) oracletypes.QuotePrice); ok {
		r0 = rf(ctx, cp)
	} else {
		r0 = ret.Get(0).(oracletypes.QuotePrice)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.CurrencyPair) error); ok {
		r1 = rf(ctx, cp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OracleKeeper_GetPriceForCurrencyPair_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPriceForCurrencyPair'
type OracleKeeper_GetPriceForCurrencyPair_Call struct {
	*mock.Call
}

// GetPriceForCurrencyPair is a helper method to define mock.On call
//   - ctx context.Context
//   - cp types.CurrencyPair
func (_e *OracleKeeper_Expecter) GetPriceForCurrencyPair(ctx interface{}, cp interface{}) *OracleKeeper_GetPriceForCurrencyPair_Call {
	return &OracleKeeper_GetPriceForCurrencyPair_Call{Call: _e.mock.On("GetPriceForCurrencyPair", ctx, cp)}
}

func (_c *OracleKeeper_GetPriceForCurrencyPair_Call) Run(run func(ctx context.Context, cp types.CurrencyPair)) *OracleKeeper_GetPriceForCurrencyPair_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.CurrencyPair))
	})
	return _c
}

func (_c *OracleKeeper_GetPriceForCurrencyPair_Call) Return(_a0 oracletypes.QuotePrice, _a1 error) *OracleKeeper_GetPriceForCurrencyPair_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OracleKeeper_GetPriceForCurrencyPair_Call) RunAndReturn(run func(context.Context, types.CurrencyPair) (oracletypes.QuotePrice, error)) *OracleKeeper_GetPriceForCurrencyPair_Call {
	_c.Call.Return(run)
	return _c
}

// NewOracleKeeper creates a new instance of OracleKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOracleKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *OracleKeeper {
	mock := &OracleKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
