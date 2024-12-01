input = "3   4
4   3
2   5
1   3
3   9
3   3"

defmodule Day1 do
    @spec part1(String.t()) :: integer
    def part1(text) do
        lines = text
            |> String.split("\n")
            |> Enum.slice(0..-2)
        left = Enum.map(lines, fn x -> String.split(x) |> List.first() |>
            String.to_integer() end) |> Enum.sort()
        right = Enum.map(lines, fn x -> String.split(x) |> List.last() |>
            String.to_integer() end) |> Enum.sort()

        Enum.zip_reduce(left, right, 0, fn l, r, acc ->
            acc + abs(l-r)
        end)
    end

    @spec part2(String.t()) :: integer
    def part2(_text) do
        0
    end
end

case File.read("day1.txt") do
    {:ok, content} ->
        IO.puts("Part1: #{Day1.part1(content)}")
        IO.puts("Part2: #{Day1.part2(content)}")
    {:error, reason} -> IO.puts("Failed to read file: #{reason}")
end
