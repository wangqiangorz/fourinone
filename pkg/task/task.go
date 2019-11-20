package task

import "github.com/wangqiangorz/fourinone/pkg/model"

type Task interface {
	Do(model.Pikachu) model.Pikachu
}
