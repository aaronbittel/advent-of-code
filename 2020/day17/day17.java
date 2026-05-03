package day17;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Arrays;
import java.util.HashSet;
import java.util.Iterator;
import java.util.List;
import java.util.Set;

import common.Common;

record Position(int x, int y, int z) { }

record Cubes(Set<Position> cubes) { 

    public Cubes nextCycle() {
        Position[] dimensions = getDimensions();
        Position minPos = dimensions[0];
        Position maxPos = dimensions[1];

        Set<Position> nextState = new HashSet<>();

        for (int z = minPos.z() - 1; z <= maxPos.z() + 1; ++z) {
            for (int y = minPos.y() - 1; y <= maxPos.y() + 1; ++y) {
                for (int x = minPos.x() - 1; x <= maxPos.x() + 1; ++x) {
                    Position pos = new Position(x, y, z);
                    int count = countNeighbours(pos);
                    if (isActive(pos) && (count == 2 || count == 3)) {
                        if (count == 2 || count == 3) {
                            nextState.add(pos);
                        }
                    } else {
                        if (count == 3) nextState.add(pos);
                    }
                }
            }
        }

        return new Cubes(nextState);
    }

    public int countActive() {
        return cubes.size();
    }

    private boolean isActive(Position pos) {
        return cubes.contains(pos);
    }

    private int countNeighbours(Position pos) {
        int count = 0;
        for (int dz = -1; dz <= 1; ++dz) {
            int z = pos.z() + dz;
            for (int dy = -1; dy <= 1; ++dy) {
                int y = pos.y() + dy;
                for (int dx = -1; dx <= 1; ++dx) {
                    if (dz == 0 && dy == 0 && dx == 0) continue;
                    int x = pos.x() + dx;
                    if (cubes.contains(new Position(x, y, z))) count++;
                }
            }
        }
        return count;
    }

    private Position[] getDimensions() {
        int minX = Integer.MAX_VALUE;
        int maxX = Integer.MIN_VALUE;
        int minY = Integer.MAX_VALUE;
        int maxY = Integer.MIN_VALUE;
        int minZ = Integer.MAX_VALUE;
        int maxZ = Integer.MIN_VALUE;
        for (Position pos : cubes) {
            if (pos.x() < minX) minX = pos.x();
            if (pos.x() > maxX) maxX = pos.x();
            if (pos.y() < minY) minY = pos.y();
            if (pos.y() > maxY) maxY = pos.y();
            if (pos.z() < minZ) minZ = pos.z();
            if (pos.z() > maxZ) maxZ = pos.z();
        }
        return new Position[]{ 
            new Position(minX, minY, minZ), new Position(maxX, maxY, maxZ)
        };
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();

        Position[] dimensions = getDimensions();
        Position minPos = dimensions[0];
        Position maxPos = dimensions[1];

        for (int z = minPos.z(); z <= maxPos.z(); ++z) {
            if (z != minPos.z()) sb.append("\n");
            sb.append("z=").append(z).append("\n");
            for (int y = minPos.y(); y <= maxPos.y(); ++y) {
                for (int x = minPos.x(); x <= maxPos.x(); ++x) {
                    Position pos = new Position(x, y, z);
                    sb.append(cubes.contains(pos) ? '#' : '.');
                }
                sb.append("\n");
            }
        }

        return sb.toString();
    }
}

class Day17 {

    private static Cubes parse(String filename) throws IOException {
        Set<Position> cubes = new HashSet<>();

        List<String> lines = Files.readAllLines(Path.of(filename));
        for (int y = 0; y < lines.size(); ++y) {
            String line = lines.get(y);
            for (int x = 0; x < line.length(); ++x) {
                if (line.charAt(x) == '#') {
                    cubes.add(new Position(x, y, 0));
                }
            }
        }

        return new Cubes(cubes);
    }

    private static int solvePart1(Cubes cubes) {
        Cubes current = cubes;
        for (int i = 0; i < 6; ++i) {
            current = current.nextCycle();
        }
        return current.countActive();
    }

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day17.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            Cubes cubes = parse(filename);

            Common.time("Part1", () -> solvePart1(cubes));
        } catch (IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
