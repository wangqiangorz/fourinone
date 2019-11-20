package common

import (
	"github.com/BurntSushi/toml"
)

var (
	confFile = "../conf/conf.toml"
	Cfg      Config
)

func InitConf() error {
	_, err := toml.DecodeFile(confFile, &Cfg)
	if err != nil {
		return err
	}
	return nil
}

func SetConfFile(confFile string) error {
	_, err := toml.DecodeFile(confFile, &Cfg)
	if err != nil {
		return err
	}
	return nil
}

type Config struct {
	Park struct {
		ServiceName string
		Servers     []string
		TryNum      int
	}
	Worker struct {
		ServiceName string
		Server      string
	}
	Cache struct {
		ServiceName string
		Servers     []string
	}
	CacheFacade struct {
		ServiceName string
		Server      string
		Trynum      int
	}
	CacheGroup struct {
		Servers [][]string
	}
	Common struct {
		ReadTimeout int
		ConnTimeout int
	}
}
