defmodule Day9 do
    alias EtsDeque, as: Deque

    @type file :: [type: :file, id: integer, size: integer]
    @type freespace :: [type: :free, size: integer]
    @type filesystem :: Deque.t(file | freespace)

    @spec parse(String.t()) :: filesystem
    defp parse(content) do
        nums = content
            |> String.graphemes()
            |> Enum.slice(0..-2//1)

        categorize(nums, 0, :file, Deque.new(:infinity))
    end

    @spec categorize([String.t()], integer, atom, filesystem) :: filesystem
    defp categorize([head | tail], id, type, result) do
        size = String.to_integer(head)
        cond do
            type == :file ->
                result = Deque.push_tail!(result, [type: :file, id: id, size: size])
                categorize(tail, id+1, :free, result)
            type == :free and size > 0 ->
                result = Deque.push_tail!(result, [type: :free, size: size])
                categorize(tail, id, :file, result)
            type == :free and size == 0 ->
                categorize(tail, id, :file, result)
        end
    end

    @spec categorize([String.t()], integer, atom, filesystem) :: filesystem
    defp categorize([], _id, _type, result), do: result

    def part1(filesystem) do
        part1_helper(filesystem, [])
        |> Enum.reverse()
        |> Enum.reduce([], fn {id, size}, acc ->
            acc ++ List.duplicate(id, size)
        end)
        |> Enum.with_index()
        |> Enum.reduce(0, fn {x, y}, acc ->
            acc + (x * y)
        end)
    end

    defp part1_helper(filesystem, result) do
        if Deque.length(filesystem) == 0 do
            result
        else
            case Deque.peek_head!(filesystem) do
                [type: :file, id: id, size: size] ->
                    {_, filesystem} = Deque.pop_head!(filesystem)
                    result = [{id, size} | result]
                    part1_helper(filesystem, result)
                _ ->
                    case Deque.peek_tail!(filesystem) do
                        [type: :free, size: _] ->
                            {_, filesystem} = Deque.pop_tail!(filesystem)
                            part1_helper(filesystem, result)
                        _ ->
                            # free space at the front && file at the back
                            {[type: :free, size: freesize], filesystem} = Deque.pop_head!(filesystem)
                            {[type: :file, id: id, size: filesize], filesystem} = Deque.pop_tail!(filesystem)

                            cond do
                                freesize > filesize ->
                                    result = [{id, filesize} | result]
                                    filesystem = Deque.push_head!(filesystem, [type: :free, size: freesize-filesize])
                                    part1_helper(filesystem, result)

                                freesize == filesize ->
                                    result = [{id, filesize} | result]
                                    part1_helper(filesystem, result)

                                freesize < filesize ->
                                    result = [{id, freesize} | result]
                                    filesystem = Deque.push_tail!(filesystem, [type: :file, id: id, size: filesize-freesize])
                                    part1_helper(filesystem, result)
                            end
                    end
            end
        end
    end

    def part2(filesystem) do
        {files, frees} = parse_part2(filesystem)
        {file, files} = Deque.pop_head!(files)
        {result, unsortable} = part2_helper(files, frees, [file], Deque.new())

        IO.inspect(result)
        IO.inspect(deque_to_list(unsortable, []))

        0
    end

    defp part2_helper(files, frees, result, unsortable) do
        case Deque.pop_tail(files) do
            {:ok, file, files} ->
                case can_fit_size?(file[:size], frees) do
                    {true, frees} ->
                        result = [file | result]
                        part2_helper(files, frees, result, unsortable)
                    false ->
                        Deque.push_tail!(unsortable, file)
                        part2_helper(files, frees, result, unsortable)
                end
            :error -> {result, unsortable}
        end
    end

    defp can_fit_size?(size, frees) do
        pos = frees
        |> Enum.with_index()
        |> Enum.filter(fn {freesize, _idx} -> freesize >= size end)

        case pos do
            [] -> false
            _ ->
                {freesize, idx} = hd(pos)
                frees = List.update_at(frees, idx, fn _elem -> freesize-size end)
                {true, frees}
        end
    end

    defp parse_part2(filesystem) do
        list = deque_to_list(filesystem, [])

        files = list
            |> Enum.filter(fn item -> item[:type] == :file end)
            |> Enum.reduce(Deque.new(), fn file, acc ->
                Deque.push_tail!(acc, file)
            end)

        frees = list
            |> Enum.filter(fn item -> item[:type] == :free end)
            |> Enum.map(fn [type: _, size: size] -> size end)
            |> IO.inspect()

        {files, frees}
    end

    defp deque_to_list(deque, result) do
        case Deque.pop_tail(deque) do
            {:ok, item, deque} ->
                result = [item | result]
                deque_to_list(deque, result)
            :error ->
                result
        end
    end

    @spec solve() :: nil
    def solve() do
        {_, content} = File.read("./input.txt")
        filesystem = parse(content)

        {took, result} = :timer.tc(fn -> part1(filesystem) end)
        IO.puts("Part1: #{result}, took: #{took / 1_000_000} seconds")

        # {took, result} = :timer.tc(fn -> part2(filesystem) end)
        # IO.puts("Part2: #{result}, took: #{took / 1_000_000} seconds")
    end
end
