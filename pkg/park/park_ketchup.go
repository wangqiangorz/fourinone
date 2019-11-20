package park

import (
	"strings"

	"github.com/wangqiangorz/fourinone/pkg/model"
	"github.com/wangqiangorz/fourinone/pkg/util"
)

type ParkKetchup struct {
	model.Ketchup
}

func NewParkKetchup() *ParkKetchup {
	parkKetchup := new(ParkKetchup)
	parkKetchup.Init()
	return parkKetchup
}
func (parkKetchup *ParkKetchup) GetNode(domain, node string) *model.Ketchup {
	if node == "" {
		return parkKetchup.GetAllNodeByDomain(domain)
	}
	ret := model.NewKetchup()
	domainNodeKey := util.GetDomainNodeKey(domain, node)
	nodeValue, isexist1 := parkKetchup.GetAny(domainNodeKey)
	sessionValue, isexist2 := parkKetchup.GetAny(SESSIONPREFIX + domainNodeKey)
	timeValue, isexist3 := parkKetchup.GetAny(TIMEPREFIX + domainNodeKey)
	if isexist1 {
		ret.SetValue(domainNodeKey, nodeValue)
	}
	if isexist2 {
		ret.SetValue(SESSIONPREFIX+domainNodeKey, sessionValue)
	}
	if isexist3 {
		ret.SetValue(TIMEPREFIX+domainNodeKey, timeValue)
	}
	return ret
}

func (parkKetchup *ParkKetchup) GetAllNodeByDomain(domain string) *model.Ketchup {
	ret := model.NewKetchup()
	for key, _ := range parkKetchup.GetALL() {
		if key == domain {
			continue
		}
		if strings.Index(key.(string), domain) == 0 {
			nodeValue, _ := parkKetchup.GetAny(key)
			ret.SetValue(key, nodeValue)
		}
	}
	return ret
}

func (parkKetchup *ParkKetchup) DeleteNode(domain, node string) *model.Ketchup {
	if node == "" {
		return parkKetchup.DeleteAllNodeByDomain(domain)
	}
	ret := model.NewKetchup()
	domainNodeKey := util.GetDomainNodeKey(domain, node)
	nodeValue, _ := parkKetchup.GetAny(domainNodeKey)
	sessionValue, _ := parkKetchup.GetString(SESSIONPREFIX + domainNodeKey)
	timeValue, _ := parkKetchup.GetAny(TIMEPREFIX + domainNodeKey)
	parkKetchup.RemoveKey(domainNodeKey)
	parkKetchup.RemoveKey(SESSIONPREFIX + domainNodeKey)
	parkKetchup.RemoveKey(TIMEPREFIX + domainNodeKey)
	ret.SetValue(domainNodeKey, nodeValue)
	ret.SetValue(SESSIONPREFIX+domainNodeKey, sessionValue)
	ret.SetValue(TIMEPREFIX+domainNodeKey, timeValue)
	return ret

}

func (parkKetchup *ParkKetchup) DeleteAllNodeByDomain(domain string) *model.Ketchup {
	ret := model.NewKetchup()
	for key, _ := range parkKetchup.GetALL() {
		if strings.Index(key.(string), domain) == 0 {
			nodeValue, _ := parkKetchup.GetAny(key)
			ret.SetValue(key, nodeValue)
			parkKetchup.RemoveKey(key)
		}
	}
	return ret
}
