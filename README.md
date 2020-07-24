# Benchmark
A simple benchmark for comparing languages "raw" speed. A simple isPrime
function is implemented in each language and the program runs a for loop from 2
to the first argument of the program (N) and checks for every loop weather the
i-th iteration is a prime. It then outputs the total number of primes from 2 to
N.

run with $ time [program] 10000000 (in seconds and real):

Programming language (compiler) | Runtime 
------------------------------- | --------
C++ (gcc)           				  | 2.251
Rust                				  | 3.220
Java                				  | 3.552
D (ldc)             				  | 3.823
Node                				  | 4.712
C++ (clang)         				  | 5.587
Go                  				  | 8.163
C# (mcs)            				  | 9.802
Python (pypy)       				  | 11.830
Python              				  | 50.530
Ruby                				  | 102.693
