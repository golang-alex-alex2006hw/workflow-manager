// Automatically generated by MockGen. DO NOT EDIT!
// Source: interface.go

package client

import (
	context "context"
	models "github.com/Clever/workflow-manager/gen-go/models"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

// Recorder for MockClient (not exported)
type _MockClientRecorder struct {
	mock *MockClient
}

func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

func (_m *MockClient) EXPECT() *_MockClientRecorder {
	return _m.recorder
}

func (_m *MockClient) HealthCheck(ctx context.Context) error {
	ret := _m.ctrl.Call(_m, "HealthCheck", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) HealthCheck(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HealthCheck", arg0)
}

func (_m *MockClient) GetJobsForWorkflowDefinition(ctx context.Context, i *models.GetJobsForWorkflowDefinitionInput) ([]models.Job, error) {
	ret := _m.ctrl.Call(_m, "GetJobsForWorkflowDefinition", ctx, i)
	ret0, _ := ret[0].([]models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) GetJobsForWorkflowDefinition(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetJobsForWorkflowDefinition", arg0, arg1)
}

func (_m *MockClient) StartJobForWorkflowDefinition(ctx context.Context, i *models.JobInput) (*models.Job, error) {
	ret := _m.ctrl.Call(_m, "StartJobForWorkflowDefinition", ctx, i)
	ret0, _ := ret[0].(*models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) StartJobForWorkflowDefinition(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartJobForWorkflowDefinition", arg0, arg1)
}

func (_m *MockClient) CancelJob(ctx context.Context, i *models.CancelJobInput) error {
	ret := _m.ctrl.Call(_m, "CancelJob", ctx, i)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) CancelJob(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CancelJob", arg0, arg1)
}

func (_m *MockClient) GetJob(ctx context.Context, jobId string) (*models.Job, error) {
	ret := _m.ctrl.Call(_m, "GetJob", ctx, jobId)
	ret0, _ := ret[0].(*models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) GetJob(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetJob", arg0, arg1)
}

func (_m *MockClient) PostStateResource(ctx context.Context, i *models.NewStateResource) (*models.StateResource, error) {
	ret := _m.ctrl.Call(_m, "PostStateResource", ctx, i)
	ret0, _ := ret[0].(*models.StateResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) PostStateResource(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PostStateResource", arg0, arg1)
}

func (_m *MockClient) DeleteStateResource(ctx context.Context, i *models.DeleteStateResourceInput) error {
	ret := _m.ctrl.Call(_m, "DeleteStateResource", ctx, i)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) DeleteStateResource(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteStateResource", arg0, arg1)
}

func (_m *MockClient) GetStateResource(ctx context.Context, i *models.GetStateResourceInput) (*models.StateResource, error) {
	ret := _m.ctrl.Call(_m, "GetStateResource", ctx, i)
	ret0, _ := ret[0].(*models.StateResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) GetStateResource(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetStateResource", arg0, arg1)
}

func (_m *MockClient) PutStateResource(ctx context.Context, i *models.PutStateResourceInput) (*models.StateResource, error) {
	ret := _m.ctrl.Call(_m, "PutStateResource", ctx, i)
	ret0, _ := ret[0].(*models.StateResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) PutStateResource(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutStateResource", arg0, arg1)
}

func (_m *MockClient) GetWorkflowDefinitions(ctx context.Context) ([]models.WorkflowDefinition, error) {
	ret := _m.ctrl.Call(_m, "GetWorkflowDefinitions", ctx)
	ret0, _ := ret[0].([]models.WorkflowDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) GetWorkflowDefinitions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetWorkflowDefinitions", arg0)
}

func (_m *MockClient) NewWorkflowDefinition(ctx context.Context, i *models.NewWorkflowDefinitionRequest) (*models.WorkflowDefinition, error) {
	ret := _m.ctrl.Call(_m, "NewWorkflowDefinition", ctx, i)
	ret0, _ := ret[0].(*models.WorkflowDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) NewWorkflowDefinition(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewWorkflowDefinition", arg0, arg1)
}

func (_m *MockClient) GetWorkflowDefinitionVersionsByName(ctx context.Context, i *models.GetWorkflowDefinitionVersionsByNameInput) ([]models.WorkflowDefinition, error) {
	ret := _m.ctrl.Call(_m, "GetWorkflowDefinitionVersionsByName", ctx, i)
	ret0, _ := ret[0].([]models.WorkflowDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) GetWorkflowDefinitionVersionsByName(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetWorkflowDefinitionVersionsByName", arg0, arg1)
}

func (_m *MockClient) UpdateWorkflowDefinition(ctx context.Context, i *models.UpdateWorkflowDefinitionInput) (*models.WorkflowDefinition, error) {
	ret := _m.ctrl.Call(_m, "UpdateWorkflowDefinition", ctx, i)
	ret0, _ := ret[0].(*models.WorkflowDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) UpdateWorkflowDefinition(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateWorkflowDefinition", arg0, arg1)
}

func (_m *MockClient) GetWorkflowDefinitionByNameAndVersion(ctx context.Context, i *models.GetWorkflowDefinitionByNameAndVersionInput) (*models.WorkflowDefinition, error) {
	ret := _m.ctrl.Call(_m, "GetWorkflowDefinitionByNameAndVersion", ctx, i)
	ret0, _ := ret[0].(*models.WorkflowDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) GetWorkflowDefinitionByNameAndVersion(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetWorkflowDefinitionByNameAndVersion", arg0, arg1)
}
