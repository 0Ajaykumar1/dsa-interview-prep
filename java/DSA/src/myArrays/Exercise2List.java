package myArrays;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

class Exercise2List {

    public static void main(String[] args) {
        String[] herosarray = { "spider man", "thor", "hulk", "iron man", "captain america" };
        List<String> heros = new ArrayList<>(Arrays.asList(herosarray));

        // 1. Length of the list;
        System.out.printf("Length of heros list is %d\n", heros.size());

        // 2. Add 'black panther' at the end of this list
        heros.add("black panther");
        System.out.print("list after adding black panther ");
        System.out.println(heros);

        // 3. You realize that you need to add 'black panther' after 'hulk',
        // so remove it from the list first and then add it after 'hulk'

        heros.remove("black panther");
        heros.add(3, "black panther");
        System.out.printf("list after adding black panther after hulk %s\n", heros);

        // 4. Now you don't like thor and hulk because they get angry easily :)
        // So you want to remove thor and hulk from list and replace them with doctor
        // strange (because he is cool).
        // Do that with one line of code.

        heros.set(1, "doctor strange");
        heros.remove(2);
        System.out.printf("list after replacing thor and hulk with doctor strange %s\n", heros);

        // 5. Sort the heros list in alphabetical order (Hint. Use dir() functions to
        // list down all functions available in list)
        Collections.sort(heros);
        System.out.println(heros);
    }
}
