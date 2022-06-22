# You are given two integers, a and b.

# Let x be the integer division and y be the decimal(float) division of a and b.

# Return the product of x and y.

# Input Format
# First Parameter - Integer a

# Second Parameter - Integer b

# Output Format
# Return the decimal product of integer division x and decimal division y of a and b.

def solve(a, b):
    # CODE HERE
    return ((a//b)*(a/b))


if __name__ == "__main__":
    a = int(input("Please enter first parameter\n"))
    b = int(input("Please enter second parameter\n"))

    print(solve(a, b))
