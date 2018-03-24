package main

type buffer struct {
	size       int
	startIndex int
	endIndeex  int
	data       []interface{}
}

func newBuffer(sz int) *buffer {

	b := new(buffer)
	b.size = sz
	b.startIndex = 0
	b.endIndeex = sz - 1
	b.data = make([]interface{}, sz, sz)

	return b
}

func (b *buffer) insert(data interface{}) {

	b.data[b.endIndeex] = data
	b.endIndeex = b.forwardIndex(b.endIndeex)

	if b.endIndeex == b.startIndex {
		b.startIndex = b.forwardIndex(b.startIndex)
	}
}

func (b *buffer) getList() []interface{} {

	r := make([]interface{}, b.size, b.size)

	i := 0
	idx := b.startIndex
	for {
		r[i] = b.data[idx]
		idx = b.forwardIndex(idx)
		if idx == b.endIndeex {
			break
		}
	}

	return r
}

func (b *buffer) forwardIndex(idx int) int {

	index := idx
	index++
	if index >= b.size {
		index = 0
	}

	return index
}
