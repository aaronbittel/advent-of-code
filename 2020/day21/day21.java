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

    private static boolean isValid(List<Meal> meals, Map<String, String> ingredientToAllergen) {
        Map<String, String> allergenToIngredient = HashMap.newHashMap(ingredientToAllergen.size());
        for (Map.Entry<String, String> entry : ingredientToAllergen.entrySet()) {
            allergenToIngredient.put(entry.getValue(), entry.getKey());
        }

        for (Meal meal : meals) {
            // count free ingredients < count free allergens
            int freeIngredientCount = 0;
            for (String ingredient : meal.ingredients()) {
                if (!ingredientToAllergen.containsKey(ingredient)) freeIngredientCount++;
            }
            int freeAllergenCount = 0;
            for (String allergen : meal.allergens()) {
                if (!allergenToIngredient.containsKey(allergen)) freeAllergenCount++;
            }
            if (freeIngredientCount < freeAllergenCount) return false;
        }

        // ingredient matches allergen, but ingredient not present in meal
        for (Meal meal : meals) {
            for (String allergen : meal.allergens()) {
                if (!allergenToIngredient.containsKey(allergen)) continue;
                String ingredient = allergenToIngredient.get(allergen);
                if (!meal.ingredients().contains(ingredient)) return false;
            }
        }

        return true;
    }

    private static boolean backtrack(List<Meal> meals, Map<String, String> ingredientToAllergen) {
        // 2. check if valid
        // 3. if yes -> continue 1.
        if (!isValid(meals, ingredientToAllergen)) return false;
        Set<String> uniqueAllergens = new HashSet<>();
        for (Meal meal : meals) {
            for (String allergen : meal.allergens()) {
                uniqueAllergens.add(allergen);
            }
        }
        if (ingredientToAllergen.size() == uniqueAllergens.size()) return true;

        // 1. choose next solution
        for (Meal meal : meals) {
            for (String ingredient : meal.ingredients()) {
                if (ingredientToAllergen.containsKey(ingredient)) continue;
                for (String allergen : meal.allergens()) {
                    if (ingredientToAllergen.containsValue(allergen)) continue;
                    ingredientToAllergen.put(ingredient, allergen);
                    if (backtrack(meals, ingredientToAllergen)) return true;
                    // 4. if no -> revert change
                    ingredientToAllergen.remove(ingredient);
                }
            }
        }

        return true;
    }

    private static int solvePart1(List<Meal> meals) {
        Map<String, String> ingredientToAllergen = new HashMap<>();
        backtrack(meals, ingredientToAllergen);

        int result = 0;
        for (Meal meal : meals) {
            for (String ingredient : meal.ingredients()) {
                if (!ingredientToAllergen.containsKey(ingredient)) result++;
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
