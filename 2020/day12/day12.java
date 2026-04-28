package day12;

import java.nio.file.Files;
import java.nio.file.Path;
import java.io.IOException;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;

import common.Common;

enum Action {
    NORTH,
    SOUTH,
    EAST,
    WEST,
    LEFT,
    RIGHT,
    FORWARD;

    public Action calculateFacing(int value) {
        return switch(this) {
            case NORTH -> valToDirection(Math.floorMod(  0 + value, 360));
            case EAST  -> valToDirection(Math.floorMod( 90 + value, 360));
            case SOUTH -> valToDirection(Math.floorMod(180 + value, 360));
            case WEST  -> valToDirection(Math.floorMod(270 + value, 360));
            default -> throw new IllegalArgumentException(
                "Can only be called when Action is one of "
                + "'NORTH', 'SOUTH', 'EAST' or 'WEST'");
        };
    }

    private static Action valToDirection(int value) {
        return switch(value) {
            case 0 -> NORTH;
            case 90 -> EAST;
            case 180 -> SOUTH;
            case 270 -> WEST;
            default -> throw new IllegalArgumentException(
                "Can only be called when Action is one of "
                + "'NORTH', 'SOUTH', 'EAST' or 'WEST'");
        };
    }
}

record Instruction(Action action, int value) {}

class Day12 {

    private static List<Instruction> parse(String filename) throws IOException {
        try(Stream<String> lines = Files.lines(Path.of(filename))) {
            return lines.map(Day12::lineToInstruction)
                        .collect(Collectors.toList());
        }
    }

    private static Instruction lineToInstruction(String line) {
        int value = Integer.parseInt(line.substring(1));
        Action action = switch (line.charAt(0)) {
            case 'N' -> Action.NORTH;
            case 'S' -> Action.SOUTH;
            case 'E' -> Action.EAST;
            case 'W' -> Action.WEST;
            case 'L' -> Action.LEFT;
            case 'R' -> Action.RIGHT;
            case 'F' -> Action.FORWARD;
            default -> throw new IllegalArgumentException("Illegal action");
        };
        return new Instruction(action, value);
    }

    private static int solvePart1(List<Instruction> instructions) {
        int north = 0;
        int east = 0;
        Action facing = Action.EAST;

        for (Instruction inst : instructions) {
            switch(inst.action()) {
                case NORTH: north += inst.value(); break;
                case SOUTH: north -= inst.value(); break;
                case EAST: east += inst.value(); break;
                case WEST: east -= inst.value(); break;
                case LEFT: facing = facing.calculateFacing(-inst.value()); break;
                case RIGHT: facing = facing.calculateFacing(inst.value()); break;
                case FORWARD: {
                    switch(facing) {
                        case NORTH: north += inst.value(); break;
                        case SOUTH: north -= inst.value(); break;
                        case EAST: east += inst.value(); break;
                        case WEST: east -= inst.value(); break;
                        default: throw new IllegalArgumentException(
                            "Can only be called when Action is one of "
                            + "'NORTH', 'SOUTH', 'EAST' or 'WEST'");
                    };
                }
            };
        }

        return Math.abs(north) + Math.abs(east);
    }

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day12.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            List<Instruction> instructions = parse(filename);

            Common.time("Part1", () -> solvePart1(instructions));
        } catch(IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
