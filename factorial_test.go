package n_choose_k 

import (
  "math/big"
  "testing"
  
  "github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
  three := new(big.Int).SetInt64(3)
  six := new(big.Int).SetInt64(6)
  ten := new(big.Int).SetInt64(10)

  zero := int(0)
  cmp1 := six.Cmp(Factorial(three))
  assert.Equal(t, cmp1, zero, "3! != 6")

  cmp2 := new(big.Int).SetInt64(720).Cmp(Factorial(six))
  assert.Equal(t, cmp2, zero, "6! != 720")

  cmp3 := new(big.Int).SetInt64(3628800).Cmp(Factorial(ten))
  assert.Equal(t, cmp3, zero, "10! != 3628800")
}
