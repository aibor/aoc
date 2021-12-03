import sys

def find_nums(data_arr, pos=0, maj=True):
    new_arr = [[], []]
    last = False

    if len(data_arr) == 1:
        return data_arr[0]

    for num in data_arr:
        last = len(num) == pos + 1
        new_arr[int(num[pos])].append(num)

    if len(new_arr[0]) == len(new_arr[1]):
        major = 1 if maj else 0
    elif len(new_arr[0]) > len(new_arr[1]):
        major = 0 if maj else 1
    else:
        major = 1 if maj else 0

    return find_nums(new_arr[major], pos + 1, maj)



data = [line.strip() for line in sys.stdin]

oxygen = find_nums(data)
carbon = find_nums(data, maj=False)

print("Part 2:", int(oxygen, 2) * int(carbon, 2))
