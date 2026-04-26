package day10;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;
import java.io.IOException;

import common.Common;

class Day10 {

    private static List<Integer> parse(String filename) throws IOException {
        try(Stream<String> lines = Files.lines(Path.of(filename))) {
            return lines.map(Integer::valueOf).sorted().collect(Collectors.toList());
        }
    }

    private static int solvePart1(List<Integer> nums) {
        int prev = 0;
        int count1 = 0;
        int count3 = 1;
        for (Integer i : nums) {
            switch (i - prev) {
                case 0:
                case 2: break;
                case 1: {
                    count1++;
                    break;
                }
                case 3: {
                    count3++;
                    break;
                }
                default: throw new IllegalStateException("Illegal state");
            }
            prev = i;
        }
        return count1 * count3;
    }

    private static long solvePart2(List<Integer> nums) {
        long result = 0;
        nums.addFirst(0);
        nums.addLast(nums.getLast() + 3);
        List<Integer> queue = new ArrayList<>(List.of(0));

        while (!queue.isEmpty()) {
            int index = queue.removeFirst();
            int current = nums.get(index);
            if (current == nums.getLast()) {
                result++;
                continue;
            }
            for (int i = index + 1; i <= index + 3 && i < nums.size(); ++i) {
                if (nums.get(i) - current > 3) break;
                queue.add(i);
            }
        }

        return result;
    }

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day10.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            List<Integer> nums = parse(filename);

            Common.time("Part1", () -> solvePart1(nums));
            Common.time("Part2", () -> solvePart2(nums));
        } catch(IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
