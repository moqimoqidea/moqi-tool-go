// Package bigcache
// @author moqi
// On 2022/10/31 11:40:11
// see: https://github.com/allegro/bigcache
package bigcache

import (
	"fmt"
	"github.com/allegro/bigcache"
	"testing"
	"time"
)

// 默认初始化
func TestInitDefaultCache(t *testing.T) {
	// 创建一个LifeWindow为5秒的cache实例
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(time.Second * 5))
	defer func(cache *bigcache.BigCache) {
		_ = cache.Close()
	}(cache)

	// 设置缓存
	err := cache.Set("key1", []byte("hello word"))
	if err != nil {
		t.Errorf("设置缓存失败: %v", err)
	}

	// 获取缓存
	data, getErr := cache.Get("key1")
	if getErr != nil {
		t.Errorf("获取缓存失败: %v", getErr)
	}
	fmt.Printf("获取结果: %s\n", data)
}
