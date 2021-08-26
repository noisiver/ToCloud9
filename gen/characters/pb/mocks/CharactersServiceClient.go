// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	pb "github.com/walkline/ToCloud9/gen/characters/pb"
)

// CharactersServiceClient is an autogenerated mock type for the CharactersServiceClient type
type CharactersServiceClient struct {
	mock.Mock
}

// AccountDataForAccount provides a mock function with given fields: ctx, in, opts
func (_m *CharactersServiceClient) AccountDataForAccount(ctx context.Context, in *pb.AccountDataForAccountRequest, opts ...grpc.CallOption) (*pb.AccountDataForAccountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.AccountDataForAccountResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.AccountDataForAccountRequest, ...grpc.CallOption) *pb.AccountDataForAccountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.AccountDataForAccountResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.AccountDataForAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CharactersToLoginByGUID provides a mock function with given fields: ctx, in, opts
func (_m *CharactersServiceClient) CharactersToLoginByGUID(ctx context.Context, in *pb.CharactersToLoginByGUIDRequest, opts ...grpc.CallOption) (*pb.CharactersToLoginByGUIDResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.CharactersToLoginByGUIDResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CharactersToLoginByGUIDRequest, ...grpc.CallOption) *pb.CharactersToLoginByGUIDResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.CharactersToLoginByGUIDResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.CharactersToLoginByGUIDRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CharactersToLoginForAccount provides a mock function with given fields: ctx, in, opts
func (_m *CharactersServiceClient) CharactersToLoginForAccount(ctx context.Context, in *pb.CharactersToLoginForAccountRequest, opts ...grpc.CallOption) (*pb.CharactersToLoginForAccountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.CharactersToLoginForAccountResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CharactersToLoginForAccountRequest, ...grpc.CallOption) *pb.CharactersToLoginForAccountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.CharactersToLoginForAccountResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.CharactersToLoginForAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
