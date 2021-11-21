### n_choose_k

a super basic math library

Copyright (c) 2021 George Carder georgercarder@gmail.com

GENERAL PUBLIC LICENSE Version 3

recall ( n k ) = n! / (k! * (n - k)!)

the factorial steps are cached for efficiency

usage

```
  six := new(big.Int).SetInt64(6)
  ten := new(big.Int).SetInt64(10)

  result := NChooseK(ten, six))

  ...

  result2 := NChooseKInt(32, 9)
  
  ...

  seventyTwo := uint64(72) // defined elsewhere
  eleven := uint64(11) // defined elsewhere
  ...
  result3 := NChooseKUint64(seventyTwo, eleven)
```

also available is `Factorial`

```

  factorialResult := Factorial(ten) // is 3628800

```
