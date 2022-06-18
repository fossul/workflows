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

func Workflow(ctx workflow.Context, config util.Config) (util.Result, error) {
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
	//var result string
	//err := workflow.ExecuteActivity(ctx, BackupCreateCmdActivity, config).Get(ctx, &result)
	err := workflow.ExecuteActivity(ctx, BackupCreateCmdActivity, config).Get(ctx, &result)
	if err != nil {
		return result, err
	}

	logger.Info("Backup workflow completed.", "result", result)

	return result, nil
}

func BackupCreateCmdActivity(ctx context.Context, config util.Config) (util.Result, error) {
	auth := SetAuth()
	var result util.Result
	logger := activity.GetLogger(ctx)
	if config.BackupCreateCmd != "" {
		result, err := client.BackupCreateCmd(auth, config)
		if err != nil {
			return result, err

		}

		for _, msg := range result.Messages {
			logger.Info(msg.Message)
		}

		return result, nil
	}

	return result, nil
}

func SetAuth() client.Auth {
	var auth client.Auth
	auth.ServerHostname = serverHostname
	auth.ServerPort = serverPort
	auth.AppHostname = appHostname
	auth.AppPort = appPort
	auth.StorageHostname = storageHostname
	auth.StoragePort = storagePort
	auth.Username = myUser
	auth.Password = myPass

	return auth
}
