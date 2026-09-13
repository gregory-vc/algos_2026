package task6

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

type node[K ordered, V any] struct {
	key   K
	value V
	left  *node[K, V]
	right *node[K, V]
}

type BinarySearchTree[K ordered, V any] struct {
	root *node[K, V]
}

func (t *BinarySearchTree[K, V]) Get(key K) *V {
	current := t.root
	for current != nil {
		if key == current.key {
			return &current.value
		}
		if key < current.key {
			current = current.left
		} else {
			current = current.right
		}
	}
	return nil
}

func (t *BinarySearchTree[K, V]) Add(key K, value V) {
	if t.root == nil {
		t.root = &node[K, V]{key: key, value: value}
		return
	}

	current := t.root
	for {
		if key == current.key {
			current.value = value
			return
		}
		if key < current.key {
			if current.left == nil {
				current.left = &node[K, V]{key: key, value: value}
				return
			}
			current = current.left
		} else {
			if current.right == nil {
				current.right = &node[K, V]{key: key, value: value}
				return
			}
			current = current.right
		}
	}
}

func (t *BinarySearchTree[K, V]) Remove(key K) *V {
	var parent *node[K, V]
	current := t.root
	for current != nil {
		if key == current.key {
			break
		}
		parent = current
		if key < current.key {
			current = current.left
		} else {
			current = current.right
		}
	}

	if current == nil {
		return nil
	}

	removed := current.value

	if current.left == nil || current.right == nil {
		var replacement *node[K, V]
		if current.left != nil {
			replacement = current.left
		} else {
			replacement = current.right
		}
		if parent == nil {
			t.root = replacement
		} else if parent.left == current {
			parent.left = replacement
		} else {
			parent.right = replacement
		}
		return &removed
	}

	successorParent := current
	successor := current.right
	for successor.left != nil {
		successorParent = successor
		successor = successor.left
	}

	current.key = successor.key
	current.value = successor.value

	if successorParent == current {
		current.right = successor.right
	} else {
		successorParent.left = successor.right
	}

	return &removed
}
