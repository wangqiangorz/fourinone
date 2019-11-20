package main

import (
	"github.com/wangqiangorz/fourinone/pkg/common"
	"github.com/wangqiangorz/fourinone/pkg/context"
	"github.com/wangqiangorz/fourinone/pkg/monitor"
)

// func main() {
// 	common.InitConf()
// 	monitor.InitLog()
// 	go context.StartPark()
// 	time.Sleep(3 * time.Second)
// 	err := context.StartWorker("test", func(input *model.Pikachu) (*model.Pikachu, error) {
// 		fmt.Println(input.GetAny("test"))
// 		input.SetValue("test", "new")
// 		return input, nil
// 	})
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	context.StartWorkerByHost("test", "127.0.0.1", 8088, func(input *model.Pikachu) (*model.Pikachu, error) {
// 		fmt.Println(input.GetAny("test"))
// 		input.SetValue("test", "new")
// 		return input, nil
// 	})
// }

func main() {
	common.InitConf()
	monitor.InitLog()
	context.StartCacheByHost("cacheService", "127.0.0.1", 8085, []string{"127.0.0.1:8085"})
	// context.StartCacheByHost("cacheService", "127.0.0.1", 8086, []string{"127.0.0.1:8086"})
	// context.StartCacheByHost("cacheService", "127.0.0.1", 8087, []string{"127.0.0.1:8087"})
}
