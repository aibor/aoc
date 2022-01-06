package main

import (
  "fmt"
  "strings"
  "strconv"

  "github.com/aibor/aoc-2021/goutils"
)

type boardnum struct {
  marked bool
  num int
}

type board struct {
  draws int
  drawnum int
  nums [5][5]*boardnum
}

func (b *boardnum) mark() {
  b.marked = true
  return
}

func (b *board) set(r, c, num int) {
  b.nums[r][c] = &boardnum{num: num, marked: false}
  return
}

func (b *board) mark(r, c int) {
  b.nums[r][c].mark()
  return
}

func (b board) unmarked_sum() (sum int) {
  for _, row := range b.nums {
    for _, bnum := range row {
      if !bnum.marked {
        sum += bnum.num
      }
    }
  }
  return
}

func (b *board) play(draws []int) {
  row_count := make([]int, 5)
  column_count := make([]int, 5)
  for draw, dnum := range draws {
    for r, row := range b.nums {
      for c, bnum := range row {
        if bnum.num == dnum {
          bnum.mark()
          row_count[r]++
          column_count[c]++
          if draw > 4 && (row_count[r] == 5 || column_count[c] == 5) {
            b.draws = draw + 1
            b.drawnum = dnum
            return
          }
        }
      }
    }
  }
  return
}

func main() {
  input, _ := goutils.ReadInput()
  draws := []int{}
  cboard := board{}
  least := board{}
  most := board{}
  row := 0

  for _, num := range strings.Split(input[0], ",") {
    numi, _ := strconv.Atoi(num)
    draws = append(draws, numi)
  }

  for _, line := range input[2:] {
    nums := strings.Fields(line)
    if len(nums) != 5 {
      cboard = board{}
      row = 0
      continue
    }

    for col, num := range nums {
      numi, _ := strconv.Atoi(num)
      cboard.set(row, col, numi)
    }

    row++

    if row == 5 {
      cboard.play(draws)
      if least.draws == 0 || cboard.draws < least.draws {
        least = cboard
      }
      if most.draws == 0 || cboard.draws > most.draws {
        most = cboard
      }
    }
  }

  fmt.Println("Part 1:", least.unmarked_sum() * least.drawnum)
  fmt.Println("Part 2:", most.unmarked_sum() * most.drawnum)
}

/*
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
*/
