package data_structures

type HashTable struct {
	data []*HashNode
	size int
}

type HashNode struct {
	key  int
	next *HashNode
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		data: make([]*HashNode, size),
		size: size,
	}
}

func (h *HashTable) hash(key int) int {
	return key & h.size
}

func (h *HashTable) Insert(key int) {
	index := h.hash(key)
	head := h.data[index]

	for n := head; n != nil; n = n.next {
		if n.key == key {
			return
		}
	}

	newNode := &HashNode{key: key}
	newNode.next = head
	h.data[index] = newNode
}

func (h *HashTable) Delete(key int) {
	index := h.hash(key)
	head := h.data[index]

	if head == nil {
		return
	}

	if head.key == key {
		h.data[index] = head.next
		return
	}

	for ; head != nil && head.next != nil; head = head.next {
		if head.next.key == key {
			head.next = head.next.next
		}
	}
}

func (h *HashTable) Search(key int) bool {
	index := h.hash(key)

	for n := h.data[index]; n != nil; n = n.next {
		if n.key == key {
			return true
		}
	}

	return false
}
