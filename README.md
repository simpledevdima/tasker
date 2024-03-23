# Tasker
Distributor for running tasks in parallel.

## Installation
```
go get github.com/skvdmt/tasker
```

## Example
```go
package main

import (
	"fmt"
	"github.com/skvdmt/tasker"
)

func main() {
	type task struct {
		id int
	}
	tsk := tasker.NewTasker(10)
	tsk.SetDebug(true)
	tsk.SetHandler(func(data interface{}) {
		t := data.(*task)
		fmt.Printf("%d\n", t.id)
	})
	for id := 1; id <= 100; id++ {
		tsk.Do(&task{
			id: id,
		})
	}
	tsk.Wait()
}
```