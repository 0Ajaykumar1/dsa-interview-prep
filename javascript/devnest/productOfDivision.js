// You are given two integers, a and b.

// Let x be the integer division and y be the decimal(float) division of a and b.

// Return the product of x and y.

// Input Format
// First Parameter - Integer a

// Second Parameter - Integer b

// Output Format
// Return the decimal product of integer division x and decimal division y of a and b.

function solve(a, b) {
    // CODE HERE
    return ((a / b) * parseInt(a / b))
}

console.log(solve(21, 5));