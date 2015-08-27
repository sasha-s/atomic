package atomic

import (
	"sync"
	"testing"
)

type counter struct{ uint64 }

func (c *counter) Inc() {
	c.uint64++
}

func TestOnce(t *testing.T) {
	var once Once
	var c counter
	var wg sync.WaitGroup
	for g := 0; g < 1000; g++ {
		wg.Add(1)
		go func() {
			once.Do(c.Inc)
			if c.uint64 != 1 {
				t.Errorf("expected 1, got %d", c.uint64)
			}
		}()
		wg.Done()
	}
	wg.Wait()
	if c.uint64 != 1 {
		t.Fatalf("expected 1, got %d", c.uint64)
	}
}

func TestFirst(t *testing.T) {
	var first First
	var c counter
	var wg sync.WaitGroup
	for g := 0; g < 1000; g++ {
		wg.Add(1)
		go func() {
			if first.First() {
				c.Inc()
			}
			if c.uint64 != 1 {
				t.Errorf("expected 1, got %d", c.uint64)
			}
		}()
		wg.Done()
	}
	wg.Wait()
	if c.uint64 != 1 {
		t.Fatalf("expected 1, got %d", c.uint64)
	}
}

func nop() {}

func BenchmarkSyncOnce(b *testing.B) {
	var once sync.Once
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			once.Do(nop)
		}
	})
}

func BenchmarkOnce(b *testing.B) {
	var once Once
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			once.Do(nop)
		}
	})
}

func BenchmarkFirst(b *testing.B) {
	var f First
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if f.First() {
				nop()
			}
		}
	})
}
