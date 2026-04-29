package day13;

import java.nio.file.Files;
import java.nio.file.Path;
import java.io.IOException;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

import common.Common;

record BusSchedule(int earliestDeparture, List<Integer> busIDs) { }
record Alignment(long timestamp, long step) { }

class Day13 {

    private static BusSchedule parse(String filename) throws IOException {
        List<String> lines = Files.readAllLines(Path.of(filename));
        int earliestDeparture = Integer.parseInt(lines.getFirst());
        List<Integer> busIDs = Arrays.stream(lines.get(1).split(","))
            .map(id -> {
                if (id.equals("x")) return null;
                return Integer.parseInt(id);
            })
            .collect(Collectors.toList());
        return new BusSchedule(earliestDeparture, busIDs);
    }

    private static int waitTimeForBus(int department, int busID) {
        double times = (double)department / busID;
        double progress = times - (int)times;
        return (int)(busID - (busID * progress));
    }

    private static int solvePart1(BusSchedule busSchedule) {
        int id = -1;
        int waitTime = Integer.MAX_VALUE;
        for (Integer busID : busSchedule.busIDs()) {
            if (busID == null) continue;
            int m = waitTimeForBus(busSchedule.earliestDeparture(), busID);
            if (m < waitTime) {
                waitTime = m;
                id = busID;
            }
        }
        return id * waitTime;
    }

    private static Alignment combineAlignment(Alignment alignment, int num, int offset) {
        for(long t = alignment.timestamp(); ; t += alignment.step()) {
            if ((t + offset) % num == 0) {
                return new Alignment(t, alignment.step() * num);
            }
        }
    }

    private static long solvePart2(List<Integer> busIDs) {
        Alignment c = new Alignment(0, busIDs.getFirst());
        for (int i = 1; i < busIDs.size(); ++i) {
            if (busIDs.get(i) == null) continue;
            c = combineAlignment(c, busIDs.get(i), i);
        }
        return c.timestamp();
    }

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day13.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            BusSchedule busSchedule = parse(filename);

            Common.time("Part1", () -> solvePart1(busSchedule));
            Common.time("Part2", () -> solvePart2(busSchedule.busIDs()));
        } catch(IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
