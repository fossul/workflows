# Welcome to Fossul Framework Workflows
Currently fossul still implementes a static workflow. Some of the major drawbacks are as follows:
* No ability to easily re-order workflow activities (steps)
* No ability to run workflow activities (steps) in parallel
* No ability to easily retry workflow activities (steps) 
* Workflow state and serialization needs to always be implemented in once place, on the controller, limiting scale as well as resilience
* Workflow logic is static and cumbersome to maintain
* Race conditions can occur where a workflow reports running but in reality it isn't (for example if workflow thread gets killed)
* Unit testing workflow is challenging

Thankfully [Temporal](https://temporal.io) has solved these problems and many more with a workflow-as-code approach. The beauty is in its simplicity, scale and resilience. 

Fossul has two workflow types today, backup and restore. The workflow and it's activities are implemented as a Temporal worker. The worker is where the Fossul workflow steps are implemented using the Temporal SDK. The worker communicates with the Temporal cluster which is responsible for orchestrating state transitions. Since the worker will execute a Fossul workflow it becomes an additional service and would be deployed as part of Fossul.

The Fossul workflow result which contains a list of steps, state and logging for each workflow is implemented inside the worker and persisted by Temporal. This means we likely no longer need to serialize that data and can now just query Temporal.

## Testing
### Build worker service
```$ scripts/fossul-backup-worker-build.sh```

### Start worker service
```$ $GOBIN/fossul-backup-worker-startup.sh```

### Run backup workflow starter
```$ go run github.com/fossul/workflows/backup/starter```

### Run restore workflow starter
```$ go run github.com/fossul/workflows/restore/starter```

