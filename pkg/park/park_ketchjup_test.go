package park

import (
	"fmt"
	"testing"

	"github.com/wangqiangorz/fourinone/pkg/common"
)

func Test(t *testing.T) {
	fmt.Println(common.Cfg)
	parkKetchup := NewParkKetchup()
	parkKetchup.SetValue("123.abc", "a")
	parkKetchup.SetValue("123.abd", "b")
	fmt.Println(parkKetchup.GetNode("123", "abc"))
	fmt.Println(parkKetchup.GetAllNodeByDomain("123"))

}
