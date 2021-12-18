package main

// 导入github上的包的时候，使用go get命令
// 不需要加schema和.git文件，下载的包会被放在GOPATH所在的路径
import (
	cm "github.com/easierway/concurrent_map"
	"testing"
)

// go get github.com/easierway/concurrent_map
func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}

// 使用第三方go依赖管理工具
// glide 依赖管理工具
func TestConcurrentMapNew(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
