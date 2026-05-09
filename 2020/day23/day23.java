package day23;

import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.stream.Collectors;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

import common.Common;

class CrabCupsGame {
    private final CupCircle cups;
    private final int max_label;
    private final int min_label;

    public CrabCupsGame(List<Integer> numbers) {
        cups = new CupCircle(numbers);
        Collections.sort(numbers);
        min_label = numbers.getFirst();
        max_label = numbers.getLast();
    }

    public String labelsAfter(int label) {
        return cups.labelsAfter(label);
    }

    public void move() {
        CupCircle.CupNode pickedUp = cups.removeNext3();

        int destinationLabel = findDestinationNode(cups.label(), pickedUp);
        cups.insertAfter(destinationLabel, pickedUp);
        cups.advance();
    }

    public int findDestinationNode(int currentLabel, CupCircle.CupNode pickedUp) {
        int destinationLabel = currentLabel - 1;
        if (destinationLabel < min_label) destinationLabel = max_label;
        for (int i = 0; i < 3; ++i) {
            if (!CupCircle.containsInChain(pickedUp, destinationLabel)) break;
            destinationLabel--;
            if (destinationLabel < min_label) destinationLabel = max_label;
        }
        return destinationLabel;
    }

    @Override
    public String toString() {
        return cups.toString();
    }
}

class CupCircle {

    private CupNode current;

    public CupCircle(List<Integer> numbers) {
        if (numbers.isEmpty()) {
            throw new IllegalArgumentException("numbers must not be empty");
        }

        current = new CupNode(numbers.getFirst());
        CupNode tail = current;

        for (int i = 1; i < numbers.size(); ++i) {

            tail.next = new CupNode(numbers.get(i));
            tail = tail.next;
        }

        tail.next = current;
    }

    public String labelsAfter(int label) {
        CupNode start = findByLabel(label);

        StringBuilder sb = new StringBuilder();

        CupNode current = start.next;

        while (current != start) {
            sb.append(current.label);
            current = current.next;
        }

        return sb.toString();
    }

    public CupNode removeNext3() {
        CupNode node = current.next;
        CupNode last = node.next.next;
        current.next = last.next;
        last.next = null;
        return node;
    }

    public void insertAfter(int label, CupNode head) {
        CupNode node = findByLabel(label);
        head.next.next.next = node.next;
        node.next = head;
    }

    public CupNode findByLabel(int label) {
        if (current == null) {
            throw new IllegalStateException("CupNode is null");
        }
        if (current.label == label) return current;
        CupNode node = current.next;
        while (node.label != current.label) {
            if (node.label == label) return node;
            node = node.next;
        }
        throw new IllegalArgumentException(label + " is not in CupNode");
    }

    public static boolean containsInChain(CupNode head, int num) {
        if (head == null) return false;
        CupNode current = head;
        while (current.next != null) {
            if (current.label == num) return true;
            current = current.next;
        }
        return current.label == num;
    }

    public void advance() {
        current = current.next;
    }

    public int label() {
        return current.label;
    }

    public static String asChainString(CupNode head) {
        if (head == null) return "<empty>";
        StringBuilder sb = new StringBuilder();
        CupNode current = head;
        sb.append(current.label);
        while (current.next != null) {
            sb.append(" -> ").append(current.next.label);
            current = current.next;
        }
        return sb.toString();
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append(current.label);
        CupNode node = current.next;
        while (node.label != current.label) {
            sb.append(" -> ").append(node);
            node = node.next;
        }
        return sb.toString();
    }

    static class CupNode {
        private final int label;
        private CupNode next;

        public CupNode(int number) {
            this.label = number;
        }

        @Override
        public String toString() {
            return String.valueOf(label);
        }
    }
}

class Day23 {

    private static CrabCupsGame parse(String filename) throws IOException {
        String input = Files.readString(Path.of(filename)).trim();
        return new CrabCupsGame(Arrays.stream(input.split(""))
            .map(Integer::parseInt)
            .collect(Collectors.toList()));
    }

    private static String solvePart1(CrabCupsGame game) {
        for (int i = 0; i < 100; ++i) {
            game.move();
        }
        return game.labelsAfter(1);
    }

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day23.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            CrabCupsGame game = parse(filename);

            Common.time("Part1", () -> solvePart1(game));
        } catch (IOException e) {
            System.err.printf("ERROR: reading file '%s': %s%n", filename, e.getMessage());
            System.exit(1);
        }
    }
}
