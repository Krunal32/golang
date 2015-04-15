##Instructions

Follow the instructions in the pdf file.


##Part 1
1. A critical sections is a block of code that contains resources which can be accessed concurrently.
2. The struct contains a pointer to a value. Access to pointers must be synchronized in critical sections because the pointer can be changed concurrently which might cause memory problems or data corruption.access violtion.
 
3. It creates a waitgroup which it adds go routines to . A operation (pop, push, len) is then selected randomly in the select block in which they are tested for access  violation. 

4. Read info on race detector

5. -race flag.

Q6. Run the TestUnsafeStack test with the data race detector enabled. Look at one of the 
warnings and inspect the stack traces. What are the two conflicting operations for your 
current run of the test? 

##Part 2

16. 
BenchmarkSafeStack	     100	  13435205 ns/op	  171085 B/op	   10038 allocs/op
BenchmarkSliceStack	     500	   4138858 ns/op	     327 B/op	       0 allocs/op
BenchmarkCspStack	      50	  48019411 ns/op	  498610 B/op	   10112 allocs/op

17.
-Arrays: usually more efficient. Can stack nodes near in contiguous memory.  Cache advantage over Linked List.
Linked List: Use less memory and doesen't occupy space that isn't currently used.
 

18. 

19.

