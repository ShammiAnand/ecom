// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	types "github.com/shammianand/ecom/types"
	mock "github.com/stretchr/testify/mock"
)

// OrderStore is an autogenerated mock type for the OrderStore type
type OrderStore struct {
	mock.Mock
}

// CancelOrder provides a mock function with given fields: _a0
func (_m *OrderStore) CancelOrder(_a0 int) (bool, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CancelOrder")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (bool, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrder provides a mock function with given fields: _a0
func (_m *OrderStore) CreateOrder(_a0 types.Order) (int, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrder")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Order) (int, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(types.Order) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(types.Order) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrderItem provides a mock function with given fields: _a0
func (_m *OrderStore) CreateOrderItem(_a0 types.OrderItem) (int, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrderItem")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(types.OrderItem) (int, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(types.OrderItem) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(types.OrderItem) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrders provides a mock function with given fields: _a0
func (_m *OrderStore) GetOrders(_a0 int) ([]types.Order, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetOrders")
	}

	var r0 []types.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]types.Order, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int) []types.Order); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOrderStore creates a new instance of OrderStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrderStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrderStore {
	mock := &OrderStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
