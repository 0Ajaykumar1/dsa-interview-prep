stocks = [298, 305, 320, 301, 292]

# Lookup/search of array O(n)
for i in range(len(stocks)):
    if(stocks[i] == 301):
        print("Searching the index in array")
        print(i)

# traversal of array O(n)
print("Traversal of array")
for stock_price in stocks:
    print(stock_price)

# adding an element to array
stocks.append(840)
# insert of array O(n)
print("Insertion in array")
stocks.insert(3, 420)
print(stocks)

# deleting an element in array O(n)
print("deleting an element in array")
stocks.remove(420)
print(stocks)
