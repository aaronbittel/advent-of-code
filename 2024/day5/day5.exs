defmodule Day5 do

    @spec parse(String.t()) :: {%{integer => [integer]}, [[integer]]}
    defp parse(filename) do
        content = File.read!(filename)
        [rules, updates] = String.split(content, "\n\n", trim: true)
        {parse_rules(rules), parse_updates(updates)}
    end

    @spec parse_rules(String.t()) :: %{integer => [integer]}
    def parse_rules(rules) do
        String.split(rules, "\n", trim: true)
        |> Enum.map(fn pair ->
            pair
            |> String.split("|", trim: true)
            |> Enum.map(&String.to_integer/1)
        end)
        |> Enum.reduce(%{}, fn [left, right], acc ->
            val = case Map.get(acc, left, []) do
                [] -> [right]
                list -> [right | list]
            end
            Map.put(acc, left, val)
        end)
    end

    @spec parse_updates(String.t()) :: [[integer]]
    def parse_updates(updates) do
        String.split(updates, "\n", trim: true)
        |> Enum.map(fn row ->
            row
            |> String.split(",", trim: true)
            |> Enum.map(&String.to_integer/1)
        end)
    end

    @spec is_correct_order?([integer], %{integer => [integer]}) :: boolean
    defp is_correct_order?(update, rules) do
        res = Enum.reduce_while(update, {true, update}, fn num, {_status, acc} ->
            [_ | tail] = acc
            if all_elements_in?(tail, Map.get(rules, num, [])) do
                {:cont, {true, tail}}
            else
                {:halt, {false, tail}}
            end
        end)

        {out, _} = res
        out
    end

    defp all_elements_in?(update, rule) do
        Enum.all?(update, fn el -> Enum.member?(rule, el) end)
    end

    @spec get_middle_value([integer]) :: integer
    defp get_middle_value(list) do
        middle_index = div(length(list), 2)
        Enum.at(list, middle_index)
    end

    @spec correct_order([integer], %{integer => [integer]}) :: [integer]
    def correct_order(update, rules) do
        Enum.reduce(update, [], fn num, acc ->
            map = Map.get(rules, num, [])
            list = List.delete(update, num)
            if all_elements_in?(list, map) do
                [num | correct_order(list, rules)]
            else
                acc
            end
        end)
    end

    @spec part1(String.t()) :: integer
    def part1(filename) do
        {rules, updates} = parse(filename)
        Enum.filter(updates, fn update -> is_correct_order?(update, rules) end)
        |> Enum.map(&get_middle_value/1)
        |> Enum.sum()
    end

    @spec part2(String.t()) :: integer
    def part2(filename) do
        {rules, updates} = parse(filename)
        Enum.reject(updates, fn update -> is_correct_order?(update, rules) end)
        |> Enum.map(fn update -> correct_order(update, rules) end)
        |> Enum.map(&get_middle_value/1)
        |> Enum.sum()
    end
end

{took, result} = :timer.tc(fn -> Day5.part1("./input.txt") end)
IO.puts("Part1: #{result}, took: #{took / 1_000_000} seconds")

{took, result} = :timer.tc(fn -> Day5.part2("./input.txt") end)
IO.puts("Part2: #{result}, took: #{took / 1_000_000} seconds")
