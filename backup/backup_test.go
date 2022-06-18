package backup

import (
	"testing"

	"github.com/fossul/fossul/src/engine/util"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_Workflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()
	var config util.Config
	config.AppPlugin = "test"

	// Mock activity implementation
	env.OnActivity(BackupCreateCmdActivity, mock.Anything, mock.Anything).Return("backup", nil)

	env.ExecuteWorkflow(Workflow, config)

	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
	var result string
	require.NoError(t, env.GetWorkflowResult(&result))
	require.Equal(t, "Backup workflow completed successfully", result)
}

func Test_Activity(t *testing.T) {
	var config util.Config
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	env.RegisterActivity(BackupCreateCmdActivity)

	val, err := env.ExecuteActivity(BackupCreateCmdActivity, config)
	require.NoError(t, err)

	var res string
	require.NoError(t, val.Get(&res))
}
