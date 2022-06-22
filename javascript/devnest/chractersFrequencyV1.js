// Given a string str1, Return the array containing the frequency of each character in the order of their occurrence.

// Input Format
// First Parameter - string str1

// Output Format
// Return the array.

// require nodejs 14 and above or ECMA 2020

function solve(str1) {
    let char_obj = {};
    for (let i of str1) {
        char_obj[i] = (char_obj[i] ?? 0) + 1; //(char_obj[i] || 0) + 1;
    };
    let arr_char_count = [];
    for (let i in char_obj) {
        arr_char_count.push(char_obj[i])
    };
    return arr_char_count;
}

console.log(solve("arfardarb"));