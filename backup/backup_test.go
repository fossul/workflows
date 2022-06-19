package backup

import (
	"testing"

	"github.com/fossul/fossul/src/engine/util"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_Workflow(t *testing.T) {
	var config util.Config
	workflowStatus := &util.Workflow{}
	var workflowStatusResult util.WorkflowStatusResult

	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()
	config.AppPlugin = "test"

	// Mock activity implementation
	env.OnActivity(PreAppQuiesceCmdActivity, mock.Anything, mock.Anything).Return(result, nil)
	env.OnActivity(AppQuiesceActivity, mock.Anything, mock.Anything).Return(result, nil)
	env.OnActivity(PostAppQuiesceCmdActivity, mock.Anything, mock.Anything).Return(result, nil)
	env.OnActivity(BackupCreateCmdActivity, mock.Anything, mock.Anything).Return(result, nil)
	env.OnActivity(BackupCreateActivity, mock.Anything, mock.Anything).Return(result, nil)
	env.OnActivity(PreAppUnQuiesceCmdActivity, mock.Anything, mock.Anything).Return(result, nil)
	env.OnActivity(AppUnQuiesceActivity, mock.Anything, mock.Anything).Return(result, nil)
	env.OnActivity(PostAppUnQuiesceCmdActivity, mock.Anything, mock.Anything).Return(result, nil)

	env.ExecuteWorkflow(Workflow, config, workflowStatus)

	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	require.NoError(t, env.GetWorkflowResult(&workflowStatusResult))
	require.Equal(t, workflowStatusResult, workflowStatusResult)
}

func Test_PreAppQuiesceCmdActivity_Activity(t *testing.T) {
	var config util.Config
	var result util.Result

	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	env.RegisterActivity(PreAppQuiesceCmdActivity)

	val, err := env.ExecuteActivity(PreAppQuiesceCmdActivity, config)
	require.NoError(t, err)

	require.NoError(t, val.Get(&result))
}

func Test_AppQuiesceActivity_Activity(t *testing.T) {
	var config util.Config
	var result util.Result

	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	env.RegisterActivity(AppQuiesceActivity)

	val, err := env.ExecuteActivity(AppQuiesceActivity, config)
	require.NoError(t, err)

	require.NoError(t, val.Get(&result))
}

func Test_PostAppQuiesceCmdActivity_Activity(t *testing.T) {
	var config util.Config
	var result util.Result

	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	env.RegisterActivity(PostAppQuiesceCmdActivity)

	val, err := env.ExecuteActivity(PostAppQuiesceCmdActivity, config)
	require.NoError(t, err)

	require.NoError(t, val.Get(&result))
}

func Test_BackupCreateCmdActivity_Activity(t *testing.T) {
	var config util.Config
	var result util.Result

	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	env.RegisterActivity(BackupCreateCmdActivity)

	val, err := env.ExecuteActivity(BackupCreateCmdActivity, config)
	require.NoError(t, err)

	require.NoError(t, val.Get(&result))
}

func Test_BackupCreateActivity_Activity(t *testing.T) {
	var config util.Config
	var result util.Result

	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	env.RegisterActivity(BackupCreateActivity)

	val, err := env.ExecuteActivity(BackupCreateActivity, config)
	require.NoError(t, err)

	require.NoError(t, val.Get(&result))
}

func Test_PreAppUnQuiesceCmdActivity_Activity(t *testing.T) {
	var config util.Config
	var result util.Result

	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	env.RegisterActivity(PreAppUnQuiesceCmdActivity)

	val, err := env.ExecuteActivity(PreAppUnQuiesceCmdActivity, config)
	require.NoError(t, err)

	require.NoError(t, val.Get(&result))
}

func Test_AppUnQuiesceActivity_Activity(t *testing.T) {
	var config util.Config
	var result util.Result

	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	env.RegisterActivity(AppUnQuiesceActivity)

	val, err := env.ExecuteActivity(AppUnQuiesceActivity, config)
	require.NoError(t, err)

	require.NoError(t, val.Get(&result))
}

func Test_PostAppUnQuiesceCmdActivity_Activity(t *testing.T) {
	var config util.Config
	var result util.Result

	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	env.RegisterActivity(PostAppUnQuiesceCmdActivity)

	val, err := env.ExecuteActivity(PostAppUnQuiesceCmdActivity, config)
	require.NoError(t, err)

	require.NoError(t, val.Get(&result))
}
