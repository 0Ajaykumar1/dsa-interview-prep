# You are given a String str. Return the String by Concatenating it with String "Hello, " and an exclamation mark at the end “!”.

# Input Format
# First Parameter - String str

# Output Format
# Return a String

def solve(str):
    # CODE HERE
    return ("Hello, " + str + "!")


if __name__ == "__main__":
    str = input("Please enter your name\n")
    print(solve(str))
