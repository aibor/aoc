BEGIN {
  prev = 0
  part_one_inc = 0
  part_two_inc = 0
}

{
  cur = int($1)

  if ( prev != 0 && cur > prev) {
    part_one_inc++
  }

  prev = cur
}

END {
  print part_one_inc
}
