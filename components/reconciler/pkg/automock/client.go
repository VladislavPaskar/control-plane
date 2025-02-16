// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/client_gen.go

// Package automock is a generated GoMock package.
package automock

import (
	context "context"
	io "io"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mothership "github.com/kyma-project/control-plane/components/reconciler/pkg"
)

// MockHttpRequestDoer is a mock of HttpRequestDoer interface.
type MockHttpRequestDoer struct {
	ctrl     *gomock.Controller
	recorder *MockHttpRequestDoerMockRecorder
}

// MockHttpRequestDoerMockRecorder is the mock recorder for MockHttpRequestDoer.
type MockHttpRequestDoerMockRecorder struct {
	mock *MockHttpRequestDoer
}

// NewMockHttpRequestDoer creates a new mock instance.
func NewMockHttpRequestDoer(ctrl *gomock.Controller) *MockHttpRequestDoer {
	mock := &MockHttpRequestDoer{ctrl: ctrl}
	mock.recorder = &MockHttpRequestDoerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpRequestDoer) EXPECT() *MockHttpRequestDoerMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockHttpRequestDoer) Do(req *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockHttpRequestDoerMockRecorder) Do(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockHttpRequestDoer)(nil).Do), req)
}

// MockClientInterface is a mock of ClientInterface interface.
type MockClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClientInterfaceMockRecorder
}

// MockClientInterfaceMockRecorder is the mock recorder for MockClientInterface.
type MockClientInterfaceMockRecorder struct {
	mock *MockClientInterface
}

// NewMockClientInterface creates a new mock instance.
func NewMockClientInterface(ctrl *gomock.Controller) *MockClientInterface {
	mock := &MockClientInterface{ctrl: ctrl}
	mock.recorder = &MockClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientInterface) EXPECT() *MockClientInterfaceMockRecorder {
	return m.recorder
}

// DeleteClustersRuntimeID mocks base method.
func (m *MockClientInterface) DeleteClustersRuntimeID(ctx context.Context, runtimeID string, reqEditors ...mothership.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, runtimeID}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteClustersRuntimeID", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteClustersRuntimeID indicates an expected call of DeleteClustersRuntimeID.
func (mr *MockClientInterfaceMockRecorder) DeleteClustersRuntimeID(ctx, runtimeID interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, runtimeID}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClustersRuntimeID", reflect.TypeOf((*MockClientInterface)(nil).DeleteClustersRuntimeID), varargs...)
}

// GetClustersRuntimeIDConfigConfigVersionStatus mocks base method.
func (m *MockClientInterface) GetClustersRuntimeIDConfigConfigVersionStatus(ctx context.Context, runtimeID, configVersion string, reqEditors ...mothership.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, runtimeID, configVersion}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetClustersRuntimeIDConfigConfigVersionStatus", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClustersRuntimeIDConfigConfigVersionStatus indicates an expected call of GetClustersRuntimeIDConfigConfigVersionStatus.
func (mr *MockClientInterfaceMockRecorder) GetClustersRuntimeIDConfigConfigVersionStatus(ctx, runtimeID, configVersion interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, runtimeID, configVersion}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClustersRuntimeIDConfigConfigVersionStatus", reflect.TypeOf((*MockClientInterface)(nil).GetClustersRuntimeIDConfigConfigVersionStatus), varargs...)
}

// GetClustersRuntimeIDStatus mocks base method.
func (m *MockClientInterface) GetClustersRuntimeIDStatus(ctx context.Context, runtimeID string, reqEditors ...mothership.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, runtimeID}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetClustersRuntimeIDStatus", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClustersRuntimeIDStatus indicates an expected call of GetClustersRuntimeIDStatus.
func (mr *MockClientInterfaceMockRecorder) GetClustersRuntimeIDStatus(ctx, runtimeID interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, runtimeID}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClustersRuntimeIDStatus", reflect.TypeOf((*MockClientInterface)(nil).GetClustersRuntimeIDStatus), varargs...)
}

// GetClustersRuntimeIDStatusChanges mocks base method.
func (m *MockClientInterface) GetClustersRuntimeIDStatusChanges(ctx context.Context, runtimeID string, reqEditors ...mothership.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, runtimeID}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetClustersRuntimeIDStatusChanges", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClustersRuntimeIDStatusChanges indicates an expected call of GetClustersRuntimeIDStatusChanges.
func (mr *MockClientInterfaceMockRecorder) GetClustersRuntimeIDStatusChanges(ctx, runtimeID interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, runtimeID}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClustersRuntimeIDStatusChanges", reflect.TypeOf((*MockClientInterface)(nil).GetClustersRuntimeIDStatusChanges), varargs...)
}

