package radish

import (
	"log"
	"reflect"
)

type Radish struct {
	listOfTasks []func()
}

func New() Radish {
	return Radish{}
}

func (r *Radish) RegisterTask(t func()) {
	r.listOfTasks = append(r.listOfTasks, t)
}

func (r *Radish) RunTaskQueue() {
	for _, task := range r.listOfTasks {
		go task()
	}
}

func Task(fn any, params ...any) {
	f := reflect.ValueOf(fn)
	if f.Kind() != reflect.Func {
		log.Fatalf("Task Is Not A Function Is: %v", f.Kind())
	}

	in := make([]reflect.Value, len(params))

	for i, param := range params {
		in[i] = reflect.ValueOf(param)
	}

	f.Call(in)

	ch := make(chan reflect.Value)

	go func() {
		ch <- f
	}()

	go func() {
		task := <-ch
		task.Call(in)
	}()
}
