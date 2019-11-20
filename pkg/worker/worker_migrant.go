package worker

import "github.com/wangqiangorz/fourinone/pkg/model"

type MigrantWorker struct {
	host     string
	port     int
	workType string
	task     func(*model.Pikachu) (*model.Pikachu, error)
	receive  func(interface{})
}

func NewWorkerMigrant(host string, port int, workType string, task func(*model.Pikachu) (*model.Pikachu, error)) *MigrantWorker {
	workerMigrant := new(MigrantWorker)
	workerMigrant.host = host
	workerMigrant.port = port
	workerMigrant.workType = workType
	workerMigrant.task = task
	return workerMigrant
}

func NewReceiver(host string, port int, workType string, receive func(interface{})) *MigrantWorker {
	receiver := new(MigrantWorker)
	receiver.host = host
	receiver.port = port
	receiver.workType = workType
	receiver.receive = receive
	return receiver
}

func (workerMigrant *MigrantWorker) Receive(input interface{}) {
	workerMigrant.receive(input)
}

func (workerMigrant *MigrantWorker) DoTask(argv *model.Pikachu) (*model.Pikachu, error) {
	return workerMigrant.task(argv)
}

func (workerMigrant *MigrantWorker) GetHost() string {
	return workerMigrant.host
}

func (workerMigrant *MigrantWorker) GetPort() int {
	return workerMigrant.port
}

func (workerMigrant *MigrantWorker) GetWorkType() string {
	return workerMigrant.workType
}

func (workerMigrant *MigrantWorker) GetTask() func(*model.Pikachu) (*model.Pikachu, error) {
	return workerMigrant.task
}

func (workerMigrant *MigrantWorker) SetHost(host string) {
	workerMigrant.host = host
}

func (workerMigrant *MigrantWorker) SetPort(port int) {
	workerMigrant.port = port
}

func (workerMigrant *MigrantWorker) SetWorkType(workType string) {
	workerMigrant.workType = workType
}

func (workerMigrant *MigrantWorker) SetTask(task func(*model.Pikachu) (*model.Pikachu, error)) {
	workerMigrant.task = task
}
