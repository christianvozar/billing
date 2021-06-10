package paasio

import (
	"io"
	"sync/atomic"
)

// New constructs a new io.Writer with associated metrics.
func NewWriterWithMetrics(w io.Writer) *WriterWithMetrics {
	return &WriterWithMetrics{w: w}
}

type WriterWithMetrics struct {
	w      io.Writer
	writes int32
	bytes  int64
}

// Write implements io.Writer.
func (wm *WriterWithMetrics) Write(b []byte) (int, error) {
	atomic.AddInt32(&wm.writes, 1)
	atomic.AddInt64(&wm.bytes, int64(len(b)))
	return wm.w.Write(b)
}

func (wm *WriterWithMetrics) WriteCount() (n int64, nops int32) {
	return wm.Bytes(), wm.Writes()
}

// Writes returns the total number of writes.
func (wm *WriterWithMetrics) Writes() int32 {
	return atomic.LoadInt32(&wm.writes)
}

// Bytes returns the number of total bytes written.
func (wm *WriterWithMetrics) Bytes() int64 {
	return atomic.LoadInt64(&wm.bytes)
}
