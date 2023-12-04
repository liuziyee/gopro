package unit

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestMinus(t *testing.T) {
	i := minus(5, 1)
	t.Logf("计算结果: %d", i)
}

func TestLong(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过该测试用例")
	}
	time.Sleep(time.Second * 5)
}

func TestDataSet(t *testing.T) {
	// 匿名结构体切片
	ds := []struct {
		a   int
		b   int
		out int
	}{
		{1, 0, 1},
		{0, 0, 0},
		{5, 3, 2},
	}

	for _, value := range ds {
		out := minus(value.a, value.b)
		if out != value.out {
			t.Errorf("预期值: %d, 函数返回值: %d", value.out, out)
		}
	}
}

const numbers = 10000

func BenchmarkStringPlus(b *testing.B) {
	b.ResetTimer()
	// N:测试次数
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str += strconv.Itoa(j)
		}
	}
	b.StopTimer()
}

func BenchmarkStringSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str = fmt.Sprintf("%s%d", str, i)
		}
	}
	b.StopTimer()
}

func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < numbers; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		_ = builder.String()
	}
	b.StopTimer()
}
