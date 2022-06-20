package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fossul/fossul/src/engine/util"
	"github.com/fossul/workflows/restore"
	"go.temporal.io/sdk/client"
)

const (
	gRPCEndpoint = "localhost:7233"
	NamespaceID  = "default"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{
		HostPort:  gRPCEndpoint,
		Namespace: NamespaceID,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "fossul_workflowID",
		TaskQueue: "fossul",
	}

	var config util.Config
	workflowStatus := &util.Workflow{}
	workflowStatus.Id = util.GetWorkflowId()
	workflowStatus.Type = "temporal"
	workflowStatus.Policy = "daily"
	workflowStatus.Status = "RUNNING"

	config.AccessWithinCluster = "false"
	config.AppPlugin = "sample-app.so"
	config.StoragePlugin = "sample-storage.so"
	//config.PreAppRestoreCmd = "echo,pre app restore command"
	//config.PostAppRestoreCmd = "echo,post app restore command"
	//config.RestoreCmd = "echo,restore command"

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, restore.RestoreWorkflow, config, workflowStatus)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	fmt.Println("test ", we)
	// Synchronously wait for the workflow completion.
	var workflowStatusResult util.WorkflowStatusResult
	err = we.Get(context.Background(), &workflowStatusResult)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", workflowStatusResult)
}
