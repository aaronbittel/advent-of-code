def to_string(filesystem):
    out = ""
    for item in filesystem:
        if type(item) is tuple:
            (id, size) = item
            out += str(id) * size
        else:
            out += "." * item
    return out


filesystem = []

with open("./test.txt") as f:
    next_idx = 0
    next_type = "file"
    for size in f.readline().strip():
        size = int(size)
        if next_type == "file":
            filesystem.append((next_idx, size))
            next_idx += 1
            next_type = "free"
        elif next_type == "free" and size > 0:
            filesystem.append(size)
            next_type = "file"
        else:
            next_type = "file"

# print(to_string(filesystem))

string = to_string(filesystem)

right = len(filesystem) - 1
while right >= 0:
    cur = filesystem[right]

    if type(cur) is tuple:
        (id, size) = cur
        for idx, elem in enumerate(filesystem):
            if idx >= right:
                break
            if type(elem) is tuple:
                continue
            if elem >= size:
                filesystem[right] = size
                filesystem[idx] -= size
                filesystem.insert(idx, cur)
                break
    right -= 1
    # print(to_string(filesystem))

res = 0
for i, c in enumerate(to_string(filesystem)):
    if c != ".":
        res += i * int(c)
print(res)
