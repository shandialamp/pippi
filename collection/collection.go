package collection

type Collection[T any] struct {
	data []T
}

func New[T any](data []T) *Collection[T] {
	return &Collection[T]{
		data: data,
	}
}

func (c *Collection[T]) ToSlice() []T {
	return c.data
}

func Map[T any, R any](c *Collection[T], method func(T) R) *Collection[R] {
	out := make([]R, len(c.data))
	for i, v := range c.data {
		out[i] = method(v)
	}
	return &Collection[R]{data: out}
}

func (c *Collection[T]) Push(v T) {
	c.data = append(c.data, v)
}

func (c *Collection[T]) Pull() T {
	last := c.data[c.Size()-1]
	c.data = c.data[:c.Size()-1]
	return last
}

func (c *Collection[T]) Size() int {
	return len(c.data)
}

func (c *Collection[T]) Each(method func(T)) {
	for _, item := range c.data {
		method(item)
	}
}
