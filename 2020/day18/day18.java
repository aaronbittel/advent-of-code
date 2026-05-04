package day18;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Deque;
import java.util.List;
import java.util.stream.Stream;

import common.Common;

enum Op {
    ADDITION,
    MULTIPLICATION
}

class EvalStack {

    private final Deque<Long> numbers;
    private final Deque<Op> operators;

    public EvalStack() {
        this.numbers = new ArrayDeque<>();
        this.operators = new ArrayDeque<>();
    }

    public long result() {
        if (numbers.size() == 1 && operators.isEmpty()) {
            return numbers.getFirst();
        }
        throw new IllegalStateException("Calling 'result' in invalid state");
    }

    public void evaluateTop() {
        long a = numbers.pop();
        long b = numbers.pop();
        Op op = operators.pop();
        long r = switch (op) {
            case ADDITION -> a + b;
            case MULTIPLICATION -> a * b;
        };
        numbers.addLast(r);
    }

    public boolean canEvaluate() {
        return numbers.size() >= 2 && !operators.isEmpty();
    }

    public void push(long number) {
        numbers.addLast(number);
    }

    public void push(Op op) {
        operators.addLast(op);
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append("Numbers:").append("\n");
        for (long num : numbers.reversed()) {
            sb.append("    ").append(num).append("\n");
        }
        sb.append("Ops:").append("\n");
        for (Op op : operators.reversed()) {
            sb.append("    ").append(op).append("\n");
        }
        return sb.toString();
    }
}

class ImParser {

    private final String source;
    private int index = 0;
    private final List<EvalStack> stacks;

    public ImParser(String source) {
        this.source = source;
        this.stacks = new ArrayList<>(List.of(new EvalStack()));
    }

    public long evaluate() {
        while (index < source.length()) {
            skipWhitespace();

            char c = current();
            if (Character.isDigit(c)) {
                long number = parseNumber();
                stacks.getLast().push(number);
                if (stacks.getLast().canEvaluate()) {
                    stacks.getLast().evaluateTop();
                }
            } else if (c == '+') {
                stacks.getLast().push(Op.ADDITION);
                advance();
            } else if (c == '*') {
                stacks.getLast().push(Op.MULTIPLICATION);
                advance();
            } else if (c == '(') {
                stacks.add(new EvalStack());
                advance();
            } else if (c == ')') {
                long r = stacks.getLast().result();
                stacks.removeLast();
                stacks.getLast().push(r);
                advance();
                if (stacks.getLast().canEvaluate()) {
                    stacks.getLast().evaluateTop();
                }
            } else {
                throw new IllegalStateException("Illegal character: " + c);
            }

        }
        return stacks.getLast().result();
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
        if (isEof())
            return;
        index++;
    }

    private char current() {
        if (isEof())
            throw new IllegalStateException("current called on EOF");
        return source.charAt(index);
    }

    private void skipWhitespace() {
        while (!isEof() && current() == ' ') {
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
        for (EvalStack stack : stacks.reversed()) {
            for (String line : stack.toString().split("\n")) {
                sb.repeat(' ', indent).append(line).append("\n");
            }
            indent += 2;
        }

        return sb.toString();
    }
}

class Day18 {

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day18.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        Path filepath = Path.of(filename);
        try (Stream<String> lines = Files.lines(filepath)) {
            Common.time("Part1", () -> lines
                    .mapToLong(line -> new ImParser(line).evaluate())
                    .sum());
        } catch (IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }

        try (Stream<String> lines = Files.lines(filepath)) {
            Common.time("Part2", () -> lines
                    .mapToLong(line -> {
                        List<Token> tokens = new Lexer(line).tokenize();
                        Parser parser = new Parser(tokens);
                        Expr expr = parser.parseExpr();
                        return expr.eval();
                    })
                    .sum());
        } catch (IOException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
