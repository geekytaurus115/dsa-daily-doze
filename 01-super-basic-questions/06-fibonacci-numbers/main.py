seqNum = int(input("Enter the sequence number > 1, to print fibonacci numbers: "))

fib0 = 0
fib1 = 1

print("First {} fibonacci numbers are: {}, {}".format(seqNum, fib0, fib1), end="")
for i in range(3, seqNum + 1):
    fibCurr = fib0 + fib1
    fib0 = fib1
    fib1 = fibCurr
    
    print(", {}".format(fibCurr), end="")
print()