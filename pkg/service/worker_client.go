package service

import (
	"net/rpc"

	"github.com/wangqiangorz/fourinone/pkg/model"
)

type WorkerClient struct {
	host     string
	port     int
	workType string
	client   *rpc.Client
}

func NewWorkerClient(workType, host string, port int) (*WorkerClient, error) {
	workerClient := new(WorkerClient)
	workerClient.workType = workType
	workerClient.host = host
	workerClient.port = port
	_, err := workerClient.getRpcClient()
	return workerClient, err
}

func (workerClient *WorkerClient) getRpcClient() (*rpc.Client, error) {
	if workerClient.client == nil {
		client, err := GetBean(workerClient.host, workerClient.port)
		return client, err
	}
	return workerClient.client, nil
}

func (workerClient *WorkerClient) Send(msg interface{}) (bool, error) {
	var ret *bool
	client, err := workerClient.getRpcClient()
	if err != nil {
		return false, err
	}
	err = client.Call(workerClient.workType+".Receive", &msg, &ret)
	return *ret, err
}

func (workerClient *WorkerClient) DoTask(pikachu *model.Pikachu) (*model.Pikachu, error) {
	var ret *model.Pikachu
	client, err := workerClient.getRpcClient()
	if err != nil {
		return nil, err
	}
	err = client.Call(workerClient.workType+".DoTask", pikachu, &ret)
	return ret, err
}

func (workerClient *WorkerClient) SetTask(task func(*model.Pikachu) (*model.Pikachu, error)) error {
	var ret *model.Pikachu
	client, err := workerClient.getRpcClient()
	if err != nil {
		return err
	}
	err = client.Call(workerClient.workType+".SetTask", task, &ret)
	return err
}

func (workerClient *WorkerClient) Close() error {
	return workerClient.client.Close()
}
