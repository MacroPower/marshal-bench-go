# marshal-bench-go

Collection of Go (un)marshalling benchmarks

Performed using an array of 1000 identical objects, containing 20 small strings. The results will look different from a lot of other benchmarks for this reason.

## Results

Sizes (bytes):

```text
Binary struct size: 420002
Binary map size: 481002
Gob struct size: 441186
Gob map size: 461061
Flate struct size: 383335
Flate map size: 383330
JSON struct size: 542001
JSON map size: 542001
```

Benchmarks:

```text
goarch: amd64
pkg: github.com/MacroPower/marshal-bench-go/benchmarks
cpu: Intel(R) Core(TM) i9-7900X CPU @ 3.30GHz
BenchmarkKelindarBinaryUnmarshalStruct-20               	    1032	   1158501 ns/op	  808136 B/op	   20003 allocs/op
BenchmarkKelindarBinaryUnmarshalMap-20                  	     199	   6052456 ns/op	 3318074 B/op	   64435 allocs/op
BenchmarkKelindarBinaryMarshalStruct-20                 	    1418	    854643 ns/op	 1387542 B/op	      16 allocs/op
BenchmarkKelindarBinaryMarshalMap-20                    	     262	   4328349 ns/op	 2528443 B/op	   42017 allocs/op
BenchmarkGobMarshalStruct-20                            	    1094	   1051566 ns/op	 2743882 B/op	      86 allocs/op
BenchmarkGobMarshalMap-20                               	     285	   4252649 ns/op	 3983802 B/op	   42044 allocs/op
BenchmarkGobUnmarshalStruct-20                          	     952	   1253947 ns/op	 1260472 B/op	   20288 allocs/op
BenchmarkGobUnmarshalMap-20                             	     291	   4103181 ns/op	 2369355 B/op	   26348 allocs/op
BenchmarkJsonUnmarshalStruct-20                         	     201	   5869754 ns/op	 1135649 B/op	   20015 allocs/op
BenchmarkJsonUnmarshalStructPtr-20                      	     176	   6644690 ns/op	 1187403 B/op	   40015 allocs/op
BenchmarkJsonUnmarshalMap-20                            	     162	   7339371 ns/op	 2684691 B/op	   24445 allocs/op
BenchmarkJsonMarshalStruct-20                           	     276	   4281003 ns/op	 1140857 B/op	      34 allocs/op
BenchmarkJsonMarshalMap-20                              	     246	   5175209 ns/op	 1139010 B/op	      35 allocs/op
BenchmarkEasyjsonUnmarshalStruct-20                     	     535	   2192070 ns/op	 1135424 B/op	   20011 allocs/op
BenchmarkEasyjsonUnmarshalStructPtr-20                  	     410	   2919791 ns/op	 1187169 B/op	   40011 allocs/op
BenchmarkEasyjsonUnmarshalStructIntern-20               	     493	   2358598 ns/op	  656577 B/op	      18 allocs/op
BenchmarkEasyjsonUnmarshalStructReaderFlate-20          	     145	   8129543 ns/op	 4090190 B/op	   20062 allocs/op
BenchmarkEasyjsonUnmarshalStructReaderFlateIntern-20    	     144	   8281942 ns/op	 3613716 B/op	      78 allocs/op
BenchmarkEasyjsonUnmarshalMap-20                        	     328	   3667487 ns/op	 2684274 B/op	   24441 allocs/op
BenchmarkEasyjsonUnmarshalMapReaderFlate-20             	     122	   9645270 ns/op	 5639315 B/op	   24493 allocs/op
BenchmarkEasyjsonMarshalStruct-20                       	    1465	    784030 ns/op	  556658 B/op	      32 allocs/op
BenchmarkEasyjsonMarshalStructWriter-20                 	    1141	   1018336 ns/op	 1900501 B/op	      46 allocs/op
BenchmarkEasyjsonMarshalStructWriterFlate-20            	     472	   2599024 ns/op	 1739511 B/op	      60 allocs/op
BenchmarkEasyjsonMarshalMap-20                          	     936	   1278219 ns/op	  557680 B/op	      32 allocs/op
```
