
function isPrime(n){
	if(n%2==0) return false;
	for(var i=3;i<Math.sqrt(n)+1;i++){
		if(n%i==0) return false;
	}
	return true;
}

for(var i=1;i<100;i++){
	if(isPrime(i)) console.log(i);
}
