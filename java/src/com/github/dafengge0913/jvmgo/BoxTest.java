package com.github.dafengge0913.jvmgo;

import java.util.ArrayList;
import java.util.List;

public class BoxTest {
    public static void main(String[] args) {
        List<Integer> list = new ArrayList<>();
        list.add(1);
        list.add(2);
        list.add(66666);
        System.out.println(list.toString());
        for (int x : list) {
            System.out.println(x);
        }
        System.out.println(sun.misc.VM.getSavedProperty("java.lang.Integer.IntegerCache.high"));
    }
}
