// Package tasker allows you to perform tasks in parallel to each other
package tasker

// NewTasker returns an interface containing a reference to the task allocator with the maximum number of concurrently executing tasks specified in the argument
func NewTasker(count uint16) Tasker {
	d := &distributor{}
	d.setCount(count)
	d.init()
	return d
}

// Tasker interface that allows parallel execution of tasks
type Tasker interface {
	Do(task interface{})
	Wait()
	SetHandler(handler func(interface{}))
	SetDebug(debug bool)
}
