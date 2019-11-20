package park

import "github.com/wangqiangorz/fourinone/pkg/model"

type Park interface {
	Create(domain, node string, nodeValue interface{}, sessionid string, heartbeat bool) (*model.Ketchup, error)
	Get(domain, node string, sessionid string) (*model.Ketchup, error)
	Update(domain, node string, nodeValue interface{}, sessionid string, heartbeat bool) (*model.Ketchup, error)
	Delete(domain, node string, sessionid string) (*model.Ketchup, error)
	GetSessionID() (string, error)
	AskMaster() (string, error)
	AskLeader() (string, error)
	HeartBeat(domain, node string, sessionID string) (bool, error)
}

type ParkArgs struct {
	Domain    string
	Node      string
	Nodevalue interface{}
	Sessionid string
	Heartbeat bool
}
