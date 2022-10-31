// Package bigcache
// @author moqi
// On 2022/10/31 11:40:11
// see: https://github.com/allegro/bigcache
package bigcache

import (
	"fmt"
	"github.com/allegro/bigcache"
	"strconv"
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

// 创建自定义缓存
func TestInitCustom(t *testing.T) {
	// 指定创建属性
	config := bigcache.Config{
		// 设置分区的数量，必须是2的整倍数
		Shards: 1024,
		// LifeWindow后,缓存对象被认为不活跃,但并不会删除对象
		LifeWindow: 5 * time.Second,
		// CleanWindow后，会删除被认为不活跃的对象，0代表不操作；
		CleanWindow: 3 * time.Second,
		// 设置最大存储对象数量，仅在初始化时可以设置
		//MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntriesInWindow: 1,
		// 缓存对象的最大字节数，仅在初始化时可以设置
		MaxEntrySize: 500,
		// 是否打印内存分配信息
		Verbose: true,
		// 设置缓存最大值(单位为MB),0表示无限制
		HardMaxCacheSize: 8192,
		// 在缓存过期或者被删除时,可设置回调函数，参数是(key、val)，默认是nil不设置
		OnRemove: callBack,
		// 在缓存过期或者被删除时,可设置回调函数，参数是(key、val,reason)默认是nil不设置
		OnRemoveWithReason: nil,
	}

	cache, err := bigcache.NewBigCache(config)
	if err != nil {
		t.Error(err)
	}

	defer func(cache *bigcache.BigCache) {
		_ = cache.Close()
	}(cache)

	// 设置缓存
	_ = cache.Set("key1", []byte("hello word"))

	// 验证CleanWindow是否生效: key 存在
	time.Sleep(1 * time.Second)
	// 获取缓存
	data, getErr := cache.Get("key1")
	if getErr != nil {
		t.Errorf("获取缓存失败: %v", getErr)
	}

	fmt.Printf("获取结果:%s\n", data)

	// 验证CleanWindow是否生效: key 已被删除
	time.Sleep(9 * time.Second)
	// 获取缓存
	data1, getErr1 := cache.Get("key1")
	if getErr1 != nil {
		t.Logf("获取缓存失败: %v", getErr1)
	}

	fmt.Printf("获取结果:%s\n", data1)
	fmt.Println("运行结束！")
}

func callBack(key string, entry []byte) {
	fmt.Printf("key: %s, with entry: %s, be removed.", key, entry)
}

func TestSetAndGet(t *testing.T) {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(time.Minute))
	// 设置缓存
	err := cache.Set("key1", []byte("php"))
	if err != nil {
		t.Errorf("设置缓存失败:%v", err)
	}
	_ = cache.Set("key2", []byte("go"))
	// 获取缓存
	for _, key := range []string{"key1", "key2"} {
		if data, err := cache.Get(key); err == nil {
			fmt.Printf("key: %s 结果:%s\n", key, data)
		}
	}
}

/** 输出
=== RUN   TestSetAndGet
key: key1 结果:php
key: key2 结果:go
--- PASS: TestSetAndGet (0.02s)
PASS
*/

func TestDelCache(t *testing.T) {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(time.Minute))
	key := "key"
	// 设置
	_ = cache.Set(key, []byte("111"))
	// 删除
	_ = cache.Delete(key)
	// 获取
	if _, err := cache.Get(key); err != nil {
		fmt.Println(err)
	}
}

/** 输出
=== RUN   TestUseCache
Entry not found
--- PASS: TestUseCache (0.02s)
PASS
*/

// 统计缓存数量和容量
func TestLenAndCap(t *testing.T) {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(time.Minute))
	_ = cache.Set("key", []byte("1"))
	_ = cache.Set("key1", []byte("1"))
	_ = cache.Set("key2", []byte("1"))
	_ = cache.Set("key3", []byte("1"))
	fmt.Printf("缓存数量: %d \n", cache.Len())
	fmt.Printf("缓存容量: %d \n", cache.Capacity())
}

/** 输出
=== RUN   TestLen
缓存数量: 4
缓存容量: 299520000
--- PASS: TestLen (0.02s)
PASS
*/

// 重置所有分区的缓存
func TestReset(t *testing.T) {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(time.Minute))
	for i := 0; i < 10; i++ {
		k := fmt.Sprintf("key%d", i)
		_ = cache.Set(k, []byte(strconv.Itoa(i)))
	}
	fmt.Printf("重置前缓存数量: %d \n", cache.Len())
	// 重置所有分区的缓存
	_ = cache.Reset()
	fmt.Printf("重置后缓存数量: %d \n", cache.Len())
}

/** 输出
=== RUN   TestReset
重置前缓存数量: 10
重置后缓存数量: 0
--- PASS: TestReset (0.02s)
PASS
*/
