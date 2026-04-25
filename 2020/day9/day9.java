package day9;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.stream.Stream;
import java.io.IOException;

import common.Common;

class Day9 {

    private static List<Long> parse(String filename) throws IOException {
        try (Stream<String> lines = Files.lines(Path.of(filename))) {
            return lines.map(Long::parseLong).toList();
        }
    }

    private static long solvePart1(List<Long> nums, int windowSize) {
        int cur = 0;
        while (cur + windowSize < nums.size()) {
            boolean found = false;
            for (int i = cur; i < cur + windowSize; ++i) {
                long a = nums.get(i);
                for (int j = i + 1; j < cur + windowSize; ++j) {
                    long b = nums.get(j);
                    if (a + b == nums.get(cur + windowSize)) {
                        cur++;
                        found = true;
                        break;
                    }
                }
                if (found) break;
            }
            if (!found) break;
        }
        return nums.get(cur + windowSize);
    }

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day9.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            List<Long> nums = parse(filename);
            Common.time("Part1", () -> solvePart1(nums, 25));
        } catch(IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
