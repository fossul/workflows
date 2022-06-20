package restore

import (
	"context"

	"github.com/fossul/fossul/src/client"
	"github.com/fossul/fossul/src/engine/util"
	"github.com/fossul/workflows/fossul"

	// TODO(cretz): Remove when tagged
	_ "go.temporal.io/sdk/contrib/tools/workflowcheck/determinism"
)

var result util.Result

func PreRestoreCmdActivity(config util.Config) (util.Result, error) {
	auth := fossul.SetAuth()

	result, err := client.PreAppRestoreCmd(auth, config)
	if err != nil {
		return result, err
	}

	return result, nil
}

func PreRestoreActivity(config util.Config) (util.Result, error) {
	auth := fossul.SetAuth()

	result, err := client.PreRestore(auth, config)
	if err != nil {
		return result, err
	}

	return result, nil
}

func RestoreCmdActivity(ctx context.Context, config util.Config) (util.Result, error) {
	auth := fossul.SetAuth()

	result, err := client.RestoreCmd(auth, config)
	if err != nil {
		return result, err
	}

	return result, nil
}

func RestoreActivity(ctx context.Context, config util.Config) (util.Result, error) {
	auth := fossul.SetAuth()

	result, err := client.Restore(auth, config)
	if err != nil {
		return result, err
	}

	return result, nil
}

func PostAppRestoreCmdActivity(ctx context.Context, config util.Config) (util.Result, error) {
	auth := fossul.SetAuth()

	result, err := client.PostAppRestoreCmd(auth, config)
	if err != nil {
		return result, err
	}

	return result, nil
}

func PostRestoreActivity(ctx context.Context, config util.Config) (util.Result, error) {
	auth := fossul.SetAuth()

	result, err := client.PostRestore(auth, config)
	if err != nil {
		return result, err
	}

	return result, nil
}
