package contractor

import (
	"fmt"
	"testing"

	"github.com/wangqiangorz/fourinone/pkg/common"
	"github.com/wangqiangorz/fourinone/pkg/model"
	"github.com/wangqiangorz/fourinone/pkg/monitor"
)

func Test(t *testing.T) {
	common.SetConfFile("../../conf/conf.toml")
	monitor.InitLog()
	c := NewContractor(func(input *model.Pikachu) (*model.Pikachu, error) {
		workers, err := GetWaitingWorkerByWorkerType("demo")
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		i := 0
		for _, worker := range workers {
			input := model.NewPikachu()
			input.SetValue("test", i)
			reply, err := worker.DoTask(input)
			i++
			if err != nil {
				fmt.Println(err.Error())
				return nil, err
			}
			fmt.Println(reply.GetAny("test"))
		}
		return nil, nil
	})
	c.Run(nil)

}
