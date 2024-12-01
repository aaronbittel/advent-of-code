# Elixir

## Learnings
- lists are implemented as linked lists => only has `hd` (head) method
    - use when the number of values are not known, e.g. `String.split()`
    returns a list because the number of elements depends on the input string
    - ```Elixir
        list = [1, 2, 3]
        list2 = [0] ++ list
        ```
        - `list2` is a new list, `list` still exists unchanged
        - this is fast because only need to go to element `[0]` and set
        next-pointer to elem `[1]`
            - `list ++ [0]` would be slower, because first need to traverse to
            the end of the first list
- tuples are implemented as continuos memory space (array-like)
    - use when the number of values are fixed, e.g. `String.split_at()` returns
    exactly 2 elements
- types are immutable
    - every operation returns a new immutable "reference"
    - values are not copied

## Questions
- atom types? `:a` => use to show if function call succeeded or not? `:ok`, `:error`
