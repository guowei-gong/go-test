package main

import "testing"

type calcCase struct {
	Name           string
	A, B, Expected int
}

func createMulTestCase(t *testing.T, c *calcCase) {
	t.Helper()
	t.Run(c.Name, func(t *testing.T) {
		if ans := Mul(c.A, c.B); ans != c.Expected {
			t.Fatalf("%d * %d expected %d, but %d got",
				c.A, c.B, c.Expected, ans)
		}
	})
}

func TestMul(t *testing.T) {
	createMulTestCase(t, &calcCase{"A", 2, 3, 6})
	createMulTestCase(t, &calcCase{"B", 2, -3, -6})
	createMulTestCase(t, &calcCase{"C", 2, 0, -1}) // 这是一个错误的预期值, 用于验证 Helper 的特性
}
