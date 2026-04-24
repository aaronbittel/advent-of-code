package day8;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.OptionalInt;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Stream;
import java.util.stream.Collectors;
import java.io.IOException;

import common.Common;

enum Operation {
    ACC,
    JMP,
    NOP
}

record Instruction(Operation operation, int argument) {

    @Override
    public String toString() {
        return operation.name() + " (" + argument + ")";
    }
}

class Day8 {

    private static List<Instruction> parse(String filename) throws IOException {
        try(Stream<String> lines = Files.lines(Path.of(filename))) {
            return lines
                    .map(Day8::lineToInstruction)
                    .collect(Collectors.toList());
        }
    }

    private static Instruction lineToInstruction(String line) {
        String[] parts = line.split(" ");
        int argument = Integer.parseInt(parts[1]);
        Operation op = Operation.valueOf(parts[0].toUpperCase());
        return new Instruction(op, argument);
    }

    private static int solvePart1(List<Instruction> instructions) {
        boolean[] visited = new boolean[instructions.size()];
        int acc = 0;
        int cur = 0;
        while (!visited[cur]) {
            Instruction inst = instructions.get(cur);
            visited[cur] = true;
            switch (inst.operation()) {
                case NOP: {
                    cur++;
                    break;
                }
                case ACC: {
                    cur++;
                    acc += inst.argument();
                    break;
                }
                case JMP: {
                    cur += inst.argument();
                    break;
                }
            }
        }
        return acc;
    }

    private static OptionalInt runAlteredProgram(
        List<Instruction> instructions,
        int acc,
        int cur,
        boolean[] visited
    ) {
        while (cur >= 0 && cur < instructions.size() && !visited[cur]) {
            Instruction inst = instructions.get(cur);
            visited[cur] = true;
            switch (inst.operation()) {
                case NOP: {
                    cur++;
                    break;
                }
                case ACC: {
                    cur++;
                    acc += inst.argument();
                    break;
                }
                case JMP: {
                    cur += inst.argument();
                    break;
                }
            }
        }
        if (cur == instructions.size()) return OptionalInt.of(acc);
        return OptionalInt.empty();
    }

    private static int solvePart2(List<Instruction> instructions) {
        boolean[] visited = new boolean[instructions.size()];
        int acc = 0;
        int curIndex = 0;
        int nextIndex = 0;
        while (!visited[curIndex]) {
            Instruction inst = instructions.get(curIndex);
            switch (inst.operation()) {
                case NOP: {
                    instructions.set(curIndex,
                        new Instruction(Operation.JMP, inst.argument()));
                    OptionalInt result = runAlteredProgram(
                        instructions,
                        acc,
                        curIndex,
                        Arrays.copyOf(visited, visited.length));
                    if (result.isPresent()) {
                        return result.getAsInt();
                    }
                    instructions.set(curIndex, inst);
                    nextIndex = curIndex + 1;
                    break;
                }
                case ACC: {
                    acc += inst.argument();
                    nextIndex = curIndex + 1;
                    break;
                }
                case JMP: {
                    instructions.set(curIndex,
                        new Instruction(Operation.NOP, inst.argument()));
                    OptionalInt result = runAlteredProgram(
                        instructions,
                        acc,
                        curIndex,
                        Arrays.copyOf(visited, visited.length));
                    if (result.isPresent()) {
                        return result.getAsInt();
                    }
                    instructions.set(curIndex, inst);
                    nextIndex = curIndex + inst.argument();
                    break;
                }
            }
            visited[curIndex] = true;
            curIndex = nextIndex;
        }
        return acc;
    }

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day8.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            List<Instruction> instructions = parse(filename);
            Common.time("Part 1", () -> solvePart1(instructions));
            Common.time("Part 2", () -> solvePart2(instructions));
        } catch(IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
