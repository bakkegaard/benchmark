import std.stdio, std.math;

bool isPrime(int n){
	if(n%2==0) return false;
	for(int i=3;i<sqrt(cast(float) n)+1;i++){
		if(n%i==0) return false;
	}
	return true;
}

void main(){
	for(int i=1;i<100;i++){
		if(isPrime(i)) writeln(i);
	}
}
