function printmap() {
  for (y = 0; y <= max_y; y++) {
    for (x = 0; x <= max_x; x++) {
      printf map[x "," y] ? "#" : "."
    }
    printf "\n"
  }
  printf "\n"
}

function fold(mark) {
  split(mark, marka, "=")
  direction = marka[1]
  line = marka[2]

  if (direction == "x") {
    for (x = line + 1; x <= max_x; x++) {
      new_x = line - (x - line)
      for (y = 0; y <= max_y; y++) {
        if (map[x "," y]) {
          map[new_x "," y]++
        }
        delete map[x "," y]
      }
    }
    max_x = line - 1
  } else {
    for (y = line + 1; y <= max_y; y++) {
      new_y = line - (y - line)
      for (x = 0; x <= max_x; x++) {
        if (map[x "," y]) {
          map[x "," new_y]++
        }
        delete map[x "," y]
      }
    }
    max_y = line - 1
  }
}


/^fold/ {
  folds[isarray(folds) ? length(folds)+1 : 1] = $3
}

/,/ {
  map[$1]++
  split($1, fields, ",")
  x = int(fields[1])
  y = int(fields[2])
  if (x > max_x) {
    max_x = x
  }
  if (y > max_y) {
    max_y = y
  }
}

END {
  fold(folds[1])

  for (c in map) {
    if (map[c]) {
      p1_dots++
    }
  }

  print "Part 1:", p1_dots

  for (f = 2; f <= length(folds); f++) {
    fold(folds[f])
  }

  print "Part 2:"
  printmap()
}
