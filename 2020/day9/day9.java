package day9;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.stream.Stream;
import java.util.stream.Collectors;
import java.io.IOException;

import common.Common;

class Day9 {

    private static List<Long> parse(String filename) throws IOException {
        try (Stream<String> lines = Files.lines(Path.of(filename))) {
            return lines.map(Long::parseLong).collect(Collectors.toList());
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

    private static long solvePart2(List<Long> nums, long invalidNumber) {
        int left = 0;
        int right = 1;
        long currentSum = nums.get(left) + nums.get(right);
        while (currentSum != invalidNumber) {
            if (currentSum < invalidNumber) {
                right += 1;
                currentSum += nums.get(right);
            } else if (currentSum > invalidNumber) {
                currentSum -= nums.get(left);
                left += 1;
            }
        }
        long min = Long.MAX_VALUE;
        long max = Long.MIN_VALUE;
        for (int i = left; i <= right; ++i) {
            long val = nums.get(i);
            if (val < min) min = val;
            if (val > max) max = val;
        }
        return min + max;
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
            Common.time("Part2", () -> {
                long invalidNumber = solvePart1(nums, 25);
                return solvePart2(nums, invalidNumber);
            });
        } catch(IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
