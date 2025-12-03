defmodule Day12 do
    @directions [{-1, 0}, {0, 1}, {1, 0}, {0, -1}]

    @type pos() :: {integer, integer}
    @type grid() :: %{pos => char}

    @spec parse(String.t()) :: grid()

    def parse(filename) do
        File.read!(filename)
        |> String.split("\n", trim: true)
        |> Enum.with_index()
        |> Enum.reduce(%{}, fn {line, y}, acc ->
            line
            |> String.graphemes()
            |> Enum.with_index()
            |> Enum.reduce(acc, fn {char, x}, inner_acc ->
                Map.put(inner_acc, {y, x}, char)
            end)
        end)
    end

    @spec plot(grid(), pos(), char(), integer(), MapSet.t(pos())) :: {MapSet.t(pos()), integer()}
    def plot(map, pos, char, perimeter, plot_area) do
        next =
            @directions
            |> Enum.map(fn {dy, dx} -> {elem(pos, 0) + dy, elem(pos, 1) + dx} end)
            |> Enum.filter(fn neighbor ->
                Map.has_key?(map, neighbor) and Map.get(map, neighbor) == char
            end)

        perimeter = perimeter + (4 - length(next))

        new_positions =
            next
            |> Enum.reject(fn n -> MapSet.member?(plot_area, n) end)

        plot_area =
            Enum.reduce(new_positions, plot_area, fn n, acc -> MapSet.put(acc, n) end)

        case new_positions do
            [] -> {plot_area, perimeter}
            _ ->
                Enum.reduce(new_positions, {plot_area, perimeter}, fn new_pos, {acc_area, acc_perim} ->
                    plot(map, new_pos, char, acc_perim, acc_area)
                end)
        end
    end

    @spec part1(grid) :: integer
    def part1(map) do
        all_positions = MapSet.new(Map.keys(map))

        process_regions(map, all_positions, [])
        |> Enum.reduce(0, fn {poses, perimeter}, acc ->
            acc + length(poses) * perimeter
        end)
    end

    defp process_regions(map, all_positions, results) do
        if MapSet.size(all_positions) == 0 do
            results
        else
            pos = hd(MapSet.to_list(all_positions))
            char = Map.get(map, pos)

            {visited, perimeter} = plot(map, pos, char, 0, MapSet.new([pos]))

            result = {MapSet.to_list(visited), perimeter}

            remaining_positions = MapSet.difference(all_positions, visited)

            process_regions(map, remaining_positions, [result | results])
        end
    end


    @spec solve(String.t()) :: nil
    def solve(filename) do
        map = parse(filename)

        {took, result} = :timer.tc(fn -> part1(map) end)
        IO.puts("Part1: #{result}, took: #{took / 1_000_000} seconds")

        # {took, result} = :timer.tc(fn -> part1(nums, 75) end)
        # IO.puts("Part2: #{result}, took: #{took / 1_000_000} seconds")
    end

end

Day12.solve("./input.txt")
