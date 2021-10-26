## 一个基本的例子

🕹 Go 语言推荐测试文件和源代码文件放一块。

😉 测试用例命名一般为 `Test` + 待测试的方法名。

🏃 执行`go test`，该 package 下所有的测试用例都会被执行。
- go test -v，显示每个用例的测试结果
- go test -cover，查看测试用例覆盖率
- go test -run TestAdd，指定运行一个测试用例