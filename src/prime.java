class prime{

	public static boolean isPrime(int n){
		if(n<2) return false;
		if(n==2) return true;
		if(n%2==0) return false;
		for(int i=3;i<Math.sqrt(n)+1;i+=2){
			if(n%i==0) return false;
		}
		return true;

	}

	public static void main(String[] args){

		if(args.length<1){
			System.out.println("Requires one argument, remember target");
			System.exit(0);
		}
		int sum=0, target=Integer.parseInt(args[0]);
		for(int i=2;i<target;i++){
			if(isPrime(i)){
				sum++;
				//System.out.println(i);
			}
		}
		System.out.println(sum);
	}

}
