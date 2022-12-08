package leetcode

type LRUCache struct {
	size     int
	capacity int
	cache    map[int]*CLinkedNode
	head     *CLinkedNode
	tail     *CLinkedNode
}

type CLinkedNode struct {
	key, value int
	prev, next *CLinkedNode
}

func NewCLinkedNode(key, value int) *CLinkedNode {
	return &CLinkedNode{
		key:   key,
		value: value,
	}
}

func (this *LRUCache) addToHead(node *CLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *CLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *CLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *CLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*CLinkedNode{},
		head:     NewCLinkedNode(0, 0),
		tail:     NewCLinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := NewCLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

/**
type LRUCache struct {
	capacity int
	datas    map[int][2]int
	age      int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		datas:    make(map[int][2]int, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	this.age++
	data, ok := this.datas[key]
	if !ok {
		return -1
	}
	this.datas[key] = [2]int{this.age, data[1]}
	return data[1]
}

func (this *LRUCache) Put(key int, value int) {
	this.age++
	if _, ok := this.datas[key]; !ok {
		if len(this.datas) == this.capacity {
			oldest := 0
			oldestAge := 0
			for k, v := range this.datas {
				if this.age-v[0] > oldestAge {
					oldestAge = this.age - v[0]
					oldest = k
				}
			}
			delete(this.datas, oldest)
		}
	}
	this.datas[key] = [2]int{this.age, value}
}

*/
