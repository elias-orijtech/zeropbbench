## Benchmarking harness for zeropb

On a MacBook Pro M1,

```sh
$ go test -bench .
goos: darwin
goarch: arm64
pkg: github.com/elias-orijtech/zeropbbench
BenchmarkZeroPBArgs/0/0/NoGC-8         	26960395	        44.28 ns/op
BenchmarkZeroPBArgs/1024/1024/GC-8     	 5228528	       226.6 ns/op
BenchmarkZeroPBArgs/1024/65536/GC-8    	  307776	      3395 ns/op
BenchmarkZeroPBArgs/65536/65536/GC-8   	  130868	      9015 ns/op
BenchmarkZeroPBArgs/1024/1024/NoGC-8   	 8179236	       146.5 ns/op
BenchmarkZeroPBArgs/1024/65536/NoGC-8  	 8118298	       146.9 ns/op
BenchmarkZeroPBArgs/65536/65536/NoGC-8 	  201465	      5879 ns/op
PASS
ok  	github.com/elias-orijtech/zeropbbench	10.160s
```
