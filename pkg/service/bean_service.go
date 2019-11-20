package service

import (
	"fmt"
	"net"
	"net/rpc"
	"time"

	"github.com/wangqiangorz/fourinone/pkg/common"
	"github.com/wangqiangorz/fourinone/pkg/util"
	logger "go.uber.org/zap"
)

func PutBean(serviceName, host string, port int, exportObj interface{}) error {

	err := rpc.RegisterName(serviceName, exportObj)
	if err != nil {
		logger.L().Warn("Register RPC service error", logger.String("host", host), logger.Int("port", port), logger.Error(err))
		return err
	}
	tcpaddr, err := net.ResolveTCPAddr("tcp4", host+fmt.Sprintf(":%d", port))
	if err != nil {
		logger.L().Warn("Resolve tcpaddr error", logger.String("host", host), logger.Int("port", port), logger.Error(err))
		return err
	}
	tcplistener, err := net.ListenTCP("tcp", tcpaddr)
	if err != nil {
		logger.L().Warn("net listenTcp error", logger.String("host", host), logger.Int("port", port), logger.Error(err))
		return err
	}
	for {
		conn, err := tcplistener.Accept()
		if err != nil {
			logger.L().Warn("Accept tcp error", logger.String("host", host), logger.Int("port", port), logger.Error(err))
			continue
		}
		go rpc.ServeConn(conn)
	}
}

func GetBean(host string, port int) (*rpc.Client, error) {
	readTimeout := time.Millisecond * 500
	connTimeout := time.Millisecond * 200
	if common.Cfg.Common.ReadTimeout != 0 {
		readTimeout = time.Duration(common.Cfg.Common.ReadTimeout) * time.Millisecond
	}
	if common.Cfg.Common.ConnTimeout != 0 {
		readTimeout = time.Duration(common.Cfg.Common.ConnTimeout) * time.Millisecond
	}
	conn, err := net.DialTimeout("tcp", util.GetStringFromHost(host, port), connTimeout)
	if err != nil {
		logger.L().Warn("dail net error", logger.String("host", host), logger.Int("port", port), logger.Error(err))
		return nil, err
	}
	err = conn.SetDeadline(time.Now().Add(readTimeout))
	client := rpc.NewClient(conn)
	return client, err
}
