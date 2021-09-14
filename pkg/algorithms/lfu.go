package algorithms

type CacheNode struct {
	key  int
	val  int
	prev *CacheNode
	next *CacheNode
	freq *FreqNode
}

func (c *CacheNode) addFreq(key, val int) *CacheNode {
	currFreq := c.freq
	var newFreq *FreqNode
	// check if next freq node exist and if its value is correct
	if currFreq.next != nil && currFreq.next.freq == currFreq.freq+1 {
		newFreq = currFreq.next
	} else {
		// creat new freq node
		newFreq = &FreqNode{
			freq: currFreq.freq + 1,
			prev: currFreq,
			next: currFreq.next,
		}
		if currFreq.next != nil {
			currFreq.next.prev = newFreq
		}
		if currFreq != nil {
			currFreq.next = newFreq
		}
	}
	newCacheNode := &CacheNode{
		key:  key,
		val:  val,
		prev: nil,
		next: newFreq.head,
		freq: newFreq,
	}
	if newFreq.head == nil {
		newFreq.tail = newCacheNode
	} else {
		newFreq.head.prev = newCacheNode
	}
	newFreq.head = newCacheNode
	c.deleteNode()
	return newCacheNode

}

//func NewFreq(prev, next *FreqNode, f int)

func (c *CacheNode) deleteNode() {
	freq := c.freq
	if c.prev == nil {
		freq.head = c.next
	} else {
		c.prev.next = c.next
	}

	if c.next == nil {
		freq.tail = c.prev
	} else {
		c.next.prev = c.prev
	}
	if freq.head == nil {
		freq.prev.next = freq.next
		if freq.next != nil {
			freq.next.prev = freq.prev
		}
	}
}

type FreqNode struct {
	freq int
	head *CacheNode
	tail *CacheNode
	prev *FreqNode
	next *FreqNode
}

type LFUCache struct {
	m    map[int]*CacheNode
	freq *FreqNode
	cap  int
}

func ConstructorLFU(capacity int) LFUCache {
	nodeMap := make(map[int]*CacheNode)
	f := &FreqNode{freq: 0}
	lfu := LFUCache{
		m:    nodeMap,
		cap:  capacity,
		freq: f,
	}
	return lfu
}

func (this *LFUCache) Get(key int) int {
	if k, ok := this.m[key]; ok {
		newNode := k.addFreq(key, k.val)
		this.m[key] = newNode
		return newNode.val
	} else {
		return -1
	}
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}
	if k, ok := this.m[key]; ok {
		newNode := k.addFreq(key, value)
		this.m[key] = newNode
	} else {
		if len(this.m) == this.cap {
			// least freq node is the second one (the first one is freq 0 with no cache node)
			dNode := this.freq.next.tail
			delete(this.m, dNode.key)
			dNode.deleteNode()
		}
		var newFreq *FreqNode
		if this.freq.next != nil && this.freq.next.freq == 1 {
			newFreq = this.freq.next
		} else {
			newFreq = &FreqNode{
				freq: this.freq.freq + 1,
				prev: this.freq,
				next: this.freq.next,
			}
			if this.freq.next != nil {
				this.freq.next.prev = newFreq
			}
			if this.freq != nil {
				this.freq.next = newFreq
			}
		}
		newCache := &CacheNode{
			key:  key,
			val:  value,
			freq: newFreq,
			prev: nil,
			next: newFreq.head,
		}
		if newFreq.head != nil {
			newFreq.head.prev = newCache
		} else {
			newFreq.tail = newCache
		}
		newFreq.head = newCache
		this.m[key] = newCache
	}
}
