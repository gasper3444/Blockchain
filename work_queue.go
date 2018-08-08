package work_queue

type Worker interface {
	Run() interface{}			//can return any data type (empty inteface)
}

type WorkQueue struct {

	Jobs         chan Worker
	Results      chan interface{}			//holds results of calling Run()
	StopRequests chan int
	NumWorkers   uint
}

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	q := new(WorkQueue)
	q.NumWorkers= nWorkers
	i := uint(0)
	q.Jobs = make(chan Worker, maxJobs)
	q.StopRequests = make(chan int,1)
	q.Results = make(chan interface{})
	//starting nWorkers worker goroutines.
	for i < nWorkers {
		go q.worker()
		i++
	}
	return q
}

// A worker goroutine that processes tasks from .Jobs unless .StopRequests has a message saying to halt now.
func (queue WorkQueue) worker() {
	running := true
	// Run tasks from the queue, unless we have been asked to stop.
	for running {
		if len(queue.StopRequests) > 0 {
					break
		}	else {
			w := <-queue.Jobs
			queue.Results <- w.Run()
		}
	}
	return //stop the execution
}

func (queue WorkQueue) Enqueue(work Worker) {				//called in MineRange
	queue.Jobs <- work //Putting the work instance in Jobs channel. (Will be retrieved in worker() and results would be put in Results)
}

func (queue WorkQueue) Shutdown() {				//called in MineRange
	messageToSend := queue.NumWorkers
	queue.StopRequests <- int(messageToSend)
}
