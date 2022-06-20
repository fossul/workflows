package backup

import (
	"time"

	"github.com/fossul/fossul/src/engine/util"
	"github.com/fossul/workflows/fossul"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"

	// TODO(cretz): Remove when tagged
	_ "go.temporal.io/sdk/contrib/tools/workflowcheck/determinism"
)

func Workflow(ctx workflow.Context, config util.Config, workflowStatus *util.Workflow) (util.WorkflowStatusResult, error) {
	retryPolicy := &temporal.RetryPolicy{
		MaximumAttempts:        1,
		InitialInterval:        time.Second * 5,
		MaximumInterval:        time.Second * 10,
		BackoffCoefficient:     1,
		NonRetryableErrorTypes: []string{"bad-bug"},
	}

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy:         retryPolicy,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("Fossul backup workflow started")

	var result util.Result
	var messages []util.Message
	var workflowStatusResult util.WorkflowStatusResult

	// PreAppQuiesceCmdActivity
	if config.PreAppQuiesceCmd != "" {
		err := workflow.ExecuteActivity(ctx, PreAppQuiesceCmdActivity, config, workflowStatus).Get(ctx, &result)
		step := fossul.StepInit(workflowStatus)
		messages = fossul.AppendMessages(result.Messages, messages)
		util.SetStepComplete(workflowStatus, step)

		if err != nil {
			workflowStatusResult = fossul.ErrorHandling(messages, result, workflowStatus, workflowStatusResult)
			return workflowStatusResult, err
		}
	}

	// AppQuiesceActivity
	if config.AppPlugin != "" {
		err := workflow.ExecuteActivity(ctx, AppQuiesceActivity, config, workflowStatus).Get(ctx, &result)
		step := fossul.StepInit(workflowStatus)
		messages = fossul.AppendMessages(result.Messages, messages)
		util.SetStepComplete(workflowStatus, step)

		if err != nil {
			workflowStatusResult = fossul.ErrorHandling(messages, result, workflowStatus, workflowStatusResult)
			return workflowStatusResult, err
		}
	}

	// PostAppQuiesceCmdActivity
	if config.PostAppQuiesceCmd != "" {
		err := workflow.ExecuteActivity(ctx, PostAppQuiesceCmdActivity, config, workflowStatus).Get(ctx, &result)
		step := fossul.StepInit(workflowStatus)
		messages = fossul.AppendMessages(result.Messages, messages)
		util.SetStepComplete(workflowStatus, step)

		if err != nil {
			workflowStatusResult = fossul.ErrorHandling(messages, result, workflowStatus, workflowStatusResult)
			return workflowStatusResult, err
		}
	}

	// BackupCreateCmdActivity
	if config.BackupCreateCmd != "" {
		err := workflow.ExecuteActivity(ctx, BackupCreateCmdActivity, config, workflowStatus).Get(ctx, &result)
		step := fossul.StepInit(workflowStatus)
		messages = fossul.AppendMessages(result.Messages, messages)
		util.SetStepComplete(workflowStatus, step)

		if err != nil {
			workflowStatusResult = fossul.ErrorHandling(messages, result, workflowStatus, workflowStatusResult)
			return workflowStatusResult, err
		}
	}

	// BackupCreateActivity
	if config.StoragePlugin != "" {
		err := workflow.ExecuteActivity(ctx, BackupCreateActivity, config, workflowStatus).Get(ctx, &result)
		step := fossul.StepInit(workflowStatus)
		messages = fossul.AppendMessages(result.Messages, messages)
		util.SetStepComplete(workflowStatus, step)

		if err != nil {
			workflowStatusResult = fossul.ErrorHandling(messages, result, workflowStatus, workflowStatusResult)
			return workflowStatusResult, err
		}
	}

	// PreAppUnQuiesceCmdActivity
	if config.PreAppUnquiesceCmd != "" {
		err := workflow.ExecuteActivity(ctx, PreAppUnQuiesceCmdActivity, config, workflowStatus).Get(ctx, &result)
		step := fossul.StepInit(workflowStatus)
		messages = fossul.AppendMessages(result.Messages, messages)
		util.SetStepComplete(workflowStatus, step)

		if err != nil {
			workflowStatusResult = fossul.ErrorHandling(messages, result, workflowStatus, workflowStatusResult)
			return workflowStatusResult, err
		}
	}

	// AppUnQuiesceActivity
	if config.AppPlugin != "" {
		err := workflow.ExecuteActivity(ctx, AppUnQuiesceActivity, config, workflowStatus).Get(ctx, &result)
		step := fossul.StepInit(workflowStatus)
		messages = fossul.AppendMessages(result.Messages, messages)
		util.SetStepComplete(workflowStatus, step)

		if err != nil {
			workflowStatusResult = fossul.ErrorHandling(messages, result, workflowStatus, workflowStatusResult)
			return workflowStatusResult, err
		}
	}

	// PostAppUnQuiesceCmdActivity
	if config.PreAppUnquiesceCmd != "" {
		err := workflow.ExecuteActivity(ctx, PostAppUnQuiesceCmdActivity, config, workflowStatus).Get(ctx, &result)
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

	logger.Info("Backup workflow completed.", "result", workflowStatusResult)

	return workflowStatusResult, nil
}
