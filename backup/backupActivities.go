package backup

import (
	"context"

	"os"

	"github.com/fossul/fossul/src/client"
	"github.com/fossul/fossul/src/engine/util"
	"go.temporal.io/sdk/activity"

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
