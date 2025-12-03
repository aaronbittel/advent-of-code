defmodule Day6 do
    @directions {
        {-1, 0}, # up
        {0, 1},  # right
        {1, 0},  # down
        {0, -1}, # left
    }

    def parse(content) do
        String.split(content, "\n", trim: :true)
        |> Enum.with_index()
        |> Enum.reduce({%{}, nil}, fn {line, row}, {map, caret_pos} ->
            line
            |> String.graphemes()
            |> Enum.with_index()
            |> Enum.reduce({map, caret_pos}, fn {char, col}, {map_acc, caret_pos_acc} ->
                new_map = Map.put(map_acc, {row, col}, char)
                new_caret_pos = if char == "^", do: {row, col}, else: caret_pos_acc
                {new_map, new_caret_pos}
            end)
        end)
    end

    defp add(pos, dir_index) do
        dir = elem(@directions, dir_index)
        {elem(pos, 0) + elem(dir, 0), elem(pos, 1) + elem(dir, 1)}
    end

    defp move(map, guard_pos, dir_index, visited, count) do
        next_pos = add(guard_pos, dir_index)
        char = Map.get(map, next_pos, "*")
        cond do
            count == 25000 -> count
            char == "*" -> visited
            char == "." or char == "^" ->
                move(map, next_pos, dir_index, MapSet.put(visited, {next_pos}), count+1)
            char == "#" -> move(map, guard_pos, rem(dir_index + 1, 4), visited, count)
        end
    end

    defp next_dir_index(dir_index) do
        rem(dir_index + 1, 4)
    end

    defp move2(map, pos, dir_index, visited) do
        next_pos = add(pos, dir_index)
        dirs = Map.get(visited, next_pos, [])
        if Enum.member?(dirs, dir_index) do
            true
        else
            char = Map.get(map, next_pos, "*")
            cond do
                char == "*" -> false
                char == "." or char == "^" ->
                    visited = Map.update(visited, next_pos, [dir_index], fn existing_value ->
                        [dir_index | existing_value]
                    end)
                    move2(map, next_pos, dir_index, visited)
                char == "#" ->
                    move2(map, pos, next_dir_index(dir_index), visited)
            end
        end
    end

    def obstacle(map, guard_pos, dir_index, obstacle_positions) do
        next_pos = add(guard_pos, dir_index)
        char = Map.get(map, next_pos, "*")
        cond do
            char == "*" -> obstacle_positions
            char == "." or char == "^" ->
                if move2(map, guard_pos, next_dir_index(dir_index), %{
                    guard_pos => [dir_index, next_dir_index(dir_index)]
                }) do
                    obstacle(map, next_pos, dir_index, MapSet.put(obstacle_positions, next_pos))
                else
                    obstacle(map, next_pos, dir_index, obstacle_positions)
                end
            char == "#" ->
                obstacle(map, guard_pos, next_dir_index(dir_index), obstacle_positions)
        end
    end

    def part1(map, guard_starting_pos) do
        move(map, guard_starting_pos, 0, MapSet.new(), -10000)
        |> MapSet.size()
    end

    def part2(map, guard_starting_pos) do
        loop_pos = obstacle(map, guard_starting_pos, 0, MapSet.new())
        no_loop_count =
            Enum.count(loop_pos, fn pos ->
                map = Map.put(map, pos, "#")

                if move(map, guard_starting_pos, 0, MapSet.new(), 0) != 25000 do
                    IO.puts("No Loop: (#{elem(pos, 0)}, #{elem(pos, 1)})")
                    false  # This element is a "No Loop"
                else
                    true # This element is not a "No Loop"
                end
            end)

IO.puts("Total No Loop positions: #{no_loop_count}")
    end

    @spec solve(String.t()) :: nil
    def solve(filename) do
        {_, content} = File.read(filename)
        {map, guard_starting_pos} = parse(content)

        # {took, result} = :timer.tc(fn -> part1(map, guard_starting_pos) end)
        # IO.puts("Part1: #{result}, took: #{took / 1_000_000} seconds")

        {took, result} = :timer.tc(fn -> part2(map, guard_starting_pos) end)
        IO.puts("Part2: #{result}, took: #{took / 1_000_000} seconds")
    end
end

Day6.solve("./input.txt")
