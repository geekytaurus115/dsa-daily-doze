### Check if a given number is prime or not.

- First, check if the number is less than 2. If it is, it cannot be prime. If it's 2, it's prime.
- If the number is **not divisible by any integer other than 1 and itself**, then it is prime.

1. For i from 2 to sqrt(n) inclusive:
   a. If n is divisible by i, return false.
2. Return true.

```
Conferring to the definition of prime number, which states that a number should have exactly two factors, but number 1 has one and only one factor. Thus 1 is not considered a Prime number.
```

The prime factorization of 72 is \( 72 = 2^3 \times 3^2 \).

When we group the prime factors in pairs, we have:

\( 72 = (2 \times 2) \times (2 \times 2) \times (3 \times 3) \).

Taking one factor from each pair, we get:

\( \sqrt{72} = 2 \times 2 \times 3 = 6\sqrt{2} \).
