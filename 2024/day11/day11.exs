defmodule Day11 do
    @spec parse(String.t()) :: [String.t()]
    def parse(filename) do
        File.read!(filename)
        |> String.split()
    end

    @spec part1([String.t()], integer) :: integer
    def part1(nums, n) do
        nums
        |> evolve_n(n)
        |> length()
    end

    @spec evolve([String.t()]) :: [String.t()]
    def evolve(nums) do
        nums
        |> Enum.reduce([], fn num, acc ->
            cond do
                num == "0" -> ["1" | acc]
                rem(String.length(num), 2) == 0 ->
                    r = div(String.length(num), 2)
                    {left, right} = String.split_at(num, r)
                    right = right
                        |> String.to_integer()
                        |> Integer.to_string()
                    [[left, right] | acc]
                true ->
                    new_num = num
                        |> String.to_integer()
                        |> (fn x -> x * 2024 end).()
                        |> Integer.to_string()
                    [new_num | acc]
            end
        end)
        |> Enum.reverse()
        |> List.flatten()
    end

    @spec evolve_n([String.t()], integer) :: [String.t()]
    defp evolve_n(nums, 0), do: nums

    @spec evolve_n([String.t()], integer) :: [String.t()]
    defp evolve_n(nums, n) do
        nums
        |> IO.inspect()
        |> evolve()
        |> evolve_n(n-1)
    end

    @spec solve(String.t()) :: nil
    def solve(filename) do
        nums = parse(filename)

        {took, result} = :timer.tc(fn -> part1(nums, 25) end)
        IO.puts("Part1: #{result}, took: #{took / 1_000_000} seconds")

        {took, result} = :timer.tc(fn -> part1(nums, 75) end)
        IO.puts("Part2: #{result}, took: #{took / 1_000_000} seconds")
    end

end

Day11.part1(["0"], 75)
