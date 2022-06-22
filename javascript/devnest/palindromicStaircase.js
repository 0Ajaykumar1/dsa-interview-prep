// Given a Positive Integer n, return an array of string containing the palindromic string of intergers.

// Input Format
// First Parameter - number n

// Output Format
// Return the array of string.

function solve(n) {
    // CODE HERE
    let arr = [];
    for (let i = 1; i <= n; i++) {
        arr.push(Math.pow(parseInt((Math.pow(10, i) - 1) / 9), 2));
    }
    return arr;
}

console.log(solve(5));