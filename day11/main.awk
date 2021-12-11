function get_adjacents(point, adj) {
  delete adj

  split(point, pointa, ",")
  x = pointa[1]
  y = pointa[2]

  if (x > 1) {
    adj[x - 1 "," y]++

    if (y > 1) {
      adj[x - 1 "," y - 1]++
    }

    if (y < cols) {
      adj[x - 1 "," y + 1]++
    }
  }

  if (y > 1) {
    adj[x "," y - 1]++
  }

  if (x < rows) {
    adj[x + 1 "," y]++

    if (y > 1) {
      adj[x + 1 "," y - 1]++
    }

    if (y < cols) {
      adj[x + 1 "," y + 1]++
    }
  }

  if (y < cols) {
    adj[x "," y + 1]++
  }
}

function flash(coord) {
  if (flashed[coord]) {
    return
  }

  flashed[coord]++
  get_adjacents(coord, adj)

  for (ad in adj) {
    map[ad]++
    if (map[ad] > 9) {
      flash(ad)
    }
  }
}


BEGIN {
  FS = ""
}

{
  for (i = 1; i <=NF; i++) {
    map[NR "," i] = $i
  }
  rows = NR
  cols = NF
}

END {
  count = length(map)

  while (++step) {
    delete queue
    delete flashed

    for (p in map) {
      map[p]++
      if (map[p] > 9) {
        queue[p]++
      }
    }

    for (p in queue) {
      flash(p)
    }

    for (p in flashed) {
      map[p] = 0
    }

    flashcount = length(flashed)

    if (!allflashed && flashcount == count) {
      allflashed = step
    }

    if (step <= 100) {
      p1_sum += flashcount
    } else if (allflashed) {
      break
    }
  }

  print "Part 1:", p1_sum
  print "Part 2:", allflashed
}
