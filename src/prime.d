import std.stdio, std.math,std.conv;

bool isPrime(int n){
	if(n<2) return false;
	if(n==2) return true;
	if(n%2==0) return false;
	for(int i=3;i<sqrt(cast(float) n)+1;i+=2){
		if(n%i==0) return false;
	}
	return true;
}

int main(char[][] args){
	if(args.length<2){
		writeln("Requires one argument, remember target");
		return 0;
	}
	int sum=0,target=to!(int)(args[1]);

	for(int i=2;i<target;i++){
		if(isPrime(i)){
			sum++;
			//writeln(i);
		}
	}
	writeln(sum);

	return 0;
}
