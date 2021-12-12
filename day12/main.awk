function pathfind(path, doublei) {
  delete newpaths
  delete seen_count
  n = split(path, caves, "-")
  cave = caves[n]
  double = doublei

  for (x in caves) {
    if (caves[x] !~ /^[[:lower:]]+$/) {
      continue
    }
    seen_count[caves[x]]++
    if (seen_count[caves[x]] > 1) {
      double = 1
      break
    }
  }

  for (c in map) {
    split(c, cn, "-")

    if (cn[1] != cave) {
      continue
    }

    if (cn[2] == "end") {
      paths[path "-" cn[2]]++
      continue
    } else if (cn[2] ~ /^[[:lower:]]+$/) {
      skip = 0
      for (seen in caves) {
        if (cn[2] == caves[seen]) {
          skip = 1
          break
        }
      }
      if (double && skip) {
        continue
      }
    }

    newpaths[path "-" cn[2]]++
  }

  for (p in newpaths) {
    pathfind(p, doublei)
  }
}

BEGIN {
  FS = "-"
}

{
  if ($1 != "end" && $2 != "start") {
    map[$1 "-" $2]++
  }
  if ($1 != "start" && $2 != "end") {
    map[$2 "-" $1]++
  }
}

END {
  delete paths

  pathfind("start", 1)

  print "Part 1:", length(paths)

  delete paths

  pathfind("start", 0)

  print "Part 2:", length(paths)
}
