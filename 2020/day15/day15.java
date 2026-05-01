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

        int[] mem = new int[limit];

        for (int i = 0; i < startingNums.size() - 1; i++) {
            mem[startingNums.get(i)] = i + 1;
        }

        int lastNumber = startingNums.get(startingNums.size() - 1);
        int turn = startingNums.size();

        for (; turn < limit; turn++) {
            int prevTurn = mem[lastNumber];
            mem[lastNumber] = turn;
            lastNumber = (prevTurn == 0) ? 0 : (turn - prevTurn);
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
