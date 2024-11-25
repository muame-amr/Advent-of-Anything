# Faster Lock Combination

## Reference

[[198] Close Up On How To Decode A Dial Combination Lock In 8 Attempts Or Less](https://www.youtube.com/watch?v=27rE5ZvWLU0)

## Step-by-step

1. Find sticky number
2. Find two guess number (0 - 11)
3. Find 1st digit by adding 5 to sticky number
4. Find 3rd digit by doing the _Math_
5. Find 8 possible 2nd digit
6. Attempt unlock

## Example

```sh
sticky number = 28
guess number 1= 6
guess number 2= 7

math

Part A: 33
Part B: 33 / 4 = 8 remaider of 1

6   16  26  36
7   17  27  37

17 / 4 = 4 remainder of 1
37 / 4 = 9 remainder of 1

Part C:

First row:  3   11  27  35
Second row: 7   23  31  39

3rd digit: 17
```
