package proxy

import (
	"fmt"
	"testing"
	"time"

	"github.com/wangqiangorz/fourinone/pkg/monitor"
	"github.com/wangqiangorz/fourinone/pkg/service"
)

func Test(t *testing.T) {
	monitor.InitLog()
	go service.PutBean("ParkService", "127.0.0.1", 8081, service.NewParkService("ParkService", "127.0.0.1", 8081, []string{"127.0.0.1:8081", "127.0.0.1:8002"}))
	time.Sleep(1 * time.Second)
	parkProxy, err := NewParkProxy("ParkService", "127.0.0.1", 8081, []string{"127.0.0.1:8081", "127.0.0.1:8002"})
	if err != nil {
		fmt.Println(err)
	}
	err = parkProxy.Create("wqa", "zll", "love", false)
	if err != nil {
		fmt.Println(err)
	}
	beanval, err := parkProxy.GetNode("wqa", "zll")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(beanval.GetNodeValue())
	err = parkProxy.Delete("wqa", "zll")
	if err != nil {
		fmt.Println(err)
	}
	err = parkProxy.Delete("wqa", "zll")
	if err != nil {
		fmt.Println(err)
	}
}
