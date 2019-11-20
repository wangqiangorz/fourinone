package context

import (
	"github.com/wangqiangorz/fourinone/pkg/model"
	"github.com/wangqiangorz/fourinone/pkg/proxy"
	"github.com/wangqiangorz/fourinone/pkg/util"
)

var (
	parkProxy *proxy.ParkProxy
)

func GetParkLocal() (*proxy.ParkProxy, error) {
	if parkProxy == nil {
		return GetPark()
	}
	return parkProxy, nil
}

func GetParkLocalByHost(host string, port int) (*proxy.ParkProxy, error) {
	if parkProxy == nil {
		return GetParkByHost(host, port)
	}
	return parkProxy, nil
}

func GetWorkerByWorkerType(workerType string) ([]*model.BeanVal, error) {
	proxy, err := GetParkLocal()
	if err != nil {
		return nil, err
	}
	return proxy.GetNodes(WORKERPRFIX + workerType)
}

func GetWorkerByWorkerTypeAndHost(host string, port int, workerType string) ([]*model.BeanVal, error) {
	proxy, err := GetParkLocalByHost(host, port)
	if err != nil {
		return nil, err
	}
	return proxy.GetNodes(WORKERPRFIX + workerType)
}

func CreateWorkType(workerType string, nodevalue interface{}) error {
	proxy, err := GetParkLocal()
	if err != nil {
		return err
	}
	return proxy.Create(WORKERPRFIX+workerType, util.GetUUid(), nodevalue, true)
}

func CreateWorkTypeByHost(host string, port int, workerType string, nodevalue interface{}) error {
	proxy, err := GetParkLocalByHost(host, port)
	if err != nil {
		return err
	}
	return proxy.Create(WORKERPRFIX+workerType, util.GetUUid(), nodevalue, true)
}
