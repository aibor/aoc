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

  for (r = 1; r <= rows; r++) {
    for (c = 1; c <= cols; c++) {
      h = map[r " " c]
      hu = map[r - 1 " " c]
      hd = map[r + 1 " " c]
      hl = map[r " " c - 1]
      hr = map[r " " c + 1]

      if (hu != "" && h >= hu || hd != "" && h >= hd || hl != "" && h >= hl || hr != "" && h >= hr) {
        continue
      }

      sum += 1 + h
    }
  }

  print "Part 1:", sum
}
