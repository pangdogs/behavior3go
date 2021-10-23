/*
从导出的工程文件加载
*/
package main

import (
	"fmt"
	b3 "github.com/pangdogs/behavior3go"
	. "github.com/pangdogs/behavior3go/config"
	. "github.com/pangdogs/behavior3go/core"
	. "github.com/pangdogs/behavior3go/examples/share"
	. "github.com/pangdogs/behavior3go/loader"
)

func main() {
	projectConfig, ok := LoadProjectCfg("project.json")
	if !ok {
		fmt.Println("LoadTreeCfg err")
		return
	}

	//自定义节点注册
	maps := b3.NewRegisterStructMaps()
	maps.Register("Log", new(LogTest))

	var firstTree *BehaviorTree
	//载入
	for _, v := range projectConfig.Trees {
		tree := CreateBevTreeFromConfig(&v, maps)
		tree.Print()
		if firstTree == nil {
			firstTree = tree
		}
	}

	//输入板
	board := NewBlackboard()
	//循环每一帧
	for i := 0; i < 5; i++ {
		firstTree.Tick(i, board)
	}
}
