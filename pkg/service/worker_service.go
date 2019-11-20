package service

import (
	"github.com/wangqiangorz/fourinone/pkg/model"
	"github.com/wangqiangorz/fourinone/pkg/worker"
)

type WorkerService struct {
	migrantWorker *worker.MigrantWorker
}

func NewWorkerService(migrantWorker *worker.MigrantWorker) *WorkerService {
	workerService := new(WorkerService)
	workerService.migrantWorker = migrantWorker
	return workerService
}

func (workService *WorkerService) SetTask(task func(*model.Pikachu) (*model.Pikachu, error), isok *bool) error {
	workService.migrantWorker.SetTask(task)
	*isok = true
	return nil
}

func (workService *WorkerService) DoTask(input *model.Pikachu, reply *model.Pikachu) error {
	ret, err := workService.migrantWorker.DoTask(input)
	*reply = *ret
	return err
}

func (workService *WorkerService) Receive(message *interface{}, reply *bool) error {
	workService.migrantWorker.Receive(*message)
	*reply = true
	return nil
}
