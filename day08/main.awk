{
  seperatorhit = 0
  for (i = 1; i <= NF; i++) {
    if ($i == "|") {
      seperatorhit = 1
      continue
    } else if (!seperatorhit) {
      continue
    }

    split($i, letters, //)

    num = length(letters)
    if (num == 2 || num == 3 || num == 4 || num == 7) {
      simple_count++
    }
  }
}

END {
  print "part 1: ", simple_count
}

