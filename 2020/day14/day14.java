package day14;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.HashMap;
import java.util.Map;
import java.io.IOException;

import common.Common;

record Mask(long maskOnes, long maskZeros) {

    public long apply(long value) {
        return (value | maskOnes) & maskZeros;
    }

    @Override
    public String toString() {
        return String.format(
            "%64s%n%64s",
            Long.toBinaryString(maskOnes),
            Long.toBinaryString(maskZeros));
    }
}

record MemWrite(int mem, long value) { }

class Day14 {

    private static List<String> parse(String filename) throws IOException {
        return Files.readAllLines(Path.of(filename));
    }

    private static Mask parseMask(String input) {
        if (input.length() != 36) {
            throw new IllegalArgumentException(
                String.format(
                    "Expected mask to be of length '36', but got '%d'",
                    input.length()));
        }
        long maskOnes = 0;
        long maskZeros = Long.MAX_VALUE;

        for (char c : input.toCharArray()) {
            maskOnes <<= 1;
            maskZeros <<= 1;
            if (c == '0') {
                // do nothing
            } else if (c == '1') {
                maskOnes |= 1;
                maskZeros |= 1;
            } else if (c == 'X') {
                maskZeros |= 1;
            } else {
                throw new IllegalArgumentException("Invalid input for mask: " + input);
            }
        }

        return new Mask(maskOnes, maskZeros);
    }

    private static MemWrite parseMemWrite(String input) {
        int memStart = input.indexOf("[");
        int memEnd = input.indexOf("]");
        if (memStart == -1 || memEnd == -1) {
            throw new IllegalArgumentException("Invalid mem line format: " + input);
        }
        int mem = Integer.parseInt(input.substring(memStart + 1, memEnd));

        int valueStart = input.lastIndexOf(" ");
        if (valueStart == -1) {
            throw new IllegalArgumentException("Invalid mem line format: " + input);
        }
        long value = Long.parseLong(input.substring(valueStart + 1));
        return new MemWrite(mem, value);
    }

    private static long solvePart1(List<String> lines) {
        Mask mask = new Mask(0, 0);
        Map<Integer, Long> memory = new HashMap<>();

        for (String line : lines) {
            if (line.startsWith("mask = ")) {
                mask = parseMask(line.substring(7));
            } else {
                MemWrite memLine = parseMemWrite(line);
                memory.put(memLine.mem(), mask.apply(memLine.value()));
            }
        }

        long result = 0;
        for (Long value : memory.values()) {
            result += value;
        }
        return result;
    }

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day14.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            List<String> lines = parse(filename);

            Common.time("Part1", () -> solvePart1(lines));
        } catch(IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
