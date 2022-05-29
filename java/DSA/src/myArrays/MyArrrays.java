package myArrays;

import java.util.Arrays;

class MyArrays {
    public static void main(String[] args) {
        // int stocks[] = new int[5];
        int stocks[] = { 298, 305, 320, 301, 292 };
        int length = stocks.length;
        int value = 301;

        // lookup/search in arrays O(n)
        for (int i = 0; i < length; i++) {
            if (stocks[i] == value) {
                System.out.println("Searching an element in array");
                System.out.println(i);
            }
        }

        // traversal in arrays O(n)
        System.out.println("Traversal of array");
        for (int stockPrice : stocks) {
            System.out.println(stockPrice);
        }

        // adding an element into array O(n)
        System.out.println("Adding an element into array");
        value = 840;
        length = stocks.length;
        int newArr[] = new int[length + 1];
        for (int i = 0; i < length; i++) {
            newArr[i] = stocks[i];
        }

        newArr[length] = value;
        System.out.println(Arrays.toString(newArr));

        stocks = newArr;
        System.out.println(Arrays.toString(stocks));

        // inserting an element into array O(n)
        System.out.println("Inserting an element into array");
        value = 420;
        length = stocks.length;
        int newStocks[] = new int[length + 1];

        for (int i = 0; i < length + 1; i++) {
            if (i < 3) {
                newStocks[i] = stocks[i];
            } else if (i == 3) {
                newStocks[i] = value;
            } else {
                newStocks[i] = stocks[i - 1];
            }
        }

        System.out.println(Arrays.toString(newStocks));
        stocks = newStocks;
        System.out.println(Arrays.toString(stocks));

        // deleting an element from array O(n)
        System.out.println("Deleting an element from array");
        int index = 3;
        length = stocks.length;
        int newDelStocks[] = new int[length - 1];
        // for (int i = 0; i < stocks.length - 1; i++) {
        // if (i >= index) {
        // newDelStocks[i] = stocks[i + 1];
        // } else {
        // newDelStocks[i] = stocks[i];
        // }
        // }
        System.arraycopy(stocks, 0, newDelStocks, 0, 3);
        System.arraycopy(stocks, index + 1, newDelStocks, index, stocks.length - index - 1);

        System.out.println(Arrays.toString(newDelStocks));
    }
}
