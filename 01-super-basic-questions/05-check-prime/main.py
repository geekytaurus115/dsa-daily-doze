import math

num = int(input("Enter a number above 2: "))

isPrime = True

for i in range(2, int(math.sqrt(num))+1):
    if num % i == 0:
        isPrime = False
        break
    
if isPrime:
    print("The number {} is Prime.".format(num))
else:
    print("The number {} is Not Prime.".format(num))