package main

import (
	"fmt"
	"os"
	"testing"
)

func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

func Test1(t *testing.T) {
	fmt.Println("I'm test1")
}

func Test2(t *testing.T) {
	fmt.Println("I'm test2")
}

func TestMain(m *testing.M) {
	setup()

	// 如果 code 不为 0 则有测试用例失败
	code := m.Run()
	teardown()

	os.Exit(code)
}
