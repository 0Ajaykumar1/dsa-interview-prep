// Given an array arr of size n, and two intervals x and y. Find the average of elements which lies between the given intervals inclusively.

// Input Format
// First Parameter - number n

// Second Parameter - array arr of size n

// Third Parameter - number x

// Fourth Parameter - number y

// Output Format
// Return the decimal value

function solve(n, arr, x, y) {
    // CODE HERE
    let sum = 0;
    let newarr = arr.slice(x, y + 1);
    for (let i = 0; i < newarr.length; i++) {
        sum += newarr[i];
    }
    return sum / newarr.length;
}

console.log(solve(6, [6, 2, 5, 4, 3, 1], 2, 5));