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

    defp move(map, guard_pos, dir_index, visited) do
        next_pos = add(guard_pos, dir_index)
        char = Map.get(map, next_pos, "*")
        cond do
            char == "*" -> visited
            char == "." or char == "^" ->
                move(map, next_pos, dir_index, MapSet.put(visited, {next_pos}))
            char == "#" -> move(map, guard_pos, rem(dir_index + 1, 4), visited)
        end
    end

    defp next_dir_index(dir_index) do
        rem(dir_index + 1, 4)
    end

    defp all_positions(map, pos, dir_index, visited) do
        next_pos = add(pos, dir_index)
        case Map.get(map, next_pos) do
            "#" -> Enum.reverse(visited)
            nil -> Enum.reverse(visited)
            _ -> all_positions(map, next_pos, dir_index, [next_pos | visited])
        end
    end

    def obstacle(map, guard_pos, dir_index, visited_dir_map, count) do
        IO.inspect(guard_pos)
        positions = all_positions(map, guard_pos, next_dir_index(dir_index), [])
        found = Enum.reduce_while(positions, false, fn pos, _acc ->
            dirs = Map.get(visited_dir_map, pos, [])
            if Enum.any?(Enum.map(dirs, fn dir -> dir == next_dir_index(dir_index) end)) do
                {:halt, true}
            else
                {:cont, false}
            end
        end)

        next_pos = add(guard_pos, dir_index)

        if found do
            IO.puts("Place obstacle at (#{elem(next_pos, 0)}, #{elem(next_pos, 1)})")
            obstacle(map, next_pos, dir_index, visited_dir_map, count+1)
        else
            char = Map.get(map, next_pos, "*")
            cond do
                char == "*" -> count
                char == "^" or char == "." ->
                    new_visited_dir_map = Map.update(
                        visited_dir_map, next_pos, [dir_index], fn existing_value ->
                            [dir_index | existing_value]
                    end)
                    obstacle(map, next_pos, dir_index, new_visited_dir_map, count)
                char == "#" ->
                    obstacle(map, guard_pos, next_dir_index(dir_index), visited_dir_map, count)
            end
        end
    end

    def part1(map, guard_starting_pos) do
        move(map, guard_starting_pos, 0, MapSet.new())
        |> MapSet.size()
    end

    def part2(map, guard_starting_pos) do
        obstacle(map, guard_starting_pos, 0, %{guard_starting_pos => [0]}, 0)
    end

    @spec solve(String.t()) :: nil
    def solve(filename) do
        {_, content} = File.read(filename)
        {map, guard_starting_pos} = parse(content)

        {took, result} = :timer.tc(fn -> part1(map, guard_starting_pos) end)
        IO.puts("Part1: #{result}, took: #{took / 1_000_000} seconds")

        {took, result} = :timer.tc(fn -> part2(map, guard_starting_pos) end)
        IO.puts("Part2: #{result}, took: #{took / 1_000_000} seconds")
    end
end

Day6.solve("./test.txt")

# {_, content} = File.read("./test.txt")
# {map, starting_pos} = Day5.parse(content)
#
# visited = %{
#     {6, 4} => [0],
#     {5, 4} => [0],
#     {4, 4} => [0],
#     {3, 4} => [0],
#     {2, 4} => [0],
#     {1, 4} => [0],
#     {1, 5} => [1],
#     {1, 6} => [1],
#     {1, 7} => [1],
#     {1, 8} => [1],
#     {2, 8} => [2],
#     {3, 8} => [2],
#     {4, 8} => [2],
#     {5, 8} => [2],
#     {6, 8} => [2],
#     {6, 7} => [3],
#     {6, 6} => [3],
#     {6, 5} => [3],
#     {6, 4} => [3],
# }
#
# IO.inspect(Day5.obstacle(map, {6, 4}, 3, visited, 0))

# Day5.solve("./test.txt")

    # defp obstacle(map, guard_pos, dir_index, visited_dir_map, count) do
    #     next_pos = add(guard_pos, dir_index)
    #     # IO.inspect(visited_dir_map)
    #
    #     last_dirs = Map.get(visited_dir_map, next_pos, [-1])
    #     valid = Enum.reduce_while(last_dirs, false, fn dir, _acc ->
    #         if next_dir_index(dir_index) == dir do
    #             {:halt, true}
    #         else
    #             {:cont, false}
    #         end
    #     end)
    #
    #     if valid do
    #         IO.inspect(add(next_pos, dir_index))
    #         obstacle(map, next_pos, dir_index, visited_dir_map, count+1)
    #     else
    #         char = Map.get(map, next_pos, "*")
    #         cond do
    #             char == "*" -> count
    #             char == "." or char == "^" ->
    #                 new_visited_dir_map = Map.update(
    #                     visited_dir_map, next_pos, [dir_index], fn existing_value ->
    #                     [dir_index | existing_value]
    #                 end)
    #                 obstacle(map, next_pos, dir_index, new_visited_dir_map, count)
    #             char == "#" ->
    #                 obstacle(map, guard_pos, next_dir_index(dir_index), visited_dir_map, count)
    #         end
    #     end
    # end
