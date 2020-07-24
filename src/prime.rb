def isPrime(n)
  if n<2 
    return false
  end
  if n == 2
    return true
  end
  if n%2==0
    return false
  end
  i=3
  while i<Math.sqrt(n)+1
    if n % i == 0
      return false
    end
    i+=2
  end

  return true
end

if ARGV.length != 1 
  puts "requires one argument, remember target"
  exit(0)
end

target = ARGV[0].to_i

counter = 0
for i in 1..target
  if isPrime(i)
    counter+=1
  end
end
puts counter
