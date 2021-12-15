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
  risk["1,1"] = 0
  queue["1,1"]

  while (length(queue)) {
    for (p in queue) {
      if (!(u in queue) || risk[p] < risk[u]) {
        u = p
      }
    }

    delete queue[u]

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
      if (v in map) {
        a = risk[u] + map[v]
        if (!risk[v] || a < risk[v]) {
          risk[v] = a
          prev[v] = u
          queue[v]
        }
      }
    }
  }

  print "Part 1:", risk[rows "," cols]
}


