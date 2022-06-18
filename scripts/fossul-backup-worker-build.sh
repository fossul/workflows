#!/bin/bash

echo "Installing Dependencies"
go mod tidy

echo "Running Unit Tests"
go test github.com/fossul/workflows/backup
if [ $? != 0 ]; then exit 1; fi

echo "Building Backup Worker"
go install github.com/fossul/workflows/backup/worker
if [ $? != 0 ]; then exit 1; fi

if [[ ! -z "${GOBIN}" ]]; then
	echo "Copying startup scripts"
	mv $GOBIN/worker $GOBIN/backupWorker
	cp scripts/fossul-backup-worker-startup.sh $GOBIN
	if [ $? != 0 ]; then exit 1; fi
fi

echo "Backup worker build completed successfully"

