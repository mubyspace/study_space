package observer

import "testing"

func TestObserver_Notify(t *testing.T) {
	observer := NewObserver()
	subject := NewSubject()
	subject.AddObserver(observer)
	subject.NotifyObserver()
}
