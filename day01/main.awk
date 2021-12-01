{
  cur = int($1)

  if (prev != 0 && cur > prev) {
    part_one_inc++
  }

  prev = cur
}

END {
  print part_one_inc
}
