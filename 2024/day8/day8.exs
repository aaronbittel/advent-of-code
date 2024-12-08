defmodule Day8 do
    @type pos() :: {integer, integer}
    @type grid() :: %{pos => char}
    @type antennas() :: [[pos]]

    @spec parse_grid(String.t()) :: grid
    def parse_grid(content) do
        String.split(content, "\n", trim: true)
        |> Enum.with_index()
        |> Enum.reduce(%{}, fn {line, row}, acc ->
            line
            |> String.graphemes()
            |> Enum.with_index()
            |> Enum.reduce(%{}, fn {char, col}, acc ->
                Map.put(acc, {row, col}, char)
            end)
            |> Map.merge(acc)
        end)
    end

    @spec parse_antennas(grid) :: antennas
    def parse_antennas(grid) do
        grid
        |> Enum.reduce(%{}, fn {pos, char}, acc ->
            if char == "." do
                acc
            else
                Map.update(acc, char, [pos], fn existing_value ->
                    [pos | existing_value]
                end)
            end
        end)
        |> Map.values()
    end

    @spec part1(grid, MapSet.t(String.t())) :: integer
    def part1(grid, antennas) do
        Enum.reduce(antennas, MapSet.new(), fn poses, acc ->
            Comb.combinations(poses)
            |> Enum.reduce(MapSet.new(), fn {a1, a2}, inner_acc ->
                diff = Vec.diff(a1, a2)
                antinode1 = Vec.add(a1, diff)
                antinode2 = Vec.add(a2, Vec.neg(diff))

                inner_acc =
                if Map.has_key?(grid, antinode1) do
                    MapSet.put(inner_acc, antinode1)
                else
                    inner_acc
                end

                if Map.has_key?(grid, antinode2) do
                    MapSet.put(inner_acc, antinode2)
                else
                    inner_acc
                end
            end)
            |> MapSet.union(acc)
        end)
        |> MapSet.size()
    end

    @spec part2(grid, MapSet.t(String.t())) :: integer
    def part2(grid, antennas) do
        Enum.reduce(antennas, MapSet.new(), fn poses, acc ->
            Comb.combinations(poses)
            |> Enum.reduce(MapSet.new(), fn {a1, a2}, inner_acc ->
                diff = Vec.diff(a1, a2)
                inner_acc = MapSet.union(MapSet.new(AntinodeGenerator.generate(grid, a1, diff)), inner_acc)
                MapSet.union(MapSet.new(AntinodeGenerator.generate(grid, a2, Vec.neg(diff))), inner_acc)
            end)
            |> MapSet.union(acc)
        end)
        |> MapSet.size()
    end


    @spec solve(String.t()) :: nil
    def solve(filename) do
        {_, content} = File.read(filename)
        grid = parse_grid(content)
        antennas = parse_antennas(grid)

        {took, result} = :timer.tc(fn -> part1(grid, antennas) end)
        IO.puts("Part1: #{result}, took: #{took / 1_000_000} seconds")

        {took, result} = :timer.tc(fn -> part2(grid, antennas) end)
        IO.puts("Part2: #{result}, took: #{took / 1_000_000} seconds")
    end

end

defmodule AntinodeGenerator do
    @spec generate(Day8.grid, Day8.pos, Day8.pos) :: [Day8.pos]
    def generate(grid, start, dir) do
        generate(grid, start, dir, [start])
    end

    @spec generate(Day8.grid, Day8.pos, Day8.pos, [Day8.pos]) :: [Day8.pos]
    defp generate(grid, cur, dir, acc) do
        next = Vec.add(cur, dir)
        if Map.has_key?(grid, next) do
            generate(grid, next, dir, [next | acc])
        else
            acc
        end
    end
end

defmodule Comb do
    @spec combinations([Day8.pos]) :: [{Day8.pos, Day8.pos}]
    def combinations([head | tail]) do
        combinations(tail) ++ Enum.map(tail, fn pos -> {head, pos} end)
    end

    def combinations([]), do: []
end

defmodule Vec do
    @spec diff(Day8.pos, Day8.pos) :: Day8.pos
    def diff({v1y, v1x}, {v2y, v2x}), do: {v1y-v2y, v1x-v2x}

    @spec add(Day8.pos, Day8.pos) :: Day8.pos
    def add({v1y, v1x}, {v2y, v2x}), do: {v1y+v2y, v1x+v2x}

    @spec neg(Day8.pos) :: Day8.pos
    def neg({v1y, v1x}), do: {-v1y, -v1x}
end

Day8.solve("./input.txt")
