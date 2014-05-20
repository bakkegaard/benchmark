#! /usr/bin/env python2
import sys

if len(sys.argv)<2:
    print "Requires one argument, remember target"
    sys.exit()

import math
def isPrime(n):
    if n%2==0:
        return False
    else:
        for e in range(3,int(math.sqrt(n)+1),2):
            if n%e==0:
                return False
        return True
s=0
target=int(sys.argv[1])

for i in range(2,target):
    if isPrime(i):
        s+=1
        #print (i)
print s
