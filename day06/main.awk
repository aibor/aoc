function simulate(state, days,    sum) {
  sum = 0

  for (day=1; day <=days; day++) {
    wrap = state[0]
    for (i=1; i<=8; i++) {
      state[i-1] = state[i]
    }
    state[6] += wrap
    state[8] = wrap
  }

  for (i=0; i<=8; i++) {
    sum += state[i]
  }

  return sum
}

BEGIN {
  FS = ","
}

{
  for (i=1; i<=NF; i++) {
    state[$i]++
  }
}

END {
  print "Part 1:", simulate(state, 80)
  print "Part 2:", simulate(state, 256-80)
}

