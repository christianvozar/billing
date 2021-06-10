package paasio

import (
	"io"
	"sync/atomic"
)

// New constructs a new io.Reader with associated metrics.
func NewReaderWithMetrics() *ReaderWithMetrics {
	return &ReaderWithMetrics{}
}

type ReaderWithMetrics struct {
	r     io.Reader
	reads int64
	bytes int64
}

func (rm *ReaderWithMetrics) Read(b []byte) {
	atomic.AddInt64(&rm.reads, 1)
	atomic.AddInt64(&rm.bytes, int64(len(b)))
	return
}

func (rm *ReaderWithMetrics) ReadCount() (n int64, nops int64) {
	return rm.Bytes(), rm.Reads()
}

// Reads returns the total number of writes.
func (rm *ReaderWithMetrics) Reads() int64 {
	return atomic.LoadInt64(&rm.reads)
}

// Bytes returns the number of total bytes written.
func (rm *ReaderWithMetrics) Bytes() int64 {
	return atomic.LoadInt64(&rm.bytes)
}
