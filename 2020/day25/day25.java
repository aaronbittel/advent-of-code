package day25;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.stream.Stream;

import common.Common;

class Day25 {

    private static List<Integer> parse(String filename) throws IOException {
        try (Stream<String> lines = Files.lines(Path.of(filename))) {
            return lines
                .map(Integer::valueOf)
                .toList();
        }
    }

    private static int reverseEngineerLoopSize(int publicKey) {
        int subjectNumber = 1;
        int loopCount = 0;
        while (subjectNumber != publicKey) {
            subjectNumber *= 7;
            subjectNumber %= 20201227;
            loopCount++;
        }
        return loopCount;
    }

    private static long transformSubjectNumber(int subjectNumber, int loopCount) {
        long value = 1;
        for (int i = 0; i < loopCount; ++i) {
            value *= subjectNumber;
            value %= 20201227;
        }
        return value;
    }

    private static long solvePart1(int cardPublicKey, int doorPublicKey) {
        int cardLoopSize = reverseEngineerLoopSize(cardPublicKey);
        return transformSubjectNumber(doorPublicKey, cardLoopSize);
    }

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day25.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];

        try {
            List<Integer> publicKeys = parse(filename);
            if (publicKeys.size() != 2) {
                throw new IllegalArgumentException("input file is malformed");
            }

            int cardPublicKey = publicKeys.get(0);
            int doorPublicKey = publicKeys.get(1);

            Common.time("Part1", () -> solvePart1(cardPublicKey, doorPublicKey));
        } catch (IOException e) {
            System.err.printf("ERROR: reading file '%s': %s%n", filename, e.getMessage());
            System.exit(1);
        }
    }
}
