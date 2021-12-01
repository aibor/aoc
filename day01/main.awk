{
  cur = int($1)
  slidwin[NR % 4] = cur

  if (NR > 1 && cur > slidwin[(NR - 1) % 4]) {
    part_one_inc++
  }

  if (NR > 3 && cur > slidwin[(NR - 3) % 4]) {
    part_two_inc++
  }
}

END {
  print "Part 1:", part_one_inc
  print "Part 2:", part_two_inc
}
