defmodule Day10 do
    @type pos() :: {integer, integer}
    @type lava_map() :: %{pos => integer}

    @directions [{-1, 0}, {0, 1}, {1, 0}, {0, -1}]

    @spec parse(String.t()) :: lava_map
    defp parse(filename) do
        filename
        |> File.read!()
        |> String.split("\n", trim: true)
        |> Enum.with_index()
        |> Enum.reduce({%{}, []}, fn {line, y}, {map, starting} ->
            line
            |> String.graphemes()
            |> Enum.map(&String.to_integer/1)
            |> Enum.with_index()
            |> Enum.reduce({map, starting}, fn {height, x}, {acc_map, acc_start} ->
                new_starting = if height == 0, do: [{y, x} | acc_start], else: acc_start
                new_map = Map.put(acc_map, {y, x}, height)
                {new_map, new_starting}
            end)
        end)
    end

    @spec part1(lava_map(), [pos()]) :: integer
    def part1(map, starts) do
        starts
        |> Enum.map(fn start -> check_path(map, start, 0, 0) end)
        |> IO.inspect()
        |> Enum.sum()
    end

    @spec check_path(lava_map(), pos(), integer, integer) :: integer
    defp check_path(map, pos, cur_height, count) do
        Enum.reduce(@directions, count, fn {dy, dx}, acc_count ->
            new_pos = {elem(pos, 0) + dy, elem(pos, 1) + dx}

            case Map.fetch(map, new_pos) do
                {:ok, height} when cur_height == 8 and height == 9 ->
                    IO.puts("mountain top")
                    IO.inspect(new_pos)
                    IO.inspect(acc_count + 1, label: "score")
                {:ok, height} when cur_height+1 == height ->
                    IO.inspect(new_pos, label: cur_height)
                    check_path(map, new_pos, height, count)
                _ ->
                    acc_count
            end
        end)
    end

    @spec solve(String.t()) :: nil
    def solve(filename) do
        {map, starts} = parse(filename)
        IO.inspect(check_path(map, {0, 0}, 0, 0))

        # {took, result} = :timer.tc(fn -> part1(map, starts) end)
        # IO.puts("Part1: #{result}, took: #{took / 1_000_000} seconds")

        # {took, result} = :timer.tc(fn -> part2(grid, antennas) end)
        # IO.puts("Part2: #{result}, took: #{took / 1_000_000} seconds")
    end

end

Day10.solve("./test_small.txt")
