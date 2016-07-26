# atomic: type-safe atomics in golang.

There is an [atomic](https://golang.org/pkg/sync/atomic/) package in standard library.

What's wrong with it?

* Interfaces not uniform: if I change a type of a field from uint32 to uint64 I have to go and fix each and every bit of code that touches it.
* Not protecting me from accidentally accessing the atomic variable in a non-safe way.

This package provides a uniform, clean interfaces and type safety.

Bonus: [Once](https://godoc.org/github.com/sasha-s/atomic#Once) which can happily replace [sync.Once](http://golang.org/pkg/sync/#Once) and it's close relative [First](https://godoc.org/github.com/sasha-s/atomic#First).

Once used less memory (just 4 bytes, single uint32) and is lock free.

Benchmarks (on MacBook Pro circa mid 2014):

```
name        time/op
SyncOnce-8  2.12ns ± 3%
Once-8      2.01ns ± 4%
First-8     1.93ns ± 2%
```

Docs are [here](https://godoc.org/github.com/sasha-s/atomic).


