package n_choose_k 

import (
  "math/big"
  "testing"
  
  "github.com/stretchr/testify/assert"
)

func TestChoose(t *testing.T) {
  three := new(big.Int).SetInt64(3)
  six := new(big.Int).SetInt64(6)
  ten := new(big.Int).SetInt64(10)

  zero := int(0)
  cmp1 := new(big.Int).SetInt64(20).Cmp(NChooseK(six, three))
  assert.Equal(t, cmp1, zero, "( 6 3 ) != 20")

  cmp2 := new(big.Int).SetInt64(120).Cmp(NChooseK(ten, three))
  assert.Equal(t, cmp2, zero, "( 10 3 ) != 120")

  cmp3 := new(big.Int).SetInt64(210).Cmp(NChooseK(ten, six))
  assert.Equal(t, cmp3, zero, "( 10 6 ) != 210")

  cmp4 := new(big.Int).SetInt64(28048800).Cmp(NChooseKInt(32, 9))
  assert.Equal(t, cmp4, zero, "( 32 9 ) != 28048800")

  cmp5 := new(big.Int).SetInt64(3022285436352).Cmp(NChooseKUint64(uint64(72), uint64(11)))
  assert.Equal(t, cmp5, zero, "( 72 11 ) != 3022285436352")
}
