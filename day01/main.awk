BEGIN {
  prev = 0
  inc = 0
}

{
  cur = int($1)

  print cur, prev
  if ( prev != 0 && cur > prev) {
    inc++
  }

  prev = cur
}

END {
  print inc
}
