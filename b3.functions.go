package behavior3go

import (
	"fmt"
	"reflect"
)

// RegisterStructMaps 定义注册结构map
type RegisterStructMaps struct {
	maps map[string]reflect.Type
}

func NewRegisterStructMaps() *RegisterStructMaps {
	return &RegisterStructMaps{make(map[string]reflect.Type)}
}

func (rsm *RegisterStructMaps) New(name string) (interface{}, error) {
	//fmt.Println("New ", name)
	var c interface{}
	var err error
	if v, ok := rsm.maps[name]; ok {
		c = reflect.New(v).Interface()
		//fmt.Println("found ", name, "  ", reflect.TypeOf(c))
		return c, nil
	} else {
		err = fmt.Errorf("not found %s struct", name)
		//fmt.Println("New no found", name, "  ", len(rsm.maps))
	}
	return nil, err
}

func (rsm *RegisterStructMaps) CheckElem(name string) bool {
	if _, ok := rsm.maps[name]; ok {
		return true
	}
	return false
}

func (rsm *RegisterStructMaps) Register(name string, c interface{}) {
	rsm.maps[name] = reflect.TypeOf(c).Elem()
}
