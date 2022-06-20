package restore

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
	env.OnActivity(PreRestoreCmdActivity, mock.Anything, mock.Anything).Return(result, nil)
	env.OnActivity(PreRestoreActivity, mock.Anything, mock.Anything).Return(result, nil)
	env.OnActivity(RestoreCmdActivity, mock.Anything, mock.Anything).Return(result, nil)
	env.OnActivity(RestoreActivity, mock.Anything, mock.Anything).Return(result, nil)
	env.OnActivity(PostAppRestoreCmdActivity, mock.Anything, mock.Anything).Return(result, nil)
	env.OnActivity(PostRestoreActivity, mock.Anything, mock.Anything).Return(result, nil)

	env.ExecuteWorkflow(RestoreWorkflow, config, workflowStatus)

	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	require.NoError(t, env.GetWorkflowResult(&workflowStatusResult))
	require.Equal(t, workflowStatusResult, workflowStatusResult)
}
