package proxy

import (
	"github.com/wangqiangorz/fourinone/pkg/model"
	"github.com/wangqiangorz/fourinone/pkg/service"
)

type WorkerProxy struct {
	workType   string
	host       string
	port       int
	workClient *service.WorkerClient
}

func NewWorkerProxy(workType, host string, port int) (*WorkerProxy, error) {
	var err error
	workProxy := new(WorkerProxy)
	workProxy.workType = workType
	workProxy.host = host
	workProxy.workClient, err = service.NewWorkerClient(workType, host, port)
	if err != nil {
		return nil, err
	}
	return workProxy, nil
}

func (workerProxy *WorkerProxy) GetWorkerProxy() string {
	return workerProxy.workType
}

func (workerProxy *WorkerProxy) GetHost() string {
	return workerProxy.host
}

func (workerProxy *WorkerProxy) GetPort() int {
	return workerProxy.port
}

func (workerProxy *WorkerProxy) GetWorkClient() *service.WorkerClient {
	return workerProxy.workClient
}

func (workerProxy *WorkerProxy) Send(input interface{}) (bool, error) {
	reply, err := workerProxy.workClient.Send(input)
	if err != nil {
		return false, err
	}
	return reply, nil
}

func (workerProxy *WorkerProxy) DoTask(pikachu *model.Pikachu) (*model.Pikachu, error) {
	reply, err := workerProxy.workClient.DoTask(pikachu)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (workerProxy *WorkerProxy) SetTask(task func(*model.Pikachu) (*model.Pikachu, error)) error {
	err := workerProxy.workClient.SetTask(task)
	return err
}

func (workerProxy *WorkerProxy) Close() error {
	err := workerProxy.workClient.Close()
	return err
}
