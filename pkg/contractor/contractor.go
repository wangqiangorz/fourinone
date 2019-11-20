package contractor

import (
	"github.com/wangqiangorz/fourinone/pkg/context"
	"github.com/wangqiangorz/fourinone/pkg/model"
	"github.com/wangqiangorz/fourinone/pkg/proxy"
	"github.com/wangqiangorz/fourinone/pkg/util"

	logger "go.uber.org/zap"
)

type Contractor struct {
	nextContractor *Contractor
	giveTask       func(*model.Pikachu) (*model.Pikachu, error)
}

func NewContractor(giveTask func(*model.Pikachu) (*model.Pikachu, error)) *Contractor {
	contractor := new(Contractor)
	contractor.giveTask = giveTask
	return contractor
}

func (contractor *Contractor) NextContractor(next *Contractor) *Contractor {
	head := contractor
	for head.nextContractor != nil {
		head = head.nextContractor
	}
	head.nextContractor = next
	return contractor
}

func (contractor *Contractor) Run(input *model.Pikachu) (*model.Pikachu, error) {
	reply, err := contractor.giveTask(input)
	if err != nil {
		return nil, err
	}
	for contractor.nextContractor != nil {
		next := contractor.nextContractor
		reply, err = next.Run(reply)
		if err != nil {
			return nil, err
		}
	}
	return reply, nil
}

func GetWaitingWorkerByWorkerType(workerType string) ([]*proxy.WorkerProxy, error) {
	workerList, err := context.GetWorkerByWorkerType(workerType)
	if err != nil {
		return nil, err
	}
	workers := make([]*proxy.WorkerProxy, 0)
	for _, v := range workerList {
		host, port := util.GetHostFromString(v.GetNodeValue().(string))
		workerClient, err := context.GetWorkerByHost(workerType, host, port)
		if err != nil {
			logger.L().Warn("getWorkerError", logger.String("workerType", workerType), logger.String("host", host), logger.Int("port", port), logger.Error(err))
		} else {
			workers = append(workers, workerClient)
		}
	}
	return workers, nil
}

func GetWaitingWorker(parkHost string, parkPort int, workerType string) ([]*proxy.WorkerProxy, error) {
	workerList, err := context.GetWorkerByWorkerTypeAndHost(parkHost, parkPort, workerType)
	if err != nil {
		return nil, err
	}
	workers := make([]*proxy.WorkerProxy, len(workerList))
	for _, v := range workerList {
		host, port := util.GetHostFromString(v.GetNodeValue().(string))
		workerClient, err := context.GetWorkerByHost(workerType, host, port)
		if err != nil {
			logger.L().Warn("getWorkerError", logger.String("workerType", workerType), logger.String("host", host), logger.Int("port", port), logger.Error(err))
		} else {
			workers = append(workers, workerClient)
		}
	}
	return workers, nil
}
