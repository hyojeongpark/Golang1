package main

import (
	"fmt"
	"time"
)

//push 설정
func push(c chan int) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		c <- i
		i++
	}
}

func main() {
	c := make(chan int) //channel생성

	go push(c)

	timerChan := time.After(10 * time.Second) //특정시간만 알림
	tick := time.Tick(2 * time.Second)        //시간의 주기적 알림
	//push 할 내용 : c, timerchan, tick
	for {
		select {
		case v := <-c: //c에 들어가는 입력값이 v가 됨
			fmt.Println(v)
		//case <-time.After(10 * time.Second): //이렇게 하면 무한루프
		case <-timerChan:
			fmt.Println("Time out")
			return
		case <-tick:
			fmt.Println("Tick")
			// default:
			// 	fmt.Println("Idle")
			// 	time.Sleep(1 * time.Second)
		}
	}
}
