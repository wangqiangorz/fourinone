package model

import (
	"github.com/wangqiangorz/fourinone/pkg/util"
)

type BeanVal struct {
	Name      string
	Domain    string
	Node      string
	Nodevalue interface{}
}

func NewBeanVal(domain, node string, nodevalue interface{}) *BeanVal {
	beanval := new(BeanVal)
	beanval.Name = util.GetDomainNodeKey(domain, node)
	beanval.Domain = domain
	beanval.Node = node
	beanval.Nodevalue = nodevalue
	return beanval
}

func (beanVal *BeanVal) GetName() string {
	return beanVal.Name
}

func (beanVal *BeanVal) GetDomain() string {
	return beanVal.Domain
}

func (beanVal *BeanVal) GetNode() string {
	return beanVal.Node
}

func (beanVal *BeanVal) GetNodeValue() interface{} {
	return beanVal.Nodevalue
}
