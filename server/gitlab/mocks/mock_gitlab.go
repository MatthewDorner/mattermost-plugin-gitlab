// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-plugin-gitlab/server/gitlab (interfaces: Gitlab)

// Package mock_gitlab is a generated GoMock package.
package mock_gitlab

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	logger "github.com/mattermost/mattermost-plugin-api/experimental/bot/logger"
	gitlab "github.com/mattermost/mattermost-plugin-gitlab/server/gitlab"
	gitlab0 "github.com/xanzy/go-gitlab"
	oauth2 "golang.org/x/oauth2"
)

// MockGitlab is a mock of Gitlab interface.
type MockGitlab struct {
	ctrl     *gomock.Controller
	recorder *MockGitlabMockRecorder
}

// MockGitlabMockRecorder is the mock recorder for MockGitlab.
type MockGitlabMockRecorder struct {
	mock *MockGitlab
}

// NewMockGitlab creates a new mock instance.
func NewMockGitlab(ctrl *gomock.Controller) *MockGitlab {
	mock := &MockGitlab{ctrl: ctrl}
	mock.recorder = &MockGitlabMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGitlab) EXPECT() *MockGitlabMockRecorder {
	return m.recorder
}

// GetCurrentUser mocks base method.
func (m *MockGitlab) GetCurrentUser(arg0 context.Context, arg1 string, arg2 oauth2.Token) (*gitlab.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(*gitlab.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentUser indicates an expected call of GetCurrentUser.
func (mr *MockGitlabMockRecorder) GetCurrentUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentUser", reflect.TypeOf((*MockGitlab)(nil).GetCurrentUser), arg0, arg1, arg2)
}

// GetGroupHooks mocks base method.
func (m *MockGitlab) GetGroupHooks(arg0 context.Context, arg1 *gitlab.UserInfo, arg2 string) ([]*gitlab.WebhookInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupHooks", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*gitlab.WebhookInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupHooks indicates an expected call of GetGroupHooks.
func (mr *MockGitlabMockRecorder) GetGroupHooks(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupHooks", reflect.TypeOf((*MockGitlab)(nil).GetGroupHooks), arg0, arg1, arg2)
}

// GetProject mocks base method.
func (m *MockGitlab) GetProject(arg0 context.Context, arg1 *gitlab.UserInfo, arg2, arg3 string) (*gitlab0.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProject", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*gitlab0.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject.
func (mr *MockGitlabMockRecorder) GetProject(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockGitlab)(nil).GetProject), arg0, arg1, arg2, arg3)
}

