package service

import (
	"net/rpc"

	"github.com/wangqiangorz/fourinone/pkg/model"
	logger "go.uber.org/zap"
)

type CacheFacadeClient struct {
	serviceName string
	host        string
	port        int
	client      *rpc.Client
}

func NewCacheFacadeClient(serviceName, host string, port int) (*CacheFacadeClient, error) {
	cacheFacadeClient := new(CacheFacadeClient)
	cacheFacadeClient.serviceName = serviceName
	cacheFacadeClient.host = host
	cacheFacadeClient.port = port
	_, err := cacheFacadeClient.getRpcClient()
	return cacheFacadeClient, err
}

func (cacheFacadeClient *CacheFacadeClient) getRpcClient() (*rpc.Client, error) {
	if cacheFacadeClient.client == nil {
		client, err := GetBean(cacheFacadeClient.host, cacheFacadeClient.port)
		if err != nil {
			logger.L()
		}
		return client, err
	}
	return cacheFacadeClient.client, nil
}

func (cacheFacadeClient *CacheFacadeClient) Add(name string, value interface{}) (string, error) {
	var ret string
	client, err := cacheFacadeClient.getRpcClient()
	if err != nil {
		return "", err
	}
	err = client.Call(cacheFacadeClient.serviceName+".Add", &model.CacheFacadeArgv{
		Name:  name,
		Value: value,
	}, &ret)
	return ret, err
}

func (cacheFacadeClient *CacheFacadeClient) Put(key, name string, value interface{}) (bool, error) {
	var ret bool
	client, err := cacheFacadeClient.getRpcClient()
	if err != nil {
		return false, err
	}
	err = client.Call(cacheFacadeClient.serviceName+".Put", &model.CacheFacadeArgv{
		Key:   key,
		Name:  name,
		Value: value,
	}, &ret)
	return ret, err
}

func (cacheFacadeClient *CacheFacadeClient) Get(key, name string) (*model.BeanVal, error) {
	var ret model.BeanVal
	client, err := cacheFacadeClient.getRpcClient()
	if err != nil {
		return nil, err
	}
	err = client.Call(cacheFacadeClient.serviceName+".Get", &model.CacheFacadeArgv{
		Key:  key,
		Name: name,
	}, &ret)
	return &ret, err
}

func (cacheFacadeClient *CacheFacadeClient) Gets(key string) ([]*model.BeanVal, error) {
	var ret []*model.BeanVal
	client, err := cacheFacadeClient.getRpcClient()
	if err != nil {
		return nil, err
	}
	err = client.Call(cacheFacadeClient.serviceName+".Gets", &model.CacheFacadeArgv{
		Key: key,
	}, &ret)
	return ret, err
}

func (cacheFacadeClient *CacheFacadeClient) Remove(key, name string) (bool, error) {
	var ret bool
	client, err := cacheFacadeClient.getRpcClient()
	if err != nil {
		return false, err
	}
	err = client.Call(cacheFacadeClient.serviceName+".Remove", &model.CacheFacadeArgv{
		Key:  key,
		Name: name,
	}, &ret)
	return ret, err
}
