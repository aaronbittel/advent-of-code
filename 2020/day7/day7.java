package day7;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Arrays;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.util.HashMap;
import java.util.Set;
import java.util.HashSet;
import java.util.Objects;
import java.util.stream.Collectors;
import java.io.IOException;

import common.Common;

class Bag {
    private int count;
    private String desc;

    public Bag(int count, String desc) {
        this.count = count;
        this.desc = desc;
    }

    public Bag(String desc) {
        this(0, desc);
    }

    public Bag() {
        this(0, "");
    }

    public boolean isEmpty() {
        return desc.isEmpty() && count == 0;
    }

    public int getCount() {
        return count;
    }

    public String getDesc() {
        return desc;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (!(o instanceof Bag)) return false;
        Bag other = (Bag)o;
        return desc.equals(other.desc);
    }

    @Override
    public int hashCode() {
        return Objects.hashCode(desc);
    }

    @Override
    public String toString() {
        if (isEmpty()) return "NONE";
        return desc + " (" + count + ")";
    }
}

class Day7 {

    private static Map<String, List<Bag>> parse(String filename) throws IOException {
        List<String> parsedLines = Files.lines(Path.of(filename))
            .map(line ->
                line.substring(0, line.length() - 1)
                    .replaceAll(" bags?", ""))
            .collect(Collectors.toList());
        Map<String, List<Bag>> rules = new HashMap<>();
        for (String line : parsedLines) {
            String[] parts = line.split(" contain ");
            String key = parts[0].strip();
            for (String val : parts[1].strip().split(", ")) {
                String[] valParts = val.split(" ", 2);
                Bag bag;
                try {
                    int count = Integer.parseInt(valParts[0]);
                    bag = new Bag(count, valParts[1].strip());
                } catch(NumberFormatException e) {
                    bag = new Bag();
                }
                rules.computeIfAbsent(key, k -> new ArrayList<>()).add(bag);
            }
        }
        return rules;
    }

    public static int solvePart1(Map<String, List<String>> rules) {
        Set<String> result = new HashSet<>();
        List<String> queue = new ArrayList<>(List.of("shiny gold"));
        while (queue.size() > 0) {
            String desc = queue.removeFirst();
            List<String> descs = rules.get(desc);
            if (descs == null) continue;
            for (String d : descs) {
                queue.add(d);
                result.add(d);
            }
        }
        return result.size();
    }

    private static int solvePart2(Map<String, List<Bag>> rules, String cur) {
        int count = 0;
        List<Bag> bags = rules.get(cur);
        if (bags == null) return count;
        for (Bag b : bags) {
            count += b.getCount() * (1 + solvePart2(rules, b.getDesc()));
        }

        return count;
    }

    public static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day7.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            Map<String, List<Bag>> rules= parse(filename);

            Common.time("Part1", () -> {
                Map<String, List<String>> part1Map = new HashMap<>();
                for (Map.Entry<String, List<Bag>> entry : rules.entrySet()) {
                    for (Bag bag : entry.getValue()) {
                        part1Map.computeIfAbsent(bag.getDesc(), k -> new ArrayList<>()).add(entry.getKey());
                    }
                }
                return solvePart1(part1Map);
            });

            Common.time("Part2", () -> solvePart2(rules, "shiny gold"));
        } catch(IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
