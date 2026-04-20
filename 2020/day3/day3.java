package day3;

import java.util.List;
import java.util.Objects;
import java.util.Set;
import java.util.HashSet;
import java.util.stream.Stream;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.function.Supplier;

import common.Common;

class Point {
    private int row;
    private int col;

    public Point() {
        this(0, 0);
    }

    public Point(int row, int col) {
        this.row = row;
        this.col = col;
    }

    public void add(Point dir) {
        row += dir.getRow();
        col += dir.getCol();
    }

    public void setRow(int row) {
        this.row = row;
    }

    public void setCol(int col) {
        this.col = col;
    }

    public int getRow() {
        return row;
    }

    public int getCol() {
        return col;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Point other = (Point) o;
        return row == other.row && col == other.col;
    }

    @Override
    public int hashCode() {
        return Objects.hash(row, col);
    }

    @Override
    public String toString() {
        return String.format("(%d, %d)", row, col);
    }
}

class Grid {
    private Set<Point> trees;
    private final int rows;
    private final int cols;

    public Grid(Set<Point> trees, int rows, int cols) {
        this.trees = trees;
        this.rows = rows;
        this.cols = cols;
    }

    public boolean isTree(Point p) {
        return trees.contains(p);
    }

    public int getRows() {
        return rows;
    }

    public int getCols() {
        return cols;
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder(rows * (cols + 1));
        for (int row = 0; row < rows; ++row) {
            for (int col = 0; col < cols; ++col) {
                sb.append(isTree(new Point(row, col)) ? '#' : '.');
            }
            sb.append('\n');
        }
        return sb.toString();
    }
}

class Day3 {

    public static int countTreesWithSlope(Grid grid, Point slope) {
        Point pos = new Point();
        int trees = 0;
        while (pos.getRow() < grid.getRows()) {
            pos.add(slope);
            pos.setCol(pos.getCol() % grid.getCols());
            if (grid.isTree(pos)) trees += 1;
        }
        return trees;
    }

    private static Grid parse(String filename) throws IOException {
        Set<Point> trees = new HashSet<>();
        List<String> lines =  Files.readAllLines(Path.of(filename));
        int rows = lines.size();
        int cols = lines.getFirst().length();
        for (int row = 0; row < rows; ++row) {
            String line = lines.get(row);
            for (int col = 0; col < cols; ++col) {
                if (line.charAt(col) == '#') {
                    trees.add(new Point(row, col));
                }
            }
        }
        return new Grid(trees, rows, cols);
    }

    public static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day3.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            Grid grid = Day3.parse(filename);

            Common.time("Part1", () -> Day3.countTreesWithSlope(grid, new Point(1, 3)));

            List<Point> slopes = List.of(
                new Point(1, 1),
                new Point(1, 3),
                new Point(1, 5),
                new Point(1, 7),
                new Point(2, 1)
            );

            Common.time("Part2", () -> slopes.stream()
                .map(slope -> (long) Day3.countTreesWithSlope(grid, slope))
                .reduce(1L, (acc, val) -> acc * val));
        } catch (IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
