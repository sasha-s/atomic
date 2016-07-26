// Package atomic provides type-safe atomics.
// See https://golang.org/src/sync/atomic/doc.go
package atomic

import (
	"runtime"
	"sync/atomic"
	"unsafe"
)

// A Uint32 is an atomic uint32.
// Once created, a Uint32 must not be copied.
type Uint32 struct{ x uint32 }

// CreateUint32 creates a Uint32.
func CreateUint32(v uint32) Uint32 {
	return Uint32{v}
}

// Load atomically loads the value.
func (a *Uint32) Load() uint32 {
	return atomic.LoadUint32(&a.x)
}

// Store atomically stores the value.
func (a *Uint32) Store(v uint32) {
	atomic.StoreUint32(&a.x, v)
}

// Swap atomically stores the new value and returns the old one.
func (a *Uint32) Swap(v uint32) uint32 {
	return atomic.SwapUint32(&a.x, v)
}

// CompareAndSwap excutes the compare-and-swap operation for *a.
func (a *Uint32) CompareAndSwap(old, new uint32) (swapped bool) {
	return atomic.CompareAndSwapUint32(&a.x, old, new)
}

// Add atomically adds δ to *a and returns the new value.
func (a *Uint32) Add(δ uint32) uint32 {
	return atomic.AddUint32(&a.x, δ)
}

// Sub atomically subtracts δ from *a and returns the new value.
func (a *Uint32) Sub(δ uint32) uint32 {
	// See https://golang.org/search?q=atomic.AddUint32
	return atomic.AddUint32(&a.x, ^(δ - 1))
}

// Dec atomically decrements *a.
func (a *Uint32) Dec() uint32 {
	// See https://golang.org/search?q=atomic.AddUint32
	return atomic.AddUint32(&a.x, ^uint32(0))
}

// Inc atomically increments *a.
func (a *Uint32) Inc() uint32 {
	return atomic.AddUint32(&a.x, 1)
}

// A Uint64 is an atomic uint32.
type Uint64 struct{ x uint64 }

// CreateUint32 creates a Uint64.
func CreateUint64(v uint64) Uint64 {
	return Uint64{v}
}

// Load atomically loads the value.
func (a *Uint64) Load() uint64 {
	return atomic.LoadUint64(&a.x)
}

// Store atomically stores the value.
func (a *Uint64) Store(v uint64) {
	atomic.StoreUint64(&a.x, v)
}

// Swap atomically stores the new value and returns the old one.
func (a *Uint64) Swap(v uint64) uint64 {
	return atomic.SwapUint64(&a.x, v)
}

// CompareAndSwap excutes the compare-and-swap operation for *a.
func (a *Uint64) CompareAndSwap(old, new uint64) (swapped bool) {
	return atomic.CompareAndSwapUint64(&a.x, old, new)
}

// Add atomically adds δ to *a and returns the new value.
func (a *Uint64) Add(δ uint64) uint64 {
	return atomic.AddUint64(&a.x, δ)
}

// Sub atomically subtracts δ from *a and returns the new value.
func (a *Uint64) Sub(δ uint64) uint64 {
	// See https://golang.org/search?q=atomic.AddUint64
	return atomic.AddUint64(&a.x, ^(δ - 1))
}

// Dec atomically decrements *a.
func (a *Uint64) Dec() uint64 {
	// See https://golang.org/search?q=atomic.AddUint64
	return atomic.AddUint64(&a.x, ^uint64(0))
}

// Inc atomically increments *a.
func (a *Uint64) Inc() uint64 {
	return atomic.AddUint64(&a.x, 1)
}

// An Int32 is an atomic int32.
// Once created, an int32 must not be copied.
type Int32 struct{ x int32 }

// CreateInt32 creates an Int32.
func CreateInt32(v int32) Int32 {
	return Int32{v}
}

// Load atomically loads the value.
func (a *Int32) Load() int32 {
	return atomic.LoadInt32(&a.x)
}

// Store atomically stores the value.
func (a *Int32) Store(v int32) {
	atomic.StoreInt32(&a.x, v)
}

// Swap atomically stores the new value and returns the old one.
func (a *Int32) Swap(v int32) int32 {
	return atomic.SwapInt32(&a.x, v)
}

// CompareAndSwap excutes the compare-and-swap operation for *a.
func (a *Int32) CompareAndSwap(old, new int32) (swapped bool) {
	return atomic.CompareAndSwapInt32(&a.x, old, new)
}

