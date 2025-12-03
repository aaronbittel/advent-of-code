defmodule Day4 do
    @directions [
        {-1, 0},
        {-1, 1},
        {-1, -1},
        {1, 0},
        {1, 1},
        {1, -1},
        {0, 1},
        {0, -1},
    ]

    @spec parse(String.t()) :: %{{integer(), integer()} => String.t()}
    defp parse(filename) do
        File.read!(filename)
        |> String.trim()
        |> String.split("\n", trim: true)
        |> Enum.with_index()
        |> Enum.reduce(%{}, fn {line, row}, acc ->
            String.codepoints(line)
            |> Enum.with_index()
            |> Enum.reduce(acc, fn {char, col}, acc ->
                Map.put(acc, {row, col}, char)
            end)
        end)
    end

    @spec part1(String.t()) :: integer()
    def part1(filename) do
        map = parse(filename)
        Map.keys(map)
        |> Enum.filter(fn key ->
            Map.get(map, key) == "X"
        end)
        |> Enum.map(fn key -> count_xmas(map, key) end)
        |> Enum.sum()
    end

    @spec part2(String.t()) :: integer()
    def part2(filename) do
        map = parse(filename)
        Map.keys(map)
        |> Enum.filter(fn key ->
            Map.get(map, key) == "A"
        end)
        |> Enum.map(fn key -> x_mas(map, key) end)
        |> Enum.sum()
    end

    @spec count_xmas(%{{integer(), integer()} => String.t()}, {integer(), integer()}) :: integer()
    defp x_mas(map, {row, col}) do
        one = Map.get(map, {row-1, col-1}, ".") <> Map.get(map, {row+1, col+1}, ".")
        two = Map.get(map, {row+1, col-1}, ".") <> Map.get(map, {row-1, col+1}, ".")
        if (one == "MS" or one == "SM") and (two == "MS" or two == "SM") do
            1
        else
            0
        end
    end

    @spec count_xmas(%{{integer(), integer()} => String.t()}, {integer(), integer()}) :: integer()
    defp count_xmas(map, {row, col}) do
        Enum.reduce(@directions, 0, fn {dr, dc}, acc ->
            acc + check_direction(map, row, col, dr, dc)
        end)
    end

    @spec check_direction(%{{integer(), integer()} => String.t()}, integer(), integer(), integer(), integer()) :: integer()
    defp check_direction(map, row, col, dr, dc) do
        Enum.reduce(0..2, true, fn i, valid ->
            valid && Map.get(map, {row + (i+1)*dr, col + (i+1)*dc}, ".") == String.at("MAS", i)
        end)
        |> (if do: 1, else: 0)
    end
end

{took, result} = :timer.tc(fn -> Day4.part1("./input.txt") end)
IO.puts("Part1: #{result}, took: #{took / 1_000_000} seconds")

{took, result} = :timer.tc(fn -> Day4.part2("./input.txt") end)
IO.puts("Part2: #{result}, took: #{took / 1_000_000} seconds")
