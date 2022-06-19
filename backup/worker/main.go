package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	backup "github.com/fossul/workflows/backup"
)

const (
	gRPCEndpoint = "localhost:7233"
	NamespaceID  = "default"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort:  gRPCEndpoint,
		Namespace: NamespaceID,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "fossul_backup", worker.Options{})

	w.RegisterWorkflow(backup.Workflow)
	w.RegisterActivity(backup.PreAppQuiesceCmdActivity)
	w.RegisterActivity(backup.AppQuiesceActivity)
	w.RegisterActivity(backup.PostAppQuiesceCmdActivity)
	w.RegisterActivity(backup.BackupCreateCmdActivity)
	w.RegisterActivity(backup.BackupCreateActivity)
	w.RegisterActivity(backup.PreAppUnQuiesceCmdActivity)
	w.RegisterActivity(backup.AppUnQuiesceActivity)
	w.RegisterActivity(backup.PostAppUnQuiesceCmdActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
