package queue

var Workers = make(chan struct{}, 8)

func Acquire() {
	Workers <- struct{}{}
}

func Release() {
	<-Workers
}
