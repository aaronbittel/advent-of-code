package common;

import java.util.function.Supplier;

public class Common {
    public static <T> void time(String label, Supplier<T> task) {
        long start = System.nanoTime();
        T result = task.get();
        long end = System.nanoTime();
        double seconds = (end - start) / 1_000_000_000.0;
        System.out.printf("%s: %s, took %.5f seconds%n", label, result, seconds);
    }

    public static <T> T timeParsing(Supplier<T> task) {
        long start = System.nanoTime();
        T result = task.get();
        long end = System.nanoTime();
        double seconds = (end - start) / 1_000_000_000.0;
        System.out.printf("Parsing took %.5f seconds%n", seconds);
        return result;
    }
}
