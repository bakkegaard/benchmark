using System;
class Prime 
{
	public static bool isPrime(int n){
		if(n<2) return false;
		if(n==2) return true;
		if(n%2==0) return false;
		for(int i=3;i<Math.Sqrt(n)+1;i+=2){
			if(n%i==0) return false;
		}
		return true;
	}
	static void Main(string[] args) 
	{
		if(args.Length<1){
			Console.WriteLine("Requires one argument, remember target");
			Environment.Exit(0);
		}
		int sum=0, target=int.Parse(args[0]);
		for(int i=2;i<target;i++){
			if(isPrime(i)){
				sum++;
			}
		}
		Console.WriteLine(sum);
	}

}
