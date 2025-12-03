defmodule Day16 do
    @directions [{-1, 0}, {0, 1}, {1, 0}, {0, -1}]

    @type pos() :: {integer(), integer()}
    @type grid() :: %{pos() => char()}

    @spec parse(String.t()) :: {grid(), pos()}
    def parse(filename) do
        map = filename
            |> File.read!()
            |> String.split("\n", trim: true)
            |> Enum.with_index()
            |> Enum.flat_map(fn {line, y} ->
                line
                |> String.graphemes()
                |> Enum.with_index()
                |> Enum.map(fn {char, x} -> {{y, x}, char} end)
            end)
            |> Map.new()

        start_pos = map
            |> Enum.find(fn {_, char} -> char == "S" end)
            |> elem(0)

        {map, start_pos}
    end

    @spec part1(grid(), pos()) :: integer()
    def part1(map, start_pos) do
        res = dfs(map, start_pos, {0, 1}, MapSet.new(), [])
        0
    end

    @spec dfs(grid(), pos(), pos(), MapSet.t(), [pos()]) :: nil
    defp dfs(map, {y, x}, dir, visited, steps) do

    end

    @spec solve(String.t()) :: nil
    def solve(filename) do
        {map, start_pos} = parse(filename)

        {took, result} = :timer.tc(fn -> part1(map, start_pos) end)
        IO.puts("Part1: #{result}, took: #{took / 1_000_000} seconds")

        # {took, result} = :timer.tc(fn -> part2(robots, size) end)
        # IO.puts("Part2: #{result}, took: #{took / 1_000_000} seconds")
    end
end

Day16.solve("./test.txt")
