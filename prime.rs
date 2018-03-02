use std::env;
use std::f32;

fn is_prime(n:u32)->bool{
    if n%2==0 {
        return false;
    }
    let mut i = 3;
    while i<((n as f32).sqrt() as u32)+1{
        if n%i==0 {
            return false;
        }
        i+=2;
    }
    return true;
}

fn main(){
    let args: Vec<String> = env::args().collect();
    if args.len()<2 {
        println!("Requires one argument, remember target");
    }
    let mut sum = 0;
    let target=args[1].parse::<u32>().unwrap();

        for i in 2..target{
            if is_prime(i) {
                sum+=1;
            }
        }
    println!("{}", sum);
}
