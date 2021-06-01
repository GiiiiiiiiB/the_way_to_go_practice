package main

import (
	"time"
	"fmt"
	"os/exec"
)

type T struct {
	a string
}

var aa chan T // 双向 channel

var bb chan <- T // 只能发送消息的 channel

var cc  <-chan T  // 只能接收消息的 channel

func send(a chan int) {
	fmt.Println("start send..")
	a <- 2
	fmt.Println("stop send ..")
}

func get(a chan int){
	fmt.Println("start  get..")
	v,ok := <- a
	fmt.Println("stop get..")
	fmt.Println("value:",v)
	fmt.Println("value:",v,"status:",ok)
}

func main(){
	a,_ := exec.LookPath("power")
	fmt.Println("-------------start------------")
	ab := make(chan int,3)
	defer close(ab)
	go send(ab)
	time.Sleep(1 * time.Second)
	go get(ab)
	time.Sleep(1*time.Second)

	go get(ab)
	time.Sleep(1*time.Second)
	close(ab)
	go get(ab)


	time.Sleep(1*time.Second)
	fmt.Println(a)
	fmt.Println("-------------end------------")
}
