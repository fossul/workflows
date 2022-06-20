package fossul

import (
	"os"

	"github.com/fossul/fossul/src/client"
	"github.com/fossul/fossul/src/engine/util"
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

func AppendMessages(newMessages, messages []util.Message) []util.Message {
	for _, msg := range newMessages {
		messages = append(messages, msg)
	}

	return messages
}

func ErrorHandling(messages []util.Message, result util.Result, workflowStatus *util.Workflow, workflowStatusResult util.WorkflowStatusResult) util.WorkflowStatusResult {
	util.SetWorkflowStatusError(workflowStatus)
	result.Messages = messages

	workflowStatusResult.Workflow = *workflowStatus
	workflowStatusResult.Result = result

	return workflowStatusResult
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
