import java.util.List;
import java.util.Collections;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.function.IntSupplier;

class Day1 {

    private static final int TARGET = 2020;

    private static List<Integer> parse(String filename) throws IOException {
        return Files.lines(Path.of(filename))
            .map(Integer::valueOf)
            .sorted()
            .toList();
    }

    private static int solvePart1(List<Integer> numbers) {
        for (int i = 0; i < numbers.size() - 1; ++i) {
            int a = numbers.get(i);
            for (int j = i + 1; j < numbers.size(); ++j) {
                int b = numbers.get(j);
                int result = TARGET - (a + b);
                if (result == 0) {
                    return a * b;
                } else if (result < 0) {
                    break;
                }
            }
        }
        throw new IllegalStateException("No solution found");
    }

    private static int solvePart2(List<Integer> numbers) {
        for (int i = 0; i < numbers.size() - 2; ++i) {
            int a = numbers.get(i);
            for (int j = i + 1; j < numbers.size() - 1; ++j) {
                int b = numbers.get(j);
                for (int k = j + 1; k < numbers.size(); ++k) {
                    int c = numbers.get(k);
                    int result = TARGET - (a + b + c);
                    if (result == 0) {
                        return a * b * c;
                    } else if (result < 0) {
                        break;
                    }
                }
            }
        }
        throw new IllegalStateException("No solution found");
    }

    private static void time(String label, IntSupplier task) {
        long start = System.nanoTime();
        int result = task.getAsInt();
        long end = System.nanoTime();
        double seconds = (end - start) / 1_000_000_000.0;
        System.out.printf("%s: %d, took %.5f seconds%n", label, result, seconds);
    }

    public static void main(String[] args) {
        if (args.length < 1) {
            System.err.println(String.format("Usage: java %s <input>", Day1.class.getSimpleName()));
            System.exit(1);
        }

        String filename = args[0];
        try {
            List<Integer> numbers = Day1.parse(filename);

            Day1.time("Part1", () -> Day1.solvePart1(numbers));
            Day1.time("Part2", () -> Day1.solvePart2(numbers));
        } catch(IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
