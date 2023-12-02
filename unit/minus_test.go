package unit

import "testing"

func TestMinus(t *testing.T) {
	i := minus(5, 1)
	t.Logf("计算结果: %d", i)
}
