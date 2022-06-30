// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/spacemeshos/go-spacemesh/common/types"
	datastore "github.com/spacemeshos/go-spacemesh/datastore"
	types0 "github.com/spacemeshos/go-spacemesh/fetch/types"
	p2p "github.com/spacemeshos/go-spacemesh/p2p"
)

// Mockfetcher is a mock of fetcher interface.
type Mockfetcher struct {
	ctrl     *gomock.Controller
	recorder *MockfetcherMockRecorder
}

// MockfetcherMockRecorder is the mock recorder for Mockfetcher.
type MockfetcherMockRecorder struct {
	mock *Mockfetcher
}

// NewMockfetcher creates a new mock instance.
func NewMockfetcher(ctrl *gomock.Controller) *Mockfetcher {
	mock := &Mockfetcher{ctrl: ctrl}
	mock.recorder = &MockfetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockfetcher) EXPECT() *MockfetcherMockRecorder {
	return m.recorder
}

// AddPeersFromHash mocks base method.
func (m *Mockfetcher) AddPeersFromHash(arg0 types.Hash32, arg1 []types.Hash32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddPeersFromHash", arg0, arg1)
}

// AddPeersFromHash indicates an expected call of AddPeersFromHash.
func (mr *MockfetcherMockRecorder) AddPeersFromHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPeersFromHash", reflect.TypeOf((*Mockfetcher)(nil).AddPeersFromHash), arg0, arg1)
}

// GetEpochATXIDs mocks base method.
func (m *Mockfetcher) GetEpochATXIDs(arg0 context.Context, arg1 types.EpochID, arg2 func([]byte, p2p.Peer), arg3 func(error)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEpochATXIDs", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetEpochATXIDs indicates an expected call of GetEpochATXIDs.
func (mr *MockfetcherMockRecorder) GetEpochATXIDs(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEpochATXIDs", reflect.TypeOf((*Mockfetcher)(nil).GetEpochATXIDs), arg0, arg1, arg2, arg3)
}

// GetHash mocks base method.
func (m *Mockfetcher) GetHash(arg0 types.Hash32, arg1 datastore.Hint, arg2 bool) chan types0.HashDataPromiseResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHash", arg0, arg1, arg2)
	ret0, _ := ret[0].(chan types0.HashDataPromiseResult)
	return ret0
}

// GetHash indicates an expected call of GetHash.
func (mr *MockfetcherMockRecorder) GetHash(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHash", reflect.TypeOf((*Mockfetcher)(nil).GetHash), arg0, arg1, arg2)
}

// GetHashes mocks base method.
func (m *Mockfetcher) GetHashes(arg0 []types.Hash32, arg1 datastore.Hint, arg2 bool) map[types.Hash32]chan types0.HashDataPromiseResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHashes", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[types.Hash32]chan types0.HashDataPromiseResult)
	return ret0
}

// GetHashes indicates an expected call of GetHashes.
func (mr *MockfetcherMockRecorder) GetHashes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHashes", reflect.TypeOf((*Mockfetcher)(nil).GetHashes), arg0, arg1, arg2)
}

// GetLayerData mocks base method.
func (m *Mockfetcher) GetLayerData(arg0 context.Context, arg1 types.LayerID, arg2 func([]byte, p2p.Peer, int), arg3 func(error, p2p.Peer, int)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLayerData", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetLayerData indicates an expected call of GetLayerData.
func (mr *MockfetcherMockRecorder) GetLayerData(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLayerData", reflect.TypeOf((*Mockfetcher)(nil).GetLayerData), arg0, arg1, arg2, arg3)
}

// RegisterPeerHashes mocks base method.
func (m *Mockfetcher) RegisterPeerHashes(arg0 p2p.Peer, arg1 []types.Hash32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterPeerHashes", arg0, arg1)
}

// RegisterPeerHashes indicates an expected call of RegisterPeerHashes.
func (mr *MockfetcherMockRecorder) RegisterPeerHashes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterPeerHashes", reflect.TypeOf((*Mockfetcher)(nil).RegisterPeerHashes), arg0, arg1)
}

// Start mocks base method.
func (m *Mockfetcher) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockfetcherMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*Mockfetcher)(nil).Start))
}

// Stop mocks base method.
func (m *Mockfetcher) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockfetcherMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*Mockfetcher)(nil).Stop))
}

// MockatxHandler is a mock of atxHandler interface.
type MockatxHandler struct {
	ctrl     *gomock.Controller
	recorder *MockatxHandlerMockRecorder
}

// MockatxHandlerMockRecorder is the mock recorder for MockatxHandler.
type MockatxHandlerMockRecorder struct {
	mock *MockatxHandler
}

