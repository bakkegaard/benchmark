# Benchmark
A simple benchmark for comparing languages "raw" speed. A simple isPrime
function is implemented in each language and the program runs a for loop from 2
to the first argument of the program (N) and checks for every loop weather the
i-th iteration is a prime. It then outputs the total number of primes from 2 to
N.

run with $ time [program] 10000000 (in seconds and real):

Programming language | Runtime (compiler instrictions)
-------------------- | -------------------------------
C++                  | 2,283 (clang++ -O3)
Java                 | 3,633 (javac)
D                    | 3,882 (ldc -O3)
rust                 | 8,992 (rustc -opt-level=3)
Javascript           | 8,988 (node)
Go                   | 9,189 (go build)
Python               | 11,992 (pypy3)
Python               | 1:22,290 (python3)
