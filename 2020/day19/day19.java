package day19;

import java.io.BufferedReader;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;

import common.Common;

sealed interface Rule permits Literal, Sequence, Alternative { }

record Literal(char c) implements Rule { }

record Sequence(List<Integer> sequence) implements Rule {
    public Sequence {
        sequence = List.copyOf(sequence);
    }
}

record Alternative(List<List<Integer>> sequences) implements Rule { }

class Day19 {

    private static Map<Integer, Rule> parseRules(BufferedReader reader) throws IOException {
        Map<Integer, Rule> rules = new HashMap<>();

        String line;
        while ((line = reader.readLine()) != null && !line.isEmpty()) {
            String[] lineParts = line.split(": ");
            int id = Integer.parseInt(lineParts[0]);
            Rule rule;
            if (lineParts[1].startsWith("\"")) {
                rule = new Literal(lineParts[1].charAt(1));
            } else if (lineParts[1].contains("|")) {
                String[] parts = lineParts[1].split(" \\| ");
                List<Integer> s1 = Arrays.stream(parts[0].split(" "))
                .map(Integer::parseInt)
                .toList();
                List<Integer> s2 = Arrays.stream(parts[1].split(" "))
                .map(Integer::parseInt)
                .toList();
                rule = new Alternative(List.of(s1, s2));
            } else {
                List<Integer> s = Arrays.stream(lineParts[1].split(" "))
                .map(Integer::parseInt)
                .toList();
                rule = new Sequence(s);
            }
            rules.put(id, rule);
        }

        return rules;
    }

    private static List<String> parseMessages(BufferedReader reader) throws IOException {
        List<String> messages = new ArrayList<>();
        String line;
        while ((line = reader.readLine()) != null) {
            messages.add(line);
        }
        return messages;
    }

    private static Set<String> isValidMessage(Map<Integer, Rule> rules, String msg, int ruleId) {
        return switch (rules.get(ruleId)) {

            case Literal l -> {
                if (!msg.isEmpty() && msg.charAt(0) == l.c()) {
                    yield Set.of(msg.substring(1));
                }
                yield Set.of();
            }

            case Sequence s -> {
                Set<String> current = Set.of(msg);

                for (Integer id : s.sequence()) {
                    Set<String> next = new HashSet<>();

                    for (String m : current) {
                        next.addAll(isValidMessage(rules, m, id));
                    }

                    current = next;

                    if (current.isEmpty()) break;
                }

                yield current;
            }

            case Alternative a -> {
                Set<String> result = new HashSet<>();

                for (List<Integer> seq : a.sequences()) {
                    Set<String> current = Set.of(msg);

                    for (Integer id : seq) {
                        Set<String> next = new HashSet<>();

                        for (String m : current) {
                            next.addAll(isValidMessage(rules, m, id));
                        }

                        current = next;

                        if (current.isEmpty()) break;
                    }

                    result.addAll(current);
                }

                yield result;
            }
        };
    }

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day19.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try (BufferedReader reader = Files.newBufferedReader(Path.of(filename))) {
            Map<Integer, Rule> rules = parseRules(reader);
            List<String> messages = parseMessages(reader);

            Common.time("Part1", () -> messages.stream()
                .filter(msg -> isValidMessage(rules, msg, 0).contains(""))
                .count()
            );

            Common.time("Part2", () -> {
                List<List<Integer>> rule8 = new ArrayList<>();
                List<List<Integer>> rule11 = new ArrayList<>();

                long last = 0;
                long current = 1;

                for (int i = 1; last != current; ++i) {
                    last = current;
                    rule8.add(Collections.nCopies(i, 42));
                    rules.put(8, new Alternative(rule8));

                    List<Integer> combined = new ArrayList<>();
                    combined.addAll(Collections.nCopies(i, 42));
                    combined.addAll(Collections.nCopies(i, 31));
                    rule11.add(combined);

                    rules.put(8, new Alternative(rule8));
                    rules.put(11, new Alternative(rule11));

                    current = messages.stream()
                    .filter(msg -> isValidMessage(rules, msg, 0).contains(""))
                    .count();
                }

                return current;
            });
        } catch (IOException e) {
            System.err.printf("ERROR: reading file '%s': %s%n", filename, e.getMessage());
            System.exit(1);
        }
    }
}
