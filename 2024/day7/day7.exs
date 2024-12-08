defmodule Day7 do
    @type eq() :: {integer, [integer]}

    @spec parse(String.t()) :: [eq]
    def parse(filename) do
        {_, content} = File.read(filename)
        content
        |> String.split("\n", trim: true)
        |> Enum.reduce([], fn eq, acc ->
            [res, nums] = eq
            |> String.split(": ", trim: true)

            res = String.to_integer(res)
            nums = nums
                |> String.split(" ", trim: true)
                |> Enum.map(&String.to_integer/1)

            [{res, nums} | acc]
        end)
    end

    @spec part1([eq]) :: integer
    def part1(eqs) do
        eqs
        |> Enum.filter(fn eq -> is_valid?(eq, [:sum, :mult]) end)
        |> Enum.map(fn eq -> elem(eq, 0) end)
        |> Enum.sum()
    end

    @spec is_valid?(eq, [atom]) :: boolean
    defp is_valid?({res, nums}, ops) do
        Comb.comb(length(nums)-1, ops)
        |> Enum.reduce_while(false, fn ops, _acc ->
            if Op.do_apply(nums, ops) == res do
                {:halt, true}
            else
                {:cont, false}
            end
        end)
    end

    @spec part2([eq]) :: integer
    def part2(eqs) do
        eqs
        |> Enum.filter(fn eq -> is_valid?(eq, [:sum, :mult, :con]) end)
        |> Enum.map(fn eq -> elem(eq, 0) end)
        |> Enum.sum()
    end

    @spec solve(String.t()) :: nil
    def solve(filename) do
        eqs = parse(filename)
        {took, result} = :timer.tc(fn -> part1(eqs) end)
        IO.puts("Part1: #{result}, took: #{took / 1_000_000} seconds")

        {took, result} = :timer.tc(fn -> part2(eqs) end)
        IO.puts("Part2: #{result}, took: #{took / 1_000_000} seconds")
    end
end

defmodule Op do
    @spec do_apply([integer], [atom]) :: integer
    def do_apply([n1, n2 | ntail], [op | otail]) do
        cond do
            op == :sum -> do_apply([n1 + n2 | ntail], otail)
            op == :mult -> do_apply([n1 * n2 | ntail], otail)
            op == :con ->
                r = String.to_integer(Integer.to_string(n1) <> Integer.to_string(n2))
                do_apply([r | ntail], otail)
        end
    end

    def do_apply([res], []), do: res
end


defmodule Comb do
    def comb(0, _), do: [[]]
    def comb(m, list) do
        for h <- list, t <- comb(m - 1, list), do: [h | t]
    end
end

Day7.solve("./input.txt")