// Add atomically adds δ to *a and returns the new value.
func (a *Int32) Add(δ int32) int32 {
	return atomic.AddInt32(&a.x, δ)
}

// Sub atomically subtracts δ from *a and returns the new value.
func (a *Int32) Sub(δ int32) int32 {
	return atomic.AddInt32(&a.x, -δ)
}

// Dec atomically decrements *a.
func (a *Int32) Dec() int32 {
	return atomic.AddInt32(&a.x, -1)
}

// Inc atomically increments *a.
func (a *Int32) Inc() int32 {
	return atomic.AddInt32(&a.x, 1)
}

// An Int64 is an atomic Int64.
type Int64 struct{ x int64 }

// CreateInt64 creates a Int64.
func CreateInt64(v int64) Int64 {
	return Int64{v}
}

// Load atomically loads the value.
func (a *Int64) Load() int64 {
	return atomic.LoadInt64(&a.x)
}

// Store atomically stores the value.
func (a *Int64) Store(v int64) {
	atomic.StoreInt64(&a.x, v)
}

// Swap atomically stores the new value and returns the old one.
func (a *Int64) Swap(v int64) int64 {
	return atomic.SwapInt64(&a.x, v)
}

// CompareAndSwap excutes the compare-and-swap operation for *a.
func (a *Int64) CompareAndSwap(old, new int64) (swapped bool) {
	return atomic.CompareAndSwapInt64(&a.x, old, new)
}

// Add atomically adds δ to *a and returns the new value.
func (a *Int64) Add(δ int64) int64 {
	return atomic.AddInt64(&a.x, δ)
}

// Sub atomically subtracts δ from *a and returns the new value.
func (a *Int64) Sub(δ int64) int64 {
	return atomic.AddInt64(&a.x, -δ)
}

// Dec atomically decrements *a.
func (a *Int64) Dec() int64 {
	return atomic.AddInt64(&a.x, -1)
}

// Inc atomically increments *a.
func (a *Int64) Inc() int64 {
	return atomic.AddInt64(&a.x, 1)
}

// A Pointer is an atomic unsafe.pointer.
type Pointer struct{ x unsafe.Pointer }

// CreatePointer creates a Pointer.
func CreatePointer(v unsafe.Pointer) Pointer {
	return Pointer{v}
}

// Load atomically loads the value.
func (a *Pointer) Load() unsafe.Pointer {
	return atomic.LoadPointer(&a.x)
}

// Store atomically stores the value.
func (a *Pointer) Store(v unsafe.Pointer) {
	atomic.StorePointer(&a.x, v)
}

// Swap atomically stores the new value and returns the old one.
func (a *Pointer) Swap(v unsafe.Pointer) unsafe.Pointer {
	return atomic.SwapPointer(&a.x, v)
}

// CompareAndSwap excutes the compare-and-swap operation for *a.
func (a *Pointer) CompareAndSwap(old, new unsafe.Pointer) (swapped bool) {
	return atomic.CompareAndSwapPointer(&a.x, old, new)
}

// A Value provides an atomic load and store of a consistently typed value.
// Values can be created as part of other data structures.
// The zero value for a Value returns nil from Load.
// Once initialized, a Value must not be copied.
// See https://golang.org/pkg/sync/atomic/#Value
type Value struct{ atomic.Value }

func CreateValue(v interface{}) Value {
	var a Value
	a.Store(v)
	return a
}

// Once is an object that will perform exactly one action.
// Similar to sync.Once, but uses less memory and is as fast.
// See http://golang.org/pkg/sync/#Once
type Once struct{ x Uint32 }

// Do calls the function f if and only if Do is being called for the first time for this instance of Once.
func (o *Once) Do(f func()) {
	for {
		switch o.x.Load() {
		case start:
			if o.x.CompareAndSwap(start, inProgress) {
				// State is inProgress now.
				f()
				// Finish.
				o.x.Swap(finished)
			}
			return
		case inProgress:
			runtime.Gosched() // In case f() is slow.
			continue          // wait for finished.
		case finished:
			return
		}
	}
}

// First is an object that will return true exactly once.
type First struct{ x Uint32 }

// First returns true if it is the first time it is called on this very instance.
// Note, this is not safe:
// var f atomic.First
// var a int
// ...
// if f.First() { a++ } // Init shared resource.
// fmt.Println(a) // Use shared resource.
func (f *First) First() bool {
	if f.x.Load() == 1 {
		return false
	}
	return f.x.Swap(1) == 0
}

const (
	start uint32 = iota
	inProgress
	finished
)
