package parallel

import "sync"

// Run executes the specified worker function multiple times in multiple
// goroutines, passing each a workerNum from 0-parallelism. This function
// returns after all instances of the function have run to completion.
func Run(parallelism int, worker func(workerNum int)) {
	allDone := sync.WaitGroup{}
	allDone.Add(parallelism)

	for workerNum := 0; workerNum < parallelism; workerNum++ {
		go func(workerNum int) {
			defer allDone.Done()
			worker(workerNum)
		}(workerNum)
	}

	allDone.Wait()
}
