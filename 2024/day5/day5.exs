defmodule Day5 do

    @spec parse(String.t()) :: {%{integer => [integer]}, [[integer]]}
    def parse(filename) do
        content = File.read!(filename)
        [rules, updates] = String.split(content, "\n\n", trim: true)
        {parse_rules(rules), parse_updates(updates)}
    end

    @spec parse_rules(String.t()) :: %{integer => [integer]}
    defp parse_rules(rules) do
        String.split(rules, "\n", trim: true)
        |> Enum.reduce(%{}, fn pair, acc ->
            [left, right] =
                String.split(pair, "|", trim: true)
                |> Enum.map(&String.to_integer/1)
            Map.update(acc, left, [right], fn existing_values -> [right | existing_values] end)
        end)
    end

    @spec parse_updates(String.t()) :: [[integer]]
    defp parse_updates(updates) do
        String.split(updates, "\n", trim: true)
        |> Enum.map(fn row ->
            row
            |> String.split(",", trim: true)
            |> Enum.map(&String.to_integer/1)
        end)
    end

    @spec is_correct_order?(%{integer => [integer]}, [integer]) :: boolean
    defp is_correct_order?(rules, update) do
         Enum.reduce_while(update, {true, update}, fn num, {_status, acc} ->
            tail = tl(acc)
            if all_elements_in?(tail, Map.get(rules, num, [])) do
                {:cont, {true, tail}}
            else
                {:halt, {false, tail}}
            end
        end)
        |> elem(0)

    end

    @spec all_elements_in?([integer], [integer]) :: boolean
    defp all_elements_in?(rule, update) do
        Enum.all?(rule, fn el -> Enum.member?(update, el) end)
    end

    @spec get_middle_value([integer]) :: integer
    defp get_middle_value(list) do
        middle_index = div(length(list), 2)
        Enum.at(list, middle_index)
    end

    @spec correct_order(%{integer => [integer]}, [integer]) :: [integer]
    def correct_order(rules, update) do
        Enum.reduce_while(update, [], fn num, acc ->
            map = Map.get(rules, num, [])
            list = List.delete(update, num)
            if all_elements_in?(list, map) do
                {:halt, [num | correct_order(rules, list)]}
            else
                {:cont, acc}
            end
        end)
    end

    @spec part1(%{integer => [integer]}, [integer]) :: integer
    def part1(rules, updates) do
        Enum.filter(updates, fn update -> is_correct_order?(rules, update) end)
        |> Enum.map(&get_middle_value/1)
        |> Enum.sum()
    end

    @spec part2(%{integer => [integer]}, [integer]) :: integer
    def part2(rules, updates) do
        Enum.reject(updates, fn update -> is_correct_order?(rules, update) end)
        |> Enum.map(fn update -> correct_order(rules, update) end)
        |> Enum.map(&get_middle_value/1)
        |> Enum.sum()
    end

    @spec solve(String.t()) :: nil
    def solve(filename) do
        {rules, updates} = parse(filename)
        {took, result} = :timer.tc(fn -> part1(rules, updates) end)
        IO.puts("Part1: #{result}, took: #{took / 1_000_000} seconds")

        {took, result} = :timer.tc(fn -> part2(rules, updates) end)
        IO.puts("Part2: #{result}, took: #{took / 1_000_000} seconds")
    end
end

Day5.solve("./input.txt")
