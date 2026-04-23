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
        return desc.equals("") ? "NONE" : desc;
    }
}

class Day7 {

    private static Map<Bag, List<Bag>> parse(String filename) throws IOException {
        List<String> lines = Files.readAllLines(Path.of(filename));

        Map<Bag, List<Bag>> rules = new HashMap<>(lines.size());
        for (String line : lines) {
            String[] parts = line.split("contain", 2);
            Bag value = new Bag(parts[0].strip().replace(" bags", ""));
            String keyBags = parts[1].substring(0, parts[1].length() - 1);
            for (String key : keyBags.split(", ")) {
                String[] bagParts = key.strip().split(" ", 2);
                Bag keyBag;
                try {
                    int count = Integer.valueOf(bagParts[0]);
                    String desc = bagParts[1].replace(count == 1 ? " bag" : " bags", "");
                    keyBag = new Bag(count, desc);
                } catch(NumberFormatException e) {
                    keyBag = new Bag();
                }
                rules.computeIfAbsent(keyBag, k -> new ArrayList<>()).add(value);
            }
        }
        return rules;
    }

    public static int solvePart1(Map<Bag, List<Bag>> rules) {
        Set<Bag> result = new HashSet<>();
        List<Bag> queue = new ArrayList<>(List.of(new Bag("shiny gold")));
        while (queue.size() > 0) {
            Bag bag = queue.removeFirst();
            List<Bag> bags = rules.get(bag);
            if (bags == null) continue;
            for (Bag b : bags) {
                queue.add(b);
                result.add(b);
            }
        }
        return result.size();
    }

    public static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day7.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            Map<Bag, List<Bag>> rules = parse(filename);

            Common.time("Part1", () -> solvePart1(rules));
        } catch(IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
