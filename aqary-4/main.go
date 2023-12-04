package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	byteGlobal    = []int{}
	ReciveChannel = make(chan int, 2)
	mapMutex      sync.Mutex
)

func main() {

	fmt.Println("Run the tests")

}

func Consumer(number int, IsFinishChannelConsumer chan bool, ctx context.Context) {

	fmt.Println("Starting Consumer... ", number)

	for {

		select {
		case <-ctx.Done():
			fmt.Println("Done Consumer... ", number)
			IsFinishChannelConsumer <- true
			return

		default:
			mapMutex.Lock()
			if len(byteGlobal) == 0 {
				mapMutex.Unlock()
				continue
			}
			// Race condition
			fmt.Println("Race Condition Consumer... ", number)
			result := byteGlobal[len(byteGlobal)-1]
			time.Sleep(time.Microsecond * 10)
			byteGlobal = byteGlobal[0 : len(byteGlobal)-1]
			fmt.Println("Consumer ", number, "result ", result, " ", byteGlobal)
			// Race condition
			mapMutex.Unlock()
		}

	}
}

func Producer(number int, ctx context.Context) {

	fmt.Println("Starting Produce... ", number, " ")

	select {
	case <-ctx.Done():
		fmt.Println("Done Consumer... ", number)
		return
	default:
		mapMutex.Lock()
		// Race condition.
		byteGlobal = append(byteGlobal, number)
		fmt.Println("Produce ", number, " ", byteGlobal)
		// Race condition.
		mapMutex.Unlock()
	}

}
