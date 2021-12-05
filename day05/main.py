from sys import stdin

vecs = [line.strip().split(" -> ") for line in stdin.readlines()]
p1_coords = {}
p2_coords = {}

for start, end in vecs:
    x1, y1 = [int(e) for e in start.split(",")]
    x2, y2 = [int(e) for e in end.split(",")]

    diag = x1 != x2 and y1 != y2

    x, y = x1, y1

    while True:
        if not diag:
            y = y1

        while True:
            key = f"{x},{y}"
            if not diag:
                p1_coords[key] = p1_coords.get(key, 0) + 1
            p2_coords[key] = p2_coords.get(key, 0) + 1

            if y == y2:
                break

            y = y + 1 if y2 > y else y - 1

            if diag:
                break

        if x == x2:
            break

        x = x + 1 if x2 > x else x - 1

p1_count = len([e for e in p1_coords.values() if e > 1])
p2_count = len([e for e in p2_coords.values() if e > 1])

print("Part 1:", p1_count)
print("Part 2:", p2_count)
