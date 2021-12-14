NR == 1 {
  formula = $1
}

/ -> / {
  map[$1] = $3
}

END {
  for (i = 0; i < 10; i++) {
    n = length(formula)
    newformula = ""
    for (j = 1; j <= n; j++){
      cur = substr(formula, j, 1)
      fol = substr(formula, j + 1, 1)
      bet = map[cur fol]
      newformula = (newformula ? newformula : cur) bet fol
    }
    formula = newformula
  }

  n = length(formula)
  for (i = 1; i <= n; i++){
    count[substr(formula, i, 1)]++
  }

  n = asort(count)

  print count[n] - count[1]
}

