arrSize = int(input("Enter the size of the array: "))

arr = []

for i in range(arrSize):
    arr.append(int(input("Enter an element: ")))
    
arr.sort()

print("Sorted array: ", arr)