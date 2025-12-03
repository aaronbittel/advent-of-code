defmodule Day1 do
  @spec parse(String.t()) :: {[integer()], [integer()]}
  def parse(filename) do
    {left, right} =
      filename
      |> File.read!()
      |> String.split("\n", trim: true)
      |> Enum.reduce({[], []}, fn pair, {left, right} ->
        [l, r] =
          pair
          |> String.split()
          |> Enum.map(&String.to_integer/1)

        {[l | left], [r | right]}
      end)

    left = left |> Enum.sort()
    right = right |> Enum.sort()
    {left, right}
  end

  @spec solve([integer()], [integer()]) :: integer()
  def solve(left, right) do
    Enum.zip_reduce(left, right, 0, fn l, r, acc ->
      acc + abs(l - r)
    end)
  end
end

{left, right} = Day1.parse("day1.txt")
IO.inspect(Day1.solve(left, right))
