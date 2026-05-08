package day22;

import java.io.BufferedReader;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Objects;
import java.util.Set;

import common.Common;

record Players(List<Integer> player1, List<Integer> player2) {
    public Players {
        player1 = List.copyOf(player1);
        player2 = List.copyOf(player2);
    }
}

class Day22 {

    private static Players parse(String filename) throws IOException {
        List<Integer> player1 = new ArrayList<>();
        List<Integer> player2 = new ArrayList<>();

        try (BufferedReader reader = new BufferedReader(Files.newBufferedReader(Path.of(filename)))) {
            reader.readLine(); // player 1
            String line;
            while (!(line = reader.readLine()).isEmpty()) {
                player1.add(Integer.valueOf(line));
            }

            reader.readLine(); // player 2
            while ((line = reader.readLine()) != null) {
                player2.add(Integer.valueOf(line));
            }
        }

        return new Players(player1, player2);
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

    /**
     * @param player1 Player1 numbers
     * @param player2 Player2 numbers
     * @return true if player1 won the game, false if player2.
     */
    private static boolean recursiveCombat(List<Integer> player1, List<Integer> player2) {
        Set<Integer> gameStates = new HashSet<>();
        while (!player1.isEmpty() && !player2.isEmpty()) {
            if (!gameStates.add(Objects.hash(player1, player2))) return true;

            int p1 = player1.removeFirst();
            int p2 = player2.removeFirst();

            if (p1 <= player1.size() && p2 <= player2.size()) {
                boolean player1Won = recursiveCombat(
                    new ArrayList<>(player1.subList(0, p1)),
                    new ArrayList<>(player2.subList(0, p2))
                );
                if (player1Won) {
                    player1.add(p1);
                    player1.add(p2);
                } else {
                    player2.add(p2);
                    player2.add(p1);
                }
            } else {
                if (p1 > p2) {
                    player1.add(p1);
                    player1.add(p2);
                } else {
                    player2.add(p2);
                    player2.add(p1);
                }
            }
        }

        return player2.isEmpty();
    }

    private static int solvePart2(List<Integer> player1, List<Integer> player2) {
        boolean player1Won = recursiveCombat(player1, player2);
        return calculatePoints(player1Won ? player1 : player2);
    }

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day22.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            Players players = parse(filename);
            List<Integer> player1 = players.player1();
            List<Integer> player2 = players.player2();

            Common.time("Part1", () -> solvePart1(new ArrayList<>(player1), new ArrayList<>(player2)));
            Common.time("Part2", () -> solvePart2(new ArrayList<>(player1), new ArrayList<>(player2)));
        } catch (IOException e) {
            System.err.printf("ERROR: reading file '%s': %s%n", filename, e.getMessage());
            System.exit(1);
        }
    }
}
