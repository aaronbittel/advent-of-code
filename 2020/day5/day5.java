package day5;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.Map;
import java.util.HashMap;
import java.util.stream.Collectors;
import java.io.IOException;

import common.Common;

class Day5 {

    private static final int ROW_LENGTH = 7;
    private static final int COL_LENGTH = 3;

    private static List<String> parse(String filename) throws IOException {
        return Files.readAllLines(Path.of(filename));
    }

    private static int calculateSeatID(String line) {
        String rowLine = line.substring(0, ROW_LENGTH);
        int row = calculateRow(rowLine);

        String colLine = line.substring(ROW_LENGTH);
        int col = calculateCol(colLine);

        return row * 8 + col;
    }

    private static int calculateRow(String rowLine) {
        int row_start = 0;
        int row_end = 127;
        for (int i = 0; i < ROW_LENGTH; i++) {
            int middle = row_start + (row_end - row_start) / 2;
            if (rowLine.charAt(i) == 'F') {
                row_end = middle;
            } else {
                row_start = middle + 1;
            }
        }
        return rowLine.charAt(ROW_LENGTH-1) == 'F' ? row_start : row_end;
    }

    private static int calculateCol(String colLine) {
        int col_start = 0;
        int col_end = 7;
        for (int i = 0; i < COL_LENGTH; i++) {
            int middle = col_start + (col_end - col_start) / 2;
            if (colLine.charAt(i) == 'L') {
                col_end = middle;
            } else {
                col_start = middle + 1;
            }
        }
        return colLine.charAt(COL_LENGTH-1) == 'L' ? col_start : col_end;
    }

    private static int solvePart2(List<String> lines) {
        List<Integer> seatIDs = lines.stream()
            .mapToInt(Day5::calculateSeatID)
            .sorted()
            .boxed()
            .collect(Collectors.toList());

        for (int id = seatIDs.getFirst(); id < seatIDs.getLast(); ++id) {
            if (!seatIDs.contains(id) && seatIDs.contains(id-1) && seatIDs.contains(id+1)) {
                return id;
            }
        }

        throw new IllegalStateException("No solution found");
    }

    public static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day5.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            List<String> lines = Day5.parse(filename);

            Common.time("Part1", () -> lines.stream()
                .mapToInt(Day5::calculateSeatID)
                .max()
                .orElseThrow());

            Common.time("Part2", () -> Day5.solvePart2(lines));

        } catch (IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
