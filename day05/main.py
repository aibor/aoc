from sys import stdin

vecs = [line.strip().split(" -> ") for line in stdin.readlines()]
coords = {}

for start, end in vecs:
    x1, y1 = [int(e) for e in start.split(",")]
    x2, y2 = [int(e) for e in end.split(",")]

    if x1 != x2 and y1 != y2:
        continue

    xw1 = x1 if x1 < x2 else x2
    xw2 = x1 if x1 > x2 else x2
    yw1 = y1 if y1 < y2 else y2
    yw2 = y1 if y1 > y2 else y2

    for i in range(xw1, xw2+1):
        for j in range(yw1, yw2+1):
            key = f"{i},{j}"
            coords[key] = coords.get(key, 0) + 1

count = len([e for e in coords.values() if e > 1])

print("Part 1:", count)
