package day12;

import java.nio.file.Files;
import java.nio.file.Path;
import java.io.IOException;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;

import common.Common;

interface Actionable {
    void apply(State state);
}

enum Direction {
    NORTH,
    EAST,
    SOUTH,
    WEST;

    public Direction turn(int degrees) {
        int steps = Math.floorMod(degrees / 90, 4);
        int idx = (ordinal() + steps) % 4;
        return values()[idx];
    }
}

record Move(Direction dir, int value) implements Actionable {
    public void apply(State state) {
        switch(dir) {
            case NORTH -> state.north += value;
            case SOUTH -> state.north -= value;
            case EAST  -> state.east += value;
            case WEST  -> state.east -= value;
        };
    }
}

record Turn(int degrees) implements Actionable {
    public void apply(State state) {
        state.facing = state.facing.turn(degrees);
    }
}

record Forward(int value) implements Actionable {
    public void apply(State state) {
        new Move(state.facing, value).apply(state);
    }
}

class State {
    int north = 0;
    int east = 0;
    Direction facing = Direction.EAST;

    @Override
    public String toString() {
        return String.format("North: %d, East: %d, Facing: %s", north, east, facing);
    }
}

class Day12 {

    private static List<Actionable> parse(String filename) throws IOException {
        try(Stream<String> lines = Files.lines(Path.of(filename))) {
            return lines.map(Day12::lineToActionable)
                        .collect(Collectors.toList());
        }
    }

    private static Actionable lineToActionable(String line) {
        int value = Integer.parseInt(line.substring(1));
        return switch (line.charAt(0)) {
            case 'N' -> new Move(Direction.NORTH, value);
            case 'S' -> new Move(Direction.SOUTH, value);
            case 'E' -> new Move(Direction.EAST, value);
            case 'W' -> new Move(Direction.WEST, value);
            case 'L' -> new Turn(-value);
            case 'R' -> new Turn(value);
            case 'F' -> new Forward(value);
            default -> throw new IllegalArgumentException("Illegal action");
        };
    }

    private static int solvePart1(List<Actionable> actions) {
        State state = new State();
        for (Actionable action : actions) {
            action.apply(state);
        }

        return Math.abs(state.north) + Math.abs(state.east);
    }

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day12.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            List<Actionable> actions = parse(filename);

            Common.time("Part1", () -> solvePart1(actions));
        } catch(IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
