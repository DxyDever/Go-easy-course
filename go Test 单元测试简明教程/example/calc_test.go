package main

import "testing"

func TestAdd(t *testing.T) {
	if ans := add(1, 2); ans != 3 {
		t.Errorf("1 + 2 预估为3， 但是%d", ans)
	}
}

/*

测试用例名称一般命名为 Test 加上待测试的方法名。
测试用的参数有且只有一个，在这里是 t *testing.T。
基准测试(benchmark)的参数是 *testing.B，TestMain 的参数是 *testing.M 类型。

运行 go test，该 package 下所有的测试用例都会被执行。


$ go test
ok      example 0.009s


或 go test -v，-v 参数会显示每个用例的测试结果，另外 -cover 参数可以查看覆盖率。

1
2
3
4
5
6
7
$ go test -v
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
=== RUN   TestMul
--- PASS: TestMul (0.00s)
PASS
ok      example 0.007s
如果只想运行其中的一个用例，例如 TestAdd，可以用 -run 参数指定，该参数支持通配符 *，和部分正则表达式，例如 ^、$。

$ go test -run TestAdd -v
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
PASS
ok      example 0.007s

*/
