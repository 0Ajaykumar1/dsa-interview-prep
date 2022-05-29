let stocks = [298, 305, 320, 301, 292];

// Lookup/search in arrays O(n)
for (let i = 0; i < stocks.length; i++) {
    if (stocks[i] === 301) {
        console.log("Searching an element in array \n", i);
    }
}

// traversal in arrays O(1)
console.log("Traversal of array");
for (let stockPrice of stocks) {
    console.log(stockPrice);
}
// adding an element to array
stocks.push(840);
// inserting an element into array O(n)
console.log("Insert an element into array");
stocks.splice(3, 0, 420);
console.log(stocks);

// deleting an element in array O(n)
console.log("Delelting an element in array");
stocks.splice(3, 1);
console.log(stocks);