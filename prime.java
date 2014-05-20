class prime{

	public static boolean isPrime(int n){
		if(n%2==0) return false;
		for(int i=3;i<Math.sqrt(n)+1;i++){
			if(n%i==0) return false;
		}
		return true;

	}

	public static void main(String[] args){
		for(int i=1;i<100;i++){
			if(isPrime(i)) System.out.println(i);
		}
	}

}
