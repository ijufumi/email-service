// Code generated by mockery. DO NOT EDIT.

package stubs

import (
	context "context"

	sesv2 "github.com/aws/aws-sdk-go-v2/service/sesv2"
	mock "github.com/stretchr/testify/mock"
)

// ListDedicatedIpPoolsAPIClient is an autogenerated mock type for the ListDedicatedIpPoolsAPIClient type
type ListDedicatedIpPoolsAPIClient struct {
	mock.Mock
}

// ListDedicatedIpPools provides a mock function with given fields: _a0, _a1, _a2
func (_m *ListDedicatedIpPoolsAPIClient) ListDedicatedIpPools(_a0 context.Context, _a1 *sesv2.ListDedicatedIpPoolsInput, _a2 ...func(*sesv2.Options)) (*sesv2.ListDedicatedIpPoolsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *sesv2.ListDedicatedIpPoolsOutput
	if rf, ok := ret.Get(0).(func(context.Context, *sesv2.ListDedicatedIpPoolsInput, ...func(*sesv2.Options)) *sesv2.ListDedicatedIpPoolsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sesv2.ListDedicatedIpPoolsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sesv2.ListDedicatedIpPoolsInput, ...func(*sesv2.Options)) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewListDedicatedIpPoolsAPIClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewListDedicatedIpPoolsAPIClient creates a new instance of ListDedicatedIpPoolsAPIClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewListDedicatedIpPoolsAPIClient(t mockConstructorTestingTNewListDedicatedIpPoolsAPIClient) *ListDedicatedIpPoolsAPIClient {
	mock := &ListDedicatedIpPoolsAPIClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
