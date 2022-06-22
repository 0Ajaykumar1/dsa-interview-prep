# Given an array arr of size n, and two intervals x and y. Find the average of elements which lies between the given intervals inclusively.

# Input Format
# First Parameter - number n

# Second Parameter - array arr of size n

# Third Parameter - number x

# Fourth Parameter - number y

# Output Format
# Return the decimal value

def solve(n, arr, x, y):
    # CODE HERE
    sum = 0

    for i in range(x, y+1):
        sum += arr[i]
    return (sum/(y-x+1))


if __name__ == "__main__":
    n = int(input("Please enter array length\n"))
    arr = input("Please provide array\n")
    x = int(input("Please provide interval start\n"))
    y = int(input("Please provide interval end\n"))
    arr = list(map(int, arr.split()))
    print(solve(n, arr, x, y))
