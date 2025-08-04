# Project planning Phase 1 [Implement scheduler and task runner]

## Main points

- This is a distributed task schedular system that run tasks and distributed them among multiple machine using ...(pick an algorithms)
- Based on the type the master will divide the tasks among the other slaves and track their progress, logging, mentoring etc...

### Schedular

- Schedular is the main process of the service which defining the master and slaves
- Schedular is one instance over the runtime program
- Schedular have the following structure:
  - tasks: a buffered channel contains the running task and buffered with the number of limited tasks to run
  - workers: the number of threads this schedular will work with
  - stopChan: to stop (workers) tasks that running right now
- We will need context to cancel a task.

```go

type Schedular struct {
  tasks chan Task
  workers int // default = 1
  stopChan chan struct{}
}

```

### Task Schema

- Create a task struct that contains the following structure:
  - Have an ID for tracking and cancellation purposes
  - Params for sending it to the ExecFunc function to execute based on it
  - Priority to prioritize the tasks
  - Timeout to give the task a time frame to work within
  - ExecFunc the function implementation to run the task with the params sent

```go
type Tasker interface {
  Execute(ctx context.Context)
}

type Task struct {
  ID string,
  params map[string]any,
  priority int
  timeout time.Duration
}
```

### Task Behavior

- Task must have the running function that will take the params and execute based on it
- Master will divide the tasks and observe their progress etc...
- User can track a task it's logging it's progress on the current running task, expected time to finish etc...
- 1. User will send a task of the defined types
- 2. Master will receive the task and act on it, assign the proper runner function and the division of the tasks to the slaves etc.. based on the distributed algorithm
- 3. Master will send the divided mini-task to each slave and observe their behavior, progress, timeout, different errors
- 4. Slaves will return the final result of mini-task to the Master, master will aggregate all the mini-task and form the user final response
