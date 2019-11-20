package context

import (
	"github.com/wangqiangorz/fourinone/pkg/common"
	"github.com/wangqiangorz/fourinone/pkg/model"
	"github.com/wangqiangorz/fourinone/pkg/proxy"
	"github.com/wangqiangorz/fourinone/pkg/service"
	"github.com/wangqiangorz/fourinone/pkg/util"
	"github.com/wangqiangorz/fourinone/pkg/worker"
	logger "go.uber.org/zap"
)

func StartPark() error {
	groupserver := common.Cfg.Park.Servers
	serviceName := common.Cfg.Park.ServiceName
	if len(groupserver) == 0 {
		return ERRORCONFIGSETTING
	}
	host, port := util.GetHostFromString(groupserver[0])
	return startPark(serviceName, host, port, groupserver)
}

func StartParkByHost(serviceName, host string, port int) error {
	groupserver := common.Cfg.Park.Servers
	if len(groupserver) == 0 {
		return ERRORCONFIGSETTING
	}
	if host == "" || port == 0 {
		host, port = util.GetHostFromString(groupserver[0])
	}
	return startPark(serviceName, host, port, groupserver)
}

func startPark(serviceName, host string, port int, groupservers []string) error {
	err := service.PutBean(serviceName, host, port, service.NewParkService(serviceName, host, port, groupservers))
	if err != nil {
		logger.L().Warn("start park error", logger.String("serviceName", serviceName), logger.String("host", host), logger.Int("port", port), logger.Strings("groupservers", groupservers), logger.Error(err))
	}
	logger.L().Info("start park success", logger.String("serviceName", serviceName), logger.String("host", host), logger.Int("port", port), logger.Strings("groupservers", groupservers))
	return err
}

func SetConfigFile(filepath string) error {
	err := common.SetConfFile(filepath)
	return err
}

func GetPark() (*proxy.ParkProxy, error) {
	groupserver := common.Cfg.Park.Servers

	if len(groupserver) == 0 || common.Cfg.Park.ServiceName == "" {
		return nil, ERRORCONFIGSETTING
	}
	host, port := util.GetHostFromString(groupserver[0])
	return getParkByHost(common.Cfg.Park.ServiceName, host, port, groupserver)
}

func getPark(serviceName, host string, port int) (*proxy.ParkProxy, error) {
	groupserver := common.Cfg.Park.Servers
	if len(groupserver) == 0 {
		return nil, ERRORCONFIGSETTING
	}
	if host == "" || port == 0 {
		host, port = util.GetHostFromString(groupserver[0])

	}
	return getParkByHost(serviceName, host, port, groupserver)
}

func GetParkByHost(host string, port int) (*proxy.ParkProxy, error) {
	groupserver := common.Cfg.Park.Servers
	if len(groupserver) == 0 || common.Cfg.Park.ServiceName == "" {
		return nil, ERRORCONFIGSETTING
	}
	return getParkByHost(common.Cfg.Park.ServiceName, host, port, groupserver)
}

func getParkByHost(serviceName, host string, port int, groupservers []string) (*proxy.ParkProxy, error) {
	parkLocal, err := proxy.NewParkProxy(serviceName, host, port, groupservers)
	return parkLocal, err
}

func StartWorker(workerType string, task func(*model.Pikachu) (*model.Pikachu, error)) error {
	if common.Cfg.Worker.Server == "" {
		return ERRORCONFIGSETTING
	}
	host, port := util.GetHostFromString(common.Cfg.Worker.Server)
	return StartWorkerByHost(workerType, host, port, task)
}

func StartWorkerByHost(workerType, host string, port int, task func(*model.Pikachu) (*model.Pikachu, error)) error {
	err := CreateWorkType(workerType, util.GetStringFromHost(host, port))
	if err != nil {
		return err
	}
	err = service.PutBean(workerType, host, port, service.NewWorkerService(worker.NewWorkerMigrant(host, port, workerType, task)))
	if err != nil {
		logger.L().Warn("start worker error", logger.String("workerType", workerType), logger.String("host", host), logger.Int("port", port), logger.Error(err))
	} else {
		logger.L().Info("start worker success", logger.String("workerType", workerType), logger.String("host", host), logger.Int("port", port))

		if err != nil {
			return err
		}
	}
	return err
}

func GetWork(workerType string) (*proxy.WorkerProxy, error) {
	if common.Cfg.Worker.Server == "" {
		return nil, ERRORCONFIGSETTING
	}
	host, port := util.GetHostFromString(common.Cfg.Worker.Server)
	return GetWorkerByHost(workerType, host, port)
}

