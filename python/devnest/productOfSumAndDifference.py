# You are given two integers, a and b.
# Compute the Sum and Difference of a and b. Return the Product of Sum and Difference.

# Input Format
# First Parameter - Integer a

# Second Parameter - Integer b

# Output Format
# Return the integer.


def solve(a, b):
    # CODE HERE
    return ((a+b)*(a-b))


if __name__ == "__main__":
    a = int(input("Please enter first parameter\n"))
    b = int(input("Please enter second parameter\n"))

    print(solve(a, b))
