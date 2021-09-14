package algorithms

type LRUNode struct {
	key  int
	val  int
	prev *LRUNode
	next *LRUNode
}

type DoubleList struct {
	head *LRUNode
	tail *LRUNode
	size int
}

func (d *DoubleList) Init() {
	d.head = &LRUNode{key: 0, val: 0}
	d.tail = &LRUNode{key: 0, val: 0}
	d.head.next = d.tail
	d.tail.prev = d.head
	d.size = 0
}

func (d *DoubleList) AddLast(node *LRUNode) {
	node.prev = d.tail.prev
	node.next = d.tail
	d.tail.prev.next = node
	d.tail.prev = node
	d.size++
}

func (d *DoubleList) Remove(node *LRUNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	d.size--
}

func (d *DoubleList) RemoveFirst() *LRUNode {
	if d.head.next == d.tail {
		return nil
	}
	first := d.head.next
	d.Remove(first)
	return first

}

type LRUCache struct {
	cacheMap   map[int]*LRUNode
	doubleList *DoubleList
	cacheCap   int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cacheCap:   capacity,
		cacheMap:   make(map[int]*LRUNode),
		doubleList: &DoubleList{},
	}
	lru.doubleList.Init()
	return lru
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cacheMap[key]; !ok {
		return -1
	}
	this.markRecent(key)
	return this.cacheMap[key].val
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cacheMap[key]; ok {
		this.deleteKey(key)
		this.addRecent(key, value)
		return
	}
	if this.cacheCap == this.doubleList.size {
		this.deleteLeastRecent()
	}
	this.addRecent(key, value)

}

func (this *LRUCache) markRecent(key int) {
	node := this.cacheMap[key]
	this.doubleList.Remove(node)
	this.doubleList.AddLast(node)
}

func (this *LRUCache) addRecent(key, val int) {
	node := &LRUNode{key: key, val: val}
	this.cacheMap[key] = node
	this.doubleList.AddLast(node)
}

func (this *LRUCache) deleteKey(key int) {
	node := this.cacheMap[key]
	this.doubleList.Remove(node)
	delete(this.cacheMap, key)
}

func (this *LRUCache) deleteLeastRecent() {
	node := this.doubleList.RemoveFirst()
	key := node.key
	delete(this.cacheMap, key)
}
