package day14;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.List;
import java.util.HashMap;
import java.util.Map;
import java.io.IOException;
import java.util.stream.Collectors;

import common.Common;

record MaskPart1(long maskOnes, long maskZeros) {

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

    private static MaskPart1 parseMaskPart1(String input) {
        if (input.length() != 36) {
            throw new IllegalArgumentException(
                String.format(
                    "Expected mask to be of length '36', but got '%d'",
                    input.length()));
        }
        long maskOnes = 0;
        long maskZeros = ~0L; // all bits 1

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

        return new MaskPart1(maskOnes, maskZeros);
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
        MaskPart1 mask = new MaskPart1(0, 0L);
        Map<Integer, Long> memory = new HashMap<>();

        for (String line : lines) {
            if (line.startsWith("mask = ")) {
                mask = parseMaskPart1(line.substring(7));
            } else {
                MemWrite memWrite = parseMemWrite(line);
                memory.put(memWrite.mem(), mask.apply(memWrite.value()));
            }
        }

        long result = 0;
        for (Long value : memory.values()) {
            result += value;
        }
        return result;
    }

    @SuppressWarnings("unused")
    private static String toBinaryString(Long value) {
        String valueStr = Long.toBinaryString(value);
        return String.format(
            "%s%s",
            "0".repeat(36 - valueStr.length()),
            valueStr
        );
    }

    private static void calculateValueVariants(
        long value,
        String valueStr,
        int index,
        List<Long> variants
    ) {
        char c;
        while (index < valueStr.length() && (c = valueStr.charAt(index)) != 'X') {
            value <<= 1;
            if (c == '1') value |= 1;
            index++;
        }

        if (index >= valueStr.length()) {
            variants.add(value);
            return;
        }

        if (valueStr.charAt(index) == 'X') {
            calculateValueVariants(value << 1, valueStr, index + 1, variants);
            calculateValueVariants((value << 1) | 1, valueStr, index + 1, variants);
        }
    }

    private static String applyMaskPart2(long value, String mask) {
        if (mask.length() != 36) {
            throw new IllegalArgumentException("mask must be of length 36");
        }
        String longStr = Long.toBinaryString(value);
        String valueStr = String.format(
            "%s%s", "0".repeat(36-longStr.length()), longStr);

        StringBuilder sb = new StringBuilder(36);

        for (int i = 0; i < 36; ++i) {
            char c = mask.charAt(i) ;
            switch(c) {
                case '0': {
                    sb.append(valueStr.charAt(i));
                    break;
                }
                case '1': {
                    sb.append('1');
                    break;
                }
                case 'X': {
                    sb.append('X');
                    break;
                }
                default:
                    throw new IllegalArgumentException(
                        String.format("Illegal char '%c' in mask", c));
            }
        }

        return sb.toString();
    }

    private static List<Long> calculateVariants(String value) {
        List<Long> variants = new ArrayList<>();
        calculateValueVariants(0L, value, 0, variants);
        return variants;
    }

    private static long solvePart2(List<String> lines) {
        String mask = "";
        Map<Long, Long> memory = new HashMap<>();

        for (String line : lines) {
            if (line.startsWith("mask = ")) {
                mask = line.substring(7);
            } else {
                MemWrite memWrite = parseMemWrite(line);
                String value = applyMaskPart2(memWrite.mem(), mask);

                for (Long variant : calculateVariants(value)) {
                    memory.put(variant, memWrite.value());
                }
            }
        }

        long result = 0;
        for (Long value : memory.values()) {
            result += value;
        }
        return result;
    }

    @SuppressWarnings("unused")
    public static List<Long> calculateVariantsStrings(String value) {
        List<String> current = new ArrayList<>(List.of(value));
        while (current.getFirst().contains("X")) {
            List<String> next = new ArrayList<>(current.size() * 2);
            for (String cur : current) {
                next.add(cur.replaceFirst("X", "0"));
                next.add(cur.replaceFirst("X", "1"));
            }
            current = next;
        }
        return current.stream()
            .map(v -> Long.valueOf(v, 2))
            .collect(Collectors.toList());
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
            Common.time("Part2", () -> solvePart2(lines));
        } catch(IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
