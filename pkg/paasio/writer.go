package paasio

import (
	"io"
	"sync/atomic"
)

// New constructs a new io.Writer with associated metrics.
func NewWriterWithMetrics() *WriterWithMetrics {
	return &WriterWithMetrics{}
}

type WriterWithMetrics struct {
	w      io.Writer
	writes int64
	bytes  int64
}

func (wm *WriterWithMetrics) Write(b []byte) {
	atomic.AddInt64(&wm.writes, 1)
	atomic.AddInt64(&wm.bytes, int64(len(b)))
	return
}

func (wm *WriterWithMetrics) WriteCount() (n int64, nops int64) {
	return wm.Bytes(), wm.Writes()
}

// Writes returns the total number of writes.
func (wm *WriterWithMetrics) Writes() int64 {
	return atomic.LoadInt64(&wm.writes)
}

// Bytes returns the number of total bytes written.
func (wm *WriterWithMetrics) Bytes() int64 {
	return atomic.LoadInt64(&wm.bytes)
}
