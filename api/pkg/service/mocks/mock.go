// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	Telegram_Market "Telegram-Market"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUser) CreateUser(userId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserMockRecorder) CreateUser(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUser)(nil).CreateUser), userId)
}

// DeleteUser mocks base method.
func (m *MockUser) DeleteUser(userId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserMockRecorder) DeleteUser(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUser)(nil).DeleteUser), userId)
}

// GetUserById mocks base method.
func (m *MockUser) GetUserById(userId int64) (Telegram_Market.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", userId)
	ret0, _ := ret[0].(Telegram_Market.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserMockRecorder) GetUserById(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUser)(nil).GetUserById), userId)
}

// GetUsers mocks base method.
func (m *MockUser) GetUsers() ([]Telegram_Market.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]Telegram_Market.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUser)(nil).GetUsers))
}

// UpdateUser mocks base method.
func (m *MockUser) UpdateUser(userId int64, users Telegram_Market.Users) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", userId, users)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserMockRecorder) UpdateUser(userId, users interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUser)(nil).UpdateUser), userId, users)
}

// MockProducts is a mock of Products interface.
type MockProducts struct {
	ctrl     *gomock.Controller
	recorder *MockProductsMockRecorder
}

// MockProductsMockRecorder is the mock recorder for MockProducts.
type MockProductsMockRecorder struct {
	mock *MockProducts
}

// NewMockProducts creates a new mock instance.
func NewMockProducts(ctrl *gomock.Controller) *MockProducts {
	mock := &MockProducts{ctrl: ctrl}
	mock.recorder = &MockProductsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducts) EXPECT() *MockProductsMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockProducts) CreateProduct(product Telegram_Market.Products) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", product)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockProductsMockRecorder) CreateProduct(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockProducts)(nil).CreateProduct), product)
}

// DeleteProduct mocks base method.
func (m *MockProducts) DeleteProduct(productId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", productId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockProductsMockRecorder) DeleteProduct(productId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockProducts)(nil).DeleteProduct), productId)
}

// GetAllProducts mocks base method.
func (m *MockProducts) GetAllProducts() ([]Telegram_Market.Products, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProducts")
	ret0, _ := ret[0].([]Telegram_Market.Products)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProducts indicates an expected call of GetAllProducts.
func (mr *MockProductsMockRecorder) GetAllProducts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProducts", reflect.TypeOf((*MockProducts)(nil).GetAllProducts))
}

