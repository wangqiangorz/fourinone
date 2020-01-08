package service

import (
	"fmt"

	"github.com/wangqiangorz/fourinone/pkg/common"
	"github.com/wangqiangorz/fourinone/pkg/util"
	logger "go.uber.org/zap"
)

type ParkLeader struct {
	serviceName string
	isMaster    bool
	host        string
	port        int
	groupServer []string
	tryNum      int
}

func NewParkLeader(serviceName, host string, port int, groupserver []string) *ParkLeader {
	parkLeader := new(ParkLeader)
	parkLeader.serviceName = serviceName
	parkLeader.isMaster = false
	parkLeader.host = host
	parkLeader.port = port
	parkLeader.groupServer = groupserver
	parkLeader.tryNum = common.Cfg.Park.TryNum
	return parkLeader
}

func (parkLeader *ParkLeader) IsMaster() string {
	if parkLeader.isMaster {
		return parkLeader.host + fmt.Sprintf(":%d", parkLeader.port)
	}
	return ""
}

func (parkLeader *ParkLeader) setMaster(isMaster bool) {
	logger.L().Info("set Park Master", logger.String("host", parkLeader.host), logger.Int("port", parkLeader.port))
	parkLeader.isMaster = isMaster
}

func (parkLeader *ParkLeader) CheckMasterPark() string {
	if parkLeader.isMaster {
		return ""
	}
	otherpark := parkLeader.getOtherMasterPark()
	if otherpark == nil {
		parkLeader.setMaster(true)
		return ""
	} else {
		defer otherpark.Close()
	}
	return otherpark.GetHost() + fmt.Sprintf(":%d", otherpark.GetPort())
}

func (parkLeader *ParkLeader) getOtherMasterPark() *ParkClient {
	parkSlice := parkLeader.getOtherPark()
	var ret *ParkClient
	for _, parkClient := range parkSlice {
		router, err := parkClient.AskMaster()
		if err != nil {
			logger.L().Warn("AskMaster error", logger.String("Host", parkClient.GetHost()), logger.Int("Port", parkClient.GetPort()), logger.Error(err))
		}
		if router != "" {
			ret = parkClient
			break
		}
	}
	if ret != nil {
		for _, parkClient := range parkSlice {
			if parkClient != ret {
				parkClient.Close()
			}
		}
	}
	return ret
}

func (parkLeader *ParkLeader) getOtherPark() []*ParkClient {
	var ret []*ParkClient
	for _, server := range parkLeader.groupServer {
		host, port := util.GetHostFromString(server)
		if host == parkLeader.host && port == parkLeader.port {
			continue
		}
		var otherPark *ParkClient
		otherPark, err := NewParkClient(parkLeader.serviceName, host, port)
		if err != nil {
			logger.L().Warn("get rpc connect error", logger.String("serviceName", parkLeader.serviceName), logger.String("host", host), logger.Int("port", port), logger.Error(err))
		} else {
			ret = append(ret, otherPark)
		}

	}
	return ret
}

func (parkLeader *ParkLeader) GetParkLeader() (*ParkClient, error) {
	client, err := parkLeader.TryGetLeader(parkLeader.host, parkLeader.port)
	if err == nil {
		return client, nil
	}
	for _, server := range parkLeader.groupServer {
		host, port := util.GetHostFromString(server)
		if host == parkLeader.host && port == parkLeader.port {
			continue
		}
		client, err = parkLeader.TryGetLeader(host, port)
		if err == nil {
			return client, nil
		}
	}
	return nil, ERRORGETLEADER
}

func (parkLeader *ParkLeader) TryGetLeader(host string, port int) (*ParkClient, error) {
	parkClient, err := NewParkClient(parkLeader.serviceName, host, port)
	if err != nil {
		logger.L().Warn("get rpc connect error", logger.String("serviceName", parkLeader.serviceName), logger.String("Host", host), logger.Int("Port", port), logger.Error(err))
		return nil, err
	}
	leader, err := parkClient.AskLeader()
	if err != nil {
		logger.L().Warn("AskLeader error", logger.String("serviceName", parkLeader.serviceName), logger.String("Host", host), logger.Int("Port", port), logger.Error(err))
		return nil, err
	}
	if leader == "" {
		return parkClient, nil
	} else {
		parkClient.Close()
		host, port := util.GetHostFromString(leader)
		ret, err := NewParkClient(parkLeader.serviceName, host, port)
		if err != nil {
			logger.L().Warn("get rpc connect error", logger.String("serviceName", parkLeader.serviceName), logger.String("Host", parkClient.GetHost()), logger.Int("Port", parkClient.GetPort()), logger.Error(err))
			return nil, err
		}
		return ret, nil
	}
}

func (parkLeader *ParkLeader) TryBeMaster() {
	logger.L().Info("park want to be master", logger.String("serviceName", parkLeader.serviceName), logger.String("host", parkLeader.host), logger.Int("port", parkLeader.port))
	otherMaster := parkLeader.getOtherMasterPark()
	if otherMaster == nil {
		parkLeader.setMaster(true)
	} else {
		logger.L().Info("groupserve already has master", logger.String("serviceName", parkLeader.serviceName), logger.String("masterhost", otherMaster.GetHost()), logger.Int("masterport", otherMaster.GetPort()))
	}
}

func (parkLeader *ParkLeader) GetHost() string {
	return parkLeader.host
}

func (parkLeader *ParkLeader) GetPort() int {
	return parkLeader.port
}

func (parkLeader *ParkLeader) GetServiceName() string {
	return parkLeader.serviceName
}
