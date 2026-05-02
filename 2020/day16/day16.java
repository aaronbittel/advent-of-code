package day16;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;

import common.Common;

record Range(int start, int end) { 
    public boolean contains(int num) {
        return num >= start && num <= end;
    }
}

record ParsedInput(List<Range> ranges, List<List<Integer>> tickets) { }

class Day16 {

    private static ParsedInput parse(String filename) throws IOException {
        List<String> lines = Files.readAllLines(Path.of(filename));
        List<Range> ranges = new ArrayList<>();

        int i = 0;
        for (; i < lines.size(); ++i) {
            String line = lines.get(i);
            if (line.isEmpty()) {
                ++i;
                break;
            }
            ranges.addAll(Arrays.asList(parseLineToRanges(line)));
        }
        ++i; // your ticket:

        List<List<Integer>> tickets = new ArrayList<>();

        for (; i < lines.size(); ++i) {
            String line = lines.get(i);
            if (line.isEmpty()) {
                i += 1;
                continue;
            }
            tickets.add(parseTicket(line));
        }

        return new ParsedInput(ranges, tickets);
    }

    private static List<Integer> parseTicket(String line) {
        return Arrays.stream(line.split(","))
            .map(Integer::parseInt)
            .collect(Collectors.toList());
    }

    private static Range[] parseLineToRanges(String line) {
        int sep = line.indexOf(": ");
        if (sep == -1) throw new IllegalArgumentException("Illegal line input");
        String[] parts = line.substring(sep + 2).split(" or ");
        return new Range[] { parseRange(parts[0]), parseRange(parts[1]) };
    }

    private static Range parseRange(String input) {
        String[] nums = input.split("-", 2);
        int start = Integer.parseInt(nums[0]);
        int end = Integer.parseInt(nums[1]);

        if (start <= end) return new Range(start, end);
        return new Range(end, start);
    }

    private static int solvePart1(ParsedInput input) {
        int result = 0;
        for (int i = 1; i < input.tickets().size(); ++i) {
            List<Integer> ticket = input.tickets().get(i);
            for (Integer num : ticket) {
                boolean found = false;
                for (Range range : input.ranges()) {
                    if (range.contains(num)) {
                        found = true;
                        break;
                    }
                }
                if (!found) result += num;
            }
        }
        return result;
    }

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day16.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            ParsedInput parsedInput = parse(filename);
            
            Common.time("Part1", () -> solvePart1(parsedInput));
        } catch (IOException e) {
        System.err.println(e.getMessage());
        System.exit(1);
        }
    }
}
