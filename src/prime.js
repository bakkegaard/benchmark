#! /usr/bin/env node

if(process.argv.length<3){
	console.log("Requires one argument, remember target");
	process.exit(0);
}

function isPrime(n){
	if(i<2) return false;
	if(i==2) return true;
	if(n%2==0) return false;
	for(var i=3;i<Math.sqrt(n)+1;i+=2){
		if(n%i==0) return false;
	}
	return true;
}

var sum=0, target=parseInt(process.argv[2]);;

for(var i=2;i<target;i++){
	if(isPrime(i)){
		sum++;
		//console.log(i);
	}
}

console.log(sum);
