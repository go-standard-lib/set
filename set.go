package set

import "github.com/go-standard-lib/map/hashmap"

type Set[V any] interface {
	Add(value *V)

	Remove(value *V)

	ToArray() []*V

	Contains(value *V) bool

	Size() int
}

type set[V any] struct {
	hash hashmap.HashMap[*V, bool]
}

func New[V any]() Set[V] {
	return &set[V]{
		hash: hashmap.New[*V, bool](),
	}
}

func (s *set[V]) Add(value *V) {
	s.hash.Put(value, true)
}

func (s *set[V]) Remove(value *V) {
	s.hash.Remove(value)
}

func (s *set[V]) Contains(value *V) bool {
	return s.hash.Contains(value)
}

func (s *set[V]) ToArray() []*V {
	return append([]*V{}, s.hash.Keys()...)
}

func (s *set[V]) Size() int {
	return s.hash.Size()
}
