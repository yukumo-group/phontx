package tasks

// Task defines a single task
type Task struct {
	Task        string
	Description string
	Function    func() error
}

// TaskList defines list of task
type TaskList struct {
}
