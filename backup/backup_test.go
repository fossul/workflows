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
	var result util.Result

	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()
	config.AppPlugin = "test"

	// Mock activity implementation
	env.OnActivity(BackupCreateCmdActivity, mock.Anything, mock.Anything).Return(result, nil)

	env.ExecuteWorkflow(Workflow, config)

	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	require.NoError(t, env.GetWorkflowResult(&result))
	require.Equal(t, result, result)
}

func Test_Activity(t *testing.T) {
	var config util.Config
	var result util.Result

	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	env.RegisterActivity(BackupCreateCmdActivity)

	val, err := env.ExecuteActivity(BackupCreateCmdActivity, config)
	require.NoError(t, err)

	require.NoError(t, val.Get(&result))
}
