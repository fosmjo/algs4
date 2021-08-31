package symboltable

type Node struct {
	Key   Key
	Value Value
	Next  *Node
}

type SequentialSearchST struct {
	head *Node
}

func NewSequentialSearchST() SymbolTable {
	return &SequentialSearchST{}
}

func (s *SequentialSearchST) Get(key Key) (value Value, exists bool) {
	for node := s.head; node != nil; node = node.Next {
		if node.Key == key {
			value = node.Value
			exists = true
			return
		}
	}
	return
}

func (s *SequentialSearchST) Put(key Key, value Value) {
	for node := s.head; node != nil; node = node.Next {
		if node.Key == key {
			node.Value = value
			return
		}
	}

	s.head = &Node{
		Key:   key,
		Value: value,
		Next:  s.head,
	}
}

func (s *SequentialSearchST) Delete(key Key) {
	var prev *Node
	for node := s.head; node != nil; prev, node = node, node.Next {
		if node.Key == key {
			if prev == nil {
				s.head = node.Next
			} else {
				prev.Next = node.Next
			}
		}
	}
}

func (s *SequentialSearchST) Contains(key Key) bool {
	for node := s.head; node != nil; node = node.Next {
		if node.Key == key {
			return true
		}
	}
	return false
}

func (s *SequentialSearchST) IsEmpty() bool {
	return s.head == nil
}

func (s *SequentialSearchST) Size() int {
	size := 0
	for node := s.head; node != nil; node = node.Next {
		size++
	}
	return size
}

func (s *SequentialSearchST) Keys() (keys []Key) {
	for node := s.head; node != nil; node = node.Next {
		keys = append(keys, node.Key)
	}
	return
}
