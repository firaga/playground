package main

import "fmt"

func test() {
	//在函数退出前，执行defer
	//捕捉异常后，程序不会异常退出
	defer func() {
		err := recover() //内置函数，可以捕捉到函数异常
		if err != nil {
			//这里是打印错误，还可以进行报警处理，例如微信，邮箱通知
			fmt.Println("err错误信息：", err)
		}
	}()

	//如果没有异常捕获，直接报错panic，运行时出错
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res结果：", res)

}

func main() {
	test()
	fmt.Println("如果程序没退出，就走我这里")
}