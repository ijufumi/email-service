// Code generated by mockery. DO NOT EDIT.

package stubs

import (
	aws "github.com/aws/aws-sdk-go-v2/aws"
	endpoints "github.com/aws/aws-sdk-go-v2/service/sesv2/internal/endpoints"
	mock "github.com/stretchr/testify/mock"
)

// EndpointResolver is an autogenerated mock type for the EndpointResolver type
type EndpointResolver struct {
	mock.Mock
}

// ResolveEndpoint provides a mock function with given fields: region, options
func (_m *EndpointResolver) ResolveEndpoint(region string, options endpoints.Options) (aws.Endpoint, error) {
	ret := _m.Called(region, options)

	var r0 aws.Endpoint
	if rf, ok := ret.Get(0).(func(string, endpoints.Options) aws.Endpoint); ok {
		r0 = rf(region, options)
	} else {
		r0 = ret.Get(0).(aws.Endpoint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, endpoints.Options) error); ok {
		r1 = rf(region, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewEndpointResolver interface {
	mock.TestingT
	Cleanup(func())
}

// NewEndpointResolver creates a new instance of EndpointResolver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEndpointResolver(t mockConstructorTestingTNewEndpointResolver) *EndpointResolver {
	mock := &EndpointResolver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
