import java.io.BufferedReader;
import java.io.FileReader;
import java.util.Map;
import java.util.HashMap;
import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;
import java.io.IOException;
import java.util.function.Supplier;

class Day4 {

    private static final List<String> fields = List.of(
        "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid");

    private static List<Map<String, String>> parse(String filename) throws IOException {
        BufferedReader reader = new BufferedReader(new FileReader(filename));

        List<Map<String, String>> batches = new ArrayList<>();
        Map<String, String> batch = new HashMap<>();

        String line;
        while ((line = reader.readLine()) != null) {
            if (line.trim().isEmpty()) {
                batches.add(batch);
                batch = new HashMap<>();
            } else {
                for (String l : line.split(" ")) {
                    String[] parts = l.strip().split(":", 2);
                    if (parts.length == 2) {
                        batch.put(parts[0], parts[1]);
                    }
                }
            }
        }

        if (!batch.isEmpty()) {
            batches.add(batch);
        }

        reader.close();
        return batches;
    }

    public static boolean isValidPassportPart1(Map<String, String> passport) {
        for (String field : fields) {
            if (field.equals("cid")) continue;
            if (!passport.containsKey(field)) return false;
        }
        return true;
    }

    public static boolean isValidPassportPart2(Map<String, String> passport) {
        for (String field : fields) {
            String value = passport.get(field);
            if (field.equals("cid")) continue;
            if (value == null) return false;
            boolean valid = switch (field) {
                case "byr" -> isValidYear(value, 1920, 2002);
                case "iyr" -> isValidYear(value, 2010, 2020);
                case "eyr" -> isValidYear(value, 2020, 2030);
                case "hgt" -> isValidHeight(value);
                case "hcl" -> value.matches("#[0-9a-f]{6}");
                case "ecl" -> isValidEyeColor(value);
                case "pid" -> value.matches("[0-9]{9}");
                default -> throw new IllegalArgumentException("Illegal field: " + field);
            };
            if (!valid) return false;
        }
        return true;
    }

    private static boolean isValidYear(String input, int from, int to) {
        if (input.length() != 4) return false;
        int year = Integer.valueOf(input);
        return year >= from && year <= to;
    }

    private static boolean isValidHeight(String input) {
        if (input.matches("\\d+cm")) {
            int height = Integer.valueOf(input.substring(0, input.length() - 2));
            return height >= 150 && height <= 193;
        } else if (input.matches("\\d+in")) {
            int height = Integer.valueOf(input.substring(0, input.length() - 2));
            return height >= 59 && height <= 76;
        }
        return false;
    }

    private static boolean isValidEyeColor(String input) {
        List<String> colors = List.of(
            "amb", "blu", "brn", "gry", "grn", "hzl", "oth");
        for (String color : colors) {
            if (input.equals(color)) return true;
        }
        return false;
    }

    private static <T> void time(String label, Supplier<T> task) {
        long start = System.nanoTime();
        T result = task.get();
        long end = System.nanoTime();
        double seconds = (end - start) / 1_000_000_000.0;
        System.out.printf("%s: %s, took %.5f seconds%n", label, result, seconds);
    }

    public static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day4.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            List<Map<String, String>> batches = Day4.parse(filename);

            time("Part1", () -> batches.stream()
                .filter(Day4::isValidPassportPart1)
                .count()
            );

            time("Part2", () -> batches.stream()
                .filter(Day4::isValidPassportPart2)
                .count()
            );

        } catch (IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
