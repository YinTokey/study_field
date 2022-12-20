package com.hello;

import groovy.lang.Singleton;

import java.util.*;
import java.lang.Character;
import java.util.stream.Collectors;

class FreshJuice {
    enum FreshJuiceSize{
        SMALL,
        MEDIUM,
        LARGE
    }
    FreshJuiceSize size;
}

/**
 *
 */
public class Hello {

    private static  Hello single = new Hello();

    private Hello() {}

    public static Hello getInstance() {
        return single;
    }

    public static int findDuplicate(String input) {
        HashMap<Character, Integer> mp = new HashMap<Character, Integer>();
        for (int i = 0; i < input.length(); i++) {
            char key = input.charAt(i);
            if (mp.get(key) != null) {
                return i;
            }
            mp.put(input.charAt(i),1);
        }
        return -1;
    }

    public static boolean twoSum(Integer[] nums, int target) {
        Set<Integer> twoSumSet = new HashSet<>();
        for (int num : nums) {
            if (twoSumSet.contains(num)) {
                return true;
            }
            twoSumSet.add(target-num);
        }
        return false;
    }

    // Try to use O(1) space
    public static String reverseString(String s) {
        StringBuilder builder = new StringBuilder();
        for (int i = s.length() -1; i >= 0 ; i--) {
            builder.append(s.charAt(i));
        }
        return builder.toString();
    }

    public static List<Integer> topK(Integer[] input, int k) {
        TreeSet<Integer> tree = new TreeSet<>();
        for (int num : input) {
            tree.add(num);
            if (tree.size() > k) {
                tree.pollFirst();
            }
        }
        System.out.println(tree);
        return tree.stream().collect(Collectors.toList());
    }

    /**
     *
     * @param args
     */
    public static void main(String[] args) {

        FreshJuice juice = new FreshJuice();
        juice.size = FreshJuice.FreshJuiceSize.LARGE;

        // System.out.println("Size: " + juice.size);
//        String name = "William";
//        name = "Willam B";
//        System.out.println(name);
//
//        StringBuilder nameStringBuilder = new StringBuilder();
//        nameStringBuilder.append("a");
//        nameStringBuilder.append(" b");
//        System.out.println(nameStringBuilder);
//
//        List<Integer> number = new ArrayList<>(Arrays.asList(1,2,3));
//        System.out.println(number);

//         String t1 = "abc";
//         Integer result = findDuplicate(t1);
//         System.out.println(result);

//         Integer li[] = new Integer[] {1,2,3,4};
//         System.out.println(twoSum(li,18));

//         String input = "hello";
//         System.out.println(reverseString(input));
         Integer li[] = new Integer[] {5,10,22,100,8};
         System.out.println(topK(li,2));
    };
}

