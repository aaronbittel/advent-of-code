package day11;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.HashMap;
import java.util.Map;
import java.util.List;
import java.io.IOException;

import common.Common;

record Position(int y, int x) {}

record Seats(Map<Position, Boolean> seats, int rows, int cols) {
    private static final char OCCUPIED = '#';
    private static final char EMPTY = 'L';
    private static final char FREESPACE = '.';

    public Seats(Seats seats) {
        Map<Position, Boolean> seatsMap = new HashMap<>();
        for (Map.Entry<Position, Boolean> entry : seats.seats.entrySet()) {
            Position key = entry.getKey();
            seatsMap.put(
                new Position(key.y(), key.x()),
                Boolean.valueOf(entry.getValue())
            );
        }
        this(seatsMap, seats.rows, seats.cols);
    }

    public int countNeighbours(Position pos) {
        requireValidPosition(pos);
        int count = 0;
        for (int dy = -1; dy <= 1; ++dy) {
            for (int dx = -1; dx <= 1; ++dx) {
                if (dy == 0 && dx == 0) continue;
                Position newPos = new Position(pos.y() + dy, pos.x() + dx);
                if (seats.containsKey(newPos) && seats.get(newPos)) count++;
            }
        }
        return count;
    }

    public int countOccupied() {
        return (int)seats.values().stream().filter(s -> s).count();
    }

    public void occupy(Position pos) {
        requireValidPosition(pos);
        seats.put(pos, true);
    }

    public void empty(Position pos) {
        requireValidPosition(pos);
        seats.put(pos, false);
    }

    private void requireValidPosition(Position pos) {
        if (!seats.containsKey(pos)) {
            throw new IllegalArgumentException("Invalid position: " + pos);
        }
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder(rows * (cols + 1));

        for (int y = 0; y < rows; ++y) {
            for (int x = 0; x < cols; ++x) {
                Position pos = new Position(y, x);
                if (seats.containsKey(pos)) {
                    sb.append(seats.get(pos) ? OCCUPIED : EMPTY);
                } else {
                    sb.append(FREESPACE);
                }
            }
            sb.append('\n');
        }

        return sb.toString();
    }
}

class Day11 {

    private static Seats parse(String filename) throws IOException {
        Map<Position, Boolean> seats = new HashMap<>();
        List<String> lines = Files.readAllLines(Path.of(filename));

        for (int y = 0; y < lines.size(); ++y) {
            String line = lines.get(y);
            for (int x = 0; x < line.length(); ++x) {
                if (line.charAt(x) == 'L') {
                    seats.put(new Position(y, x), false);
                }
            }
        }

        return new Seats(seats, lines.size(), lines.getFirst().length());
    }

    private static int solvePart1(Seats seats) {
        Seats current = new Seats(seats);
        Seats next = applyRules(current);
        while (!current.equals(next)) {
            current = next;
            next = applyRules(current);
        }
        return next.countOccupied();
    }

    private static Seats applyRules(Seats current) {
        Seats next = new Seats(current);
        for (Map.Entry<Position, Boolean> seat : current.seats().entrySet()) {
            Position pos = seat.getKey();
            Boolean isOccupied = seat.getValue();
            int count = current.countNeighbours(pos);
            if (!isOccupied && count == 0) next.occupy(pos);
            else if (isOccupied && count >= 4) next.empty(pos);
        }
        return next;
    }

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day11.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            Seats seats = parse(filename);

            Common.time("Part1", () -> solvePart1(seats));
        } catch(IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
