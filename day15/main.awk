function printmap(max_x, max_y) {
  for (sx = 1; sx <= max_x; sx++) {
    for (sy = 1; sy <= max_y; sy++) {
      printf getmap(sx "," sy)
    }
    printf "\n"
  }
  printf "\n"
}


function walk(target, risk, sizer, sizec) {
  delete risk
  delete queue

  risk["1,1"] = 0
  queue["1,1"]
  u = "1,1"

  while (length(queue)) {
    for (p in queue) {
      if (!(u in queue) || risk[p] < risk[u]) {
        u = p
      }
    }

    delete queue[u]

    if (u == target) {
      break
    }

    split(u, ua, ",")
    r = int(ua[1])
    c = int(ua[2])

    delete neigh
    if (r > 1)
      neigh["u"] = r - 1 "," c
    if (r < sizer)
      neigh["d"] = r + 1 "," c
    if (c > 1)
      neigh["l"] = r "," c - 1
    if (c < sizec)
      neigh["r"] = r "," c + 1

    for (key in neigh) {
      v = neigh[key]

      a = risk[u] + getmap(v)
      if (!risk[v] || a < risk[v]) {
        risk[v] = a
        prev[v] = u
        queue[v]
      }
    }
  }
}

function getmap(point) {
  if (point in map) {
    return map[point]
  }

  split(point, pointa, ",")
  x = int(pointa[1])
  y = int(pointa[2])

  if (x < 1 || y < 1) {
    return
  }

  basex = x % rows
  basey = y % cols
  addx = int((x -1)/ rows)
  addy = int((y -1) / cols)

  basep = (basex ? basex : rows) "," (basey ? basey : cols)

  newv = map[basep] + addx + addy
  return newv > 9 ? (newv % 10) + 1 : newv
}


BEGIN {
  FS = ""
}

{
  for (i = 1; i <= NF; i++) {
    map[NR "," i] = int($i)
  }
  rows = NR
  cols = NF
}

END {
  targetnode = rows "," cols
  walk(targetnode, part1, rows, cols)

  print "Part 1:", part1[targetnode]

  targetnode = rows * 5"," cols * 5
  walk(targetnode, part2, rows * 5, cols * 5)

  print "Part 2:", part2[targetnode]
}


