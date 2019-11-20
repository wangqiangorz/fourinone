package service

import (
	"fmt"
	"testing"

	"github.com/wangqiangorz/fourinone/pkg/model"
	"github.com/wangqiangorz/fourinone/pkg/park"
)

func Test(t *testing.T) {
	parkService := NewParkService("parkService", "localhost", 2018, []string{"localhost:8081"})
	var reply model.Ketchup
	parkService.Create(&park.ParkArgs{
		Domain:    "test",
		Node:      "123",
		Nodevalue: "localhost:8081",
	}, &reply)
	parkService.Get(&park.ParkArgs{
		Domain:    "test",
		Node:      "123",
		Nodevalue: "localhost:8081",
	}, &reply)
	// parkService.Delete(&park.ParkArgs{
	// 	Domain:    "test",
	// 	Node:      "123",
	// 	Nodevalue: "localhost:8081",
	// }, &reply)
	// fmt.Println(reply)
	parkService.Get(&park.ParkArgs{
		Domain:    "test",
		Node:      "123",
		Nodevalue: "localhost:8081",
	}, &reply)
	fmt.Println(reply)
	parkService.Get(&park.ParkArgs{
		Domain:    "test",
		Node:      "123",
		Nodevalue: "localhost:8081",
	}, &reply)
	fmt.Println(reply)
	var test string
	parkService.AskMaster(&park.ParkArgs{
		Domain:    "test",
		Node:      "123",
		Nodevalue: "localhost:8081",
	}, &test)
	fmt.Println(test)
	parkService.GetSessionID(&park.ParkArgs{
		Domain:    "test",
		Node:      "123",
		Nodevalue: "localhost:8081",
	}, &test)
	fmt.Println(test)

}
