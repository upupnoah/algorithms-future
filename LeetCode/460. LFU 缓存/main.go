package _60__LFU_缓存

import "container/list"

// 思路：维护 n 堆书
// LRU：维护一堆书
// LFU：维护 n 堆书，并且维护一个 minFreq（表示latest frequently used)
type entry struct {
	key, value, freq int // freq 表示频率（这本书被看的次数
}

type LFUCache struct {
	capacity   int
	minFreq    int
	keyToNode  map[int]*list.Element
	freqToList map[int]*list.List
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity:   capacity,
		minFreq:    0,
		keyToNode:  map[int]*list.Element{},
		freqToList: map[int]*list.List{},
	}
}

func (c *LFUCache) pushFront(e *entry) {
	if _, ok := c.freqToList[e.freq]; !ok {
		c.freqToList[e.freq] = list.New()
	}
	c.keyToNode[e.key] = c.freqToList[e.freq].PushFront(e)
}

func (c *LFUCache) getEntry(key int) *entry {
	node := c.keyToNode[key]
	if node == nil {
		return nil
	}
	// 抽出
	e := node.Value.(*entry)
	lst := c.freqToList[e.freq]
	lst.Remove(node)
	if lst.Len() == 0 { // 这堆书空了
		delete(c.freqToList, e.freq)
		if c.minFreq == e.freq {
			c.minFreq++ // 更新 minFreq
		}
	}
	e.freq++
	c.pushFront(e)
	return e
}

func (c *LFUCache) Get(key int) int {
	if e := c.getEntry(key); e != nil {
		return e.value
	}
	return -1
}

func (c *LFUCache) Put(key int, value int) {
	if e := c.getEntry(key); e != nil { // 有则更新
		e.value = value
		return
	}
	if len(c.keyToNode) == c.capacity { // 已经有 capacity 个元素了，删除 lfu（最近最不常用节点）
		lst := c.freqToList[c.minFreq]
		delete(c.keyToNode, lst.Remove(lst.Back()).(*entry).key)
		if lst.Len() == 0 {
			delete(c.freqToList, c.minFreq)
		}
	}
	// 插入
	c.pushFront(&entry{
		key:   key,
		value: value,
		freq:  1,
	})
	c.minFreq = 1 // 更新
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
