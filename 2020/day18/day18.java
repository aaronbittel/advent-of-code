package day18;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.List;
import java.util.stream.Stream;

import common.Common;

enum Op {
    ADDITION,
    MULTIPLICATION
};

class Stack {

    private List<Long> numbers;
    private List<Op> operations;

    public Stack() {
        this.numbers = new ArrayList<>();
        this.operations = new ArrayList<>();
    }

    public long result() {
        if (numbers.size() == 1 && operations.isEmpty()) {
            return numbers.getFirst();
        }
        throw new IllegalStateException("Calling 'result' in invalid state");
    }

    public boolean isEmpty() {
        return numbers.isEmpty() && operations.isEmpty();
    }

    public int sizeNum() {
        return numbers.size();
    }

    public int sizeOp() {
        return operations.size();
    }

    public long popNum() {
        return numbers.removeLast();
    }

    public Op popOp() {
        return operations.removeLast();
    }

    public void push(long number) {
        numbers.add(number);
    }

    public void push(Op op) {
        operations.add(op);
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append("Numbers:").append("\n");
        for (long num : numbers.reversed()) {
            sb.append("    ").append(num).append("\n");
        }
        sb.append("Ops:").append("\n");
        for (Op op : operations.reversed()) {
            sb.append("    ").append(op).append("\n");
        }
        return sb.toString();
    }
}

class Parser {

    private final String source;
    private int index = 0;
    private List<Stack> stacks;

    public Parser(String source) {
        this.source = source;
        this.index = 0;
        this.stacks = new ArrayList<>(List.of(new Stack()));
    }

    public long evaluate() {
        while(index < source.length()) {
            skipWhitespace();

            char c = current();
            if (Character.isDigit(c)) {
                long number = parseNumber();
                stacks.getLast().push(number);
                if (canEvalulate()) pushEvalResult();
            } else if (c == '+') {
                stacks.getLast().push(Op.ADDITION);
                advance();
            } else if (c == '*') {
                stacks.getLast().push(Op.MULTIPLICATION);
                advance();
            } else if (c == '(') {
                stacks.add(new Stack());
                advance();
            } else if (c == ')') {
                long r = stacks.getLast().result();
                stacks.removeLast();
                stacks.getLast().push(r);
                advance();
                if (canEvalulate()) pushEvalResult();
            } else {
                throw new IllegalStateException("Illegal character: " + c);
            }

        }
        return stacks.getLast().result();
    }

    private void pushEvalResult() {
        long a = stacks.getLast().popNum();
        long b = stacks.getLast().popNum();
        Op op = stacks.getLast().popOp();
        long r = switch (op) {
            case ADDITION -> a + b;
            case MULTIPLICATION -> a * b;
        };
        stacks.getLast().push(r);
    }

    private boolean canEvalulate() {
        return stacks.getLast().sizeNum() >= 2 && stacks.getLast().sizeOp() >= 1;
    }

    private int parseNumber() {
        int start = index;
        while (!isEof() && Character.isDigit(current())) {
            advance();
        }
        int end = index;
        return Integer.parseInt(source.substring(start, end));
    }

    private void advance() {
        if (isEof()) return;
        index++;
    }

    private char current() {
        if (isEof()) throw new IllegalStateException("current called on EOF");
        return source.charAt(index);
    }

    private void skipWhitespace() {
        while(!isEof() && current() == ' ') {
            index++;
        }
    }

    private boolean isEof() {
        return index >= source.length();
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();

        int indent = 0;
        for (Stack stack : stacks.reversed()) {
            for (String line : stack.toString().split("\n")) {
                sb.repeat(' ', indent).append(line).append("\n");
            }
            indent += 2;
        }

        return sb.toString();
    }
}

class Day18 {

    private static final String simple = "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2";

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day18.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try (Stream<String> lines = Files.lines(Path.of(filename))) {
            Common.time("Part1", () -> {
                return lines
                        .mapToLong(line -> new Parser(line).evaluate())
                        .sum();
            });
        } catch (IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
