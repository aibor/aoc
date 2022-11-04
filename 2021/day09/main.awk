function explore_basin(x, y,    queue, explored, keys) {
  queue[x"," y]++
  explored[x "," y]++
  size = 1

  while (length(queue)) {
    for (node in queue) {
      break
    }
    delete queue[node]
    split(node, coords, ",")
    a = coords[1]
    b = coords[2]

    if (a > 1) {
      keys[a - 1 "," b]++
    }
    if (a < rows) {
      keys[a + 1 "," b]++
    }
    if (b > 1) {
      keys[a "," b - 1]++
    }
    if (b < cols) {
      keys[a "," b + 1]++
    }

    for (key in keys) {
      if (!explored[key]) {
        explored[key]++
        if (map[key] != 9) {
          size++
          queue[key]++
        }
      }
    }
  }

  return size
}

BEGIN {
  FS = ""
}

{
  for (i = 1; i <= NF; i++) {
    map[NR "," i] = $i
  }
  rows = NR
  cols = NF
}

END {
  sum = 0

  for (r = 1; r <= rows; r++) {
    for (c = 1; c <= cols; c++) {
      h = map[r "," c]
      hu = 9
      hd = 9
      hl = 9
      hr = 9

      if (r > 1) {
        hu = map[r - 1 "," c]
      }
      if (r < rows) {
        hd = map[r + 1 "," c]
      }
      if (c > 1) {
        hl = map[r "," c - 1]
      }
      if (c < cols) {
        hr = map[r "," c + 1]
      }

      if (h < hu && h < hd && h < hl && h < hr) {
        sum += 1 + h
        basins[r "," c] = explore_basin(r, c)
      }
    }
  }

  n = asort(basins, basins, "@val_num_desc")

  print "Part 1:", sum
  print "Part 2:", basins[1] * basins[2] * basins[3]
}
