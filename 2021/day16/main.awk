function bin2dec(binstr,    d) {
  n = split(binstr, bina, "")

  for (k = 0; k < n; k++) {
    d += bina[n-k] * 2 ** k
  }

  return d
}

function hex2bin(hexnum,    b) {
  d = strtonum(hexnum)

  while (d) {
    b = d%2 b
    d = int(d/2)
  }

  while (length(b) < 4) {
    b = 0 b
  }

  return b
}


BEGIN {
  FS = ""
}

{
  for (i = 1; i <= NF; i++) {
    input = input hex2bin("0x" $i)
  }
}

END {
  i = 1
  version = ""
  type = ""
  stackid = 0

  while (i <= length(input)) {
    if (substr(input, i) ~ /^0*$/) {
      break
    }

    if (version == "") {
      if (stackid > 0) {
        if (stacklid[stackid] == 0) {
          if (i >= stack[stackid]) {
            stackid--
          }
        } else {
          if (--stack[stackid] <= 0) {
            stackid--
          }
        }
      }

      version = bin2dec(substr(input, i, 3))
      i += 3
      versionsum += version
    } else if (type == "") {
      type = bin2dec(substr(input, i, 3))
      i += 3
    } else if (type == 4) {
    # literal
      going = 1
      litval = ""

      while (going == 1) {
        going = substr(input, i, 1)
        litval = litval substr(input, i + 1, 4)
        i += 5
      }
      val = bin2dec(litval)

      print "lit:", i, version, type, litval, val
      version = ""
      type = ""
    } else {
    # operator
      lid = substr(input, i++, 1)
      lenbits = (lid  == 0) ? 15 : 11
      len = bin2dec(substr(input, i, lenbits))
      i += lenbits

      stacklid[++stackid] = lid
      stack[stackid] = len

      print "opt:", i, version, type, lid, len
      version = ""
      type = ""
    }
  }

  print "Part 1:", versionsum
}
