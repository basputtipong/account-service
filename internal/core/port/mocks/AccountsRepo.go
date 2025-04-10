// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	port "account-service/internal/core/port"

	mock "github.com/stretchr/testify/mock"
)

// AccountsRepo is an autogenerated mock type for the AccountsRepo type
type AccountsRepo struct {
	mock.Mock
}

// GetByUserId provides a mock function with given fields: userId
func (_m *AccountsRepo) GetByUserId(userId string) ([]port.AccountRepoRes, error) {
	ret := _m.Called(userId)

	if len(ret) == 0 {
		panic("no return value specified for GetByUserId")
	}

	var r0 []port.AccountRepoRes
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]port.AccountRepoRes, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(string) []port.AccountRepoRes); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]port.AccountRepoRes)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrentMainAccountByUserId provides a mock function with given fields: userId
func (_m *AccountsRepo) GetCurrentMainAccountByUserId(userId string) (port.AccountRepoRes, error) {
	ret := _m.Called(userId)

	if len(ret) == 0 {
		panic("no return value specified for GetCurrentMainAccountByUserId")
	}

	var r0 port.AccountRepoRes
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (port.AccountRepoRes, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(string) port.AccountRepoRes); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(port.AccountRepoRes)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFlagByAccountId provides a mock function with given fields: accountIds
func (_m *AccountsRepo) GetFlagByAccountId(accountIds []string) ([]port.Flag, error) {
	ret := _m.Called(accountIds)

	if len(ret) == 0 {
		panic("no return value specified for GetFlagByAccountId")
	}

	var r0 []port.Flag
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]port.Flag, error)); ok {
		return rf(accountIds)
	}
	if rf, ok := ret.Get(0).(func([]string) []port.Flag); ok {
		r0 = rf(accountIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]port.Flag)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(accountIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAccountById provides a mock function with given fields: req
func (_m *AccountsRepo) UpdateAccountById(req port.UpdateAccountRepoReq) error {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAccountById")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(port.UpdateAccountRepoReq) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAccountsRepo creates a new instance of AccountsRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountsRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountsRepo {
	mock := &AccountsRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
