defmodule Day18 do
    @directions [{0, -1}, {1, 0}, {0, 1}, {-1, 0}]
    @test_count 12
    @input_count 1024

    @type pos() :: {integer(), integer()}

    @spec parse(String.t()) :: [pos()]
    def parse(filename) do
        filename
        |> File.stream!()
        |> Stream.map(fn line ->
            line
            |> String.trim()
            |> String.split(",")
            |> Enum.map(&String.to_integer/1)
            |> List.to_tuple()
        end)
    end

    @spec part1([pos()], integer(), boolean()) :: integer()
    def part1(obstacles, count, is_input) do
        obstacles = obstacles
            |> Enum.take(count)
            |> MapSet.new()

        start = {0, 0}
        goal = if is_input, do: {70, 70}, else: {6, 6}

        queue = [{start, 0}]
        visited = MapSet.new([start])

        bfs(queue, visited, goal, obstacles, is_input)
    end

    @spec bfs([{pos(), integer()}], MapSet.t(pos()), pos(), MapSet.t(pos()), boolean()) :: integer()
    def bfs([{current, steps} | rest], visited, goal, obstacles, is_input) do
        if current == goal do
            steps
        else
            next_positions = for {dx, dy} <- @directions do
                {x, y} = current
                {x + dx, y + dy}
            end

            valid_positions = next_positions
                |> Enum.filter(fn {x, y} -> valid_position?({x, y}, obstacles, visited, is_input) end)

            new_visited = Enum.reduce(valid_positions, visited, &MapSet.put(&2, &1))
            new_queue = rest ++ Enum.map(valid_positions, &{&1, steps + 1})

            bfs(new_queue, new_visited, goal, obstacles, is_input)
        end
    end

    defp valid_position?({x, y}, obstacles, visited, is_input) do
        bounds = if is_input, do: 70, else: 6
        x >= 0 and x <= bounds and y >= 0 and y <= bounds and
            not MapSet.member?(obstacles, {x, y}) and
            not MapSet.member?(visited, {x, y})
    end

    @spec show([pos()], MapSet.t(pos())) :: nil
    defp show(obstacles, path, hl \\ {-1, -1}) do
        obstacles = MapSet.new(obstacles)
        for y <- 0..70 do
            for x <- 0..70 do
                cond do
                    {x, y} == hl -> IO.ANSI.green() <> "H" <> IO.ANSI.reset()
                    MapSet.member?(obstacles, {x, y}) -> "#"
                    MapSet.member?(path, {x, y}) -> IO.ANSI.blue() <> "O" <> IO.ANSI.reset()
                    true -> "."
                end
                |> IO.write()
            end
            IO.puts("")
        end
        IO.puts("")
    end

    def part2(obstacles, count, is_input) do
        {first_list, second_list} = obstacles |> Enum.split(count)
        first = MapSet.new(first_list)

        start = {0, 0}
        goal = if is_input, do: {70, 70}, else: {6, 6}
        queue = [{start, [start]}]
        visited = MapSet.new([start])

        path = find_path(queue, visited, goal, first, is_input)

        second_list
            |> Enum.reduce_while({first, path}, fn obstacle, {all_obstacles, current_path} ->
                all_obstacles = MapSet.put(all_obstacles, obstacle)

                if not MapSet.member?(current_path, obstacle) do
                    {:cont, {all_obstacles, current_path}}
                else
                    case find_path(queue, visited, goal, all_obstacles, is_input) do
                        :no_path -> {:halt, obstacle}
                        new_path -> {:cont, {all_obstacles, new_path}}
                    end
                end
            end)
    end

    defp find_path([], _, _, _, _), do: :no_path
    defp find_path([{current, path} | rest], visited, goal, obstacles, is_input) do
        if current == goal do
            MapSet.new(path)
        else
            next_positions = for {dx, dy} <- @directions do
                {x, y} = current
                {x + dx, y + dy}
            end

            valid_positions = next_positions
                |> Enum.filter(fn {x, y} -> valid_position?({x, y}, obstacles, visited, is_input) end)


            new_visited = Enum.reduce(valid_positions, visited, &MapSet.put(&2, &1))
            new_queue = rest ++ Enum.map(valid_positions, &{&1, path ++ [&1]})

            find_path(new_queue, new_visited, goal, obstacles, is_input)
        end
    end

    @spec solve(String.t()) :: nil
    def solve(filename) do
        {count, is_input} = if filename == "./test.txt" do
            {@test_count, false}
        else
            {@input_count, true}
        end

        obstacles = parse(filename)

        {took, result} = :timer.tc(fn -> part1(obstacles, count, is_input) end)
        IO.puts("Part1: #{result}, took: #{took / 1_000_000} seconds")

        {took, result} = :timer.tc(fn -> part2(obstacles, count, is_input) end)
        IO.puts("Part2: \"#{elem(result, 0)},#{elem(result, 1)}\", took: #{took / 1_000_000} seconds")
    end
end

Day18.solve("./input.txt")
