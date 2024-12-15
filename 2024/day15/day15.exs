defmodule Day15 do
    @type pos() :: {integer(), integer()}
    @type grid() :: %{pos() => char()}

    @spec parse(String.t()) :: {grid(), pos(), String.t()}
    def parse(filename) do
        [map, steps] = filename
        |> File.read!()
        |> String.split("\n\n", trim: true)

        {map, pos} = parse_map(map)
        steps = parse_steps(steps)

        {map, pos, steps}
    end

    @spec parse_map(String.t()) :: {grid(), pos()}
    defp parse_map(map) do
        map
        |> String.split("\n", trim: true)
        |> Enum.with_index()
        |> Enum.flat_map_reduce(nil, fn {line, y}, acc ->
            res = line
                |> String.graphemes()
                |> Enum.with_index()
                |> Enum.map(fn {char, x} ->
                    pos = if char == "@", do: {y, x}, else: nil
                    {{y, x}, char, pos}
                end)

            new_acc = res
                |> Enum.reduce(acc, fn {_, _, pos}, acc ->
                    if pos, do: pos, else: acc
                end)

            {res, new_acc}
        end)
        |> then(fn {parsed_map, robot} ->
            map = parsed_map
            |> Enum.map(fn {pos, char, _} -> {pos, char} end)
            |> Map.new()

            {map, robot}
        end)
    end

    @spec parse_steps(String.t()) :: String.t()
    defp parse_steps(steps) do
        steps
        |> String.split("\n", trim: true)
        |> Enum.join("")
        |> String.graphemes()
    end

    @spec part1(grid(), pos(), String.t()) :: integer()
    def part1(map, pos, steps) do
        {final_map, _} =
            steps
            |> Enum.reduce({map, pos}, fn step, {current_map, current_pos} ->
                {new_map, new_pos} = move(current_map, current_pos, step)
                {new_map, new_pos}
            end)

        final_map
        |> Enum.filter(fn {_, char} -> char == "O" end)
        |> Enum.map(fn {{y, x}, _} ->
            y*100 + x
        end)
        |> Enum.sum()
    end

    defp print_map(map, dir \\ "") do
        if dir == "", do: IO.puts(""), else: IO.puts("dir: #{dir}")
        for y <- 0..50 do
            for x <- 0..50 do
                IO.write("#{Map.get(map, {y, x})}")
            end
            IO.puts("")
        end
    end

    defp check_pos(map, {y,x}, {d_y, d_x}) do
        case Map.get(map, {y, x}) do
            "O" -> check_pos(map, {y+d_y, x+d_x}, {d_y, d_x})
            "." -> {:ok, {y, x}}
            "#" -> :no_push
        end
    end

    defp move(map, {y, x}, step) do
        {d_y, d_x} = case step do
            "^" -> {-1, 0}
            ">" -> {0, 1}
            "v" -> {1, 0}
            "<" -> {0, -1}
        end

        new_pos = {y+d_y, x+d_x}

        case check_pos(map, new_pos, {d_y, d_x}) do
            {:ok, next_free_pos} ->
                map = Map.put(map, {y,x}, ".")
                map = Map.put(map, next_free_pos, "O")
                {Map.put(map, new_pos, "@"), new_pos}
            :no_push -> {map, {y, x}}
        end
    end

    @spec solve(String.t()) :: nil
    def solve(filename) do
        {map, pos, steps} = parse(filename)

        {took, result} = :timer.tc(fn -> part1(map, pos, steps) end)
        IO.puts("Part1: #{result}, took: #{took / 1_000_000} seconds")

        # {took, result} = :timer.tc(fn -> part2(robots, size) end)
        # IO.puts("Part2: #{result}, took: #{took / 1_000_000} seconds")
    end
end

Day15.solve("./input.txt")
