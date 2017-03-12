package utils

import "reflect"

type AsyncTask struct {
	handler reflect.Value
	params  []reflect.Value
}

func NewAsyncTask(handler interface{}, params ...interface{}) *AsyncTask {

	handlerValue := reflect.ValueOf(handler)

	if handlerValue.Kind() == reflect.Func {
		task := AsyncTask{
			handler: handlerValue,
			params:  make([]reflect.Value, 0),
		}
		if paramNum := len(params); paramNum > 0 {
			task.params = make([]reflect.Value, paramNum)
			for index, v := range params {
				task.params[index] = reflect.ValueOf(v)
			}
		}
		return &task
	}
	panic("handler not func")
}

func (p *AsyncTask) Do() []reflect.Value {
	return p.handler.Call(p.params)
}
