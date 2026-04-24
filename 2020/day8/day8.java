package day8;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.Set;
import java.util.HashSet;
import java.io.IOException;

import common.Common;

class Day8 {

    private static List<String> parse(String filename) throws IOException {
        return Files.readAllLines(Path.of(filename));
    }

    private static int solvePart1(List<String> lines) {
        Set<Integer> visited = new HashSet<>(lines.size());
        int acc = 0;
        int cur = 0;
        while (!visited.contains(cur)) {
            String[] parts = lines.get(cur).split(" ");
            String instruction = parts[0];
            int value = Integer.parseInt(parts[1]);
            visited.add(cur);
            switch (instruction) {
                case "nop": {
                    cur++;
                    break;
                }
                case "acc": {
                    cur++;
                    acc += value;
                    break;
                }
                case "jmp": {
                    cur += value;
                    break;
                }
            }
        }
        return acc;
    }

    public static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day8.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            List<String> lines = parse(filename);
            Common.time("Part 1", () -> solvePart1(lines));
        } catch(IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
