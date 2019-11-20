package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/wangqiangorz/fourinone/pkg/monitor"
)

func Test_PutBean(t *testing.T) {
	monitor.InitLog()
	go PutBean("ParkService", "127.0.0.1", 8081, NewParkService("ParkService", "127.0.0.1", 8081, []string{"127.0.0.1:8081", "127.0.0.1:8002"}))
	time.Sleep(5 * time.Second)

	//pl := NewParkLeader("ParkService", "127.0.0.1", 8081, []string{"127.0.0.1:8081", "127.0.0.1:8002"})
	// leader, err := pl.GetParkLeader()
	// fmt.Println(leader.GetHost(), leader.GetPort())
	parkClient, err := NewParkClient("ParkService", "127.0.0.1", 8081)
	ret, err := parkClient.Create("test", "fabu", "xxx", "123123", true)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ret)
	ret, err = parkClient.Get("test", "fabu", "123123")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ret)
	ret, err = parkClient.Update("test", "fabu", "fdsfvds", "123123", true)
	if err != nil {
		fmt.Println(err.Error())
	}
	ret1, err1 := parkClient.AskMaster()
	if err1 != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ret1)

	ret1, err1 = parkClient.GetSessionID()
	if err1 != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ret1)

	ret, err = parkClient.Delete("test", "fabu", "fdsfvds")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ret)
	ret, err = parkClient.Get("test", "fabu", "123123")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ret)
	ret1, err = parkClient.AskMaster()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ret1)
	ret1, err = parkClient.AskLeader()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ret1)
}
