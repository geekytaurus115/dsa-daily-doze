n = int(input("Enter the size of the arry: "))

arr = []

# for _ in range(n):
#     arr.append(int(input("Enter element: ")))

arr = [int(input("Enter element: ")) for _ in range(n)]
    
max_ele = max(arr)
print(max_ele)
