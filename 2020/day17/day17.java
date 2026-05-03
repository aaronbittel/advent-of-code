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

record Position3D(int x, int y, int z) { }

record Position4D(int x, int y, int z, int w) { }

record Cubes3D(Set<Position3D> cubes) {

    public Cubes3D nextCycle() {
        Position3D[] dimensions = getDimensions();
        Position3D minPos = dimensions[0];
        Position3D maxPos = dimensions[1];

        Set<Position3D> nextState = new HashSet<>();

        for (int z = minPos.z() - 1; z <= maxPos.z() + 1; ++z) {
            for (int y = minPos.y() - 1; y <= maxPos.y() + 1; ++y) {
                for (int x = minPos.x() - 1; x <= maxPos.x() + 1; ++x) {
                    Position3D pos = new Position3D(x, y, z);
                    int count = countNeighbours(pos);
                    if (isActive(pos)) {
                        if (count == 2 || count == 3) {
                            nextState.add(pos);
                        }
                    } else {
                        if (count == 3) nextState.add(pos);
                    }
                }
            }
        }

        return new Cubes3D(nextState);
    }

    public int countActive() {
        return cubes.size();
    }

    private boolean isActive(Position3D pos) {
        return cubes.contains(pos);
    }

    private int countNeighbours(Position3D pos) {
        int count = 0;
        for (int dz = -1; dz <= 1; ++dz) {
            int z = pos.z() + dz;
            for (int dy = -1; dy <= 1; ++dy) {
                int y = pos.y() + dy;
                for (int dx = -1; dx <= 1; ++dx) {
                    if (dz == 0 && dy == 0 && dx == 0) continue;
                    int x = pos.x() + dx;
                    if (cubes.contains(new Position3D(x, y, z))) count++;
                }
            }
        }
        return count;
    }

    private Position3D[] getDimensions() {
        int minX = Integer.MAX_VALUE;
        int maxX = Integer.MIN_VALUE;
        int minY = Integer.MAX_VALUE;
        int maxY = Integer.MIN_VALUE;
        int minZ = Integer.MAX_VALUE;
        int maxZ = Integer.MIN_VALUE;
        for (Position3D pos : cubes) {
            if (pos.x() < minX) minX = pos.x();
            if (pos.x() > maxX) maxX = pos.x();
            if (pos.y() < minY) minY = pos.y();
            if (pos.y() > maxY) maxY = pos.y();
            if (pos.z() < minZ) minZ = pos.z();
            if (pos.z() > maxZ) maxZ = pos.z();
        }
        return new Position3D[]{
            new Position3D(minX, minY, minZ), new Position3D(maxX, maxY, maxZ)
        };
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();

        Position3D[] dimensions = getDimensions();
        Position3D minPos = dimensions[0];
        Position3D maxPos = dimensions[1];

        for (int z = minPos.z(); z <= maxPos.z(); ++z) {
            if (z != minPos.z()) sb.append("\n");
            sb.append("z=").append(z).append("\n");
            for (int y = minPos.y(); y <= maxPos.y(); ++y) {
                for (int x = minPos.x(); x <= maxPos.x(); ++x) {
                    Position3D pos = new Position3D(x, y, z);
                    sb.append(cubes.contains(pos) ? '#' : '.');
                }
                sb.append("\n");
            }
        }

        return sb.toString();
    }
}

record Cubes4D(Set<Position4D> cubes) {

    public Cubes4D nextCycle() {
        Position4D[] dimensions = getDimensions();
        Position4D minPos = dimensions[0];
        Position4D maxPos = dimensions[1];

        Set<Position4D> nextState = new HashSet<>();

        for (int w = minPos.w() - 1; w <= maxPos.w() + 1; ++w) {
            for (int z = minPos.z() - 1; z <= maxPos.z() + 1; ++z) {
                for (int y = minPos.y() - 1; y <= maxPos.y() + 1; ++y) {
                    for (int x = minPos.x() - 1; x <= maxPos.x() + 1; ++x) {
                        Position4D pos = new Position4D(x, y, z, w);
                        int count = countNeighbours(pos);
                        if (isActive(pos)) {
                            if (count == 2 || count == 3) {
                                nextState.add(pos);
                            }
                        } else {
                            if (count == 3) nextState.add(pos);
                        }
                    }
                }
            }
        }

        return new Cubes4D(nextState);
    }

