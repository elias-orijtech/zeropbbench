package gcbench

import (
	"fmt"
	"testing"
)

func BenchmarkZeroPBArgs(b *testing.B) {
	benches := []struct {
		bufSize int
		argSize int
		alloc   bool
	}{
		{}, // baseline
		{
			bufSize: 1 * 1024,
			argSize: 1 * 1024,
			alloc:   true,
		},
		{
			bufSize: 64 * 1024,
			argSize: 1 * 1024,
			alloc:   true,
		},
		{
			bufSize: 64 * 1024,
			argSize: 64 * 1024,
			alloc:   true,
		},
		{
			bufSize: 1 * 1024,
			argSize: 1 * 1024,
			alloc:   false,
		},
		{
			bufSize: 64 * 1024,
			argSize: 1 * 1024,
			alloc:   false,
		},
		{
			bufSize: 64 * 1024,
			argSize: 64 * 1024,
			alloc:   false,
		},
	}
	for _, bench := range benches {
		alloc := "GC"
		if !bench.alloc {
			alloc = "NoGC"
		}
		name := fmt.Sprintf("%d/%d/%s", bench.argSize, bench.bufSize, alloc)
		b.Run(name, func(b *testing.B) {
			var buf []uint32
			if !bench.alloc {
				buf = make([]uint32, bench.bufSize/4)
			}
			for i := 0; i < b.N; i++ {
				if bench.alloc {
					buf = make([]uint32, bench.bufSize/4)
				}
				args := buf[:bench.argSize/4]
				for j := range args {
					args[j] = uint32(b.N)
				}
				foreign_func(args)
			}
		})
	}
}