// GetProductById mocks base method.
func (m *MockProducts) GetProductById(productId int64) (Telegram_Market.Products, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductById", productId)
	ret0, _ := ret[0].(Telegram_Market.Products)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductById indicates an expected call of GetProductById.
func (mr *MockProductsMockRecorder) GetProductById(productId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductById", reflect.TypeOf((*MockProducts)(nil).GetProductById), productId)
}

// UpdateProduct mocks base method.
func (m *MockProducts) UpdateProduct(productId int64, product Telegram_Market.UpdateProducts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", productId, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockProductsMockRecorder) UpdateProduct(productId, product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockProducts)(nil).UpdateProduct), productId, product)
}

// MockLocations is a mock of Locations interface.
type MockLocations struct {
	ctrl     *gomock.Controller
	recorder *MockLocationsMockRecorder
}

// MockLocationsMockRecorder is the mock recorder for MockLocations.
type MockLocationsMockRecorder struct {
	mock *MockLocations
}

// NewMockLocations creates a new mock instance.
func NewMockLocations(ctrl *gomock.Controller) *MockLocations {
	mock := &MockLocations{ctrl: ctrl}
	mock.recorder = &MockLocationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocations) EXPECT() *MockLocationsMockRecorder {
	return m.recorder
}

// CreateLocation mocks base method.
func (m *MockLocations) CreateLocation(location Telegram_Market.Locations) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLocation", location)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLocation indicates an expected call of CreateLocation.
func (mr *MockLocationsMockRecorder) CreateLocation(location interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLocation", reflect.TypeOf((*MockLocations)(nil).CreateLocation), location)
}

// DeleteLocation mocks base method.
func (m *MockLocations) DeleteLocation(locationId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLocation", locationId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLocation indicates an expected call of DeleteLocation.
func (mr *MockLocationsMockRecorder) DeleteLocation(locationId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLocation", reflect.TypeOf((*MockLocations)(nil).DeleteLocation), locationId)
}

// GetAllLocations mocks base method.
func (m *MockLocations) GetAllLocations() ([]Telegram_Market.Locations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllLocations")
	ret0, _ := ret[0].([]Telegram_Market.Locations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllLocations indicates an expected call of GetAllLocations.
func (mr *MockLocationsMockRecorder) GetAllLocations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllLocations", reflect.TypeOf((*MockLocations)(nil).GetAllLocations))
}

// GetLocationById mocks base method.
func (m *MockLocations) GetLocationById(locationId int64) (Telegram_Market.Locations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocationById", locationId)
	ret0, _ := ret[0].(Telegram_Market.Locations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocationById indicates an expected call of GetLocationById.
func (mr *MockLocationsMockRecorder) GetLocationById(locationId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocationById", reflect.TypeOf((*MockLocations)(nil).GetLocationById), locationId)
}

// UpdateLocations mocks base method.
func (m *MockLocations) UpdateLocations(locationId int64, location Telegram_Market.UpdateLocations) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLocations", locationId, location)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLocations indicates an expected call of UpdateLocations.
func (mr *MockLocationsMockRecorder) UpdateLocations(locationId, location interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLocations", reflect.TypeOf((*MockLocations)(nil).UpdateLocations), locationId, location)
}

// MockWallets is a mock of Wallets interface.
type MockWallets struct {
	ctrl     *gomock.Controller
	recorder *MockWalletsMockRecorder
}

// MockWalletsMockRecorder is the mock recorder for MockWallets.
type MockWalletsMockRecorder struct {
	mock *MockWallets
}

// NewMockWallets creates a new mock instance.
func NewMockWallets(ctrl *gomock.Controller) *MockWallets {
	mock := &MockWallets{ctrl: ctrl}
	mock.recorder = &MockWalletsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWallets) EXPECT() *MockWalletsMockRecorder {
	return m.recorder
}

// CreateWallet mocks base method.
func (m *MockWallets) CreateWallet(wallet Telegram_Market.Wallet) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWallet", wallet)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWallet indicates an expected call of CreateWallet.
func (mr *MockWalletsMockRecorder) CreateWallet(wallet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWallet", reflect.TypeOf((*MockWallets)(nil).CreateWallet), wallet)
}

// DeleteWallet mocks base method.
func (m *MockWallets) DeleteWallet(WalletID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWallet", WalletID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWallet indicates an expected call of DeleteWallet.
func (mr *MockWalletsMockRecorder) DeleteWallet(WalletID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWallet", reflect.TypeOf((*MockWallets)(nil).DeleteWallet), WalletID)
}

// GetAllWallets mocks base method.
func (m *MockWallets) GetAllWallets() ([]Telegram_Market.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllWallets")
	ret0, _ := ret[0].([]Telegram_Market.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllWallets indicates an expected call of GetAllWallets.
func (mr *MockWalletsMockRecorder) GetAllWallets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllWallets", reflect.TypeOf((*MockWallets)(nil).GetAllWallets))
}

// GetWalletById mocks base method.
func (m *MockWallets) GetWalletById(WalletID int64) (*Telegram_Market.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletById", WalletID)
	ret0, _ := ret[0].(*Telegram_Market.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletById indicates an expected call of GetWalletById.
func (mr *MockWalletsMockRecorder) GetWalletById(WalletID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletById", reflect.TypeOf((*MockWallets)(nil).GetWalletById), WalletID)
}

// UpdateWallet mocks base method.
func (m *MockWallets) UpdateWallet(walletID int64, Wallet Telegram_Market.Wallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWallet", walletID, Wallet)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWallet indicates an expected call of UpdateWallet.
func (mr *MockWalletsMockRecorder) UpdateWallet(walletID, Wallet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWallet", reflect.TypeOf((*MockWallets)(nil).UpdateWallet), walletID, Wallet)
}