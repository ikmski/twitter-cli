package main

type buffer struct {
	size       int
	startIndex int
	endIndeex  int
	data       []interface{}
}

func newBuffer(size int) *buffer {

	b := new(buffer)
	b.size = size
	b.startIndex = 0
	b.endIndeex = 0
	b.data = make([]interface{}, size, size)

	return b
}

func (b *buffer) push(data interface{}) {

	b.data[b.endIndeex] = data
	b.endIndeex = b.getNextIndex(b.endIndeex)

	if b.endIndeex == b.startIndex {
		b.startIndex = b.getNextIndex(b.startIndex)
	}
}

func (b *buffer) toList() []interface{} {

	r := make([]interface{}, b.size, b.size)

	i := 0
	idx := b.startIndex
	for {
		r[i] = b.data[idx]
		idx = b.getNextIndex(idx)
		if idx == b.endIndeex {
			break
		}
		i++
	}

	return r
}

func (b *buffer) getNextIndex(idx int) int {

	index := idx
	index++
	if index >= b.size {
		index = 0
	}

	return index
}
