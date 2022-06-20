#!/bin/bash

echo "Installing Dependencies"
go mod tidy

echo "Running Unit Tests"
go test -v github.com/fossul/workflows/backup
if [ $? != 0 ]; then exit 1; fi

echo "Building Fossul Worker"
go install github.com/fossul/workflows/worker
if [ $? != 0 ]; then exit 1; fi

if [[ ! -z "${GOBIN}" ]]; then
	echo "Copying startup scripts"
	cp scripts/fossul-backup-worker-startup.sh $GOBIN
	if [ $? != 0 ]; then exit 1; fi
fi

echo "Fossul worker build completed successfully"

