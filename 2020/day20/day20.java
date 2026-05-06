package day20;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;

import common.Common;

class ImageData {
    private final char[][] data;

    public ImageData(char[][] data) {
        int size = data.length;
        for (int y = 0; y < size; ++y) {
            if (data[y].length != size) {
                throw new IllegalArgumentException("ImageData must be a square");
            }
        }
        char[][] copied = new char[size][];
        for (int y = 0; y < size; ++y) {
            copied[y] = Arrays.copyOf(data[y], size);
        }
        this.data = copied;
    }

    public ImageData rotate90() {
        int size = data.length;
        char[][] rotated = new char[size][size];
        for (int y = 0; y < size; ++y) {
            for (int x = 0; x < size; ++x) {
                rotated[x][size - 1 - y] = data[y][x];
            }
        }
        return new ImageData(rotated);
    }

    public ImageData flipH() {
        int size = data.length;
        char[][] flipped = new char[size][size];
        for (int y = 0; y < size; ++y) {
            flipped[size - 1 - y] = Arrays.copyOf(data[y], size);
        }
        return new ImageData(flipped);
    }

    public char[] north() {
        return Arrays.copyOf(data[0], data.length);
    }

    public char[] east() {
        int size = data.length;
        char[] column = new char[size];
        for (int y = 0; y < size; ++y) {
            column[y] = data[y][size - 1];
        }
        return column;
    }

    public char[] south() {
        return Arrays.copyOf(data[data.length - 1], data.length);
    }

    public char[] west() {
        int size = data.length;
        char[] column = new char[size];
        for (int x = 0; x < size; ++x) {
            column[x] = data[0][x];
        }
        return column;
    }

    @Override
    public String toString() {
        int height = data.length;
        int width = data[0].length;

        StringBuilder sb = new StringBuilder(height * (width + 1));
        for (int y = 0; y < height; ++y) {
            for (int x = 0; x < width; ++x) {
                sb.append(data[y][x]);
            }
            sb.append("\n");
        }
        return sb.toString();
    }
}

class Tile {
    private final int id;
    private final ImageData data;

    public Tile(int id, ImageData data) {
        this.id = id;
        this.data = data;
    }

    public boolean matches(Tile other) {
        return Arrays.equals(data.north(), other.data.south())
            || Arrays.equals(data.east(), other.data.west())
            || Arrays.equals(data.south(), other.data.north())
            || Arrays.equals(data.west(), other.data.east());
    }

    public Tile[] allOrientations() {
        Tile[] tiles = new Tile[8];
        Tile cur = this;
        int index = 0;
        for (int i = 0; i < 4; ++i) {
            tiles[index++] = cur;
            cur = cur.rotate90();
        }
        cur = cur.flipH();
        for (int i = 0; i < 4; ++i) {
            tiles[index++] = cur;
            cur = cur.rotate90();
        }
        return tiles;
    }

    public Tile rotate90() {
        return new Tile(id(), data.rotate90());
    }

    public Tile flipH() {
        return new Tile(id(), data.flipH());
    }

    public int id() {
        return id;
    }

    @Override
    public String toString() {
        return "Tile: %d%n%s".formatted(id, data);
    }
}

class Day20 {

    private static List<Tile> parse(String filename) throws IOException {
        List<Tile> tiles = new ArrayList<>();
        String data = Files.readString(Path.of(filename));
        for (String batch : data.split("\n\n")) {
            tiles.add(parseTile(batch));
        }
        return tiles;
    }

    private static Tile parseTile(String batch) {
        String[] lines = batch.split("\n");
        if (lines.length <= 1) {
            throw new IllegalArgumentException("Batch must contain 'id' and 'data'");
        }
        String idLine = lines[0];
        if (!idLine.matches("Tile \\d+:")) {
            throw new IllegalArgumentException("IdLine has unknown format");
        }
        int idStart = idLine.indexOf(' ');
        int idEnd = idLine.indexOf(':');
        int id = Integer.parseInt(idLine.substring(idStart + 1, idEnd));

        return new Tile(id, parseImageData(Arrays.copyOfRange(lines, 1, lines.length)));
    }

    private static ImageData parseImageData(String[] lines) {
        if (lines.length == 0) {
            throw new IllegalArgumentException("ImageData must not be empty");
        }

        int height = lines.length;
        int width = lines[0].length();

        char[][] data = new char[height][width];
        for (int y = 0; y < height; ++y) {
            String line = lines[y];
            for (int x = 0; x < width; ++x) {
                data[y][x] = line.charAt(x);
            }
        }
        return new ImageData(data);
    }

    private static Map<Integer, List<Integer>> findTileMatches(List<Tile> tiles) {
        Map<Integer, List<Integer>> tileMatches = new HashMap<>();
        for (Tile tile : tiles) {
            tileMatches.put(tile.id(), new ArrayList<>(4));
        }

        for (int i = 0; i < tiles.size(); ++i) {
            Tile base = tiles.get(i);
            for (int j = 0; j < tiles.size(); ++j) {
                if (i == j) continue;
                Tile other = tiles.get(j);
                for (Tile o1 : base.allOrientations()) {
                    boolean found = false;
                    for (Tile o2 : other.allOrientations()) {
                        if (o1.matches(o2)) {
                            found = true;
                            tileMatches.get(base.id()).add(other.id());
                            break;
                        }
                    }
                    if (found) break;
                }
            }
        }

        return tileMatches;
    }

    private static long solvePart1(Map<Integer, List<Integer>> tileMatches) {
        return tileMatches.entrySet().stream()
            .filter(entry -> entry.getValue().size() == 2)
            .mapToLong(entry -> entry.getKey())
            .reduce(1, (acc, elem) -> acc * elem);
    }

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day20.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            List<Tile> tiles = parse(filename);

            Common.time("Part1", () -> {
                Map<Integer, List<Integer>> tileMatches = findTileMatches(tiles);
                return solvePart1(tileMatches);
            });
        } catch (IOException e) {
            System.err.printf("ERROR: reading file '%s': %s%n", filename, e.getMessage());
            System.exit(1);
        }
    }
}
