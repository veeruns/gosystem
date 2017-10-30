package gosystem

import (
	"fmt"
	"time"
)

func Routines() {
	qsignal := make(chan bool)
	go Report(qsignal)
	fmt.Println("From a normal stuff")
	vs := <-qsignal
	fmt.Println(vs)
	in := make(chan int)
	go PeriodicSend(in)
	for i := range in {
		fmt.Println(i)
	}
}

func Report(qs chan bool) {
	fmt.Println("Hello From Routine")
	qs <- false
}

func PeriodicSend(ic chan int) {
	i := 0
	for i <= 10 {
		i++
		ic <- i
		time.Sleep(1 * time.Second)
	}
	close(ic)
}
