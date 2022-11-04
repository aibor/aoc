BEGIN {
  FS = ","
}

{
  for (i = 1; i <= NF; i++) {
    pos[i] = $i
    sum += $i
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

  avg = int(sum / count)

  for (i in pos) {
    dist_p1 = pos[i] - med
    dist_p1 *= dist_p1 >= 0 ? 1 : -1
    fuelsum_p1 += dist_p1

    dist_p2 = pos[i] - avg
    dist_p2 *= dist_p2 >= 0 ? 1 : -1

    for (j = 1; j <= dist_p2; j++) {
      fuelsum_p2 += j
    }
  }

  print "Part 1:", fuelsum_p1
  print "Part 2:", fuelsum_p2
}
