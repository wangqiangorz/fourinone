package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/wangqiangorz/fourinone/pkg/worker"
)

func Test_a(t *testing.T) {
	go PutBean("workerServer", "127.0.0.1", 8081, NewWorkerService(worker.NewReceiver("127.0.0.1", 8081, "workServer", func(input interface{}) {
		fmt.Println(input)
	})))
	time.Sleep(time.Second * 3)

	client, err := NewWorkerClient("workerServer", "127.0.0.1", 8081)
	if err != nil {
		fmt.Println(err.Error())
	}
	client.Send(1)
}