    public int countActive() {
        return cubes.size();
    }

    private boolean isActive(Position4D pos) {
        return cubes.contains(pos);
    }

    private int countNeighbours(Position4D pos) {
        int count = 0;
        for (int dw = -1; dw <= 1; ++dw) {
            int w = pos.w() + dw;
            for (int dz = -1; dz <= 1; ++dz) {
                int z = pos.z() + dz;
                for (int dy = -1; dy <= 1; ++dy) {
                    int y = pos.y() + dy;
                    for (int dx = -1; dx <= 1; ++dx) {
                        if (dw == 0 && dz == 0 && dy == 0 && dx == 0) continue;
                        int x = pos.x() + dx;
                        if (cubes.contains(new Position4D(x, y, z, w))) count++;
                    }
                }
            }
        }
        return count;
    }

    private Position4D[] getDimensions() {
        int minX = Integer.MAX_VALUE;
        int maxX = Integer.MIN_VALUE;
        int minY = Integer.MAX_VALUE;
        int maxY = Integer.MIN_VALUE;
        int minZ = Integer.MAX_VALUE;
        int maxZ = Integer.MIN_VALUE;
        int minW = Integer.MAX_VALUE;
        int maxW = Integer.MIN_VALUE;
        for (Position4D pos : cubes) {
            if (pos.x() < minX) minX = pos.x();
            if (pos.x() > maxX) maxX = pos.x();
            if (pos.y() < minY) minY = pos.y();
            if (pos.y() > maxY) maxY = pos.y();
            if (pos.z() < minZ) minZ = pos.z();
            if (pos.z() > maxZ) maxZ = pos.z();
            if (pos.w() < minW) minW = pos.w();
            if (pos.w() > maxW) maxW = pos.w();
        }
        return new Position4D[]{
            new Position4D(minX, minY, minZ, minW), new Position4D(maxX, maxY, maxZ, maxW)
        };
    }

    @Override
    public String toString() {
        return "TODO!";
    }
}

class Day17 {

    private static Cubes3D parsePart1(String filename) throws IOException {
        Set<Position3D> cubes = new HashSet<>();

        List<String> lines = Files.readAllLines(Path.of(filename));
        for (int y = 0; y < lines.size(); ++y) {
            String line = lines.get(y);
            for (int x = 0; x < line.length(); ++x) {
                if (line.charAt(x) == '#') {
                    cubes.add(new Position3D(x, y, 0));
                }
            }
        }

        return new Cubes3D(cubes);
    }

    private static Cubes4D parsePart2(String filename) throws IOException {
        Set<Position4D> cubes = new HashSet<>();

        List<String> lines = Files.readAllLines(Path.of(filename));
        for (int y = 0; y < lines.size(); ++y) {
            String line = lines.get(y);
            for (int x = 0; x < line.length(); ++x) {
                if (line.charAt(x) == '#') {
                    cubes.add(new Position4D(x, y, 0, 0));
                }
            }
        }

        return new Cubes4D(cubes);
    }

    private static int solvePart1(Cubes3D cubes) {
        Cubes3D current = cubes;
        for (int i = 0; i < 6; ++i) {
            current = current.nextCycle();
        }
        return current.countActive();
    }

    private static int solvePart2(Cubes4D cubes) {
        Cubes4D current = cubes;
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
            Cubes3D cubesPart1 = parsePart1(filename);
            Cubes4D cubesPart2 = parsePart2(filename);

            Common.time("Part1", () -> solvePart1(cubesPart1));
            Common.time("Part2", () -> solvePart2(cubesPart2));
        } catch (IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
