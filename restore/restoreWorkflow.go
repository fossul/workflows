package restore

import (
	"time"

	"github.com/fossul/fossul/src/engine/util"
	"github.com/fossul/workflows/fossul"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"

	// TODO(cretz): Remove when tagged
	_ "go.temporal.io/sdk/contrib/tools/workflowcheck/determinism"
)

func RestoreWorkflow(ctx workflow.Context, config util.Config, workflowStatus *util.Workflow) (util.WorkflowStatusResult, error) {
	retryPolicy := &temporal.RetryPolicy{
		MaximumAttempts:        3,
		InitialInterval:        time.Second,
		MaximumInterval:        time.Second * 10,
		BackoffCoefficient:     2,
		NonRetryableErrorTypes: []string{"bad-bug"},
	}

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy:         retryPolicy,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("Fossul restore workflow started")

	var result util.Result
	var messages []util.Message
	var workflowStatusResult util.WorkflowStatusResult

	// PreRestoreCmdActivity
	if config.PreAppRestoreCmd != "" {
		err := workflow.ExecuteActivity(ctx, PreRestoreCmdActivity, config, workflowStatus).Get(ctx, &result)
		step := fossul.StepInit(workflowStatus)
		messages = fossul.AppendMessages(result.Messages, messages)
		util.SetStepComplete(workflowStatus, step)

		if err != nil {
			workflowStatusResult = fossul.ErrorHandling(messages, result, workflowStatus, workflowStatusResult)
			return workflowStatusResult, err
		}
	}

	// PreRestoreActivity
	if config.AppPlugin != "" {
		err := workflow.ExecuteActivity(ctx, PreRestoreActivity, config, workflowStatus).Get(ctx, &result)
		step := fossul.StepInit(workflowStatus)
		messages = fossul.AppendMessages(result.Messages, messages)
		util.SetStepComplete(workflowStatus, step)

		if err != nil {
			workflowStatusResult = fossul.ErrorHandling(messages, result, workflowStatus, workflowStatusResult)
			return workflowStatusResult, err
		}
	}

	// RestoreCmdActivity
	if config.RestoreCmd != "" {
		err := workflow.ExecuteActivity(ctx, RestoreCmdActivity, config, workflowStatus).Get(ctx, &result)
		step := fossul.StepInit(workflowStatus)
		messages = fossul.AppendMessages(result.Messages, messages)
		util.SetStepComplete(workflowStatus, step)

		if err != nil {
			workflowStatusResult = fossul.ErrorHandling(messages, result, workflowStatus, workflowStatusResult)
			return workflowStatusResult, err
		}
	}

	// RestoreActivity
	if config.StoragePlugin != "" {
		err := workflow.ExecuteActivity(ctx, RestoreActivity, config, workflowStatus).Get(ctx, &result)
		step := fossul.StepInit(workflowStatus)
		messages = fossul.AppendMessages(result.Messages, messages)
		util.SetStepComplete(workflowStatus, step)

		if err != nil {
			workflowStatusResult = fossul.ErrorHandling(messages, result, workflowStatus, workflowStatusResult)
			return workflowStatusResult, err
		}
	}

	// PostAppRestoreCmdActivity
	if config.PostAppRestoreCmd != "" {
		err := workflow.ExecuteActivity(ctx, PostAppRestoreCmdActivity, config, workflowStatus).Get(ctx, &result)
		step := fossul.StepInit(workflowStatus)
		messages = fossul.AppendMessages(result.Messages, messages)
		util.SetStepComplete(workflowStatus, step)

		if err != nil {
			workflowStatusResult = fossul.ErrorHandling(messages, result, workflowStatus, workflowStatusResult)
			return workflowStatusResult, err
		}
	}

	// PostRestoreActivity
	if config.AppPlugin != "" {
		err := workflow.ExecuteActivity(ctx, PostRestoreActivity, config, workflowStatus).Get(ctx, &result)
		step := fossul.StepInit(workflowStatus)
		messages = fossul.AppendMessages(result.Messages, messages)
		util.SetStepComplete(workflowStatus, step)

		if err != nil {
			workflowStatusResult = fossul.ErrorHandling(messages, result, workflowStatus, workflowStatusResult)
			return workflowStatusResult, err
		}
	}

	util.SetWorkflowStatusEnd(workflowStatus)
	workflowStatusResult.Workflow = *workflowStatus
	result.Messages = messages
	workflowStatusResult.Result = result

	logger.Info("Restore workflow completed.", "result", workflowStatusResult)

	return workflowStatusResult, nil
}
