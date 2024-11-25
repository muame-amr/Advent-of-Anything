import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Scanner;

public class Solution {
    public static void main(String[] args) throws FileNotFoundException {
        int total = 0;
        ArrayList<String> inputs = readInput("input.txt");
        for (String s : inputs) {
            total += extractSum(s);
        }

        System.out.println("[Answer]: " + total); // Answer: 54667
    }

    private static ArrayList<String> readInput(String filename) throws FileNotFoundException {
        Scanner s = new Scanner(new File(filename));
        ArrayList<String> list = new ArrayList<>();

        while (s.hasNext()) {
            list.add(s.next());
        }

        s.close();

        return list;
    }

    private static int extractSum(String s) {

        String digitOnlyString = s.replaceAll("[^0-9]", "");

        char firstDigit = digitOnlyString.charAt(0);
        char lastDigit = digitOnlyString.charAt(digitOnlyString.length() - 1);
        String concat = "" + firstDigit + lastDigit;

        return Integer.parseInt(concat);
    }
}