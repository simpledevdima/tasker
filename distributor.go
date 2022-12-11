package tasker

import (
	"fmt"
	"sync"
)

// distributor data structure that allows you to distribute tasks for their parallel execution
type distributor struct {
	debug      bool
	wg         *sync.WaitGroup
	count      uint16
	containers chan uint16
	tasks      chan interface{}
	handler    func(interface{})
}

// Wait method of waiting for the completion of all tasks
func (d *distributor) Wait() {
	d.wg.Wait()
}

// SetHandler sets the handler function of each task
func (d *distributor) SetHandler(handler func(interface{})) {
	d.handler = handler
}

// Do sending information about the task to the general task channel
func (d *distributor) Do(task interface{}) {
	d.wg.Add(1)
	d.tasks <- task
}

// SetDebug enable/disable debugging
func (d *distributor) SetDebug(dbg bool) {
	d.debug = dbg
}

// open task execution
func (d *distributor) open(container uint16) {
	defer d.wg.Done()
	tsk := <-d.tasks
	if d.debug {
		fmt.Printf("tasker exec container %d\n", container)
	}
	d.handler(tsk)
	d.containers <- container
}

// check for an empty container and executing a new task in it
func (d *distributor) check() {
	for {
		select {
		case container := <-d.containers:
			if container <= d.count {
				go d.open(container)
			}
		}
	}
}

// init preparation for operation of the new distributor
func (d *distributor) init() {
	d.containers = make(chan uint16)
	d.tasks = make(chan interface{})
	d.wg = &sync.WaitGroup{}
	go d.check()
	for container := uint16(1); container <= d.count; container++ {
		d.containers <- container
	}
}

// setCount setting the maximum number of parallel running tasks
func (d *distributor) setCount(count uint16) {
	d.count = count
}
