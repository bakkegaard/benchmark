#include <iostream>
#include <math.h>
#include <stdlib.h>

using namespace std;

bool isPrime(int i){
	if(i%2==0) return false;
	for(int j=3;j<((int) sqrt(i))+1;j+=2){
		if(i%j==0) return false;
	}
	return true;
}

int main(int argc, char * argv[]){

	if(argc==1){
		cout << "Requires one argument, remember target" << endl;
		return 0;
	}

	int sum=0, target=atoi(argv[1]);

	for (int i=2;i<=target;i++){
		if(isPrime(i)){
			
			sum++;

			//cout << i << endl;
		}
	
	}
	
	cout << sum << endl;
}
