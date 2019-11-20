package service

import (
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/wangqiangorz/fourinone/pkg/model"
	"github.com/wangqiangorz/fourinone/pkg/park"
	"github.com/wangqiangorz/fourinone/pkg/util"
	logger "go.uber.org/zap"
)

type ParkService struct {
	parkKetchup *park.ParkKetchup
	m           *sync.RWMutex
	parkLeader  *ParkLeader
	hbinfo      *model.Ketchup
}

func NewParkService(serviceName, host string, port int, servers []string) *ParkService {
	parkService := new(ParkService)
	parkService.parkKetchup = park.NewParkKetchup()
	parkService.m = new(sync.RWMutex)
	parkService.parkLeader = NewParkLeader(serviceName, host, port, servers)
	parkService.parkLeader.TryBeMaster()
	return parkService
}

func (parkService *ParkService) Create(args *park.ParkArgs, reply *model.Ketchup) error {
	if args.Domain != "" && args.Node != "" {
		parkService.m.Lock()
		domainNodeKey := util.GetDomainNodeKey(args.Domain, args.Node)
		if !parkService.parkKetchup.ContainsKey(domainNodeKey) {
			if !parkService.parkKetchup.ContainsKey(args.Domain) {
				parkService.parkKetchup.SetValue(args.Domain, 0)
			}
			parkService.parkKetchup.SetValue(park.SESSIONPREFIX+domainNodeKey, args.Sessionid)
			parkService.parkKetchup.SetValue(park.TIMEPREFIX+domainNodeKey, util.GetTimeStamp())
			parkService.parkKetchup.SetValue(domainNodeKey, args.Nodevalue)
			parkService.m.Unlock()
			logger.L().Info("Create Node Success!",
				logger.String("domain", args.Domain),
				logger.String("node", args.Node),
				logger.Any("nodeValue", args.Nodevalue))
			parkService.Get(args, reply)

		} else {
			logger.L().Info("Create Node Failed! Park already contains key",
				logger.String("Key", domainNodeKey))
			return KEYEXIST
		}
	}
	return nil
}

func (parkService *ParkService) Update(args *park.ParkArgs, reply *model.Ketchup) error {
	if args.Domain != "" && args.Node != "" {
		parkService.m.Lock()
		domainNodeKey := util.GetDomainNodeKey(args.Domain, args.Node)
		if parkService.parkKetchup.ContainsKey(domainNodeKey) {
			parkService.parkKetchup.SetValue(park.SESSIONPREFIX+domainNodeKey, args.Sessionid)
			parkService.parkKetchup.SetValue(park.TIMEPREFIX+domainNodeKey, util.GetTimeStamp())
			parkService.parkKetchup.SetValue(domainNodeKey, args.Nodevalue)
			parkService.m.Unlock()
			logger.L().Info("Update Node Success!",
				logger.String("domain", args.Domain),
				logger.String("node", args.Node))
			parkService.Get(args, reply)
		} else {
			logger.L().Info("Update Node Failed! Park doesn't contains key",
				logger.String("Key", domainNodeKey))
			return KEYNOTEXIST
		}
	}
	return nil
}

func (parkService *ParkService) Get(args *park.ParkArgs, reply *model.Ketchup) error {
	if args.Domain == "" {
		return nil
	}
	parkService.m.RLock()
	defer parkService.m.RUnlock()
	domainNodeKey := util.GetDomainNodeKey(args.Domain, args.Node)
	if parkService.parkKetchup.ContainsKey(domainNodeKey) {
		*reply = *parkService.parkKetchup.GetNode(args.Domain, args.Node)
		logger.L().Info("Get Node Success!",
			logger.String("domain", args.Domain),
			logger.String("node", args.Node))
		return nil
	} else {
		logger.L().Warn("Get Node Failed! The node is not exist!",
			logger.String("domain", args.Domain),
			logger.String("node", args.Node))
		return KEYNOTEXIST
	}
}

func (parkService *ParkService) Delete(args *park.ParkArgs, reply *model.Ketchup) error {
	parkService.m.Lock()
	defer parkService.m.Unlock()
	domainNodeKey := util.GetDomainNodeKey(args.Domain, args.Node)
	if parkService.parkKetchup.ContainsKey(domainNodeKey) {
		*reply = *parkService.parkKetchup.DeleteNode(args.Domain, args.Node)
		logger.L().Info("Delete Node Success!",
			logger.String("domain", args.Domain),
			logger.String("node", args.Node))
		return nil
	} else {
		logger.L().Warn("Delete Node Failed! The node is not exist!",
			logger.String("domain", args.Domain),
			logger.String("node", args.Node))
		return KEYNOTEXIST
	}
}

func (parkService *ParkService) GetSessionID(args *park.ParkArgs, reply *string) error {
	*reply = park.SESSIONPREFIX + uuid.NewV4().String()
	logger.L().Info("Get Sessionid Success!",
		logger.String("Sessionid", *reply))
	return nil
}

func (parkService *ParkService) AskMaster(args *park.ParkArgs, reply *string) error {
	*reply = parkService.parkLeader.IsMaster()
	if *reply != "" {
		logger.L().Info("AskMaster request",
			logger.String("serviceName", parkService.parkLeader.GetServiceName()),
			logger.String("host", parkService.parkLeader.GetHost()),
			logger.Int("port", parkService.parkLeader.GetPort()),
			logger.Bool("isMaster", true))
	}
	logger.L().Info("AskMaster request",
		logger.String("serviceName", parkService.parkLeader.GetServiceName()),
		logger.String("host", parkService.parkLeader.GetHost()),
		logger.Int("port", parkService.parkLeader.GetPort()),
		logger.Bool("isMaster", false))
	return nil
}

func (parkService *ParkService) AskLeader(args *park.ParkArgs, reply *string) error {
	*reply = parkService.parkLeader.CheckMasterPark()
	if *reply != "" {
		host, port := util.GetHostFromString(*reply)
		logger.L().Info("AskLeader request",
			logger.String("leaderhost", host),
			logger.Int("leaderport", port))
	}
	logger.L().Info("AskLeader request",
		logger.String("leaderhost", parkService.parkLeader.GetHost()),
		logger.Int("leaderport", parkService.parkLeader.GetPort()))
	return nil
}

func (parkService *ParkService) HearBeat(args *park.ParkArgs, reply *bool) error {
	domainNodeKey := util.GetDomainNodeKey(args.Domain, args.Node)
	parkService.hbinfo.SetValue(domainNodeKey, time.Now().Unix())
	return nil
}
