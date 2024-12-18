// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/itsLeonB/go-mate/internal/entity"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// SubscriptionRepository is an autogenerated mock type for the SubscriptionRepository type
type SubscriptionRepository struct {
	mock.Mock
}

// FindByUserIDandModel provides a mock function with given fields: ctx, userID, model
func (_m *SubscriptionRepository) FindByUserIDandModel(ctx context.Context, userID uuid.UUID, model string) ([]*entity.UserSubscription, error) {
	ret := _m.Called(ctx, userID, model)

	if len(ret) == 0 {
		panic("no return value specified for FindByUserIDandModel")
	}

	var r0 []*entity.UserSubscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) ([]*entity.UserSubscription, error)); ok {
		return rf(ctx, userID, model)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) []*entity.UserSubscription); ok {
		r0 = rf(ctx, userID, model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.UserSubscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, string) error); ok {
		r1 = rf(ctx, userID, model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: ctx, subscription
func (_m *SubscriptionRepository) Insert(ctx context.Context, subscription *entity.UserSubscription) error {
	ret := _m.Called(ctx, subscription)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.UserSubscription) error); ok {
		r0 = rf(ctx, subscription)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSubscriptionRepository creates a new instance of SubscriptionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSubscriptionRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *SubscriptionRepository {
	mock := &SubscriptionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
