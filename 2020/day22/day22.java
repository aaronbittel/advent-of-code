package day22;

import java.io.BufferedReader;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.List;

import common.Common;

class Day22 {

    private static List<List<Integer>> parse(String filename) throws IOException {
        List<List<Integer>> players = new ArrayList<>();
        try (BufferedReader reader = new BufferedReader(Files.newBufferedReader(Path.of(filename)))) {
            reader.readLine(); // player 1
            String line;
            List<Integer> player1 = new ArrayList<>();
            while (!(line = reader.readLine()).isEmpty()) {
                player1.add(Integer.valueOf(line));
            }
            players.add(player1);

            reader.readLine(); // player 2
            List<Integer> player2 = new ArrayList<>();
            while ((line = reader.readLine()) != null) {
                player2.add(Integer.valueOf(line));
            }
            players.add(player2);
        }

        if (players.size() != 2) {
            throw new IllegalStateException("there must be 2 players");
        }
        return players;
    }

    private static int calculatePoints(List<Integer> nums) {
        int result = 0;
        int multiplier = nums.size();
        for (Integer n : nums) {
            result += n * multiplier;
            multiplier--;
        }
        return result;
    }

    private static int solvePart1(List<Integer> player1, List<Integer> player2) {
        while (!player1.isEmpty() && !player2.isEmpty()) {
            int p1 = player1.removeFirst();
            int p2 = player2.removeFirst();
            if (p1 > p2) {
                player1.add(p1);
                player1.add(p2);
            } else {
                player2.add(p2);
                player2.add(p1);
            }
        }

        return calculatePoints(player1.isEmpty() ? player2 : player1);
    }

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day22.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            List<List<Integer>> players = parse(filename);
            List<Integer> player1 = players.get(0);
            List<Integer> player2 = players.get(1);

            Common.time("Part1", () -> solvePart1(player1, player2));
        } catch (IOException e) {
            System.err.printf("ERROR: reading file '%s': %s%n", filename, e.getMessage());
            System.exit(1);
        }
    }
}
