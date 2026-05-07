package day21;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Arrays;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.stream.Stream;

import common.Common;

record Meal (List<String> ingredients, List<String> allergens) {
    public Meal {
        ingredients = List.copyOf(ingredients);
        allergens = List.copyOf(allergens);
    }
}

class Day21 {

    private static List<Meal> parse(String filename) throws IOException {
        try (Stream<String> lines = Files.lines(Path.of(filename))) {
            return lines
                .map(Day21::lineToMeal)
                .toList();
        }
    }

    private static Meal lineToMeal(String line) {
        int delimIdx = line.indexOf(" (contains ");
        if (delimIdx == -1) {
            throw new IllegalArgumentException("Line has no allergens");
        }

        List<String> ingredients = Arrays.asList(line.substring(0, delimIdx).split(" "));
        List<String> allergens = Arrays.asList(
            line.substring(delimIdx + 11, line.length() - 1).split(", "));
        return new Meal(ingredients, allergens);
    }

    private static Map<String, Set<String>> candidatesByAllergen(List<Meal> meals) {
        Map<String, Set<String>> candidatesByAllergen = new HashMap<>();

        for (Meal meal : meals) {
            for (String allergen : meal.allergens()) {
                if (!candidatesByAllergen.containsKey(allergen)) {
                    candidatesByAllergen.put(allergen, new HashSet<String>(meal.ingredients()));
                } else {
                    candidatesByAllergen.get(allergen).retainAll(meal.ingredients());
                }
            }
        }

        return candidatesByAllergen;
    }

    private static int solvePart1(List<Meal> meals) {
        Map<String, Set<String>> candidatesByAllergen = candidatesByAllergen(meals);

        Set<String> resolvedIngredients = new HashSet<>();

        while (true) {
            for (Set<String> candidates : candidatesByAllergen.values()) {
                if (candidates.size() == 1) {
                    resolvedIngredients.addAll(candidates);
                }
            }

            for (Set<String> candidate : candidatesByAllergen.values()) {
                if (candidate.size() == 1) continue;
                candidate.removeAll(resolvedIngredients);
            }

            boolean allResolved = true;
            for (Set<String> candidate : candidatesByAllergen.values()) {
                if (candidate.size() > 1) {
                    allResolved = false;
                    break;
                }
            }

            if (allResolved) break;
        }

        Set<String> allergenic = new HashSet<>();
        for (Set<String> ingredientSet : candidatesByAllergen.values()) {
            allergenic.addAll(ingredientSet);
        }

        int result = 0;
        for (Meal meal : meals) {
            for (String ingredient : meal.ingredients()) {
                if (!allergenic.contains(ingredient)) result++;
            }
        }
        return result;
    }

    static void main(String[] args) {
        if (args.length < 1) {
            System.err.printf("Usage: java %s <input>%n", Day21.class.getSimpleName());
            System.exit(1);
        }

        String filename = args[0];
        try {
            List<Meal> meals = parse(filename);

            Common.time("Part1", () -> solvePart1(meals));
        } catch (IOException e) {
            System.err.printf("ERROR: reading file '%s': %s%n", filename, e.getMessage());
            System.exit(1);
        }
    }
}
