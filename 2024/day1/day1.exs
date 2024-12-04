defmodule Day1 do

    @spec parse(String.t()) :: {[integer], [integer]}
    def parse(filepath) do
        File.stream!(filepath)
        |> Stream.map(fn line ->
            line
            |> String.trim()
            |> String.split()
            |> Enum.map(fn w -> String.to_integer(w) end)
        end)
            |> Enum.to_list()
            |> (fn list ->
                flatten = List.flatten(list)
                [Enum.slice(flatten, 0..-1//2), Enum.slice(flatten, 1..-1//2)]
            end).()
    end

    @spec part1([integer], [integer]) :: integer
    def part1(left, right) do
        Enum.zip_reduce(left, right, 0, fn l, r, acc ->
            acc + abs(l-r)
        end)
    end

    @spec part2([integer], [integer]) :: integer
    def part2(left, right) do
        right_counter = Enum.frequencies(right)

        Enum.reduce(left, 0, fn num, acc ->
            acc + num * Map.get(right_counter, num, 0)
        end)
    end
end

[left, right] = Day1.parse("day1.txt")
{took, result} = :timer.tc(fn -> Day1.part1(left, right) end)
IO.puts("Part1: #{result}, took: #{took} ms")
{took, result} = :timer.tc(fn -> Day1.part2(left, right) end)
IO.puts("Part2: #{result}, took: #{took} ms")
