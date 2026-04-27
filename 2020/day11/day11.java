package day11;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.List;
import java.util.function.Function;
import java.io.IOException;

import common.Common;

enum SeatState {
    FLOOR('.'),
    OCCUPIED('#'),
    EMPTY('L');

    private final char symbol;

    SeatState(char symbol) {
        this.symbol = symbol;
    }

    public char getSymbol() {
        return symbol;
    }
}

record Seats(SeatState[][] seats) {
    public Seats(Seats original) {
        SeatState[][] copy = new SeatState[original.seats.length][];

        for (int i = 0; i < original.seats.length; i++) {
            copy[i] = original.seats[i].clone();
        }

        this(copy);
    }

    public int countCloseNeighbours(int y, int x) {
        int count = 0;
        for (int dy = -1; dy <= 1; ++dy) {
            int newY = y + dy;
            if (newY < 0 || newY >= getHeight()) continue;
            for (int dx = -1; dx <= 1; ++dx) {
                if (dy == 0 && dx == 0) continue;
                int newX = x + dx;
                if (newX < 0 || newX >= getWidth()) continue;
                if (isOccupied(newY, newX)) count++;
            }
        }
        return count;
    }

    public boolean isOccupied(int y, int x) {
        return seats[y][x] == SeatState.OCCUPIED;
    }

    public boolean isEmpty(int y, int x) {
        return seats[y][x] == SeatState.EMPTY;
    }

    public boolean isFloor(int y, int x) {
        return seats[y][x] == SeatState.FLOOR;
    }

    public int getHeight() {
        return seats.length;
    }

    public int getWidth() {
        if (seats.length > 0) {
            return seats[0].length;
        }
        return 0;
    }

    public int countNeighbours(int y, int x) {
        int count = 0;
        for (int dy = -1; dy <= 1; ++dy) {
            int newY = y + dy;
            if (newY < 0 || newY >= getHeight()) continue;
            for (int dx = -1; dx <= 1; ++dx) {
                if (dy == 0 && dx == 0) continue;
                int newX = x + dx;
                if (newX < 0 || newX >= getWidth()) continue;
                newY = y + dy; // reset because newY is mutated during directional scan
                while(isValidPosition(newY, newX) && isFloor(newY, newX)) {
                    newY += dy;
                    newX += dx;
                }
                if (isValidPosition(newY, newX) && isOccupied(newY, newX)) count++;
            }
        }
        return count;
    }

    public int countOccupied() {
        int count = 0;
        for (int y = 0; y < getHeight(); ++y) {
            for (int x = 0; x < getWidth(); ++x) {
                if (isOccupied(y, x)) count++;
            }
        }
        return count;
    }

    public void occupy(int y, int x) {
        seats[y][x] = SeatState.OCCUPIED;
    }

    public void empty(int y, int x) {
        seats[y][x] = SeatState.EMPTY;
    }

    public boolean isValidPosition(int y, int x) {
        return y >= 0 && y < getHeight() && x >= 0 && x < getWidth();
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (!(o instanceof Seats)) return false;
        Seats other = (Seats)o;
        return Arrays.deepEquals(this.seats, other.seats);
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder(getHeight() * (getWidth() + 1));

        for (int y = 0; y < getHeight(); ++y) {
            for (int x = 0; x < getWidth(); ++x) {
                sb.append(seats[y][x].getSymbol());
            }
            sb.append('\n');
        }

        return sb.toString();
    }
}

class Day11 {

    private static Seats parse(String filename) throws IOException {
        List<String> lines = Files.readAllLines(Path.of(filename));
        SeatState[][] seats = new SeatState[lines.size()][];

        for (int y = 0; y < lines.size(); ++y) {
            String line = lines.get(y);
            SeatState[] row = new SeatState[line.length()];
            for (int x = 0; x < line.length(); ++x) {
                row[x] = line.charAt(x) == 'L' ? SeatState.EMPTY : SeatState.FLOOR;
            }
            seats[y] = row;
        }

        return new Seats(seats);
    }

    private static int solve(Seats seats, Function<Seats, Seats> applyRules) {
        Seats current = new Seats(seats);
        Seats next = applyRules.apply(current);

        while (!current.equals(next)) {
            current = next;
            next = applyRules.apply(current);
        }

        return next.countOccupied();
    }

    private static Seats applyRulesPart1(Seats current) {
        Seats next = new Seats(current);
        for (int y = 0; y < current.getHeight(); ++y) {
            for (int x = 0; x < current.getWidth(); ++x) {
                if (current.isFloor(y, x)) continue;
                boolean isOccupied = current.isOccupied(y, x);
                int count = current.countCloseNeighbours(y, x);
                if (!isOccupied && count == 0) next.occupy(y, x);
                else if (isOccupied && count >= 4) next.empty(y, x);
            }
        }
        return next;
    }

    private static Seats applyRulesPart2(Seats current) {
        Seats next = new Seats(current);
        for (int y = 0; y < current.getHeight(); ++y) {
            for (int x = 0; x < current.getWidth(); ++x) {
                if (current.isFloor(y, x)) continue;
                boolean isOccupied = current.isOccupied(y, x);
                int count = current.countNeighbours(y, x);
                if (!isOccupied && count == 0) next.occupy(y, x);
                else if (isOccupied && count >= 5) next.empty(y, x);
            }
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

            Common.time("Part1", () -> solve(seats, Day11::applyRulesPart1));
            Common.time("Part2", () -> solve(seats, Day11::applyRulesPart2));
        } catch(IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
