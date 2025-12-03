defmodule Day13 do
    @type dir() :: {integer(), integer()}
    @type press() :: {integer(), integer(), integer()}
    @type game() :: {dir(), dir(), dir()}

    @spec parse_position(String.t(), String.t(), String.t()) :: dir()
    defp parse_position(str, prefix, seperator) do
        str
        |> String.replace_prefix(prefix, "")
        |> String.split(seperator)
        |> Enum.map(&String.to_integer/1)
        |> List.to_tuple()
    end

    @spec parse_dir(String.t(), atom()) :: dir()
    defp parse_dir(str, :a), do: parse_position(str, "Button A: X+", ", Y+")

    @spec parse_dir(String.t(), atom()) :: dir()
    defp parse_dir(str, :b), do: parse_position(str, "Button B: X+", ", Y+")

    @spec parse_dir(String.t(), atom()) :: dir()
    defp parse_dir(str, :price), do: parse_position(str, "Prize: X=", ", Y=")

    @spec parse_game(String.t()) :: game()
    defp parse_game(str) do
        [button_a, button_b, price] = String.split(str, "\n", trim: true)
        {parse_dir(button_a, :a), parse_dir(button_b, :b), parse_dir(price, :price)}
    end

    @spec parse(String.t()) :: [game()]
    def parse(filename) do
        filename
        |> File.read!()
        |> String.split("\n\n", trim: true)
        |> Enum.map(&parse_game/1)
    end

    @spec part1([game()]) :: integer()
    def part1(games) do
        games
        |> Enum.map(fn game ->
            {{a_x, a_y}, {b_x, b_y}, {p_x, p_y}} = game

            count_a = 0
            count_b = min(min(div(p_x, b_x), 100), min(div(p_y, b_y), 100))

            part1_helper({count_a, a_x, a_y}, {count_b, b_x, b_y}, {p_x, p_y})
        end)
        |> Enum.reject(fn res -> res == {-1, -1} end)
        |> Enum.reduce(0, fn {count_a, count_b}, acc ->
            acc + count_a * 3 + count_b
        end)
    end

    @spec part1_helper(press(), press(), dir()) :: dir()
    defp part1_helper({count_a, a_x, a_y}, {count_b, b_x, b_y}, {p_x, p_y}) do
        {res_x, res_y} = calculate({count_a, a_x, a_y}, {count_b, b_x, b_y})

        cond do
            count_a > 100 or count_b < 0 -> {-1, -1}
            res_x == p_x and res_y == p_y -> {count_a, count_b}
            res_x > p_x or res_y > p_y ->
                part1_helper({count_a, a_x, a_y}, {count_b-1, b_x, b_y}, {p_x, p_y})
            res_x < p_x or res_y < p_y ->
                part1_helper({count_a+1, a_x, a_y}, {count_b, b_x, b_y}, {p_x, p_y})
            true -> IO.puts("what happend")
        end
    end

    @spec calculate(press(), press()) :: dir()
    defp calculate({count_a, a_x, a_y}, {count_b, b_x, b_y}) do
        {count_a * a_x + count_b * b_x, count_a * a_y + count_b * b_y}
    end

    @spec solve(String.t()) :: nil
    def solve(filename) do
        games = parse(filename)

        {took, result} = :timer.tc(fn -> part1(games) end)
        IO.puts("Part1: #{result}, took: #{took / 1_000_000} seconds")

        # {took, result} = :timer.tc(fn -> part1(nums, 75) end)
        # IO.puts("Part2: #{result}, took: #{took / 1_000_000} seconds")
    end

end

Day13.solve("./input.txt")