// GetReconciliations mocks base method.
func (m *MockClientInterface) GetReconciliations(ctx context.Context, params *mothership.GetReconciliationsParams, reqEditors ...mothership.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetReconciliations", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReconciliations indicates an expected call of GetReconciliations.
func (mr *MockClientInterfaceMockRecorder) GetReconciliations(ctx, params interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReconciliations", reflect.TypeOf((*MockClientInterface)(nil).GetReconciliations), varargs...)
}

// GetReconciliationsSchedulingIDInfo mocks base method.
func (m *MockClientInterface) GetReconciliationsSchedulingIDInfo(ctx context.Context, schedulingID string, reqEditors ...mothership.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, schedulingID}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetReconciliationsSchedulingIDInfo", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReconciliationsSchedulingIDInfo indicates an expected call of GetReconciliationsSchedulingIDInfo.
func (mr *MockClientInterfaceMockRecorder) GetReconciliationsSchedulingIDInfo(ctx, schedulingID interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, schedulingID}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReconciliationsSchedulingIDInfo", reflect.TypeOf((*MockClientInterface)(nil).GetReconciliationsSchedulingIDInfo), varargs...)
}

// PostClusters mocks base method.
func (m *MockClientInterface) PostClusters(ctx context.Context, body mothership.PostClustersJSONRequestBody, reqEditors ...mothership.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostClusters", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostClusters indicates an expected call of PostClusters.
func (mr *MockClientInterfaceMockRecorder) PostClusters(ctx, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostClusters", reflect.TypeOf((*MockClientInterface)(nil).PostClusters), varargs...)
}

// PostClustersWithBody mocks base method.
func (m *MockClientInterface) PostClustersWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...mothership.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, contentType, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostClustersWithBody", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostClustersWithBody indicates an expected call of PostClustersWithBody.
func (mr *MockClientInterfaceMockRecorder) PostClustersWithBody(ctx, contentType, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, contentType, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostClustersWithBody", reflect.TypeOf((*MockClientInterface)(nil).PostClustersWithBody), varargs...)
}

// PostOperationsSchedulingIDCorrelationIDStop mocks base method.
func (m *MockClientInterface) PostOperationsSchedulingIDCorrelationIDStop(ctx context.Context, schedulingID, correlationID string, body mothership.PostOperationsSchedulingIDCorrelationIDStopJSONRequestBody, reqEditors ...mothership.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, schedulingID, correlationID, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostOperationsSchedulingIDCorrelationIDStop", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostOperationsSchedulingIDCorrelationIDStop indicates an expected call of PostOperationsSchedulingIDCorrelationIDStop.
func (mr *MockClientInterfaceMockRecorder) PostOperationsSchedulingIDCorrelationIDStop(ctx, schedulingID, correlationID, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, schedulingID, correlationID, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostOperationsSchedulingIDCorrelationIDStop", reflect.TypeOf((*MockClientInterface)(nil).PostOperationsSchedulingIDCorrelationIDStop), varargs...)
}

// PostOperationsSchedulingIDCorrelationIDStopWithBody mocks base method.
func (m *MockClientInterface) PostOperationsSchedulingIDCorrelationIDStopWithBody(ctx context.Context, schedulingID, correlationID, contentType string, body io.Reader, reqEditors ...mothership.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, schedulingID, correlationID, contentType, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostOperationsSchedulingIDCorrelationIDStopWithBody", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostOperationsSchedulingIDCorrelationIDStopWithBody indicates an expected call of PostOperationsSchedulingIDCorrelationIDStopWithBody.
func (mr *MockClientInterfaceMockRecorder) PostOperationsSchedulingIDCorrelationIDStopWithBody(ctx, schedulingID, correlationID, contentType, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, schedulingID, correlationID, contentType, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostOperationsSchedulingIDCorrelationIDStopWithBody", reflect.TypeOf((*MockClientInterface)(nil).PostOperationsSchedulingIDCorrelationIDStopWithBody), varargs...)
}

// PutClusters mocks base method.
func (m *MockClientInterface) PutClusters(ctx context.Context, body mothership.PutClustersJSONRequestBody, reqEditors ...mothership.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutClusters", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutClusters indicates an expected call of PutClusters.
func (mr *MockClientInterfaceMockRecorder) PutClusters(ctx, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutClusters", reflect.TypeOf((*MockClientInterface)(nil).PutClusters), varargs...)
}

// PutClustersRuntimeIDStatus mocks base method.
func (m *MockClientInterface) PutClustersRuntimeIDStatus(ctx context.Context, runtimeID string, body mothership.PutClustersRuntimeIDStatusJSONRequestBody, reqEditors ...mothership.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, runtimeID, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutClustersRuntimeIDStatus", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutClustersRuntimeIDStatus indicates an expected call of PutClustersRuntimeIDStatus.
func (mr *MockClientInterfaceMockRecorder) PutClustersRuntimeIDStatus(ctx, runtimeID, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, runtimeID, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutClustersRuntimeIDStatus", reflect.TypeOf((*MockClientInterface)(nil).PutClustersRuntimeIDStatus), varargs...)
}

// PutClustersRuntimeIDStatusWithBody mocks base method.
func (m *MockClientInterface) PutClustersRuntimeIDStatusWithBody(ctx context.Context, runtimeID, contentType string, body io.Reader, reqEditors ...mothership.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, runtimeID, contentType, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutClustersRuntimeIDStatusWithBody", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutClustersRuntimeIDStatusWithBody indicates an expected call of PutClustersRuntimeIDStatusWithBody.
func (mr *MockClientInterfaceMockRecorder) PutClustersRuntimeIDStatusWithBody(ctx, runtimeID, contentType, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, runtimeID, contentType, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutClustersRuntimeIDStatusWithBody", reflect.TypeOf((*MockClientInterface)(nil).PutClustersRuntimeIDStatusWithBody), varargs...)
}

// PutClustersWithBody mocks base method.
func (m *MockClientInterface) PutClustersWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...mothership.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, contentType, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutClustersWithBody", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutClustersWithBody indicates an expected call of PutClustersWithBody.
func (mr *MockClientInterfaceMockRecorder) PutClustersWithBody(ctx, contentType, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, contentType, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutClustersWithBody", reflect.TypeOf((*MockClientInterface)(nil).PutClustersWithBody), varargs...)
}

// MockClientWithResponsesInterface is a mock of ClientWithResponsesInterface interface.
type MockClientWithResponsesInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClientWithResponsesInterfaceMockRecorder
}

// MockClientWithResponsesInterfaceMockRecorder is the mock recorder for MockClientWithResponsesInterface.
type MockClientWithResponsesInterfaceMockRecorder struct {
	mock *MockClientWithResponsesInterface
}

// NewMockClientWithResponsesInterface creates a new mock instance.
func NewMockClientWithResponsesInterface(ctrl *gomock.Controller) *MockClientWithResponsesInterface {
	mock := &MockClientWithResponsesInterface{ctrl: ctrl}
	mock.recorder = &MockClientWithResponsesInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientWithResponsesInterface) EXPECT() *MockClientWithResponsesInterfaceMockRecorder {
	return m.recorder
}

// DeleteClustersRuntimeIDWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) DeleteClustersRuntimeIDWithResponse(ctx context.Context, runtimeID string, reqEditors ...mothership.RequestEditorFn) (*mothership.DeleteClustersRuntimeIDResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, runtimeID}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteClustersRuntimeIDWithResponse", varargs...)
	ret0, _ := ret[0].(*mothership.DeleteClustersRuntimeIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteClustersRuntimeIDWithResponse indicates an expected call of DeleteClustersRuntimeIDWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) DeleteClustersRuntimeIDWithResponse(ctx, runtimeID interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, runtimeID}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClustersRuntimeIDWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).DeleteClustersRuntimeIDWithResponse), varargs...)
}

// GetClustersRuntimeIDConfigConfigVersionStatusWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) GetClustersRuntimeIDConfigConfigVersionStatusWithResponse(ctx context.Context, runtimeID, configVersion string, reqEditors ...mothership.RequestEditorFn) (*mothership.GetClustersRuntimeIDConfigConfigVersionStatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, runtimeID, configVersion}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetClustersRuntimeIDConfigConfigVersionStatusWithResponse", varargs...)
	ret0, _ := ret[0].(*mothership.GetClustersRuntimeIDConfigConfigVersionStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClustersRuntimeIDConfigConfigVersionStatusWithResponse indicates an expected call of GetClustersRuntimeIDConfigConfigVersionStatusWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) GetClustersRuntimeIDConfigConfigVersionStatusWithResponse(ctx, runtimeID, configVersion interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, runtimeID, configVersion}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClustersRuntimeIDConfigConfigVersionStatusWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).GetClustersRuntimeIDConfigConfigVersionStatusWithResponse), varargs...)
}

// GetClustersRuntimeIDStatusChangesWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) GetClustersRuntimeIDStatusChangesWithResponse(ctx context.Context, runtimeID string, reqEditors ...mothership.RequestEditorFn) (*mothership.GetClustersRuntimeIDStatusChangesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, runtimeID}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetClustersRuntimeIDStatusChangesWithResponse", varargs...)
	ret0, _ := ret[0].(*mothership.GetClustersRuntimeIDStatusChangesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClustersRuntimeIDStatusChangesWithResponse indicates an expected call of GetClustersRuntimeIDStatusChangesWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) GetClustersRuntimeIDStatusChangesWithResponse(ctx, runtimeID interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, runtimeID}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClustersRuntimeIDStatusChangesWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).GetClustersRuntimeIDStatusChangesWithResponse), varargs...)
}

// GetClustersRuntimeIDStatusWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) GetClustersRuntimeIDStatusWithResponse(ctx context.Context, runtimeID string, reqEditors ...mothership.RequestEditorFn) (*mothership.GetClustersRuntimeIDStatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, runtimeID}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetClustersRuntimeIDStatusWithResponse", varargs...)
	ret0, _ := ret[0].(*mothership.GetClustersRuntimeIDStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClustersRuntimeIDStatusWithResponse indicates an expected call of GetClustersRuntimeIDStatusWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) GetClustersRuntimeIDStatusWithResponse(ctx, runtimeID interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, runtimeID}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClustersRuntimeIDStatusWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).GetClustersRuntimeIDStatusWithResponse), varargs...)
}

// GetReconciliationsSchedulingIDInfoWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) GetReconciliationsSchedulingIDInfoWithResponse(ctx context.Context, schedulingID string, reqEditors ...mothership.RequestEditorFn) (*mothership.GetReconciliationsSchedulingIDInfoResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, schedulingID}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetReconciliationsSchedulingIDInfoWithResponse", varargs...)
	ret0, _ := ret[0].(*mothership.GetReconciliationsSchedulingIDInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReconciliationsSchedulingIDInfoWithResponse indicates an expected call of GetReconciliationsSchedulingIDInfoWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) GetReconciliationsSchedulingIDInfoWithResponse(ctx, schedulingID interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, schedulingID}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReconciliationsSchedulingIDInfoWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).GetReconciliationsSchedulingIDInfoWithResponse), varargs...)
}

// GetReconciliationsWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) GetReconciliationsWithResponse(ctx context.Context, params *mothership.GetReconciliationsParams, reqEditors ...mothership.RequestEditorFn) (*mothership.GetReconciliationsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetReconciliationsWithResponse", varargs...)
	ret0, _ := ret[0].(*mothership.GetReconciliationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReconciliationsWithResponse indicates an expected call of GetReconciliationsWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) GetReconciliationsWithResponse(ctx, params interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReconciliationsWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).GetReconciliationsWithResponse), varargs...)
}

// PostClustersWithBodyWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) PostClustersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...mothership.RequestEditorFn) (*mothership.PostClustersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, contentType, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostClustersWithBodyWithResponse", varargs...)
	ret0, _ := ret[0].(*mothership.PostClustersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostClustersWithBodyWithResponse indicates an expected call of PostClustersWithBodyWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) PostClustersWithBodyWithResponse(ctx, contentType, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, contentType, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostClustersWithBodyWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).PostClustersWithBodyWithResponse), varargs...)
}

// PostClustersWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) PostClustersWithResponse(ctx context.Context, body mothership.PostClustersJSONRequestBody, reqEditors ...mothership.RequestEditorFn) (*mothership.PostClustersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostClustersWithResponse", varargs...)
	ret0, _ := ret[0].(*mothership.PostClustersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostClustersWithResponse indicates an expected call of PostClustersWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) PostClustersWithResponse(ctx, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostClustersWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).PostClustersWithResponse), varargs...)
}

