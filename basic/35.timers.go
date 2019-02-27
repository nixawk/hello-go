package main

import "time"
import "fmt"

func main() {

	timer1 := time.NewTimer(2 * time.Second) /* time.Second == 1s */

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}

/* Timers are for when you want to do something once in the future. */
/* https://gobyexample.com/timers */
