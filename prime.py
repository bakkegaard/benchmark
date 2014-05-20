#! /usr/bin/env python2

import math
def isPrime(n):
    if n%2==0:
        return False
    else:
        for e in range(3,int(math.sqrt(n)+1)):
            if n%e==0:
                return False
        return True

for i in range(1,100):
    if isPrime(i):
        print (i)
