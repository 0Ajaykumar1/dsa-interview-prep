# Given a string str1, Return the array containing the frequency of each character in the order of their occurrence.

# Input Format
# First Parameter - string str1

# Output Format
# Return the array.

def solve(str1):
    # CODE HERE
    freq = {}
    for i in str1:
        if i in freq:
            freq[i] += 1
        else:
            freq[i] = 1
    return freq.values()


if __name__ == "__main__":
    str = input("Please enter string\n")
    print(solve(str))
