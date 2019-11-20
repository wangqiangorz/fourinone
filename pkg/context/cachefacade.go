package context

import (
	"github.com/wangqiangorz/fourinone/pkg/model"
	"github.com/wangqiangorz/fourinone/pkg/proxy"
	"github.com/wangqiangorz/fourinone/pkg/util"

	logger "go.uber.org/zap"
)

type CacheFacade struct {
	serviceName     string
	cacheGroups     [][]string
	cacheServerList []*proxy.ParkProxy
	tryNum          int
}

func NewCacheFacade(serviceName string, tryNum int, cacheGroups [][]string) *CacheFacade {
	cacheFacade := new(CacheFacade)
	cacheFacade.serviceName = serviceName
	cacheFacade.tryNum = tryNum
	cacheFacade.cacheGroups = cacheGroups
	cacheFacade.cacheServerList = make([]*proxy.ParkProxy, len(cacheGroups))
	return cacheFacade
}

func (cacheFacade *CacheFacade) getCacheFromKey(key string) (*proxy.ParkProxy, error) {
	index := util.HashCode(key) % len(cacheFacade.cacheGroups)
	cacheServerGroup := cacheFacade.cacheGroups[index]
	if cacheFacade.cacheServerList[index] != nil {
		return cacheFacade.cacheServerList[index], nil
	}
	return GetCahceByGroupServers(cacheFacade.serviceName, cacheServerGroup)
}

func (cacheFacade *CacheFacade) Add(argv *model.CacheFacadeArgv, reply *string) error {
	key := ""
	i := 0
	var err error
	for i <= cacheFacade.tryNum {
		key = util.GetUUid()
		proxy, err := cacheFacade.getCacheFromKey(key)
		if err != nil {
			logger.L().Warn("error getCacheFromkey", logger.Error(err))
			continue
		}
		err = proxy.Create(key, argv.Name, argv.Value, true)
		if err != nil {
			logger.L().Warn("error create cache", logger.String("host", proxy.GetServiceHost()), logger.Int("port", proxy.GetServicePort()), logger.Error(err))
			continue
		} else {
			break
		}
	}
	if i > cacheFacade.tryNum {
		*reply = ""
		return err
	}
	*reply = key
	return nil
}

func (cacheFacade *CacheFacade) Put(argv *model.CacheFacadeArgv, reply *bool) error {
	proxy, err := cacheFacade.getCacheFromKey(argv.Key)
	err = proxy.Update(argv.Key, argv.Name, argv.Value, true)
	if err != nil {
		logger.L().Warn("error put cache", logger.String("host", proxy.GetServiceHost()), logger.Int("port", proxy.GetServicePort()), logger.Error(err))
		*reply = false
		return err
	}
	return nil
}

func (cacheFacade *CacheFacade) Get(argv *model.CacheFacadeArgv, reply *model.BeanVal) error {
	proxy, err := cacheFacade.getCacheFromKey(argv.Key)
	resp, err := proxy.GetNode(argv.Key, argv.Name)
	if err != nil {
		logger.L().Warn("error get cache", logger.String("host", proxy.GetServiceHost()), logger.Int("port", proxy.GetServicePort()), logger.Error(err))
		return err
	}
	*reply = *resp
	return nil
}

func (cacheFacade *CacheFacade) Gets(argv *model.CacheFacadeArgv, reply *[]*model.BeanVal) error {
	proxy, err := cacheFacade.getCacheFromKey(argv.Key)
	resp, err := proxy.GetNodes(argv.Key)
	if err != nil {
		logger.L().Warn("error get cache", logger.String("host", proxy.GetServiceHost()), logger.Int("port", proxy.GetServicePort()), logger.Error(err))
		return err
	}
	*reply = resp
	return nil
}

func (cacheFacade *CacheFacade) Remove(argv *model.CacheFacadeArgv, reply *bool) error {
	proxy, err := cacheFacade.getCacheFromKey(argv.Key)
	err = proxy.Delete(argv.Key, argv.Name)
	if err != nil {
		logger.L().Warn("error get cache", logger.String("host", proxy.GetServiceHost()), logger.Int("port", proxy.GetServicePort()), logger.Error(err))
		*reply = false
		return err
	}
	*reply = true
	return nil
}
