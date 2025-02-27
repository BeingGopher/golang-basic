package main

import(
	"fmt"
	"time"
)

func main(){

	timer1 := time.NewTimer(2*time.Second)

	<-timer1.C

	fmt.Println("Time 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func(){
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	//time.Sleep(1600 * time.Millisecond)
	stop2 := timer2.Stop()
	if stop2{
		fmt.Println("timer 2 stopped")
	}

	time.Sleep(2 * time.Second)

}