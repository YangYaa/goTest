package channel

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

var InvalidPoolCapErr = errors.New("invalid pool cap")
var InvalidPoolStatus = errors.New("pool is already closed")

const (
	STOP    = 0
	RUNNING = 1
)

type Pool struct {
	capacity       uint32
	runningWorkers uint32
	status         uint32
	taskChan       chan *Task
	sync.Mutex
}

func NewPool(cap uint32) (*Pool, error) {
	if cap <= 0 {
		return nil, InvalidPoolCapErr
	}
	return &Pool{
		capacity: cap,
		status:   RUNNING,
		taskChan: make(chan *Task, cap),
	}, nil
}
func (p *Pool) Put(task *Task) error {
	var err error
	for try := 0; try < 3; try++ {
		err = p.put(task)
		if err == nil {
			break
		}
		time.Sleep(1e9)
	}
	return err
}

func (p *Pool) put(task *Task) error {
	p.incRunning()
	if p.status == STOP {
		return InvalidPoolStatus
	}

	if p.GetRunningWorks() < p.GetCap() {
		p.asyncRun()
	}

	if p.status == RUNNING {
		p.taskChan <- task
	}
	return nil
}

func (p *Pool) incRunning() {
	atomic.AddUint32(&p.runningWorkers, 1)
}
func (p *Pool) decRunning() {
	atomic.AddUint32(&p.runningWorkers, ^uint32(0))
}

func (p *Pool) asyncRun() {
	go func() {
		for {
			select {
			case task, ok := <-p.taskChan:
				if !ok {
					return
				}
				result, err := task.Handler(task.Params...)
				if task.ErrorHandle != nil {
					task.ErrorHandle(err)
				}
				task.Output <- &TaskResult{OutPut: result, Err: err}
				p.decRunning()
			}
		}
	}()
}

func (p *Pool) GetRunningWorks() uint32 {
	return atomic.LoadUint32(&p.runningWorkers)
}

func (p *Pool) GetCap() uint32 {
	return p.capacity
}
