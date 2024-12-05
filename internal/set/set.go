package set

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return make(Set[T])
}

func (set *Set[T]) Add(v T) *Set[T] {
	(*set)[v] = struct{}{}
	return set
}

func (set *Set[T]) Remove(v T) *Set[T] {
	delete(*set, v)
	return set
}

func (set Set[T]) Contains(v T) bool {
	_, ok := set[v]
	return ok
}
