defmodule Day14 do
    @test_size {11, 7}
    @default_size {101, 103}
    @solution_part2 6493

    @type dir() :: {integer(), integer()}
    @type robot() :: {dir(), dir()}

    @spec parse(String.t()) :: [robot()]
    def parse(filename) do
        filename
        |> File.stream!()
        |> Stream.map(&String.trim/1)
        |> Enum.map(&parse_robot/1)
    end

    @spec parse_robot(String.t()) :: robot()
    defp parse_robot(robot) do
        robot
        |> String.split(" ", trim: true)
        |> Enum.map(&parse_dir/1)
        |> List.to_tuple()
    end

    @spec parse_dir(String.t()) :: dir()
    defp parse_dir(dir) do
        dir
        |> String.slice(2..-1//1)
        |> String.split(",")
        |> Enum.map(&String.to_integer/1)
        |> List.to_tuple()
    end

    @spec quadrants([robot()], dir()) :: %{integer() => integer()}
    defp quadrants(robots, {width, height}) do
        border_x  = div(width,2)
        border_y = div(height,2)

        robots
        |> Enum.reduce(%{0 => 0, 1 => 0, 2 => 0, 3 => 0}, fn {{x, y}, _}, acc ->
            if x == border_x or y == border_y do
                acc
            else
                quadrant = cond do
                    x < border_x and y < border_y -> 0
                    x > border_x and y < border_y -> 1
                    x < border_x and y > border_y -> 2
                    x > border_x and y > border_y -> 3
                end

                Map.update!(acc, quadrant, &(&1 + 1))
            end
        end)
    end

    @spec part1([robot()], dir()) :: integer()
    def part1(robots, size) do
        move(robots, size, 100)
        |> quadrants(size)
        |> Map.values()
        |> Enum.reduce(1, fn v, acc -> v * acc end)
    end

    @spec move([robot()], dir(), integer) :: [robot()]
    defp move(robots, _, 0), do: robots
    defp move(robots, {width, height}, count) do
        robots = robots
            |> Enum.map(fn {{r_x, r_y}, {v_x, v_y}} ->
                {new_x, new_y} = {r_x + v_x, r_y + v_y}

                new_x = rem(new_x, width)
                new_y = rem(new_y, height)

                new_x = if new_x < 0, do: new_x + width, else: new_x
                new_y = if new_y < 0, do: new_y + height, else: new_y

                {{new_x, new_y}, {v_x, v_y}}
            end)

        move(robots, {width, height}, count - 1)
    end

    def part2(robots, size) do
        robots = move(robots, size, @solution_part2)
        print_robots(robots, size)
        nil
    end

    defp loop(robots, size, count \\ 0) do
        print_robots(robots, size)
        IO.puts("Loop count: #{count}")
        IO.gets("Press Enter to move robots or Ctrl+C to exit...")
        robots = move(robots, size, 1)
        loop(robots, size, count + 1)
    end

    def print_robots(robots, {width, height}) do
        robot_set = robots
        |> Enum.map(fn {{x, y}, _} -> {x, y} end)
        |> MapSet.new()

        IO.puts("")
        for y <- 0..(height - 1) do
            for x <- 0..(width - 1) do
                if MapSet.member?(robot_set, {x, y}) do
                    IO.write("*")
                else
                    IO.write(" ")
                end
            end
            IO.puts("")
        end
    end

    @spec solve(String.t()) :: nil
    def solve(filename) do
        robots = parse(filename)
        size = if filename == "./test.txt", do: @test_size, else: @default_size

        {took, _} = :timer.tc(fn -> part2(robots, size) end)
        IO.puts("Part2 took: #{took / 1_000_000} seconds")

        {took, result} = :timer.tc(fn -> part1(robots, size) end)
        IO.puts("Part1: #{result}, took: #{took / 1_000_000} seconds")

    end

end

Day14.solve("./input.txt")
