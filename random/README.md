# Random

## RandDistinct


### Test
```text
=== RUN   TestRandDistinct
--- PASS: TestRandDistinct (0.01s)
PASS
coverage: 100.0% of statements in ../../algogo/...

Process finished with exit code 0
```

### Bench
```text
goos: windows
goarch: amd64
pkg: github.com/guitarpawat/algogo/random
BenchmarkRandDistinct10-4        	 3000000	       334 ns/op
BenchmarkRandDistinct100-4       	  500000	      3618 ns/op
BenchmarkRandDistinct1000-4      	   50000	     33889 ns/op
BenchmarkRandDistinct10000-4     	    5000	    341875 ns/op
BenchmarkRandDistinct100000-4    	     300	   3696667 ns/op
BenchmarkRandDistinct1000000-4   	      20	  87402900 ns/op
PASS

Process finished with exit code 0
```