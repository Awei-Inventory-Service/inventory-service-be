// Code generated by MockGen. DO NOT EDIT.
// Source: internal/resource/item_purchase_chain/type.go
//
// Generated by this command:
//
//	mockgen -source=internal/resource/item_purchase_chain/type.go -destination=mocks/app/resource/item_purchase_chain/function.go
//

// Package mock_itempurchasechain is a generated GoMock package.
package mock_itempurchasechain

import (
	context "context"
	reflect "reflect"

	model "github.com/inventory-service/app/model"
	error_wrapper "github.com/inventory-service/lib/error_wrapper"
	gomock "go.uber.org/mock/gomock"
)

// MockItemPurchaseChainResource is a mock of ItemPurchaseChainResource interface.
type MockItemPurchaseChainResource struct {
	ctrl     *gomock.Controller
	recorder *MockItemPurchaseChainResourceMockRecorder
	isgomock struct{}
}

// MockItemPurchaseChainResourceMockRecorder is the mock recorder for MockItemPurchaseChainResource.
type MockItemPurchaseChainResourceMockRecorder struct {
	mock *MockItemPurchaseChainResource
}

// NewMockItemPurchaseChainResource creates a new mock instance.
func NewMockItemPurchaseChainResource(ctrl *gomock.Controller) *MockItemPurchaseChainResource {
	mock := &MockItemPurchaseChainResource{ctrl: ctrl}
	mock.recorder = &MockItemPurchaseChainResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemPurchaseChainResource) EXPECT() *MockItemPurchaseChainResourceMockRecorder {
	return m.recorder
}

// BulkUpdate mocks base method.
func (m *MockItemPurchaseChainResource) BulkUpdate(ctx context.Context, payload []model.ItemPurchaseChainGet) *error_wrapper.ErrorWrapper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkUpdate", ctx, payload)
	ret0, _ := ret[0].(*error_wrapper.ErrorWrapper)
	return ret0
}

// BulkUpdate indicates an expected call of BulkUpdate.
func (mr *MockItemPurchaseChainResourceMockRecorder) BulkUpdate(ctx, payload any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkUpdate", reflect.TypeOf((*MockItemPurchaseChainResource)(nil).BulkUpdate), ctx, payload)
}

// Create mocks base method.
func (m *MockItemPurchaseChainResource) Create(ctx context.Context, itemID, branchID string, purchase model.Purchase) *error_wrapper.ErrorWrapper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, itemID, branchID, purchase)
	ret0, _ := ret[0].(*error_wrapper.ErrorWrapper)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockItemPurchaseChainResourceMockRecorder) Create(ctx, itemID, branchID, purchase any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockItemPurchaseChainResource)(nil).Create), ctx, itemID, branchID, purchase)
}

// Get mocks base method.
func (m *MockItemPurchaseChainResource) Get(ctx context.Context, payload model.ItemPurchaseChain) ([]model.ItemPurchaseChainGet, *error_wrapper.ErrorWrapper) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, payload)
	ret0, _ := ret[0].([]model.ItemPurchaseChainGet)
	ret1, _ := ret[1].(*error_wrapper.ErrorWrapper)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockItemPurchaseChainResourceMockRecorder) Get(ctx, payload any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockItemPurchaseChainResource)(nil).Get), ctx, payload)
}

// Update mocks base method.
func (m *MockItemPurchaseChainResource) Update(ctx context.Context, id string, payload model.ItemPurchaseChain) *error_wrapper.ErrorWrapper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, payload)
	ret0, _ := ret[0].(*error_wrapper.ErrorWrapper)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockItemPurchaseChainResourceMockRecorder) Update(ctx, id, payload any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockItemPurchaseChainResource)(nil).Update), ctx, id, payload)
}