// NewMockatxHandler creates a new mock instance.
func NewMockatxHandler(ctrl *gomock.Controller) *MockatxHandler {
	mock := &MockatxHandler{ctrl: ctrl}
	mock.recorder = &MockatxHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockatxHandler) EXPECT() *MockatxHandlerMockRecorder {
	return m.recorder
}

// HandleAtxData mocks base method.
func (m *MockatxHandler) HandleAtxData(arg0 context.Context, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleAtxData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleAtxData indicates an expected call of HandleAtxData.
func (mr *MockatxHandlerMockRecorder) HandleAtxData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleAtxData", reflect.TypeOf((*MockatxHandler)(nil).HandleAtxData), arg0, arg1)
}

// MockblockHandler is a mock of blockHandler interface.
type MockblockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockblockHandlerMockRecorder
}

// MockblockHandlerMockRecorder is the mock recorder for MockblockHandler.
type MockblockHandlerMockRecorder struct {
	mock *MockblockHandler
}

// NewMockblockHandler creates a new mock instance.
func NewMockblockHandler(ctrl *gomock.Controller) *MockblockHandler {
	mock := &MockblockHandler{ctrl: ctrl}
	mock.recorder = &MockblockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockblockHandler) EXPECT() *MockblockHandlerMockRecorder {
	return m.recorder
}

// HandleBlockData mocks base method.
func (m *MockblockHandler) HandleBlockData(arg0 context.Context, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleBlockData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleBlockData indicates an expected call of HandleBlockData.
func (mr *MockblockHandlerMockRecorder) HandleBlockData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleBlockData", reflect.TypeOf((*MockblockHandler)(nil).HandleBlockData), arg0, arg1)
}

// MockballotHandler is a mock of ballotHandler interface.
type MockballotHandler struct {
	ctrl     *gomock.Controller
	recorder *MockballotHandlerMockRecorder
}

// MockballotHandlerMockRecorder is the mock recorder for MockballotHandler.
type MockballotHandlerMockRecorder struct {
	mock *MockballotHandler
}

// NewMockballotHandler creates a new mock instance.
func NewMockballotHandler(ctrl *gomock.Controller) *MockballotHandler {
	mock := &MockballotHandler{ctrl: ctrl}
	mock.recorder = &MockballotHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockballotHandler) EXPECT() *MockballotHandlerMockRecorder {
	return m.recorder
}

// HandleBallotData mocks base method.
func (m *MockballotHandler) HandleBallotData(arg0 context.Context, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleBallotData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleBallotData indicates an expected call of HandleBallotData.
func (mr *MockballotHandlerMockRecorder) HandleBallotData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleBallotData", reflect.TypeOf((*MockballotHandler)(nil).HandleBallotData), arg0, arg1)
}

// MockproposalHandler is a mock of proposalHandler interface.
type MockproposalHandler struct {
	ctrl     *gomock.Controller
	recorder *MockproposalHandlerMockRecorder
}

// MockproposalHandlerMockRecorder is the mock recorder for MockproposalHandler.
type MockproposalHandlerMockRecorder struct {
	mock *MockproposalHandler
}

// NewMockproposalHandler creates a new mock instance.
func NewMockproposalHandler(ctrl *gomock.Controller) *MockproposalHandler {
	mock := &MockproposalHandler{ctrl: ctrl}
	mock.recorder = &MockproposalHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockproposalHandler) EXPECT() *MockproposalHandlerMockRecorder {
	return m.recorder
}

// HandleProposalData mocks base method.
func (m *MockproposalHandler) HandleProposalData(arg0 context.Context, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleProposalData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleProposalData indicates an expected call of HandleProposalData.
func (mr *MockproposalHandlerMockRecorder) HandleProposalData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleProposalData", reflect.TypeOf((*MockproposalHandler)(nil).HandleProposalData), arg0, arg1)
}

// MocktxHandler is a mock of txHandler interface.
type MocktxHandler struct {
	ctrl     *gomock.Controller
	recorder *MocktxHandlerMockRecorder
}

// MocktxHandlerMockRecorder is the mock recorder for MocktxHandler.
type MocktxHandlerMockRecorder struct {
	mock *MocktxHandler
}

// NewMocktxHandler creates a new mock instance.
func NewMocktxHandler(ctrl *gomock.Controller) *MocktxHandler {
	mock := &MocktxHandler{ctrl: ctrl}
	mock.recorder = &MocktxHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocktxHandler) EXPECT() *MocktxHandlerMockRecorder {
	return m.recorder
}

