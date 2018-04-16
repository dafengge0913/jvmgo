package com.github.dafengge0913.jvmgo;

public class MethodInvokeTest implements Runnable {


    public static void main(String[] args) {
        new MethodInvokeTest().test();
    }

    public void test() {
        MethodInvokeTest.staticMethod();
        MethodInvokeTest test = new MethodInvokeTest();
        test.instanceMethod();
        super.equals(null);
        this.run();
        ((Runnable) test).run();
    }

    public static void staticMethod() {
    }

    private void instanceMethod() {
    }

    @Override
    public void run() {
    }
}
