defmodule RC do
    def comb(0, _), do: [[]]
    def comb(m, list) do
        for h <- list, t <- comb(m - 1, list), do: [h | t]
    end
end

defmodule Day7 do

    def part1(filename) do
        ops = [:sum, :mult]

        {_, content} = File.read(filename)
        String.split(content, "\n", trim: :true)
        |> Enum.map(fn line ->
            [res, eq] = line
            |> String.split(": ", trim: :true)

            res = String.to_integer(res)
            eq = String.split(eq, " ")
            |> Enum.map(&String.to_integer/1)

            combinations = RC.comb(length(eq)-1, ops)

            if Enum.reduce_while(combinations, false, fn comb, _acc ->
                result = apply_operations(eq, comb)
                if result == res do
                    {:halt, true}
                else
                    {:cont, false}
                end
            end) do
                res
            end
        end)
        |> Enum.reject(&is_nil/1)
        |> Enum.sum()
    end

    defp apply_operations([num], []), do: num
    defp apply_operations([num1, num2 | rest_nums], [op | rest_ops]) do
        result = apply_op(num1, num2, op)

        apply_operations([result | rest_nums], rest_ops)
    end

    defp apply_op(num1, num2, :mult), do: num1 * num2
    defp apply_op(num1, num2, :sum), do: num1 + num2
    defp apply_op(_, _, op), do: raise "Unsupported operation: #{inspect(op)}"

    @spec solve(String.t()) :: nil
    def solve(filename) do
        {took, result} = :timer.tc(fn -> part1(filename) end)
        IO.puts("Part1: #{result}, took: #{took / 1_000_000} seconds")

        # {took, result} = :timer.tc(fn -> part2(map, guard_starting_pos) end)
        # IO.puts("Part2: #{result}, took: #{took / 1_000_000} seconds")
    end

end

Day7.solve("./input.txt")

