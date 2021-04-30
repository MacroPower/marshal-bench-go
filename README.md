# marshal-bench-go

Collection of Go (un)marshalling benchmarks

## Results

```
goarch: amd64
pkg: github.com/MacroPower/marshal-bench-go/json
cpu: Intel(R) Core(TM) i9-7900X CPU @ 3.30GHz
BenchmarkJsonUnmarshalStruct-20        	  173648	      6977 ns/op	     984 B/op	      23 allocs/op
BenchmarkJsonUnmarshalMap-20           	  142856	      8354 ns/op	    2907 B/op	      28 allocs/op
BenchmarkJsonMarshalStruct-20          	  275514	      4356 ns/op	    2654 B/op	       8 allocs/op
BenchmarkJsonMarshalMap-20             	  244900	      4941 ns/op	    2333 B/op	       7 allocs/op
BenchmarkEasyjsonUnmarshalStruct-20    	  521784	      2343 ns/op	     480 B/op	      20 allocs/op
BenchmarkEasyjsonUnmarshalMap-20       	  315848	      3761 ns/op	    2668 B/op	      24 allocs/op
BenchmarkEasyjsonMarshalStruct-20      	  800047	      1498 ns/op	    2186 B/op	       8 allocs/op
BenchmarkEasyjsonMarshalMap-20         	  616256	      1931 ns/op	    1865 B/op	       7 allocs/op
```
