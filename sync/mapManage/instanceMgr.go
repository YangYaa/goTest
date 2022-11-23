package mapManage

import (
	"fmt"
	"sync"
)

type InstanceData struct {
	name        string
	callBackUrl string
}

type InstanceManager struct {
	InstanceMap sync.Map
}

var InstanceMgr *InstanceManager

func NewManager() *InstanceManager {
	mgrInstance := new(InstanceManager)
	return mgrInstance
}

func (mgr *InstanceManager) AddInstance(data ...string) bool {
	name := data[0]
	callBackUrl := data[1]
	instance := InstanceData{}
	instance.name = name
	instance.callBackUrl = callBackUrl
	mgr.InstanceMap.Store(name, instance)
	return true
}

func (mgr *InstanceManager) GetInstance(name string) string {
	url := ""
	v, ok := mgr.InstanceMap.Load(name)
	if ok {
		if v.(InstanceData).callBackUrl != "" {
			url = v.(InstanceData).callBackUrl
		}
	}
	return url
}

func (mgr *InstanceManager) DelInstance(name string) bool {
	if _, ok := mgr.InstanceMap.Load(name); ok == true {
		mgr.InstanceMap.Delete(name)
		return true
	} else {
		fmt.Println("The instance not have this name")
		return false
	}
}
