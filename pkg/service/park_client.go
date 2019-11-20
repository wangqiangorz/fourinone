package service

import (
	"net/rpc"

	"github.com/wangqiangorz/fourinone/pkg/model"
	"github.com/wangqiangorz/fourinone/pkg/park"
	logger "go.uber.org/zap"
)

type ParkClient struct {
	park.Park
	serviceName string
	host        string
	port        int
	client      *rpc.Client
}

func NewParkClient(serviceName, host string, port int) (*ParkClient, error) {
	parkClient := new(ParkClient)
	parkClient.serviceName = serviceName
	parkClient.host = host
	parkClient.port = port
	_, err := parkClient.getRpcClient()
	return parkClient, err
}

func (parkClient *ParkClient) getRpcClient() (*rpc.Client, error) {
	if parkClient.client == nil {
		client, err := GetBean(parkClient.host, parkClient.port)
		return client, err
	}
	return parkClient.client, nil
}

func (parkClient *ParkClient) GetHost() string {
	return parkClient.host
}

func (parkClient *ParkClient) GetPort() int {
	return parkClient.port
}

func (parkClient *ParkClient) Create(domain, node string, nodeValue interface{}, sessionid string, heartbeat bool) (*model.Ketchup, error) {
	var ret *model.Ketchup
	client, err := parkClient.getRpcClient()
	if err != nil {
		return nil, err
	}
	err = client.Call(parkClient.serviceName+".Create", &park.ParkArgs{
		Domain:    domain,
		Node:      node,
		Nodevalue: nodeValue,
		Sessionid: sessionid,
		Heartbeat: heartbeat,
	}, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (parkClient *ParkClient) Get(domain, node string, sessionid string) (*model.Ketchup, error) {
	var ret *model.Ketchup
	client, err := parkClient.getRpcClient()
	if err != nil {
		return nil, err
	}
	err = client.Call(parkClient.serviceName+".Get", &park.ParkArgs{
		Domain:    domain,
		Node:      node,
		Sessionid: sessionid,
	}, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (parkClient *ParkClient) Update(domain, node string, nodeValue interface{}, sessionid string, heartbeat bool) (*model.Ketchup, error) {
	var ret *model.Ketchup
	client, err := parkClient.getRpcClient()
	if err != nil {
		return nil, err
	}
	err = client.Call(parkClient.serviceName+".Update", &park.ParkArgs{
		Domain:    domain,
		Node:      node,
		Nodevalue: nodeValue,
		Sessionid: sessionid,
		Heartbeat: heartbeat,
	}, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (parkClient *ParkClient) Delete(domain, node string, sessionid string) (*model.Ketchup, error) {
	var ret *model.Ketchup
	client, err := parkClient.getRpcClient()
	if err != nil {
		return nil, err
	}
	err = client.Call(parkClient.serviceName+".Delete", &park.ParkArgs{
		Domain:    domain,
		Node:      node,
		Sessionid: sessionid,
	}, &ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}

func (parkClient *ParkClient) GetSessionID() (string, error) {
	var ret string
	client, err := parkClient.getRpcClient()
	if err != nil {
		return "", err
	}
	err = client.Call(parkClient.serviceName+".GetSessionID", &park.ParkArgs{}, &ret)
	if err != nil {
		return "", err
	}
	return ret, err
}

func (parkClient *ParkClient) AskMaster() (string, error) {
	var ret string
	client, err := parkClient.getRpcClient()
	if err != nil {
		return "", err
	}
	err = client.Call(parkClient.serviceName+".AskMaster", &park.ParkArgs{}, &ret)
	if err != nil {
		return "", err
	}
	return ret, nil
}

func (parkClient *ParkClient) AskLeader() (string, error) {
	var ret string
	client, err := parkClient.getRpcClient()
	if err != nil {
		return "", err
	}
	err = client.Call(parkClient.serviceName+".AskLeader", &park.ParkArgs{}, &ret)
	if err != nil {
		logger.L().Warn("test2", logger.Error(err))
		return "", err
	}
	return ret, nil
}

func (parkClient *ParkClient) HeartBeat(domain, node string, sessionID string) (bool, error) {
	var ret bool
	client, err := parkClient.getRpcClient()
	if err != nil {
		return false, err
	}
	err = client.Call(parkClient.serviceName+".HeartBeat", &park.ParkArgs{
		Domain:    domain,
		Node:      node,
		Sessionid: sessionID,
	}, &ret)
	if err != nil {
		return false, err
	}
	return ret, nil
}

func (parkClient *ParkClient) Close() error {
	if parkClient.client != nil {
		err := parkClient.client.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
