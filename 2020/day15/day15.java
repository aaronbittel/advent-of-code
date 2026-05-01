package day15;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import common.Common;

class Day15 {

    private static List<Integer> parse(String filename) throws IOException {
        List<Integer> nums = new ArrayList<>();
        for (String s : Files.readString(Path.of(filename)).stripTrailing().split(",")) {
            nums.add(Integer.parseInt(s));
        }
        return nums;
    }

    private static int solve(List<Integer> startingNums, int limit) {
        if (startingNums.isEmpty()) {
            throw new IllegalArgumentException("startingNums must not be empty");
        }

        Map<Integer, Integer> mem = new HashMap<>();
        for (int i = 0; i < startingNums.size() - 1; i++) {
            int num = startingNums.get(i);
            mem.put(num, i+1);
        }

        int turn = startingNums.size();
        int lastNumber = startingNums.getLast();

        for (; turn < limit; ++turn) {
            if (mem.containsKey(lastNumber)) {
                int lastNumberTurn = mem.get(lastNumber);
                mem.put(lastNumber, turn);
                lastNumber = turn - lastNumberTurn;
            } else {
                mem.put(lastNumber, turn);
                lastNumber = 0;
            }
        }

        return lastNumber;
    }

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day15.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        System.out.println(filename);
        try {
            List<Integer> startingNums = parse(filename);

            Common.time("Part1", () -> solve(startingNums, 2020));
            Common.time("Part2", () -> solve(startingNums, 30000000));
        } catch(IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
