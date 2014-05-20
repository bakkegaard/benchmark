#include <iostream>
#include <math.h>

using namespace std;

bool isPrime(int i){
	if(i%2==0) return false;
	for(int j=3;j<((int) sqrt(i))+1;j++){
		if(i%j==0) return false;
	}
	return true;
}

int main(){

	for (int i=1;i<=100;i++){
		if(isPrime(i)) cout << i << endl;
	}
}
