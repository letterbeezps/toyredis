package set

type (
	Set[T comparable] struct {
		data map[string]set[T]
	}

	set[T comparable] map[T]struct{}
)

var placeholder = struct{}{}

// set methods
func newset[T comparable]() set[T] {
	return make(set[T])
}

func (s set[T]) add(items ...T) {
	if len(items) == 0 {
		return
	}

	for _, v := range items {
		s[v] = placeholder
	}
}

func (s set[T]) remove(items ...T) {
	if len(items) == 0 {
		return
	}

	for _, v := range items {
		delete(s, v)
	}
}

func (s set[T]) has(items ...T) bool {
	if len(items) == 0 {
		return false
	}

	for _, v := range items {
		if _, ok := s[v]; !ok {
			return false
		}
	}
	return true
}

func (s set[T]) each(f func(item T) bool) {
	for v := range s {
		if !f(v) {
			break
		}
	}
}

func (s set[T]) copy() set[T] {
	ret := newset[T]()
	for v := range s {
		ret.add(v)
	}
	return ret
}

func (s set[T]) list() []T {
	list := make([]T, 0, len(s))
	for v := range s {
		list = append(list, v)
	}
	return list
}

func (s set[T]) size() int {
	return len(s)
}

func (s set[T]) Merge(t set[T]) {
	t.each(func(item T) bool {
		s[item] = placeholder
		return true
	})
}

func (s set[T]) Separate(t set[T]) {
	s.remove(t.list()...)
}

func union[T comparable](set1, set2 set[T], sets ...set[T]) set[T] {
	ret := set1.copy()

	set2.each(func(item T) bool {
		ret.add(item)
		return true
	})

	for _, subSet := range sets {
		subSet.each(func(item T) bool {
			ret.add(item)
			return true
		})
	}

	return ret
}

func difference[T comparable](set1, set2 set[T], sets ...set[T]) set[T] {
	ret := set1.copy()

	ret.Separate(set2)

	for _, subSet := range sets {
		ret.Separate(subSet)
	}

	return ret
}

func intersection[T comparable](set1, set2 set[T], sets ...set[T]) set[T] {
	all := union(set1, set1, sets...)
	ret := union(set1, set2, sets...)

	all.each(func(item T) bool {
		if !set1.has(item) || !set2.has(item) {
			ret.remove(item)
		}

		for _, subSet := range sets {
			if !subSet.has(item) {
				ret.remove(item)
			}
		}

		return true
	})

	return ret
}

// Set methods
func New[T comparable]() *Set[T] {
	return &Set[T]{
		data: map[string]set[T]{},
	}
}