// HandleBlockTransaction mocks base method.
func (m *MocktxHandler) HandleBlockTransaction(arg0 context.Context, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleBlockTransaction", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleBlockTransaction indicates an expected call of HandleBlockTransaction.
func (mr *MocktxHandlerMockRecorder) HandleBlockTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleBlockTransaction", reflect.TypeOf((*MocktxHandler)(nil).HandleBlockTransaction), arg0, arg1)
}

// HandleProposalTransaction mocks base method.
func (m *MocktxHandler) HandleProposalTransaction(arg0 context.Context, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleProposalTransaction", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleProposalTransaction indicates an expected call of HandleProposalTransaction.
func (mr *MocktxHandlerMockRecorder) HandleProposalTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleProposalTransaction", reflect.TypeOf((*MocktxHandler)(nil).HandleProposalTransaction), arg0, arg1)
}

// MockpoetHandler is a mock of poetHandler interface.
type MockpoetHandler struct {
	ctrl     *gomock.Controller
	recorder *MockpoetHandlerMockRecorder
}

// MockpoetHandlerMockRecorder is the mock recorder for MockpoetHandler.
type MockpoetHandlerMockRecorder struct {
	mock *MockpoetHandler
}

// NewMockpoetHandler creates a new mock instance.
func NewMockpoetHandler(ctrl *gomock.Controller) *MockpoetHandler {
	mock := &MockpoetHandler{ctrl: ctrl}
	mock.recorder = &MockpoetHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockpoetHandler) EXPECT() *MockpoetHandlerMockRecorder {
	return m.recorder
}

// ValidateAndStoreMsg mocks base method.
func (m *MockpoetHandler) ValidateAndStoreMsg(data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateAndStoreMsg", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateAndStoreMsg indicates an expected call of ValidateAndStoreMsg.
func (mr *MockpoetHandlerMockRecorder) ValidateAndStoreMsg(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateAndStoreMsg", reflect.TypeOf((*MockpoetHandler)(nil).ValidateAndStoreMsg), data)
}

// MockmeshProvider is a mock of meshProvider interface.
type MockmeshProvider struct {
	ctrl     *gomock.Controller
	recorder *MockmeshProviderMockRecorder
}

// MockmeshProviderMockRecorder is the mock recorder for MockmeshProvider.
type MockmeshProviderMockRecorder struct {
	mock *MockmeshProvider
}

// NewMockmeshProvider creates a new mock instance.
func NewMockmeshProvider(ctrl *gomock.Controller) *MockmeshProvider {
	mock := &MockmeshProvider{ctrl: ctrl}
	mock.recorder = &MockmeshProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockmeshProvider) EXPECT() *MockmeshProviderMockRecorder {
	return m.recorder
}

// ProcessedLayer mocks base method.
func (m *MockmeshProvider) ProcessedLayer() types.LayerID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessedLayer")
	ret0, _ := ret[0].(types.LayerID)
	return ret0
}

// ProcessedLayer indicates an expected call of ProcessedLayer.
func (mr *MockmeshProviderMockRecorder) ProcessedLayer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessedLayer", reflect.TypeOf((*MockmeshProvider)(nil).ProcessedLayer))
}

// SetZeroBlockLayer mocks base method.
func (m *MockmeshProvider) SetZeroBlockLayer(arg0 types.LayerID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetZeroBlockLayer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetZeroBlockLayer indicates an expected call of SetZeroBlockLayer.
func (mr *MockmeshProviderMockRecorder) SetZeroBlockLayer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetZeroBlockLayer", reflect.TypeOf((*MockmeshProvider)(nil).SetZeroBlockLayer), arg0)
}

// Mockhost is a mock of host interface.
type Mockhost struct {
	ctrl     *gomock.Controller
	recorder *MockhostMockRecorder
}

// MockhostMockRecorder is the mock recorder for Mockhost.
type MockhostMockRecorder struct {
	mock *Mockhost
}

// NewMockhost creates a new mock instance.
func NewMockhost(ctrl *gomock.Controller) *Mockhost {
	mock := &Mockhost{ctrl: ctrl}
	mock.recorder = &MockhostMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockhost) EXPECT() *MockhostMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *Mockhost) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockhostMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*Mockhost)(nil).Close))
}

// GetPeers mocks base method.
func (m *Mockhost) GetPeers() []p2p.Peer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeers")
	ret0, _ := ret[0].([]p2p.Peer)
	return ret0
}

// GetPeers indicates an expected call of GetPeers.
func (mr *MockhostMockRecorder) GetPeers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeers", reflect.TypeOf((*Mockhost)(nil).GetPeers))
}