// GetProjectHooks mocks base method.
func (m *MockGitlab) GetProjectHooks(arg0 context.Context, arg1 *gitlab.UserInfo, arg2, arg3 string) ([]*gitlab.WebhookInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectHooks", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*gitlab.WebhookInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectHooks indicates an expected call of GetProjectHooks.
func (mr *MockGitlabMockRecorder) GetProjectHooks(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectHooks", reflect.TypeOf((*MockGitlab)(nil).GetProjectHooks), arg0, arg1, arg2, arg3)
}

// GetReviews mocks base method.
func (m *MockGitlab) GetReviews(arg0 context.Context, arg1 *gitlab.UserInfo) ([]*gitlab.MergeRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReviews", arg0, arg1)
	ret0, _ := ret[0].([]*gitlab.MergeRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReviews indicates an expected call of GetReviews.
func (mr *MockGitlabMockRecorder) GetReviews(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReviews", reflect.TypeOf((*MockGitlab)(nil).GetReviews), arg0, arg1)
}

// GetUnreads mocks base method.
func (m *MockGitlab) GetUnreads(arg0 context.Context, arg1 *gitlab.UserInfo) ([]*gitlab0.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnreads", arg0, arg1)
	ret0, _ := ret[0].([]*gitlab0.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnreads indicates an expected call of GetUnreads.
func (mr *MockGitlabMockRecorder) GetUnreads(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnreads", reflect.TypeOf((*MockGitlab)(nil).GetUnreads), arg0, arg1)
}

// GetUserDetails mocks base method.
func (m *MockGitlab) GetUserDetails(arg0 context.Context, arg1 *gitlab.UserInfo) (*gitlab0.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDetails", arg0, arg1)
	ret0, _ := ret[0].(*gitlab0.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDetails indicates an expected call of GetUserDetails.
func (mr *MockGitlabMockRecorder) GetUserDetails(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDetails", reflect.TypeOf((*MockGitlab)(nil).GetUserDetails), arg0, arg1)
}

// GetYourAssignments mocks base method.
func (m *MockGitlab) GetYourAssignments(arg0 context.Context, arg1 *gitlab.UserInfo) ([]*gitlab.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetYourAssignments", arg0, arg1)
	ret0, _ := ret[0].([]*gitlab.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetYourAssignments indicates an expected call of GetYourAssignments.
func (mr *MockGitlabMockRecorder) GetYourAssignments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetYourAssignments", reflect.TypeOf((*MockGitlab)(nil).GetYourAssignments), arg0, arg1)
}

// GetYourPrDetails mocks base method.
func (m *MockGitlab) GetYourPrDetails(arg0 context.Context, arg1 logger.Logger, arg2 *gitlab.UserInfo, arg3 []*gitlab.PRDetails) ([]*gitlab.PRDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetYourPrDetails", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*gitlab.PRDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetYourPrDetails indicates an expected call of GetYourPrDetails.
func (mr *MockGitlabMockRecorder) GetYourPrDetails(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetYourPrDetails", reflect.TypeOf((*MockGitlab)(nil).GetYourPrDetails), arg0, arg1, arg2, arg3)
}

// GetYourPrs mocks base method.
func (m *MockGitlab) GetYourPrs(arg0 context.Context, arg1 *gitlab.UserInfo) ([]*gitlab.MergeRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetYourPrs", arg0, arg1)
	ret0, _ := ret[0].([]*gitlab.MergeRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetYourPrs indicates an expected call of GetYourPrs.
func (mr *MockGitlabMockRecorder) GetYourPrs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetYourPrs", reflect.TypeOf((*MockGitlab)(nil).GetYourPrs), arg0, arg1)
}

// NewGroupHook mocks base method.
func (m *MockGitlab) NewGroupHook(arg0 context.Context, arg1 *gitlab.UserInfo, arg2 string, arg3 *gitlab.AddWebhookOptions) (*gitlab.WebhookInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewGroupHook", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*gitlab.WebhookInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewGroupHook indicates an expected call of NewGroupHook.
func (mr *MockGitlabMockRecorder) NewGroupHook(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewGroupHook", reflect.TypeOf((*MockGitlab)(nil).NewGroupHook), arg0, arg1, arg2, arg3)
}

// NewProjectHook mocks base method.
func (m *MockGitlab) NewProjectHook(arg0 context.Context, arg1 *gitlab.UserInfo, arg2 interface{}, arg3 *gitlab.AddWebhookOptions) (*gitlab.WebhookInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewProjectHook", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*gitlab.WebhookInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewProjectHook indicates an expected call of NewProjectHook.
func (mr *MockGitlabMockRecorder) NewProjectHook(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewProjectHook", reflect.TypeOf((*MockGitlab)(nil).NewProjectHook), arg0, arg1, arg2, arg3)
}

// ResolveNamespaceAndProject mocks base method.
func (m *MockGitlab) ResolveNamespaceAndProject(arg0 context.Context, arg1 *gitlab.UserInfo, arg2 string, arg3 bool) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveNamespaceAndProject", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ResolveNamespaceAndProject indicates an expected call of ResolveNamespaceAndProject.
func (mr *MockGitlabMockRecorder) ResolveNamespaceAndProject(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveNamespaceAndProject", reflect.TypeOf((*MockGitlab)(nil).ResolveNamespaceAndProject), arg0, arg1, arg2, arg3)
}

// TriggerProjectPipeline mocks base method.
func (m *MockGitlab) TriggerProjectPipeline(arg0 *gitlab.UserInfo, arg1, arg2 string) (*gitlab.PipelineInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TriggerProjectPipeline", arg0, arg1, arg2)
	ret0, _ := ret[0].(*gitlab.PipelineInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TriggerProjectPipeline indicates an expected call of TriggerProjectPipeline.
func (mr *MockGitlabMockRecorder) TriggerProjectPipeline(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerProjectPipeline", reflect.TypeOf((*MockGitlab)(nil).TriggerProjectPipeline), arg0, arg1, arg2)
}
