package main

import (
  "fmt"
  "strings"
  "strconv"

  "github.com/aibor/aoc/goutils"
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
