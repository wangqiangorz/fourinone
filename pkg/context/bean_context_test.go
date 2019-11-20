package context

import (
	"fmt"
	"testing"

	"github.com/wangqiangorz/fourinone/pkg/common"
	"github.com/wangqiangorz/fourinone/pkg/model"
	"github.com/wangqiangorz/fourinone/pkg/monitor"
)

func test(t *testing.T) {
	monitor.InitLog()
	err := common.SetConfFile("../../conf/conf.toml")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// go StartWorker("Test", func(input *model.Pikachu) (*model.Pikachu, error) {
	// 	fmt.Println(input.GetAny("test"))
	// 	return input, nil
	// })
	// time.Sleep(3 * time.Second)
	client, err := GetWork("test")
	if err != nil {
		fmt.Println(err.Error())
	}
	pikachu := model.NewPikachu()
	pikachu.SetValue("test", "t")
	reply, err := client.DoTask(pikachu)
	fmt.Println(reply.GetAny("test"))
	// // go StartPark()
	// client, err := GetPark()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// client.Create("Test", "a", "b", true)
	// val, err := client.GetNode("Test", "a")
	// fmt.Println(val)

}
