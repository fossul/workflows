package backup

import (
	"testing"

	"github.com/fossul/fossul/src/engine/util"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_Workflow(t *testing.T) {
	var result util.Result
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
