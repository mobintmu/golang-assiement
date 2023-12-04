package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestM8N2(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	IsFinishChannelConsumer := make(chan bool, 8)

	go Consumer(1, IsFinishChannelConsumer, ctx)
	go Consumer(2, IsFinishChannelConsumer, ctx)
	go Producer(1, ctx)
	go Consumer(3, IsFinishChannelConsumer, ctx)
	go Producer(2, ctx)
	go Consumer(4, IsFinishChannelConsumer, ctx)
	go Consumer(5, IsFinishChannelConsumer, ctx)
	go Consumer(6, IsFinishChannelConsumer, ctx)
	go Consumer(7, IsFinishChannelConsumer, ctx)
	go Consumer(8, IsFinishChannelConsumer, ctx)

	time.Sleep(2 * time.Second)

	cancel()

	for counter := 1; counter <= 8; counter++ {
		<-IsFinishChannelConsumer
	}

	fmt.Println("_______ Finish __________")
	fmt.Println(byteGlobal)

	assert.Equal(t, len(byteGlobal), 0)
}

func TestM8N8(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	IsFinishChannelConsumer := make(chan bool, 8)

	go Consumer(1, IsFinishChannelConsumer, ctx)
	go Consumer(2, IsFinishChannelConsumer, ctx)
	go Producer(1, ctx)
	go Consumer(3, IsFinishChannelConsumer, ctx)
	go Producer(2, ctx)
	go Consumer(4, IsFinishChannelConsumer, ctx)
	go Consumer(5, IsFinishChannelConsumer, ctx)
	go Consumer(6, IsFinishChannelConsumer, ctx)
	go Consumer(7, IsFinishChannelConsumer, ctx)
	go Consumer(8, IsFinishChannelConsumer, ctx)
	go Producer(3, ctx)
	go Producer(4, ctx)
	go Producer(5, ctx)
	go Producer(6, ctx)
	go Producer(7, ctx)
	go Producer(8, ctx)

	time.Sleep(2 * time.Second)

	cancel()

	for counter := 1; counter <= 8; counter++ {
		<-IsFinishChannelConsumer
	}

	fmt.Println("_______ Finish __________")
	fmt.Println(byteGlobal)

	assert.Equal(t, len(byteGlobal), 0)
}

func TestM8N16(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	IsFinishChannelConsumer := make(chan bool, 8)

	for counter := 1; counter <= 8; counter++ {
		go Consumer(counter, IsFinishChannelConsumer, ctx)
	}
	for counter := 1; counter <= 16; counter++ {
		go Producer(counter, ctx)
	}

	time.Sleep(2 * time.Second)

	cancel()

	for counter := 1; counter <= 8; counter++ {
		<-IsFinishChannelConsumer
	}

	fmt.Println("_______ Finish __________")
	fmt.Println(byteGlobal)

	assert.Equal(t, len(byteGlobal), 0)
}

func TestM2N8(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	IsFinishChannelConsumer := make(chan bool, 2)

	for counter := 1; counter <= 2; counter++ {
		go Consumer(counter, IsFinishChannelConsumer, ctx)
	}
	for counter := 1; counter <= 16; counter++ {
		go Producer(counter, ctx)
	}

	time.Sleep(2 * time.Second)

	cancel()

	for counter := 1; counter <= 2; counter++ {
		<-IsFinishChannelConsumer
	}

	fmt.Println("_______ Finish __________")
	fmt.Println(byteGlobal)

	assert.Equal(t, len(byteGlobal), 0)
}
