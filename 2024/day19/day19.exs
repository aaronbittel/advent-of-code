defmodule Day19 do

    @spec parse(String.t()) :: {[String.t()], [String.t()]}
    def parse(filename) do
        [patterns, designs] = filename
        |> File.read!()
        |> String.split("\n\n", trim: true)

        {parse_patterns(patterns), parse_designs(designs)}
    end

    @spec parse_patterns(String.t()) :: [String.t()]
    defp parse_patterns(patterns) do
        patterns
            |> String.split(", ")
    end

    @spec parse_designs(String.t()) :: [String.t()]
    defp parse_designs(designs) do
        designs
        |> String.split("\n", trim: true)
    end

    @spec possible?(String.t(), [String.t()], [String.t()]) :: boolean()
    defp possible?(_, _, []), do: false

    @spec possible?(String.t(), [String.t()], [String.t()]) :: boolean()
    defp possible?(design, patterns, [next_pattern | rest]) do
        # IO.puts("\tCurrent design: #{design}")
        # IO.puts("\tNext pattern: #{next_pattern}, Remaining patterns: #{rest}")

        if design == "" do
            true
        else
            cond do
                not String.starts_with?(design, next_pattern) ->
                    # IO.puts("\tDoes not start with: #{next_pattern}")
                    possible?(design, patterns, rest)

                true ->
                    new_design = String.replace_prefix(design, next_pattern, "")
                    # IO.puts("\tNew design after removing #{next_pattern}: #{new_design}")
                    possible?(new_design, patterns, patterns) or possible?(design, patterns, rest)
            end
        end
    end

    def part1(patterns, designs) do
        designs
        |> Enum.filter(fn d ->
            possible_patterns = Enum.filter(patterns, fn pat ->
                String.contains?(d, pat)
            end)
            # IO.puts("CHECKING: #{d} with patterns: #{Enum.join(possible_patterns, ", ")}")
            result = possible?(d, patterns, possible_patterns)
            # if result, do: IO.puts("#{d} is POSSIBLE!")
            # IO.puts("")
            result
        end)
        |> Enum.count()
    end

    @spec solve(String.t()) :: nil
    def solve(filename) do
        {patterns, designs} = parse(filename)

        {took, result} = :timer.tc(fn -> part1(patterns, designs) end)
        IO.puts("Part1: #{result}, took: #{took / 1_000_000} seconds")

        # {took, result} = :timer.tc(fn -> part2(obstacles, count, is_input) end)
        # IO.puts("Part1: #{result}, took: #{took / 1_000_000} seconds")
    end
end

Day19.solve("./input.txt")
