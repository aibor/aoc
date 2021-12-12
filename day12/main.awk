function pathfind(path) {
  delete newpaths
  n = split(path, caves, "-")
  cave = caves[n]

  i = 0
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
      if (skip) {
        continue
      }
    }

    newpaths[path "-" cn[2]]++
  }

  for (p in newpaths) {
    pathfind(p)
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

  pathfind("start")

  print "Part 1:", length(paths)
}
