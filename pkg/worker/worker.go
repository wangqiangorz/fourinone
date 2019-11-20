package worker

import "github.com/wangqiangorz/fourinone/pkg/model"

type Worker interface {
	SetTask(task func(*model.Pikachu) (*model.Pikachu, error)) error
	DoTask(*model.Pikachu) (*model.Pikachu, error)
}
