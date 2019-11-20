package proxy

import (
	"github.com/wangqiangorz/fourinone/pkg/model"
	"github.com/wangqiangorz/fourinone/pkg/service"
	"github.com/wangqiangorz/fourinone/pkg/util"
	logger "go.uber.org/zap"
)

type ParkProxy struct {
	sid         string
	parkLeader  *service.ParkLeader
	parkClient  *service.ParkClient
	serviceName string
	host        string
	port        int
	groupserver []string
}

func NewParkProxy(serviceName, host string, port int, groupserver []string) (parkProxy *ParkProxy, err error) {
	parkProxy = new(ParkProxy)
	parkProxy.host = host
	parkProxy.port = port
	parkProxy.groupserver = groupserver
	parkProxy.serviceName = serviceName
	parkProxy.parkLeader = service.NewParkLeader(serviceName, host, port, groupserver)
	parkProxy.parkClient, err = parkProxy.parkLeader.GetParkLeader()
	if err != nil {
		logger.L().Info("get park leader error", logger.String("serviceName", serviceName), logger.String("host", host), logger.Int("port", port), logger.Strings("groupservers", groupserver), logger.Error(err))
		return nil, err
	}
	err = parkProxy.init()
	return parkProxy, err
}

func (parkProxy *ParkProxy) init() (err error) {
	if parkProxy.sid == "" {
		parkProxy.sid, err = parkProxy.parkClient.GetSessionID()
		return err
	}
	return nil
}

func (parkProxy *ParkProxy) Close() error {
	return parkProxy.parkClient.Close()
}

func (parkProxy *ParkProxy) Create(domain, node string, nodeValue interface{}, heartbeat bool) error {
	if domain != "" && node != "" {
		_, err := parkProxy.parkClient.Create(domain, node, nodeValue, parkProxy.sid, heartbeat)
		return err
	}
	return PARAMETERABSENT
}

func (parkProxy *ParkProxy) Update(domain, node string, nodeValue interface{}, heartbeat bool) error {
	if domain != "" && node != "" {
		_, err := parkProxy.parkClient.Update(domain, node, nodeValue, parkProxy.sid, heartbeat)
		return err
	}
	return PARAMETERABSENT
}

func (parkProxy *ParkProxy) Delete(domain, node string) error {
	if domain != "" && node != "" {
		_, err := parkProxy.parkClient.Delete(domain, node, parkProxy.sid)
		return err
	}
	return PARAMETERABSENT
}

func (parkProxy *ParkProxy) GetNode(domain, node string) (*model.BeanVal, error) {
	if domain != "" && node != "" {
		reply, err := parkProxy.parkClient.Get(domain, node, parkProxy.sid)
		if err != nil {
			return nil, err
		}
		return Ketchup2BeanVal(reply, domain, node), nil
	}
	return nil, PARAMETERABSENT
}

func (parkProxy *ParkProxy) GetNodes(domain string) ([]*model.BeanVal, error) {
	if domain != "" {
		reply, err := parkProxy.parkClient.Get(domain, "", parkProxy.sid)
		if err != nil {
			return nil, err
		}
		return Ketchup2BeanValList(reply), nil
	}
	return nil, PARAMETERABSENT
}

func (parkProxy *ParkProxy) GetServiceName() string {
	return parkProxy.serviceName
}

func (parkProxy *ParkProxy) GetServiceHost() string {
	return parkProxy.host
}

func (parkProxy *ParkProxy) GetServicePort() int {
	return parkProxy.port
}

func (parkProxy *ParkProxy) GetServiceGroupServer() []string {
	return parkProxy.groupserver
}

func Ketchup2BeanVal(ketchup *model.Ketchup, domain, node string) *model.BeanVal {
	nodevalue, ok := ketchup.GetAny(util.GetDomainNodeKey(domain, node))
	if !ok {
		return nil
	}
	beanval := model.NewBeanVal(domain, node, nodevalue)
	return beanval
}

func Ketchup2BeanValList(ketchup *model.Ketchup) []*model.BeanVal {
	ret := make([]*model.BeanVal, 0)
	for _, domainNodeKey := range ketchup.GetAllKey() {
		domain, node := util.GetDomainNodeFromKey(domainNodeKey.(string))
		ret = append(ret, Ketchup2BeanVal(ketchup, domain, node))
	}
	return ret
}
