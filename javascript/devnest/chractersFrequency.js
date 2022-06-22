// Given a string str1, Return the array containing the frequency of each character in the order of their occurrence.

// Input Format
// First Parameter - string str1

// Output Format
// Return the array.

function solve(str1) {
    let freq = {};
    for (let i of str1) {
        if (i in freq) {
            freq[i] += 1;
        } else {
            freq[i] = 1;
        }
    }
    return Object.values(freq);
}

console.log(solve("arfardarb"));