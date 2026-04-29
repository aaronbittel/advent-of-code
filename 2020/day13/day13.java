package day13;

import java.nio.file.Files;
import java.nio.file.Path;
import java.io.IOException;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Stream;
import java.util.stream.Collectors;

import common.Common;

record TimeTable(int department, List<Integer> buses) { }

class Day13 {

    private static TimeTable parse(String filename) throws IOException {
        List<String> lines = Files.readAllLines(Path.of(filename));
        int department = Integer.parseInt(lines.getFirst());
        List<Integer> buses = Arrays.stream(lines.get(1).split(","))
            .filter(busID -> !busID.equals("x"))
            .map(Integer::parseInt)
            .collect(Collectors.toList());
        return new TimeTable(department, buses);
    }

    private static int minutesTillNextDepartment(int department, int busID) {
        double times = (double)department / busID;
        double progress = times - (int)times;
        return (int)(busID - (busID * progress));
    }

    private static int solvePart1(TimeTable tt) {
        int id = -1;
        int waitTime = Integer.MAX_VALUE;
        for (int busID : tt.buses()) {
            int m = minutesTillNextDepartment(tt.department(), busID);
            if (m < waitTime) {
                waitTime = m;
                id = busID;
            }
        }
        return id * waitTime;
    }

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day13.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            TimeTable tt = parse(filename);

            Common.time("Part1", () -> solvePart1(tt));
        } catch(IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
