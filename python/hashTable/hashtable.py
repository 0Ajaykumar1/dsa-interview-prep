stock_prices = []
with open("stock_prices.csv", "r") as f:
    for line in f:
        tokens = line.split(',')1
        day = tokens[0]
        price = float(tokens[1])
        stock_prices.append([day, price])

if __name__ == "__main__":
    stock_prices
    stock_prices[0]
