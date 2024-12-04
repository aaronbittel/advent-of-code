defmodule Day3 do
    defp mul_to_res(mul) do
        mul
        |> String.replace("mul(", "")
        |> String.replace(")", "")
        |> String.split(",")
        |> Enum.map(fn num -> String.to_integer(num) end)
        |> (fn [a, b] -> a * b end).()
    end

    @spec part1(String.t()) :: integer
    def part1(filepath) do
        {_, content} = File.read(filepath)

        pattern = ~r/mul\([0-9]{1,3},[0-9]{1,3}\)/
        Regex.scan(pattern, content)
        |> Enum.map(fn [mul] -> mul_to_res(mul) end)
        |> Enum.sum()
    end

    @spec part2(String.t()) :: integer
    def part2(filepath) do
        {_, content} = File.read(filepath)

        pattern = ~r/(mul\([0-9]{1,3},[0-9]{1,3}\)|don't\(\)|do\(\))/

        my_res = Regex.scan(pattern, content)
        |> Enum.flat_map(fn [match, _] -> [match] end)
        |> Enum.reduce({0, true}, fn item, {result, state} ->
            cond do
                state && String.starts_with?(item, "mul") -> {mul_to_res(item) + result, state}
                item == "don't()" -> {result, false}
                item == "do()" -> {result, true}
                # mult but state == false
                true -> {result, state}
            end
        end)
        elem(my_res, 0)
    end
end

{took, result} = :timer.tc(fn -> Day3.part1("./input.txt") end)
IO.puts("Part1: #{result}, took: #{took} ms")

{took, result} = :timer.tc(fn -> Day3.part2("./input.txt") end)
IO.puts("Part2: #{result}, took: #{took} ms")
