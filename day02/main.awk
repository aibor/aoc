$1 == "down" {
  aim += $2
}

$1 == "up" {
  aim -= $2
}

$1 == "forward" {
  fwd += $2
  depth += $2 * aim
}

{
  mv[$1] += $2
}

END {
  print "Part 1:", mv["forward"] * (mv["down"] - mv["up"])
  print "Part 2:", fwd * depth
}
