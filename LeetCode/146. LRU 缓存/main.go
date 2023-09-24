package _46__LRU_缓存

import (
	"container/list"
)

type entry struct {
	key, value int
}

type LRUCache struct {
	capacity  int
	list      *list.List // 双向链表
	keyToNode map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:  capacity,
		list:      list.New(),
		keyToNode: map[int]*list.Element{},
	}
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.keyToNode[key]; ok {
		c.list.MoveToFront(node)
		return node.Value.(entry).value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.keyToNode[key]; ok {
		node.Value = entry{key, value} // 更新
		c.list.MoveToFront(node)
		return
	}
	c.keyToNode[key] = c.list.PushFront(entry{key, value}) // 新书，放在最上面
	if len(c.keyToNode) > c.capacity {
		delete(c.keyToNode, c.list.Remove(c.list.Back()).(entry).key) // 去掉最后一本书
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
