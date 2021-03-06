package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int)  {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func consumer(in <-chan int)  {
	for data := range in {
		fmt.Println(data)
	}
}

func main()  {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)

	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		fmt.Println("loop")
	}
}