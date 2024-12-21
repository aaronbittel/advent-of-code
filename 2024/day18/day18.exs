defmodule Day18 do
    @directions [{0, -1}, {1, 0}, {0, 1}, {-1, 0}]
    @test_count 12
    @input_count 1024

    @type dir() :: {integer(), integer()}

    @spec parse(String.t(), integer()) :: [dir()]
    def parse(filename, count) do
        filename
        |> File.stream!()
        |> Stream.take(count)
        |> Stream.map(fn line ->
            line
            |> String.trim()
            |> String.split(",")
            |> Enum.map(&String.to_integer/1)
            |> List.to_tuple()
        end)
        |> MapSet.new()
    end

    def part1(obstacles, is_input) do
        start = {0, 0}
        goal = if is_input, do: {70, 70}, else: {6, 6}

        IO.puts("Calling with goal = (#{elem(goal, 0)}, #{elem(goal, 1)})")
        IO.puts("#obstacles: #{Enum.count(obstacles)}")

        show(obstacles)

        queue = [{start, 0}]
        visited = MapSet.new([start])

        bfs(queue, visited, goal, obstacles, is_input)
    end

    def bfs([{current, steps} | rest], visited, goal, obstacles, is_input) do
        # IO.puts("#{elem(current, 0)}, #{elem(current, 1)}: #{steps}")
        if current == goal do
            steps
        else
            next_positions = for {dx, dy} <- @directions do
                {x, y} = current
                {x + dx, y + dy}
            end

            valid_positions = next_positions
                |> Enum.filter(fn {x, y} -> valid_position?({x, y}, obstacles, visited, is_input) end)

            # valid_positions
            # |> Enum.each(fn pos -> IO.puts("X=#{elem(pos, 0)}, Y=#{elem(pos, 1)}") end)

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

    @spec show(MapSet.t()) :: nil
    defp show(obstacles) do
        for y <- 0..6 do
            for x <- 0..6 do
                char = if MapSet.member?(obstacles, {x, y}), do: "#", else: "."
                IO.write(char)
            end
            IO.puts("")
        end
    end

    @spec solve(String.t()) :: nil
    def solve(filename) do
        {count, is_input} = if filename == "./test.txt" do
            {@test_count, false}
        else
            {@input_count, true}
        end

        obstacles = parse(filename, count)

        {took, result} = :timer.tc(fn -> part1(obstacles, is_input) end)
        IO.puts("Part1: #{result}, took: #{took / 1_000_000} seconds")

        # {took, result} = :timer.tc(fn -> part2(robots, size) end)
        # IO.puts("Part2: #{result}, took: #{took / 1_000_000} seconds")
    end
end

Day18.solve("./input.txt")
