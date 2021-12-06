BEGIN {
  FS = ","
  days = 80
}

{
  for (i=1; i<=NF; i++) {
    state[$i]++
  }
}

END {
  for (day=1; day<=days; day++) {
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

  print "Part 1:", sum
}

