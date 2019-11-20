package context

import (
	"fmt"
	"testing"
	"time"

	"github.com/wangqiangorz/fourinone/pkg/common"
	"github.com/wangqiangorz/fourinone/pkg/monitor"
)

func Test(t *testing.T) {
	monitor.InitLog()
	err := common.SetConfFile("../../conf/conf.toml")
	// go StartCacheByHost("cacheService", "127.0.0.1", 8085, []string{"127.0.0.1:8085"})
	// go StartCacheByHost("cacheService", "127.0.0.1", 8086, []string{"127.0.0.1:8086"})
	// go StartCacheByHost("cacheService", "127.0.0.1", 8087, []string{"127.0.0.1:8087"})
	go StartCacheFacade()
	time.Sleep(3 * time.Second)
	proxy, err := GetCacheFacade()
	if err != nil {
		fmt.Println(err.Error())
	}
	key, err := proxy.Add("123", "456")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(key)
	reply, err := proxy.Get(key, "123")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(reply)
}
