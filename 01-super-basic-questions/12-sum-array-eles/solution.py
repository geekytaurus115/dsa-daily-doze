arrSize = int(input("Enter the size of the array: "))
arr = [int(input("Enter ele: ")) for _ in range(arrSize)]
print(arr, end="\n")

sum = 0

for i in range(arrSize):
    sum += arr[i]
    
print("Sum of array's elements: "+str(sum))