# SortedSearch

### Bench

```text
goos: windows
goarch: amd64
pkg: github.com/guitarpawat/algogo/sorted_search
BenchmarkLinearSearch1-4         	200000000	         5.72 ns/op
BenchmarkLinearSearch10-4        	20000000	        59.7 ns/op
BenchmarkLinearSearch100-4       	  500000	      3706 ns/op
BenchmarkLinearSearch1000-4      	    5000	    225703 ns/op
BenchmarkLinearSearch10000-4     	     100	  18151483 ns/op
BenchmarkLinearSearch100000-4    	       1	1913602300 ns/op
BenchmarkBinarySearch1-4         	200000000	         6.50 ns/op
BenchmarkBinarySearch10-4        	20000000	       103 ns/op
BenchmarkBinarySearch100-4       	 1000000	      1995 ns/op
BenchmarkBinarySearch1000-4      	   30000	     52945 ns/op
BenchmarkBinarySearch10000-4     	    2000	    787602 ns/op
BenchmarkBinarySearch100000-4    	     200	   8941080 ns/op
BenchmarkBinarySearch1000000-4   	      20	  97987985 ns/op
PASS

Process finished with exit code 0
```