BEGIN {
  FS = ","
}

{
  for (i = 1; i <= NF; i++) {
    pos[i] = $i
  }
  count = NF
}

END {
  asort(pos)

  if (count % 2) {
    med = pos[(count + 1) / 2]
  } else {
    med = (pos[count / 2] + pos[(count / 2) + 1]) / 2.0
  }

  for (i in pos) {
    fuel = pos[i] - med
    fuelsum += fuel * (fuel >= 0 ? 1 : -1)
  }

  print "Part 1:", fuelsum
}
