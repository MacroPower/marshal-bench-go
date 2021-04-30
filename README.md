# marshal-bench-go

Collection of Go (un)marshalling benchmarks

## Results

```
goarch: amd64
pkg: github.com/MacroPower/marshal-bench-go/json
cpu: Intel(R) Core(TM) i9-7900X CPU @ 3.30GHz
BenchmarkJsonUnmarshalStruct-20              	  177123	      6827 ns/op	     984 B/op	      23 allocs/op
BenchmarkJsonUnmarshalStructPtr-20           	  164353	      7585 ns/op	    1144 B/op	      43 allocs/op
BenchmarkJsonUnmarshalMap-20                 	  146149	      8216 ns/op	    2909 B/op	      28 allocs/op
BenchmarkJsonMarshalStruct-20                	  281032	      4347 ns/op	    2654 B/op	       8 allocs/op
BenchmarkJsonMarshalMap-20                   	  237324	      4835 ns/op	    2334 B/op	       7 allocs/op
BenchmarkEasyjsonUnmarshalStruct-20          	  521766	      2341 ns/op	     480 B/op	      20 allocs/op
BenchmarkEasyjsonUnmarshalStructPtr-20       	  400022	      3033 ns/op	     800 B/op	      40 allocs/op
BenchmarkEasyjsonUnmarshalStructIntern-20    	  480001	      2540 ns/op	       0 B/op	       0 allocs/op
BenchmarkEasyjsonUnmarshalMap-20             	  307677	      3768 ns/op	    2668 B/op	      24 allocs/op
BenchmarkEasyjsonMarshalStruct-20            	  855968	      1514 ns/op	    2186 B/op	       8 allocs/op
BenchmarkEasyjsonMarshalMap-20               	  631621	      1963 ns/op	    1865 B/op	       7 allocs/op
```
