num = int(input("Enter a number: "))

factResult = 1

def factorial(n):
    if n == 0:
        return 1
    return n * factorial(n - 1)


for i in range(num, 1, -1):
    factResult *= i
# print("Factorial of ", num, "is ", factResult)
print("Factorial of {} is {}".format(num, factResult))
print("Factorial of {} is {} using recurssion".format(num, factResult))