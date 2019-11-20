package cache

import "github.com/wangqiangorz/fourinone/pkg/model"

type Cache interface {
	Add(string, interface{}) (string, error)
	Put(string, string, interface{}) (*model.BeanVal, error)
	Get(string, string) (*model.BeanVal, error)
	Gets(string) ([]*model.BeanVal, error)
	Remove(string, string) (*model.BeanVal, error)
	Removes(string) ([]*model.BeanVal, error)
}
