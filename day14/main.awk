NR == 1 {
  formula = $1
}

/ -> / {
  map[$1] = $3
}

END {
  steps = 40
  low = steps
  n = split(formula, elema, "")
  count[elema[n]]++

  for (e = 1; e < n; e++) {
    cur = elema[e]
    fol = elema[e + 1]
    count[cur]++
    pairs[cur fol]++
  }

  for (i = 1; i <= steps; i++) {
    delete newpairs
    for (p in pairs) {
      newpairs[p] = pairs[p]
    }

    for (pair in pairs) {
      b = map[pair]
      count[b] += pairs[pair]
      split(pair, paira, "")
      newpairs[pair] -= pairs[pair]
      newpairs[paira[1] b] += pairs[pair]
      newpairs[b paira[2]] += pairs[pair]
    }

    delete pairs
    for (p in newpairs) {
      pairs[p] = newpairs[p]
    }
  }

  n = asort(count)

  print count[n] - count[1]
}
