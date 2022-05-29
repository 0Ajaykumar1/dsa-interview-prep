package myArrays;
// Let us say your expense for every month are listed below,

import java.util.Arrays;

// January - 2200
// February - 2350
// March - 2600
// April - 2130
// May - 2190
// Create a list to store these monthly expenses and using that find out,

// 1. In Feb, how many dollars you spent extra compare to January?
// 2. Find out your total expense in first quarter (first three months) of the year.
// 3. Find out if you spent exactly 2000 dollars in any month
// 4. June month just finished and your expense is 1980 dollar. Add this item to our monthly expense list
// 5. You returned an item that you bought in a month of April and
// got a refund of 200$. Make a correction to your monthly expense list
// based on this
class Exercise1 {
    public static void main(String[] args) {
        int[] expenses = { 2200, 2350, 2600, 2130, 2190 };

        // 1. In Feb, how many dollars you spent extra compare to January?

        System.out.printf("Excess amount spent in February is %d \n", Math.abs(expenses[0] - expenses[1]));

        // 2. Find out your total expense in first quarter (first three months) of the
        // year.

        System.out.printf("First quarter expenses total is %d\n", expenses[0] + expenses[1] + expenses[2]);

        // 3. Find out if you spent exactly 2000 dollars in any month
        boolean test = false;
        for (int expense : expenses) {
            if (expense == 2000) {
                test = true;
                break;
            }
        }
        System.out.printf("Is any month expenses are $2000 %b\n", test);

        // 4. June month just finished and your expense is 1980 dollar. Add this item to
        // our monthly expense list
        int[] newExpenses = new int[expenses.length + 1];
        int i;
        for (i = 0; i < expenses.length; i++) {
            newExpenses[i] = expenses[i];
        }
        newExpenses[i] = 1980;

        System.out.println(Arrays.toString(newExpenses));

        // 5. You returned an item that you bought in a month of April and
        // got a refund of 200$. Make a correction to your monthly expense list
        // based on this

        newExpenses[3] = newExpenses[3] - 200;
        System.out.println(Arrays.toString(newExpenses));
    }
}