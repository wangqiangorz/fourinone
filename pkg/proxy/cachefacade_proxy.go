package proxy

import "github.com/wangqiangorz/fourinone/pkg/service"

type CacheFacadeProxy struct {
	host        string
	port        int
	serviceName string
	client      *service.CacheFacadeClient
}

func NewCacheFacadeProxy(serviceName, host string, port int) (*CacheFacadeProxy, error) {
	facadeProxy := new(CacheFacadeProxy)
	facadeProxy.host = host
	facadeProxy.port = port
	facadeProxy.serviceName = serviceName
	client, err := service.NewCacheFacadeClient(serviceName, host, port)
	if err != nil {
		return nil, err
	}
	facadeProxy.client = client
	return facadeProxy, nil
}

func (cacheFacadeProxy *CacheFacadeProxy) GetServiceName() string {
	return cacheFacadeProxy.serviceName
}

func (cacheFacadeProxy *CacheFacadeProxy) GetServiceHost() string {
	return cacheFacadeProxy.host
}

func (cacheFacadeProxy *CacheFacadeProxy) GetServicePort() int {
	return cacheFacadeProxy.port
}

func (cacheFacadeProxy *CacheFacadeProxy) Add(name string, value interface{}) (string, error) {
	key, err := cacheFacadeProxy.client.Add(name, value)
	if err != nil {
		return "", err
	}
	return key, nil
}

func (cacheFacadeProxy *CacheFacadeProxy) Put(key, name string, value interface{}) error {
	_, err := cacheFacadeProxy.client.Put(key, name, value)
	return err
}

func (cacheFacadeProxy *CacheFacadeProxy) Get(key, name string) (interface{}, error) {
	resp, err := cacheFacadeProxy.client.Get(key, name)
	if err != nil {
		return nil, err
	}
	return resp.GetNodeValue(), nil
}

func (cacheFacadeProxy *CacheFacadeProxy) Gets(key string) ([]interface{}, error) {
	resp, err := cacheFacadeProxy.client.Gets(key)
	if err != nil {
		return nil, err
	}
	ret := make([]interface{}, 0)
	for _, item := range resp {
		ret = append(ret, item.GetNodeValue())
	}
	return ret, nil
}

func (cacheFacadeProxy *CacheFacadeProxy) Remove(key, name string) error {
	_, err := cacheFacadeProxy.client.Remove(key, name)
	if err != nil {
		return err
	}
	return nil
}
