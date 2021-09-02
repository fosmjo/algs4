package symboltable

type BinarySearchST struct {
	n      int
	keys   []Key
	values []Value
}

func NewBinarySearchST(cap int) OrderedSymbolTable {
	return &BinarySearchST{
		keys:   make([]Key, cap),
		values: make([]Value, cap),
	}
}

func (s *BinarySearchST) Put(key Key, value Value) {
	r := s.Rank(key)
	if r < s.n && s.keys[r] == key {
		s.values[r] = value
		return
	}

	if s.n+1 > len(s.keys) {
		panic("out of capacity")
	}

	for j := s.n; j > r; j-- {
		s.keys[j] = s.keys[j-1]
		s.values[j] = s.values[j-1]
	}

	s.keys[r] = key
	s.values[r] = value
	s.n++
}

func (s *BinarySearchST) Get(key Key) (value Value, exists bool) {
	r := s.Rank(key)
	if r < s.n && s.keys[r] == key {
		value = s.values[r]
		exists = true
		return
	}
	return
}

func (s *BinarySearchST) Delete(key Key) {
	r := s.Rank(key)
	if r < s.n && s.keys[r] == key {
		for j := r; j < s.n-1; j++ {
			s.keys[j] = s.keys[j+1]
			s.values[j] = s.values[j+1]
		}
		s.n--
	}

}

func (s *BinarySearchST) Contains(key Key) bool {
	r := s.Rank(key)
	return r < s.n && s.keys[r] == key
}

func (s *BinarySearchST) IsEmpty() bool {
	return s.Size() == 0
}

func (s *BinarySearchST) Size() int {
	return s.n
}

func (s *BinarySearchST) Keys() []string {
	return s.keys[0:s.n]
}

func (s *BinarySearchST) Min() (key Key, exists bool) {
	exists = s.n > 0
	if !exists {
		return
	}

	key = s.keys[0]
	return
}

func (s *BinarySearchST) Max() (key Key, exists bool) {
	exists = s.n > 0
	if !exists {
		return
	}

	key = s.keys[s.n-1]
	return
}

func (s *BinarySearchST) Floor(key Key) (outKey Key, exists bool) {
	r := s.Rank(key)
	if r <= 0 && s.keys[r] != key {
		exists = false
		return
	}

	if r < s.n && s.keys[r] == key {
		outKey = key
	} else {
		outKey = s.keys[r-1]
	}
	exists = true
	return
}

func (s *BinarySearchST) Ceiling(key Key) (outKey Key, exists bool) {
	r := s.Rank(key)
	if r > s.n-1 || (r == s.n-1 && key != s.keys[r]) {
		exists = false
		return
	}

	if s.keys[r] == key {
		outKey = key
	} else {
		outKey = s.keys[r+1]
	}
	exists = true
	return
}

func (s *BinarySearchST) Rank(key Key) int {
	return s.rank(key, 0, s.n-1)
}

func (s *BinarySearchST) rank(key Key, lo, hi int) int {
	// NOTE
	if hi < lo {
		return lo
	}

	mid := lo + (hi-lo)>>1
	if key < s.keys[mid] {
		return s.rank(key, lo, mid-1)
	} else if key > s.keys[mid] {
		return s.rank(key, mid+1, hi)
	} else {
		return mid
	}
}

func (s *BinarySearchST) Select(k int) (key Key, exists bool) {
	exists = s.n > k
	if !exists {
		return
	}

	key = s.keys[k]
	return
}

func (s *BinarySearchST) DeleteMin() {
	min, exists := s.Min()
	if !exists {
		return
	}
	s.Delete(min)
}

func (s *BinarySearchST) DeleteMax() {
	max, exists := s.Max()
	if !exists {
		return
	}
	s.Delete(max)
}

func (s *BinarySearchST) SizeBetween(lo, hi Key) int {
	r1 := s.Rank(lo)
	r2 := s.Rank(hi)

	if r1 < s.n && s.keys[r1] == lo {
		r1--
	}
	if r2 < s.n && s.keys[r2] == hi {
		r2++
	}

	return r2 - r1 - 1
}

func (s *BinarySearchST) KeysBetween(lo, hi Key) []Key {
	r1 := s.Rank(lo)
	r2 := s.Rank(hi)

	if r1 < s.n && s.keys[r1] == lo && r1 > 0 {
		r1--
	}
	if r2 < s.n && s.keys[r2] == hi {
		r2++
	}

	return s.keys[r1:r2]
}
