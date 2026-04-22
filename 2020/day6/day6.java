package day6;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Arrays;
import java.util.List;
import java.util.Set;
import java.util.HashSet;
import java.util.stream.Collectors;
import java.io.IOException;

import common.Common;

class Day6 {

    private static List<String> parse(String filename) throws IOException {
        return Arrays.asList(Files.readString(Path.of(filename)).split("\\R\\R"));
    }

    private static long countDistinctLetters(String group) {
        return group.chars()
            .distinct()
            .filter(c -> c >= 'a' && c <= 'z')
            .count();
    }

    private static int countCommonAnswers(String group) {
        return group.lines()
            .map(line -> line.chars()
                .boxed()
                .collect(Collectors.toSet())
            )
            .reduce((a, b) -> {
                Set<Integer> res = new HashSet<>(a);
                res.retainAll(b);
                return res;
            })
            .map(Set::size)
            .orElse(0);
    }

    public static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day6.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            List<String> groups = Day6.parse(filename);

            Common.time("Part1", () -> groups.stream()
                    .mapToLong(Day6::countDistinctLetters)
                    .sum()
            );

            Common.time("Part2", () -> groups.stream()
                .mapToInt(Day6::countCommonAnswers)
                .sum()
            );
        } catch (IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
