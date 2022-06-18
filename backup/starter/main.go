package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fossul/fossul/src/engine/util"
	"github.com/fossul/workflows/backup"
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
		ID:        "fossul_backup_workflowID",
		TaskQueue: "fossul_backup",
	}

	var config util.Config
	config.AccessWithinCluster = "false"
	config.BackupCreateCmd = "echo,backup create command"
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, backup.Workflow, config)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	fmt.Println("test ", we)
	// Synchronously wait for the workflow completion.
	var result util.Result
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}
