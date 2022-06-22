# Given a Positive Integer n, return an array of string containing the palindromic string of intergers.

# Input Format
# First Parameter - number n

# Output Format
# Return the array of string.

def solve(n):
    # CODE HERE
    arr = []
    for i in range(1, n+1):
        arr.append(((10**i-1)//9)**2)
    return arr


if __name__ == "__main__":
    n = int(input("Please enter string\n"))
    print(solve(n))
