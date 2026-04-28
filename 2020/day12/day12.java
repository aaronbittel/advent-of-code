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
        state.move(dir, value);
    }
}

record Turn(int degrees) implements Actionable {
    public void apply(State state) {
        state.turn(degrees);
    }
}

record Forward(int value) implements Actionable {
    public void apply(State state) {
        state.forward(value);
    }
}

interface State {
    void move(Direction dir, int value);
    void forward(int value);
    void turn(int degrees);
    int manhattenDistance();
}

class ShipState implements State {
    int north = 0;
    int east = 0;
    Direction facing = Direction.EAST;

    public void move(Direction dir, int value) {
        switch(dir) {
            case NORTH -> north += value;
            case SOUTH -> north -= value;
            case EAST  -> east += value;
            case WEST  -> east -= value;
        }
    }

    public void forward(int value) {
        move(facing, value);
    }

    public void turn(int degrees) {
        facing = facing.turn(degrees);
    }

    public int manhattenDistance() {
        return Math.abs(north) + Math.abs(east);
    }

    @Override
    public String toString() {
        return String.format("North: %d, East: %d, Facing: %s", north, east, facing);
    }
}

class WaypointState implements State {
    int waypointNorth = 1;
    int waypointEast = 10;

    int north = 0;
    int east = 0;

    public void move(Direction dir, int value) {
        switch(dir) {
            case NORTH -> waypointNorth += value;
            case SOUTH -> waypointNorth -= value;
            case EAST  -> waypointEast += value;
            case WEST  -> waypointEast -= value;
        }
    }

    public void forward(int value) {
        north += waypointNorth * value;
        east  += waypointEast  * value;
    }

    public void turn(int degrees) {
        int times = Math.floorMod(degrees / 90, 4);
        for (int i = 0; i < times; i++) {
            int tmp = waypointNorth;
            waypointNorth = -waypointEast;
            waypointEast = tmp;
        }
    }

    public int manhattenDistance() {
        return Math.abs(north) + Math.abs(east);
    }

    @Override
    public String toString() {
        return String.format(
            "Ship: North: %d, East: %d; Waypont: North %d, East: %d",
                north, east, waypointNorth, waypointEast);
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
        State ship = new ShipState();
        for (Actionable action : actions) {
            action.apply(ship);
        }

        return ship.manhattenDistance();
    }

    private static int solvePart2(List<Actionable> actions) {
        State waypoint = new WaypointState();

        for (Actionable action : actions) {
            action.apply(waypoint);
        }

        return waypoint.manhattenDistance();
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
            Common.time("Part2", () -> solvePart2(actions));
        } catch(IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
