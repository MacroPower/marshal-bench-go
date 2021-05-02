# marshal-bench-go

Collection of Go (un)marshalling benchmarks

## Results

```
goarch: amd64
pkg: github.com/MacroPower/marshal-bench-go/benchmarks
cpu: Intel(R) Core(TM) i9-7900X CPU @ 3.30GHz
BenchmarkKelindarBinaryUnmarshalStruct-20       	  809617	      1299 ns/op	     800 B/op	      21 allocs/op
BenchmarkKelindarBinaryUnmarshalStructPtr-20    	  406520	      3055 ns/op	    1793 B/op	      62 allocs/op
BenchmarkKelindarBinaryUnmarshalMap-20          	  196681	      6108 ns/op	    3365 B/op	      66 allocs/op
BenchmarkKelindarBinaryMarshalStruct-20         	  922657	      1283 ns/op	    1617 B/op	       6 allocs/op
BenchmarkKelindarBinaryMarshalMap-20            	  259143	      4564 ns/op	    2449 B/op	      47 allocs/op
BenchmarkJsonUnmarshalStruct-20                 	  173958	      7009 ns/op	     984 B/op	      23 allocs/op
BenchmarkJsonUnmarshalStructPtr-20              	  156043	      7619 ns/op	    1144 B/op	      43 allocs/op
BenchmarkJsonUnmarshalMap-20                    	  144578	      8335 ns/op	    2907 B/op	      28 allocs/op
BenchmarkJsonMarshalStruct-20                   	  243346	      5042 ns/op	    2077 B/op	       7 allocs/op
BenchmarkJsonMarshalMap-20                      	  224697	      5491 ns/op	    1756 B/op	       6 allocs/op
BenchmarkEasyjsonUnmarshalStruct-20             	  521491	      2338 ns/op	     480 B/op	      20 allocs/op
BenchmarkEasyjsonUnmarshalStructPtr-20          	  399906	      3062 ns/op	     800 B/op	      40 allocs/op
BenchmarkEasyjsonUnmarshalStructIntern-20       	  480710	      2522 ns/op	       0 B/op	       0 allocs/op
BenchmarkEasyjsonUnmarshalMap-20                	  292749	      3851 ns/op	    2667 B/op	      24 allocs/op
BenchmarkEasyjsonMarshalStruct-20               	  855596	      1371 ns/op	    1609 B/op	       7 allocs/op
BenchmarkEasyjsonMarshalMap-20                  	  631602	      1826 ns/op	    1289 B/op	       6 allocs/op
```
