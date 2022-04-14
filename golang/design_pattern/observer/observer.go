package observer

import "fmt"

/*
多个对象间存在一对多关系，当一个对象发生改变时，把这种改变通知给其他多个对象，从而影响其他对象的行为

观察者模式用于触发联动。
subject的改变会触发他所有观察者的相关动作
*/

type IObserver interface {
	Notify()
}

type ISubject interface {
	AddObserver(observer ...IObserver)
	NotifyObserver()
}

type Observer struct{}

type Subject struct {
	OArr []IObserver
}

func NewObserver() *Observer {
	return &Observer{}
}

func NewSubject() *Subject {
	return &Subject{}
}

func (o *Observer) Notify() {
	fmt.Println("执行指令")
}

func (s *Subject) AddObserver(observer ...IObserver) {
	s.OArr = append(s.OArr, observer...)
}

func (s *Subject) NotifyObserver() {
	for _, v := range s.OArr {
		v.Notify()
	}
}
