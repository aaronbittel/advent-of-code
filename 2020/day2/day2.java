import java.util.List;
import java.util.stream.Stream;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.function.Supplier;
import java.util.regex.Pattern;
import java.util.regex.Matcher;

class Range {
    private int start;
    private int end;

    public Range(int start, int end) {
        this.start = start;
        this.end = end;
    }

    public int getStart() {
        return start;
    }

    public int getEnd() {
        return end;
    }
}

class PasswordLine {
    private Range range;
    private char c;
    private String password;

    public PasswordLine(Range range, char c, String password) {
        this.range = range;
        this.c = c;
        this.password = password;
    }

    public boolean isValidPart1() {
        int count = 0;
        for (char c : password.toCharArray()) {
            if (c == this.c) {
                count += 1;
            }
        }
        return count >= range.getStart() && count <= range.getEnd();
    }

    public boolean isValidPart2() {
        if (range.getEnd() > password.length()) {
            throw new IllegalArgumentException("Range End greater then password length");
        }
        char left = password.charAt(range.getStart() - 1);
        char right = password.charAt(range.getEnd() - 1);
        return (left == c && right != c) || (left != c && right == c);
    }

    @Override
    public String toString() {
        return String.format("%d-%d %c: %s", range.getStart(), range.getEnd(), c, password);
    }
}

class Day2 {

    private static final Pattern pattern = Pattern.compile("(\\d+)-(\\d+) ([a-zA-Z]): (\\w+)");

    private static List<PasswordLine> parse(String filename) throws IOException {
        return Files.lines(Path.of(filename)).map(Day2::parsePasswordLine).toList();
    }

    private static PasswordLine parsePasswordLine(String line) {
        Matcher m = Day2.pattern.matcher(line);
        if (m.matches()) {
            return new PasswordLine(
                new Range(Integer.parseInt(m.group(1)), Integer.parseInt(m.group(2))),
                m.group(3).charAt(0),
                m.group(4)
            );
        }
        throw new IllegalArgumentException("Illegal Input");
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
            System.err.printf("Usage: java %s <input>%n", Day2.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            List<PasswordLine> passwordLines = Day2.parse(filename);
            Day2.time("Part1", () -> passwordLines.stream()
                .filter(PasswordLine::isValidPart1)
                .count());
            Day2.time("Part2", () -> passwordLines.stream()
                .filter(PasswordLine::isValidPart2)
                .count());
        } catch(IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
