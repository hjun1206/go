// lru 算法 简单实现
// 效果: 实现缓存，缓存大小固定为某个值
// 设置key的时候，需要判断内存大小,超过内存值，则淘汰最长时间未使用 key
// 获取key的时候，需要更新key排序

package main

import (
	"container/list"
	"fmt"
	"sync"
)

type CacheNode struct {
	Key, Value interface{}
}

type MemCache struct {
	maxSize int                           // 缓存大小
	keyList *list.List                    // 缓存的key值
	data    map[interface{}]*list.Element // 数据

	mx sync.Mutex // 加锁
}

// 初始化缓存
func NewMemCache(size int) *MemCache {
	return &MemCache{
		maxSize: size,
		keyList: list.New(),
		data:    make(map[interface{}]*list.Element, size),
	}
}

// Set 添加数据
func (n *MemCache) Set(key string, value interface{}) {
	// 加锁
	n.mx.Lock()
	defer n.mx.Unlock()

	// 判断key是否存在
	if v, ok := n.data[key]; ok {
		// 赋值
		n.data[key].Value = &CacheNode{Key: key, Value: value}
		// 调整key顺序(调到最前面)
		n.keyList.MoveToFront(v)
	} else {
		// 判断是否满了，删除
		if len(n.data) >= n.maxSize {
			back := n.keyList.Back()
			delete(n.data, back.Value.(*CacheNode).Key)
			n.keyList.Remove(back)
		}
		// 添加数据
		n.data[key] = n.keyList.PushFront(&CacheNode{Key: key, Value: value})
	}
}

// Get 获取数据
func (n *MemCache) Get(key string) (interface{}, bool) {
	// 读取数据
	if ele, ok := n.data[key]; ok {
		// 移动数据
		n.keyList.MoveToFront(ele)
		return ele.Value, true
	}
	return nil, false
}

// Size 获取数据大小
func (n *MemCache) Size() int {
	return n.keyList.Len()
}

func main() {
	lru := NewMemCache(3)
	lru.Set("10", "value1")
	lru.Set("20", "value2")
	lru.Set("30", "value3")
	lru.Set("10", "value4")
	lru.Set("50", "value5")

	fmt.Println("LRU Size:", lru.Size())
	if v, ret := lru.Get("10"); ret {
		fmt.Println("Get(10) : ", v)
	}
	if v, ret := lru.Get("20"); ret {
		fmt.Println("Get(20) : ", v)
	}
	if v, ret := lru.Get("30"); ret {
		fmt.Println("Get(30) : ", v)
	}
	if v, ret := lru.Get("40"); ret {
		fmt.Println("Get(40) : ", v)
	}
	if v, ret := lru.Get("50"); ret {
		fmt.Println("Get(50) : ", v)
	}
}