func GetWorkerByHost(workerType, host string, port int) (*proxy.WorkerProxy, error) {
	workerProxy, err := proxy.NewWorkerProxy(workerType, host, port)
	if err != nil {
		logger.L().Info("get worker fail", logger.String("workerType", workerType), logger.String("host", host), logger.Int("port", port), logger.Error(err))
	} else {
		logger.L().Info("get worker success", logger.String("workerType", workerType), logger.String("host", host), logger.Int("port", port))
	}

	return workerProxy, err
}

func StartCacheByHost(serviceName, host string, port int, cacheGroups []string) error {
	if host == "" || port == 0 {
		host, port = util.GetHostFromString(cacheGroups[0])
	}
	return startPark(serviceName, host, port, cacheGroups)
}

func StartCache() error {
	serviceName := common.Cfg.Cache.ServiceName
	groupserver := common.Cfg.Cache.Servers
	if len(groupserver) == 0 || common.Cfg.Cache.ServiceName == "" {
		return ERRORCONFIGSETTING
	}
	host, port := util.GetHostFromString(groupserver[0])
	return startPark(serviceName, host, port, groupserver)
}

func GetCache() (*proxy.ParkProxy, error) {
	groupserver := common.Cfg.Cache.Servers
	if len(groupserver) == 0 || common.Cfg.Cache.ServiceName == "" {
		return nil, ERRORCONFIGSETTING
	}
	host, port := util.GetHostFromString(groupserver[0])
	return getParkByHost(common.Cfg.Park.ServiceName, host, port, groupserver)
}

func GetCacheByHost(host string, port int) (*proxy.ParkProxy, error) {
	groupserver := common.Cfg.Park.Servers
	if len(groupserver) == 0 || common.Cfg.Park.ServiceName == "" {
		return nil, ERRORCONFIGSETTING
	}
	return getCacheByHost(common.Cfg.Park.ServiceName, host, port, groupserver)
}

func GetCahceByGroupServers(serviceName string, groupserver []string) (*proxy.ParkProxy, error) {
	if common.Cfg.Cache.ServiceName == "" {
		return nil, ERRORCONFIGSETTING
	}
	if len(groupserver) == 0 {
		return nil, ERRORGETGROUPSERVER
	}
	host, port := util.GetHostFromString(groupserver[0])
	return getCacheByHost(common.Cfg.Cache.ServiceName, host, port, groupserver)
}

func getCacheByHost(serviceName, host string, port int, groupservers []string) (*proxy.ParkProxy, error) {
	parkLocal, err := proxy.NewParkProxy(serviceName, host, port, groupservers)
	return parkLocal, err
}

func StartCacheFacade() error {
	serviceName := common.Cfg.CacheFacade.ServiceName
	if serviceName == "" || common.Cfg.CacheFacade.Server == "" {
		return ERRORCONFIGSETTING
	}
	cacheGroups := common.Cfg.CacheGroup.Servers
	if len(cacheGroups) == 0 {
		return ERRORCONFIGSETTING
	}
	host, port := util.GetHostFromString(common.Cfg.CacheFacade.Server)
	return startCacheFacadeByHost(serviceName, host, port, cacheGroups)
}

func startCacheFacadeByHost(serviceName, host string, port int, cacheGroups [][]string) error {
	tryNum := common.Cfg.CacheFacade.Trynum
	groupServers := common.Cfg.CacheGroup.Servers
	if len(groupServers) == 0 {
		return ERRORCONFIGSETTING
	}
	err := service.PutBean(serviceName, host, port, NewCacheFacade(serviceName, tryNum, groupServers))
	if err != nil {
		logger.L().Warn("start worker error", logger.String("serviceName", serviceName), logger.String("host", host), logger.Int("port", port), logger.Error(err))
	} else {
		logger.L().Info("start worker success", logger.String("serviceName", serviceName), logger.String("host", host), logger.Int("port", port))

	}
	return err
}

func GetCacheFacadeByHost(serviceName, host string, port int) (*proxy.CacheFacadeProxy, error) {
	proxy, err := proxy.NewCacheFacadeProxy(serviceName, host, port)
	return proxy, err
}

func GetCacheFacade() (*proxy.CacheFacadeProxy, error) {
	serviceName := common.Cfg.CacheFacade.ServiceName
	if serviceName == "" || common.Cfg.CacheFacade.Server == "" {
		return nil, ERRORCONFIGSETTING
	}
	host, port := util.GetHostFromString(common.Cfg.CacheFacade.Server)
	return GetCacheFacadeByHost(serviceName, host, port)
}
