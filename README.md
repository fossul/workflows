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

## Sample Fossul Workflow Output Through Temporal
```
[
  {
    "workflow": {
      "id": 1383,
      "status": "COMPLETED",
      "type": "temporal",
      "policy": "daily",
      "steps": [
        {
          "id": 0,
          "status": "COMPLETE",
          "label": "Step 0"
        },
        {
          "id": 1,
          "status": "COMPLETE",
          "label": "Step 1"
        },
        {
          "id": 2,
          "status": "COMPLETE",
          "label": "Step 2"
        },
        {
          "id": 3,
          "status": "COMPLETE",
          "label": "Step 3"
        },
        {
          "id": 4,
          "status": "COMPLETE",
          "label": "Step 4"
        }
      ]
    },
    "result": {
      "messages": [
        {
          "time": 1655609847,
          "level": "INFO",
          "message": "Performing post unquiesce command [echo,post app unquiesce command]"
        },
        {
          "time": 1655609847,
          "level": "CMD",
          "message": "Executing command [echo post app unquiesce command]"
        },
        {
          "time": 1655609847,
          "level": "INFO",
          "message": "post app unquiesce command\n"
        },
        {
          "time": 1655609847,
          "level": "INFO",
          "message": "Command [echo post app unquiesce command] completed successfully"
        },
        {
          "time": 1655609847,
          "level": "INFO",
          "message": "Performing post quiesce command [echo,post quiesce command]"
        },
        {
          "time": 1655609847,
          "level": "CMD",
          "message": "Executing command [echo post quiesce command]"
        },
        {
          "time": 1655609847,
          "level": "INFO",
          "message": "post quiesce command\n"
        },
        {
          "time": 1655609847,
          "level": "INFO",
          "message": "Command [echo post quiesce command] completed successfully"
        },
        {
          "time": 1655609847,
          "level": "INFO",
          "message": "Performing backup create command [echo,backup create command]"
        },
        {
          "time": 1655609847,
          "level": "CMD",
          "message": "Executing command [echo backup create command]"
        },
        {
          "time": 1655609847,
          "level": "INFO",
          "message": "backup create command\n"
        },
        {
          "time": 1655609847,
          "level": "INFO",
          "message": "Command [echo backup create command] completed successfully"
        },
        {
          "time": 1655609847,
          "level": "INFO",
          "message": "Performing pre unquiesce command [echo,pre app unquiesce command]"
        },
        {
          "time": 1655609847,
          "level": "CMD",
          "message": "Executing command [echo pre app unquiesce command]"
        },
        {
          "time": 1655609847,
          "level": "INFO",
          "message": "pre app unquiesce command\n"
        },
        {
          "time": 1655609847,
          "level": "INFO",
          "message": "Command [echo pre app unquiesce command] completed successfully"
        },
        {
          "time": 1655609847,
          "level": "INFO",
          "message": "Performing post unquiesce command [echo,post app unquiesce command]"
        },
        {
          "time": 1655609847,
          "level": "CMD",
          "message": "Executing command [echo post app unquiesce command]"
        },
        {
          "time": 1655609847,
          "level": "INFO",
          "message": "post app unquiesce command\n"
        },
        {
          "time": 1655609847,
          "level": "INFO",
          "message": "Command [echo post app unquiesce command] completed successfully"
        }
      ],
      "backup": {}
    }
  }
]
```

## Testing
### Build worker service
```$ scripts/fossul-backup-worker-build.sh```

### Start worker service
```$ $GOBIN/fossul-backup-worker-startup.sh```

### Run workflow starter
```$ go run github.com/fossul/workflows/backup/starter```

