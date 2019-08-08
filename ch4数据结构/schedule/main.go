package main

import (
	"DataStructre/ch4数据结构/queue"
	"fmt"
)

func main() {
	type mission struct {
		name string
		time int
	}

	missions := []mission{{"m1", 100}, {"m2", 400}, {"m3", 400}, {"m4", 600}, {"m5", 4050}}

	var q queue.Queue
	for _, v := range missions {
		q.Enqueue(v)
	}
	fmt.Println(q)
}
