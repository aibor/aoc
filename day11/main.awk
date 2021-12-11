function get_adjacents(x, y, adj) {
  delete adj

  if (x > 1) {
    adj[x - 1 "," y]++

    if (y > 1) {
      adj[x - 1 "," y - 1]++
    }

    if (y < cols) {
      adj[x - 1 "," y + 1]++
    }
  }

  if (y > 1)
    adj[x "," y - 1]++

  if (x < rows) {
    adj[x + 1 "," y]++

    if (y > 1) {
      adj[x + 1 "," y - 1]++
    }

    if (y < cols) {
      adj[x + 1 "," y + 1]++
    }
  }

  if (y < cols)
    adj[x "," y + 1]++
}

function flash(a, b) {
  if (flashed[a "," b]) {
    return
  }

  flashed[a "," b]++
  get_adjacents(a, b, adj)

  for (ad in adj) {
    map[ad]++
    if (map[ad] > 9 && !flashed[ad]) {
      split(ad, adc, ",")
      flash(adc[1], adc[2])
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
  count = length(map)
}

END {
  for (step = 1; step <= 1000; step++) {
    delete queue
    delete flashed

    print "step", step

    for (r = 1; r <= rows; r++) {
      for (c = 1; c <= cols; c++) {
        map[r "," c]++
      }
    }

    for (r = 1; r <= rows; r++) {
      for (c = 1; c <= cols; c++) {
        if (map[r "," c] > 9) {
          flash(r, c)
        }
      }
    }

    for (o in flashed) {
      map[o] = 0
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