// PostOperationsSchedulingIDCorrelationIDStopWithBodyWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) PostOperationsSchedulingIDCorrelationIDStopWithBodyWithResponse(ctx context.Context, schedulingID, correlationID, contentType string, body io.Reader, reqEditors ...mothership.RequestEditorFn) (*mothership.PostOperationsSchedulingIDCorrelationIDStopResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, schedulingID, correlationID, contentType, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostOperationsSchedulingIDCorrelationIDStopWithBodyWithResponse", varargs...)
	ret0, _ := ret[0].(*mothership.PostOperationsSchedulingIDCorrelationIDStopResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostOperationsSchedulingIDCorrelationIDStopWithBodyWithResponse indicates an expected call of PostOperationsSchedulingIDCorrelationIDStopWithBodyWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) PostOperationsSchedulingIDCorrelationIDStopWithBodyWithResponse(ctx, schedulingID, correlationID, contentType, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, schedulingID, correlationID, contentType, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostOperationsSchedulingIDCorrelationIDStopWithBodyWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).PostOperationsSchedulingIDCorrelationIDStopWithBodyWithResponse), varargs...)
}

// PostOperationsSchedulingIDCorrelationIDStopWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) PostOperationsSchedulingIDCorrelationIDStopWithResponse(ctx context.Context, schedulingID, correlationID string, body mothership.PostOperationsSchedulingIDCorrelationIDStopJSONRequestBody, reqEditors ...mothership.RequestEditorFn) (*mothership.PostOperationsSchedulingIDCorrelationIDStopResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, schedulingID, correlationID, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostOperationsSchedulingIDCorrelationIDStopWithResponse", varargs...)
	ret0, _ := ret[0].(*mothership.PostOperationsSchedulingIDCorrelationIDStopResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostOperationsSchedulingIDCorrelationIDStopWithResponse indicates an expected call of PostOperationsSchedulingIDCorrelationIDStopWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) PostOperationsSchedulingIDCorrelationIDStopWithResponse(ctx, schedulingID, correlationID, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, schedulingID, correlationID, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostOperationsSchedulingIDCorrelationIDStopWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).PostOperationsSchedulingIDCorrelationIDStopWithResponse), varargs...)
}

// PutClustersRuntimeIDStatusWithBodyWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) PutClustersRuntimeIDStatusWithBodyWithResponse(ctx context.Context, runtimeID, contentType string, body io.Reader, reqEditors ...mothership.RequestEditorFn) (*mothership.PutClustersRuntimeIDStatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, runtimeID, contentType, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutClustersRuntimeIDStatusWithBodyWithResponse", varargs...)
	ret0, _ := ret[0].(*mothership.PutClustersRuntimeIDStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutClustersRuntimeIDStatusWithBodyWithResponse indicates an expected call of PutClustersRuntimeIDStatusWithBodyWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) PutClustersRuntimeIDStatusWithBodyWithResponse(ctx, runtimeID, contentType, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, runtimeID, contentType, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutClustersRuntimeIDStatusWithBodyWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).PutClustersRuntimeIDStatusWithBodyWithResponse), varargs...)
}

// PutClustersRuntimeIDStatusWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) PutClustersRuntimeIDStatusWithResponse(ctx context.Context, runtimeID string, body mothership.PutClustersRuntimeIDStatusJSONRequestBody, reqEditors ...mothership.RequestEditorFn) (*mothership.PutClustersRuntimeIDStatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, runtimeID, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutClustersRuntimeIDStatusWithResponse", varargs...)
	ret0, _ := ret[0].(*mothership.PutClustersRuntimeIDStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutClustersRuntimeIDStatusWithResponse indicates an expected call of PutClustersRuntimeIDStatusWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) PutClustersRuntimeIDStatusWithResponse(ctx, runtimeID, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, runtimeID, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutClustersRuntimeIDStatusWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).PutClustersRuntimeIDStatusWithResponse), varargs...)
}

// PutClustersWithBodyWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) PutClustersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...mothership.RequestEditorFn) (*mothership.PutClustersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, contentType, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutClustersWithBodyWithResponse", varargs...)
	ret0, _ := ret[0].(*mothership.PutClustersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutClustersWithBodyWithResponse indicates an expected call of PutClustersWithBodyWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) PutClustersWithBodyWithResponse(ctx, contentType, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, contentType, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutClustersWithBodyWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).PutClustersWithBodyWithResponse), varargs...)
}

// PutClustersWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) PutClustersWithResponse(ctx context.Context, body mothership.PutClustersJSONRequestBody, reqEditors ...mothership.RequestEditorFn) (*mothership.PutClustersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutClustersWithResponse", varargs...)
	ret0, _ := ret[0].(*mothership.PutClustersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutClustersWithResponse indicates an expected call of PutClustersWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) PutClustersWithResponse(ctx, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutClustersWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).PutClustersWithResponse), varargs...)
}
