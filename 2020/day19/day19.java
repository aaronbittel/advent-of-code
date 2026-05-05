package day19;

import java.io.BufferedReader;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import common.Common;

sealed interface Rule permits Literal, Sequence, Alternative { }

record Literal(char c) implements Rule { }

record Sequence(List<Integer> sequence) implements Rule {
    public Sequence {
        sequence = List.copyOf(sequence);
    }
}

record Alternative(List<Integer> s1, List<Integer> s2) implements Rule {
    public Alternative {
        s1 = List.copyOf(s1);
        s2 = List.copyOf(s2);
    }
}

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
                rule = new Alternative(s1, s2);
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

    private static String isValidMessage(Map<Integer, Rule> rules, String msg, int ruleId) {
        return switch (rules.get(ruleId)) {
            case Literal l -> (!msg.isEmpty() && msg.charAt(0) == l.c()) ? msg.substring(1) : null;
            case Sequence s -> {
                for (Integer id : s.sequence()) {
                    String newMsg = isValidMessage(rules, msg, id);
                    if (newMsg == null) yield null;
                    msg = newMsg;
                }
                yield msg;
            }
            case Alternative a -> {
                boolean foundOne = true;
                String tryOne = msg;
                for (Integer id : a.s1()) {
                    tryOne = isValidMessage(rules, tryOne, id);
                    if (tryOne == null) {
                        foundOne = false;
                        break;
                    }
                }

                if (foundOne) yield tryOne;

                boolean foundTwo = true;
                String tryTwo = msg;
                for (Integer id : a.s2()) {
                    tryTwo = isValidMessage(rules, tryTwo, id);
                    if (tryTwo == null) {
                        foundTwo = false;
                        break;
                    }
                }

                if (foundTwo) yield tryTwo;
                yield null;
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
                    .filter(msg -> {
                        String result = isValidMessage(rules, msg, 0);
                        return result != null && result.isEmpty();
                    })
                    .count()
            );
        } catch (IOException e) {
            System.err.printf("ERROR: reading file '%s': %s%n", filename, e.getMessage());
            System.exit(1);
        }
    }
}
