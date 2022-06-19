package backup

import (
	"context"
	"time"

	"os"

	"github.com/fossul/fossul/src/client"
	"github.com/fossul/fossul/src/engine/util"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"

	// TODO(cretz): Remove when tagged
	_ "go.temporal.io/sdk/contrib/tools/workflowcheck/determinism"
)

var result util.Result
var port string = os.Getenv("FOSSUL_SERVER_SERVICE_PORT")
var configDir string = os.Getenv("FOSSUL_SERVER_CONFIG_DIR")
var dataDir string = os.Getenv("FOSSUL_SERVER_DATA_DIR")
var myUser string = os.Getenv("FOSSUL_USERNAME")
var myPass string = os.Getenv("FOSSUL_PASSWORD")
var serverHostname string = os.Getenv("FOSSUL_SERVER_CLIENT_HOSTNAME")
var serverPort string = os.Getenv("FOSSUL_SERVER_CLIENT_PORT")
var appHostname string = os.Getenv("FOSSUL_APP_CLIENT_HOSTNAME")
var appPort string = os.Getenv("FOSSUL_APP_CLIENT_PORT")
var storageHostname string = os.Getenv("FOSSUL_STORAGE_CLIENT_HOSTNAME")
var storagePort string = os.Getenv("FOSSUL_STORAGE_CLIENT_PORT")
var debug string = os.Getenv("FOSSUL_SERVER_DEBUG")

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
		return workflowStatusResult, err
	}

	step = StepInit(workflowStatus)
	err = workflow.ExecuteActivity(ctx, PostAppUnQuiesceCmdActivity, config, workflowStatus).Get(ctx, &result)
	messages = util.PrependMessages(messages, result.Messages)
	result.Messages = messages

	util.SetStepComplete(workflowStatus, step)
	workflowStatusResult.Workflow = *workflowStatus
	workflowStatusResult.Result = result
	if err != nil {
		return workflowStatusResult, err
	}

	logger.Info("Backup workflow completed.", "result", workflowStatusResult)

	return workflowStatusResult, nil
}

func PreAppQuiesceCmdActivity(ctx context.Context, config util.Config) (util.Result, error) {
	auth := SetAuth()
	logger := activity.GetLogger(ctx)

	if config.PreAppQuiesceCmd != "" {
		result, err := client.PreQuiesceCmd(auth, config)

		if err != nil {
			return result, err
		}

		// print debug messages
		for _, msg := range result.Messages {
			logger.Debug(msg.Message)
		}

		return result, nil
	}

	return result, nil
}

func AppQuiesceActivity(ctx context.Context, config util.Config) (util.Result, error) {
	auth := SetAuth()
	logger := activity.GetLogger(ctx)

	if config.AppPlugin != "" {
		result, err := client.Quiesce(auth, config)

		if err != nil {
			return result, err
		}

		// print debug messages
		for _, msg := range result.Messages {
			logger.Debug(msg.Message)
		}

		return result, nil
	}

	return result, nil
}

func PostAppQuiesceCmdActivity(ctx context.Context, config util.Config) (util.Result, error) {
	auth := SetAuth()
	logger := activity.GetLogger(ctx)

	if config.PostAppQuiesceCmd != "" {
		result, err := client.PostQuiesceCmd(auth, config)

		if err != nil {
			return result, err
		}

		// print debug messages
		for _, msg := range result.Messages {
			logger.Debug(msg.Message)
		}

		return result, nil
	}

	return result, nil
}

func BackupCreateCmdActivity(ctx context.Context, config util.Config) (util.Result, error) {
	auth := SetAuth()
	logger := activity.GetLogger(ctx)

	if config.BackupCreateCmd != "" {
		result, err := client.BackupCreateCmd(auth, config)

		if err != nil {
			return result, err
		}

		// print debug messages
		for _, msg := range result.Messages {
			logger.Debug(msg.Message)
		}

		return result, nil
	}

	return result, nil
}

func BackupCreateActivity(ctx context.Context, config util.Config) (util.Result, error) {
	auth := SetAuth()
	logger := activity.GetLogger(ctx)

	if config.StoragePlugin != "" {
		result, err := client.Backup(auth, config)

		if err != nil {
			return result, err
		}

		// print debug messages
		for _, msg := range result.Messages {
			logger.Debug(msg.Message)
		}

		return result, nil
	}

	return result, nil
}

func PreAppUnQuiesceCmdActivity(ctx context.Context, config util.Config) (util.Result, error) {
	auth := SetAuth()
	logger := activity.GetLogger(ctx)

	if config.PreAppUnquiesceCmd != "" {
		result, err := client.PreUnquiesceCmd(auth, config)

		if err != nil {
			return result, err
		}

		// print debug messages
		for _, msg := range result.Messages {
			logger.Debug(msg.Message)
		}

		return result, nil
	}

	return result, nil
}

func AppUnQuiesceActivity(ctx context.Context, config util.Config) (util.Result, error) {
	auth := SetAuth()
	logger := activity.GetLogger(ctx)

	if config.AppPlugin != "" {
		result, err := client.Unquiesce(auth, config)

		if err != nil {
			return result, err
		}

		// print debug messages
		for _, msg := range result.Messages {
			logger.Debug(msg.Message)
		}

		return result, nil
	}

	return result, nil
}

func PostAppUnQuiesceCmdActivity(ctx context.Context, config util.Config) (util.Result, error) {
	auth := SetAuth()
	logger := activity.GetLogger(ctx)

	if config.PreAppUnquiesceCmd != "" {
		result, err := client.PostUnquiesceCmd(auth, config)

		if err != nil {
			return result, err
		}

		// print debug messages
		for _, msg := range result.Messages {
			logger.Debug(msg.Message)
		}

		return result, nil
	}

	return result, nil
}
