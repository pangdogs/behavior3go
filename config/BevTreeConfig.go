package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

// BTNodeCfg 行为树节点配置
type BTNodeCfg struct {
	Id           string                 `json:"id"`
	Name         string                 `json:"name"`
	CategoryName string                 `json:"category"`
	Title        string                 `json:"title"`
	Description  string                 `json:"description"`
	Children     []string               `json:"children"`
	Child        string                 `json:"child"`
	Parameters   map[string]interface{} `json:"parameters"`
	Properties   map[string]interface{} `json:"properties"`
}

func (btnc *BTNodeCfg) GetPropertyAsFloat64(name string) float64 {
	v, ok := btnc.Properties[name]
	if !ok {
		return 0
	}
	return v.(float64)
}

func (btnc *BTNodeCfg) GetPropertyAsInt64(name string) int64 {
	v, ok := btnc.Properties[name]
	if !ok {
		return 0
	}
	return int64(v.(float64))
}

func (btnc *BTNodeCfg) GetPropertyAsBool(name string) bool {
	v, ok := btnc.Properties[name]
	if !ok {
		return false
	}

	b, ok := v.(bool)
	if !ok {
		if str, ok := v.(string); ok {
			return strings.ToLower(str) == "true"
		}
		return false
	}

	return b
}

func (btnc *BTNodeCfg) GetPropertyAsString(name string) string {
	v, ok := btnc.Properties[name]
	if !ok {
		return ""
	}
	return v.(string)
}

// BTTreeCfg 行为树配置
type BTTreeCfg struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Root        string                 `json:"root"`
	Properties  map[string]interface{} `json:"properties"`
	Nodes       map[string]BTNodeCfg   `json:"nodes"`
}

// LoadTreeCfg 加载行为树
func LoadTreeCfg(path string) (*BTTreeCfg, error) {
	var tree BTTreeCfg

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("LoadTreeCfg failed, %v", err)
	}

	err = json.Unmarshal(file, &tree)
	if err != nil {
		return nil, fmt.Errorf("LoadTreeCfg failed, %v", err)
	}

	return &tree, nil
}
