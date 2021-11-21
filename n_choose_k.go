package n_choose_k

import (
  "math/big"
)

// recall ( n k ) = n! / (k! * (n - k)!)

func NChooseK(n, k *big.Int) (ret *big.Int) {
  kFactorial := Factorial(k)
  differenceFactorial := Factorial(new(big.Int).Sub(n, k))
  denominator := new(big.Int).Mul(kFactorial, differenceFactorial)
  // numerator done second since it is larger than k and n-k so can embrace caching
  numerator := Factorial(n)   
  ret = new(big.Int).Div(numerator, denominator)
  return
}

func NChooseKUint64(n, k uint64) (ret *big.Int) {
  ret = NChooseK(new(big.Int).SetUint64(n), new(big.Int).SetUint64(k))
  return
}

func NChooseKInt(n, k int) (ret *big.Int) {
  ret = NChooseK(new(big.Int).SetInt64(int64(n)), new(big.Int).SetInt64(int64(k)))
  return
}
