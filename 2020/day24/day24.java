package day24;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

import common.Common;

record Position(int x, int y) { }

enum HexDirection {
    EAST(1, 0),
    SOUTHEAST(0, 1),
    SOUTHWEST(-1, 1),
    WEST(-1, 0),
    NORTHWEST(0, -1),
    NORTHEAST(1, -1);

    private final int x;
    private final int y;

    HexDirection(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public int x() {
        return x;
    }

    public int y() {
        return y;
    }
}

class Day24 {

    private static List<List<HexDirection>> parse(String filename) throws IOException {
        return Files.lines(Path.of(filename))
            .map(Day24::lineToHexDirections)
            .toList();
    }

    private static List<HexDirection> lineToHexDirections(String line) {
        int index = 0;
        List<HexDirection> dirs = new ArrayList<>();
        while (index < line.length()) {
            char c = line.charAt(index);
            HexDirection dir = switch (c) {
                case 'e' -> {
                    index++;
                    yield HexDirection.EAST;
                }
                case 'w' -> {
                    index++;
                    yield HexDirection.WEST;
                }
                case 's' -> {
                    char next = line.charAt(++index);
                    index++;
                    if (next == 'e') yield HexDirection.SOUTHEAST;
                    if (next == 'w') yield HexDirection.SOUTHWEST;
                    throw new IllegalArgumentException("Unknown direction: s" + next);
                }
                case 'n' -> {
                    char next = line.charAt(++index);
                    index++;
                    if (next == 'e') yield HexDirection.NORTHEAST;
                    if (next == 'w') yield HexDirection.NORTHWEST;
                    throw new IllegalArgumentException("Unknown direction: n" + next);
                }
                default -> throw new IllegalArgumentException("Unknown direction: " + c);
            };
            dirs.add(dir);
        }
        return dirs;
    }

    private static int solvePart1(List<List<HexDirection>> tilePaths) {
        Set<Position> flippedTiles = new HashSet<>();
        for (List<HexDirection> path : tilePaths) {
            int x = 0;
            int y = 0;
            for (HexDirection dir : path) {
                x += dir.x();
                y += dir.y();
            }
            Position pos = new Position(x, y);
            if (!flippedTiles.add(pos)) {
                flippedTiles.remove(pos);
            }
        }
        return flippedTiles.size();
    }

    private static int solvePart2(List<List<HexDirection>> tilePaths) {
        Set<Position> flippedTiles = new HashSet<>();
        for (List<HexDirection> path : tilePaths) {
            int x = 0;
            int y = 0;
            for (HexDirection dir : path) {
                x += dir.x();
                y += dir.y();
            }
            Position pos = new Position(x, y);
            if (!flippedTiles.add(pos)) {
                flippedTiles.remove(pos);
            }
        }

        for (int i = 0; i < 100; ++i) {
            Set<Position> nextDay = new HashSet<>();

            for (Position pos : flippedTiles) {
                int count = 0;
                for (HexDirection dir : HexDirection.values()) {
                    int nx = pos.x() + dir.x();
                    int ny = pos.y() + dir.y();
                    if (flippedTiles.contains(new Position(nx, ny))) count++;
                }

                if (count == 1 || count == 2) nextDay.add(pos);

                for (HexDirection dir : HexDirection.values()) {
                    int nx = pos.x() + dir.x();
                    int ny = pos.y() + dir.y();
                    Position neighborPosition = new Position(nx, ny);
                    if (flippedTiles.contains(neighborPosition)) continue;
                    int countFlipped = 0;
                    for (HexDirection d : HexDirection.values()) {
                        if (flippedTiles.contains(new Position(nx + d.x(), ny + d.y()))) countFlipped++;
                    }
                    if (countFlipped == 2) nextDay.add(neighborPosition);
                }
            }

            flippedTiles = nextDay;
        }

        return flippedTiles.size();
    }

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day24.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            List<List<HexDirection>> tilePaths = parse(filename);

            Common.time("Part1", () -> solvePart1(tilePaths));
            Common.time("Part2", () -> solvePart2(tilePaths));
        } catch (IOException e) {
            System.err.printf("ERROR: reading file '%s': %s%n", filename, e.getMessage());
            System.exit(1);
        }
    }
}
