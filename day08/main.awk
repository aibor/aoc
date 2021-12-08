function sortstr(str) {
  split(str, letters, "")
  elem = asort(letters)
  out = ""

  for (v = 1; v <= elem; v++) {
    out = out letters[v]
  }

  return out
}


function arrdiff(longer, shorter) {
  split(longer, bigger, "")
  split(shorter, smaller, "")
  diffx = ""

  for (l in bigger) {
    found = 0

    for (m in smaller) {
      if (bigger[l] == smaller[m]) {
        found = 1
      }
    }

    if (!found) {
      diffx = diffx bigger[l]
    }
  }

  return diffx
}


function getmap(input, nums, map) {
  delete map
  delete seg
  seg["a"] = arrdiff(nums[3], nums[2])

  for (code in input) {
    diff = ""
    diffn = ""
    if (input[code] == 6) {
      diff = arrdiff(code, nums[4])
      diffn = arrdiff(code, nums[3])
      if (length(diff) == 2) {
        sub(seg["a"], "", diff)
        seg["g"] = diff
        seg["e"] = arrdiff(nums[7], code)
      } else if (length(diffn) == 3) {
        seg["d"] = arrdiff(nums[7], code)
      } else {
        seg["c"] = arrdiff(nums[7], code)
      }
    }
  }

  str = nums[2]
  sub(seg["c"], "", str)
  seg["f"] = str

  for (code in input) {
    diff = ""
    if (input[code] == 5) {
      diff = arrdiff(code, seg["a"] seg["c"] seg["e"] seg["d"] seg["f"] seg["g"])
      if (length(diff) == 1) {
        seg["b"] = diff
      }
    }
  }

  map[sortstr(seg["a"] seg["b"] seg["c"] seg["e"] seg["f"] seg["g"])] = "0"
  map[sortstr(seg["c"] seg["f"])] = "1"
  map[sortstr(seg["a"] seg["c"] seg["d"] seg["e"] seg["g"])] = "2"
  map[sortstr(seg["a"] seg["c"] seg["d"] seg["f"] seg["g"])] = "3"
  map[sortstr(seg["b"] seg["c"] seg["d"] seg["f"])] = "4"
  map[sortstr(seg["a"] seg["b"] seg["d"] seg["f"] seg["g"])] = "5"
  map[sortstr(seg["a"] seg["b"] seg["d"] seg["e"] seg["f"] seg["g"])] = "6"
  map[sortstr(seg["a"] seg["c"] seg["f"])] = "7"
  map[sortstr(seg["a"] seg["b"] seg["c"] seg["d"] seg["e"] seg["f"] seg["g"])] = "8"
  map[sortstr(seg["a"] seg["b"] seg["c"] seg["d"] seg["f"] seg["g"])] = "9"
}


{
  delete input
  delete nums
  delete map
  seperatorhit = 0
  numstr = ""

  for (i = 1; i <= NF; i++) {
    if ($i == "|") {
      seperatorhit = 1

      getmap(input, nums, map)
    } else if (!seperatorhit) {
      len = length($i)
      if (len == 2 || len == 3 || len == 4 || len == 7) {
        nums[len] = sortstr($i)
      }
      input[sortstr($i)] = len
    } else {

      split($i, letters, //)

      num = length(letters)
      if (num == 2 || num == 3 || num == 4 || num == 7) {
        simple_count++
      }

      numstr = numstr map[sortstr($i)]
    }
  }
  sum += int(numstr)
}

END {
  print "Part 1: ", simple_count
  print "Part 2: ", sum
}
