package backup

import (
	"github.com/fossul/fossul/src/client"
	"github.com/fossul/fossul/src/engine/util"
)

func StepInit(workflow *util.Workflow) util.Step {
	step := util.CreateStep(workflow)
	util.SetWorkflowStep(workflow, step)

	return step
}

func SetWorkflowStep(workflow *util.Workflow, step util.Step) *util.Workflow {
	steps := workflow.Steps
	steps = append(steps, step)
	workflow.Steps = steps

	return workflow
}

func SetLastMessage(result util.Result, workflow *util.Workflow) *util.Workflow {

	if result.Messages != nil {
		if len(result.Messages) > 0 {
			lastMessage := result.Messages[len(result.Messages)-1]

			lastMessageText := lastMessage.Message
			workflow.LastMessage = lastMessageText
		}
	}

	return workflow
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
