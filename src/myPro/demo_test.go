/**
 测试用例（文件名以_test结尾，函数名以Test开头）
运行测试用例的需要以go test运行test文件
go test -v // 测试功能
go test -bench="." // 压力测试
go test -cover // 覆盖率测试
 */
package main

import "testing"
/**
 功能测试
 */
func TestGetAr(t *testing.T)  {
	a := GetArea(100, 200)
	if a < 200 {
		t.Error("测试失败")
	}
}

/**
 性能测试
 */
func BenchmarkGetAr(t *testing.B)  {
	for i := 0; i < t.N; i++ {
		GetArea(100, 20)
	}
}