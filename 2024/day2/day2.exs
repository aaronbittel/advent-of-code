defmodule Day2 do

    @spec is_safe([integer], atom | nil) :: boolean
    defp is_safe(numbers, dir \\ nil) do
        dir = if dir == nil do
            [first, second | _] = numbers
            if first > second, do: :desc, else: :asc
        else
            dir
        end

        case numbers do
            [] -> true
            [_] -> true
            [a, b | tail] when dir == :desc and a - b in 1..3 ->
                is_safe([b | tail], :desc)
            [a, b | tail] when dir == :asc and b - a in 1..3 ->
                is_safe([b | tail], :asc)
            _ ->
                false
        end
    end


    @spec exclude_indices([integer]) :: [[interger]]
    defp exclude_indices(list) do
        for i <- 0..(length(list) - 1), do: List.delete_at(list, i)
    end

    @spec part1(String.t()) :: integer
    def part1(filename) do
        File.stream!(filename)
        |> Stream.map(fn line ->
            String.split(line)
            |> Enum.map(fn num -> String.to_integer(num) end)
        end)
        |> Enum.filter(fn nums -> is_safe(nums) end)
        |> Enum.count()
    end

    @spec part2(String.t()) :: integer
    def part2(filename) do
        File.stream!(filename)
        |> Stream.map(fn line ->
            String.split(line)
            |> Enum.map(fn num -> String.to_integer(num) end)
        end)
        |> Enum.filter(fn nums ->
            Enum.any?(exclude_indices(nums), fn nums -> is_safe(nums) end)
        end)
        |> Enum.count()
    end

end


{took, result} = :timer.tc(fn -> Day2.part1("./input.txt") end)
IO.puts("Part1: #{result}, took: #{took} ms")
#
{took, result} = :timer.tc(fn -> Day2.part2("./input.txt") end)
IO.puts("Part2: #{result}, took: #{took} ms")
