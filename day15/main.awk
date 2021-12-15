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
  for (p in map) {
    queue[p]++
    risk[p] = 100000000000000
  }
  risk["1,1"] = 0

  while (length(queue)) {
    u = "1,1"
    for (p in queue) {
      if (!queue[u] || (prev[p] && risk[p] < risk[u])) {
        u = p
      }
    }

    delete queue[u]
    print u

    if (u == rows "," cols) {
      break
    }

    split(u, ua, ",")
    r = int(ua[1])
    c = int(ua[2])

    neigh["u"] = r - 1 "," c
    neigh["d"] = r + 1 "," c
    neigh["l"] = r "," c - 1
    neigh["r"] = r "," c + 1

    for (key in neigh) {
      v = neigh[key]
      for (p in queue) {
        if (p == v) {
          a = risk[u] + map[v]
          if (a <= risk[v]) {
            risk[v] = a
            prev[v] = u
          }
        }
      }
    }
  }

  print "Part 1:", risk[rows "," cols]
}


