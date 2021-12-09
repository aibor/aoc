BEGIN {
  FS = ""
}

{
  for (i = 1; i <= NF; i++) {
    map[NR " " i] = $i
  }
  rows = NR
  cols = NF
}

END {
  sum = 0
  basinnum = 1

  for (r = 1; r <= rows; r++) {
    for (c = 1; c <= cols; c++) {
      delete b
      h = map[r " " c]
      hu = 9
      hd = 9
      hl = 9
      hr = 9

      if (r > 1) {
        hu = map[r - 1 " " c]
        b["u"] = basinmap[r - 1 " " c]
      }
      if (r < rows) {
        hd = map[r + 1 " " c]
      }
      if (c > 1) {
        hl = map[r " " c - 1]
        b["l"] = basinmap[r " " c - 1]
      }
      if (c < cols) {
        hr = map[r " " c + 1]
      }

      if (h < hu && h < hd && h < hl && h < hr) {
        sum += 1 + h
      }

      if (h < 9) {
        bnum = 0
        for (dir in b) {
          if (b[dir]) {
            bnum = b[dir]
            break
          }
        }

        if (bnum == 0) {
          if (hu == 9 && r > 1) {
            for (j = c + 1; j <= cols; j++) {
              if (map[r " " j] == 9) {
                break
              } else if (map[r - 1 " " j] != 9) {
                bnum = basinmap[r - 1 " " j]
                break
              }
            }
          }
        }

        if (bnum == 0) {
          bnum = basinnum++
        }

        basinmap[r " " c] = bnum
        basins[bnum]++
      }
    }
  }

  n = asort(basins, basins, "@val_num_desc")

  print "Part 1:", sum
  print "Part 2:", basins[1] * basins[2] * basins[3]
}
