package backup

import (
	"context"

	"github.com/fossul/fossul/src/client"
	"github.com/fossul/fossul/src/engine/util"
	"github.com/fossul/workflows/fossul"

	// TODO(cretz): Remove when tagged
	_ "go.temporal.io/sdk/contrib/tools/workflowcheck/determinism"
)

func PreAppQuiesceCmdActivity(ctx context.Context, config util.Config) (util.Result, error) {
	var result util.Result
	auth := fossul.SetAuth()

	result, err := client.PreQuiesceCmd(auth, config)

	if err != nil {
		return result, err
	}

	return result, nil
}

func AppQuiesceActivity(ctx context.Context, config util.Config) (util.Result, error) {
	var result util.Result
	auth := fossul.SetAuth()

	result, err := client.Quiesce(auth, config)
	if err != nil {
		return result, err
	}

	return result, nil
}

func PostAppQuiesceCmdActivity(ctx context.Context, config util.Config) (util.Result, error) {
	var result util.Result
	auth := fossul.SetAuth()

	result, err := client.PostQuiesceCmd(auth, config)
	if err != nil {
		return result, err
	}

	return result, nil
}

func BackupCreateCmdActivity(ctx context.Context, config util.Config) (util.Result, error) {
	var result util.Result
	auth := fossul.SetAuth()

	result, err := client.BackupCreateCmd(auth, config)
	if err != nil {
		return result, err
	}

	return result, nil
}

func BackupCreateActivity(ctx context.Context, config util.Config) (util.Result, error) {
	var result util.Result
	auth := fossul.SetAuth()

	result, err := client.Backup(auth, config)
	if err != nil {
		return result, err
	}

	return result, nil
}

func PreAppUnQuiesceCmdActivity(ctx context.Context, config util.Config) (util.Result, error) {
	var result util.Result
	auth := fossul.SetAuth()

	result, err := client.PreUnquiesceCmd(auth, config)
	if err != nil {
		return result, err
	}

	return result, nil
}

func AppUnQuiesceActivity(ctx context.Context, config util.Config) (util.Result, error) {
	var result util.Result
	auth := fossul.SetAuth()

	result, err := client.Unquiesce(auth, config)
	if err != nil {
		return result, err
	}

	return result, nil
}

func PostAppUnQuiesceCmdActivity(ctx context.Context, config util.Config) (util.Result, error) {
	var result util.Result
	auth := fossul.SetAuth()

	result, err := client.PostUnquiesceCmd(auth, config)
	if err != nil {
		return result, err
	}

	return result, nil
}
