num = int(input("Enter a number: "))

original = num
reverse = 0

while num != 0:
    digit = num % 10
    reverse = reverse * 10 + digit
    num //= 10
    
# print("Reverse: ", reverse)

if reverse == original:
    print("Palindrom")
else:
    print("Not Palindrom")