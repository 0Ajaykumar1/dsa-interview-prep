let heros = ['spider man', 'thor', 'hulk', 'iron man', 'captain america'];

// 1. Length of the list

console.log("Length of heros list is ", heros.length);

// 2. Add 'black panther' at the end of this list

heros.push("black panther");
console.log("list after adding black panther ", heros);

// 3. You realize that you need to add 'black panther' after 'hulk',
// so remove it from the list first and then add it after 'hulk'

heros.pop();
heros.splice(3, 0, "black panther");
console.log("list after adding black panther after hulk", heros);

// 4. Now you don't like thor and hulk because they get angry easily :)
//    So you want to remove thor and hulk from list and replace them with doctor strange (because he is cool).
//    Do that with one line of code.

heros.splice(1, 2, "doctor strange");
console.log("list after replacing thor and hulk with doctor strange", heros);

// 5. Sort the heros list in alphabetical order (Hint. Use dir() functions to list down all functions available in list)

heros.sort();
console.log(heros);