package n_choose_k

import (
  "math/big"

  lm "github.com/georgercarder/lockless-map"
)

var bOne = newBigOne() // don't use for arithmetic.. only comparisons

var g_factorialCache = lm.NewLocklessMap()

func Factorial(b *big.Int) (ret *big.Int) {
  raw := raw(b)
  t := g_factorialCache.Take(string(raw))
  if t != nil {
    ret = t.(*big.Int)
    return
  }
  if b.Cmp(bOne) < 1 { // x <= y
    return newBigOne() 
  }
  ret = new(big.Int).Mul(b, Factorial(subtractOneFrom(b))) 
  g_factorialCache.Put(string(raw), ret)
  return
}

func newBigOne() (ret *big.Int) {
  ret = new(big.Int).SetInt64(1)
  return
}

func subtractOneFrom(b *big.Int) (ret *big.Int) {
  ret = new(big.Int).Sub(b, newBigOne()) 
  return
}

func raw(b *big.Int) (ret []byte) {
  if (b.Sign() < 0) {
    ret = append(ret, []byte{0x1}...)
  } else {
    ret = append(ret, []byte{0x0}...)
  }
  ret = append(ret, b.Bytes()...)
  return
}
