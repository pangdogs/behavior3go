package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// BTProjectCfg 工程json类型
type BTProjectCfg struct {
	ID     string       `json:"id"`
	Select string       `json:"selectedTree"`
	Scope  string       `json:"scope"`
	Trees  []*BTTreeCfg `json:"trees"`
}

// LoadProjectCfg 加载BT工程
func LoadProjectCfg(path string) (*BTProjectCfg, error) {
	var project BTProjectCfg

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("LoadProjectCfg failed, %v", err)
	}

	err = json.Unmarshal(file, &project)
	if err != nil {
		return nil, fmt.Errorf("LoadProjectCfg failed, %v", err)
	}

	return &project, nil
}
