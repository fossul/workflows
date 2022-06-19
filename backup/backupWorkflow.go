package backup

import (
	"time"

	"github.com/fossul/fossul/src/engine/util"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"

	// TODO(cretz): Remove when tagged
	_ "go.temporal.io/sdk/contrib/tools/workflowcheck/determinism"
)

func Workflow(ctx workflow.Context, config util.Config, workflowStatus *util.Workflow) (util.WorkflowStatusResult, error) {
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
	logger.Info("Fossul backup workflow started")

	var result util.Result
	var messages []util.Message
	var workflowStatusResult util.WorkflowStatusResult

	step := StepInit(workflowStatus)
	err := workflow.ExecuteActivity(ctx, PreAppQuiesceCmdActivity, config, workflowStatus).Get(ctx, &result)
	messages = util.PrependMessages(messages, result.Messages)
	result.Messages = messages

	util.SetStepComplete(workflowStatus, step)
	workflowStatusResult.Workflow = *workflowStatus
	workflowStatusResult.Result = result
	if err != nil {
		util.SetWorkflowStatusError(workflowStatus)
		workflowStatusResult.Workflow = *workflowStatus
		return workflowStatusResult, err
	}

	step = StepInit(workflowStatus)
	err = workflow.ExecuteActivity(ctx, PostAppQuiesceCmdActivity, config, workflowStatus).Get(ctx, &result)
	messages = util.PrependMessages(messages, result.Messages)
	result.Messages = messages

	util.SetStepComplete(workflowStatus, step)
	workflowStatusResult.Workflow = *workflowStatus
	workflowStatusResult.Result = result
	if err != nil {
		util.SetWorkflowStatusError(workflowStatus)
		workflowStatusResult.Workflow = *workflowStatus
		return workflowStatusResult, err
	}

	step = StepInit(workflowStatus)
	err = workflow.ExecuteActivity(ctx, BackupCreateCmdActivity, config, workflowStatus).Get(ctx, &result)
	messages = util.PrependMessages(messages, result.Messages)
	result.Messages = messages

	util.SetStepComplete(workflowStatus, step)
	workflowStatusResult.Workflow = *workflowStatus
	workflowStatusResult.Result = result
	if err != nil {
		util.SetWorkflowStatusError(workflowStatus)
		workflowStatusResult.Workflow = *workflowStatus
		return workflowStatusResult, err
	}

	step = StepInit(workflowStatus)
	err = workflow.ExecuteActivity(ctx, PreAppUnQuiesceCmdActivity, config, workflowStatus).Get(ctx, &result)
	messages = util.PrependMessages(messages, result.Messages)
	result.Messages = messages

	util.SetStepComplete(workflowStatus, step)
	workflowStatusResult.Workflow = *workflowStatus
	workflowStatusResult.Result = result
	if err != nil {
		util.SetWorkflowStatusError(workflowStatus)
		workflowStatusResult.Workflow = *workflowStatus
		return workflowStatusResult, err
	}

	step = StepInit(workflowStatus)
	err = workflow.ExecuteActivity(ctx, PostAppUnQuiesceCmdActivity, config, workflowStatus).Get(ctx, &result)
	messages = util.PrependMessages(messages, result.Messages)
	result.Messages = messages

	util.SetStepComplete(workflowStatus, step)
	util.SetWorkflowStatusEnd(workflowStatus)

	workflowStatusResult.Workflow = *workflowStatus
	workflowStatusResult.Result = result
	if err != nil {
		util.SetWorkflowStatusError(workflowStatus)
		workflowStatusResult.Workflow = *workflowStatus
		return workflowStatusResult, err
	}

	logger.Info("Backup workflow completed.", "result", workflowStatusResult)

	return workflowStatusResult, nil
}
