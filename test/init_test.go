package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestAdd 测试 Add 函数
func TestAdd(t *testing.T) {
	// 定义测试用例
	testCases := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
		{-1, -1, -2},
	}

	// 遍历测试用例并执行测试
	for _, tc := range testCases {
		result := Add(tc.a, tc.b)
		assert.Equal(t, tc.expected, result, "Add(%d, %d) should be %d", tc.a, tc.b, tc.expected)
	}
}
