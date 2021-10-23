package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// RawProjectCfg 原生工程json类型
type RawProjectCfg struct {
	Name string       `json:"name"`
	Data BTProjectCfg `json:"data"`
	Path string       `json:"path"`
}

// LoadRawProjectCfg 加载原生工程
func LoadRawProjectCfg(path string) (*RawProjectCfg, error) {
	var project RawProjectCfg

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("LoadRawProjectCfg failed, %v", err)
	}

	err = json.Unmarshal(file, &project)
	if err != nil {
		return nil, fmt.Errorf("LoadRawProjectCfg failed, %v", err)
	}

	return &project, nil
}
