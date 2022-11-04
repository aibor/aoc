BEGIN {
  FS = ""
}

{
  for (i=0; i<NF; i++) {
    bits[i] += $(NF-i)
  }
}

END {
  for (i in bits) {
    gamma += (bits[i] > NR/2) ? 2**i : 0
    epsilon += (bits[i] < NR/2) ? 2**i : 0
  }

  print "Part 1 :", gamma * epsilon
}
