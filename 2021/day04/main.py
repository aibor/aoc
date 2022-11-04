import sys

def unmarked_sum(board):
    return sum([e for row in board for e in row if e is not None])


draws = [int(e) for e in sys.stdin.readline().strip().split(",")]
least_draws = None
least_draws_result = 0
most_draws = None
most_draws_result = 0

for line in sys.stdin:
    if line.strip() == "":
        board = []
        continue

    board.append([int(e) for e in line.strip().split()])

    if len(board) < 5:
        continue

    solved = False

    for draw, num in enumerate(draws):
        if solved:
            break

        column_marked = [0, 0, 0, 0, 0]

        for row in board:
            if solved:
                break

            row_marked = 0

            for i in range(0,5):
                if solved:
                    break

                if row[i] == num:
                    row[i] = None
                if row[i] is None:
                    row_marked += 1
                    column_marked[i] += 1
                if draw > 4 and (row_marked == 5 or column_marked[i] == 5):
                    solved = True

                    if least_draws is None or draw < least_draws:
                        least_draws = draw
                        least_draws_result = unmarked_sum(board) * num

                    if most_draws is None or draw > most_draws:
                        most_draws = draw
                        most_draws_result = unmarked_sum(board) * num

print("Part 1:", least_draws_result)
print("Part 2:", most_draws_result)
