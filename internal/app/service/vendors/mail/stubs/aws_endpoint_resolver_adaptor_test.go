// Code generated by mockery. DO NOT EDIT.

package stubs

import (
	aws "github.com/aws/aws-sdk-go-v2/aws"
	mock "github.com/stretchr/testify/mock"
)

// awsEndpointResolverAdaptor is an autogenerated mock type for the awsEndpointResolverAdaptor type
type awsEndpointResolverAdaptor struct {
	mock.Mock
}

// Execute provides a mock function with given fields: service, region
func (_m *awsEndpointResolverAdaptor) Execute(service string, region string) (aws.Endpoint, error) {
	ret := _m.Called(service, region)

	var r0 aws.Endpoint
	if rf, ok := ret.Get(0).(func(string, string) aws.Endpoint); ok {
		r0 = rf(service, region)
	} else {
		r0 = ret.Get(0).(aws.Endpoint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(service, region)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewAwsEndpointResolverAdaptor interface {
	mock.TestingT
	Cleanup(func())
}

// newAwsEndpointResolverAdaptor creates a new instance of awsEndpointResolverAdaptor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newAwsEndpointResolverAdaptor(t mockConstructorTestingTnewAwsEndpointResolverAdaptor) *awsEndpointResolverAdaptor {
	mock := &awsEndpointResolverAdaptor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
