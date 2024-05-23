arrSize = int(input("Enter the size of the array: "))

arr = [int(input("Enter elements of the array: ")) for _ in range(arrSize)]

def reverse_array(arr):
    for i in range(len(arr)//2):
        arr[i], arr[len(arr) - 1 - i] = arr[len(arr) - 1 - i], arr[i]
        
    return arr

print("Reversed array: ", reverse_array(arr))