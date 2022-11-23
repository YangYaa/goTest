package channel

import "errors"

type TaskResult struct {
	OutPut interface{}
	Err    error
}
type Task struct {
	Handler     func(v ...interface{}) (interface{}, error)
	Params      []interface{}
	Output      chan *TaskResult
	ErrorHandle func(err error)
}

func NewTask(handler func(v ...interface{}) (interface{}, error), params []interface{}, output chan *TaskResult) *Task {
	return &Task{
		Handler: handler,
		Params:  params,
		Output:  output,
	}
}

func NotifyTask(p ...interface{}) (interface{}, error) {

	val := p[2].(int)
	if val%2 != 0 {
		return "notify successfully", nil
	} else {
		err := errors.New("error to notify")
		return "failed to notify", err
	}

}
