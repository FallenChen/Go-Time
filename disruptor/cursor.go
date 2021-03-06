package disruptor

import "sync/atomic"

// prevent false sharing of the sequence cursor by padding the CPU cache line with 64 *bytes* of data
type Cursor [8]int64 

const defaultCursorValue = -1

func NewCursor() *Cursor {
	var this Cursor
	this[0] = defaultCursorValue
	return &this
}

func (this *Cursor) Store(value int64) {
	atomic.StoreInt64(&this[0], value)
}

func (this *Cursor) Load() int64 {
	return atomic.LoadInt64(&this[0])
}