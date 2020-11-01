//
// Created by 嘟嘟 on 2020/10/6.
//



package stack

import (
	"strconv"
)

//用顺序栈实现计算一个数学表达式

//实现将中缀表达式转成后缀表达式
func Transition(str string) []byte {
	//一:初始化一个栈,用于保存暂时还不能确定的运算顺序的运算符
	//二:从左到右处理各个元素,直到末尾,可能遇到三种情况:
		//1.遇到操作数,直接加入后缀表达式.
		//2.遇到界线符,遇到 "(" 直接入栈,遇到 ")" 则依次弹出栈内运算符并加入后缀表达式,直到弹出 "(" 为止. 注."(" 不加入后缀表达式.
		//3.遇到运算符. 依次弹出栈中优先级高于或等于当前运算符的所有运算符,并加入后缀表达式,若碰到 "(" 或栈空则停止.之后在把当前运算符入栈
	//按上述方式处理完表达式后,若栈中还有运算符,则将剩余运算符依次弹出,并加入后缀表达式.

	//将中缀表达式转换成字节切片
	strSlice := make([]byte,5)
	strSlice = []byte(str)
	//fmt.Println("中缀字节切片:",strSlice)

	//保存后缀表达式,以字节切片的形式保存
	PostfixExpression := make([]byte,0)

	//符号栈,用于保存暂时不能确定运算顺序的运算符
	stack := InitSequentialStack()

	//fmt.Println("栈内数据:",stack.data)

	for i := 0; i < len(strSlice); i++ {

		switch string(strSlice[i]) {
			case "(":
				//直接入栈
				stack.Push(int(strSlice[i]))

				//fmt.Println("(")
			case ")":
				//fmt.Println(stack.data)
				for !stack.IsEmpty() {
					//先判断栈顶元素是否为 "(", 如果为 "(" 则弹出,并结束弹出符号
					if string(stack.Get()) == "(" {
						stack.Pull()
						break
					}

					//fmt.Println("-----",string(stack.Get()))
					topEle, _ := stack.Pull()

					PostfixExpression = append(PostfixExpression,  byte(topEle))
				}


				//fmt.Println(")")
			case "*":

				if stack.IsEmpty() {
					//fmt.Println("+")
					stack.Push(int(strSlice[i]))
					continue
				}

				//扫描到乘号,依次弹出栈中优先级高于和等于乘号的运算符,并加入后缀表达式
				for !stack.IsEmpty() && ( string(stack.Get()) == "*" || string(stack.Get()) == "/" ) {
					topEle, _ := stack.Pull()

					PostfixExpression = append(PostfixExpression,  byte(topEle))
				}

				//最后将乘号(码值)入栈
				stack.Push(int(strSlice[i]))

				//fmt.Println("*")
			case "/":

				if stack.IsEmpty() {
					//fmt.Println("+")
					stack.Push(int(strSlice[i]))
					continue
				}

				//扫描到除号,依次弹出栈中优先级高于和等于除号的运算符,并加入后缀表达式
				for !stack.IsEmpty() && ( string(stack.Get()) == "*" || string(stack.Get()) == "/" ) {
					topEle, _ := stack.Pull()
					PostfixExpression = append(PostfixExpression,  byte(topEle))
				}

				//最后将除号(码值)入栈
				stack.Push(int(strSlice[i]))

				//fmt.Println("/")
			case "+":
				if stack.IsEmpty() {
					//fmt.Println("+")
					stack.Push(int(strSlice[i]))
					continue
				}

				//扫描到加号,依次弹出栈中优先级高于和等于加号的运算符,并加入后缀表达式
				for !stack.IsEmpty() && (string(stack.Get()) == "+" || string(stack.Get()) == "-" ) {
					topEle, _ := stack.Pull()
					PostfixExpression = append(PostfixExpression,  byte(topEle))
				}

				//最后将加号(码值)入栈
				stack.Push(int(strSlice[i]))

				//fmt.Println("+")
			case "-":

				if stack.IsEmpty() {
					stack.Push(int(strSlice[i]))
					continue
				}

				//扫描到减号,依次弹出栈中优先级高于和等于减号的运算符,并加入后缀表达式
				for !stack.IsEmpty() && (string(stack.Get()) == "+" || string(stack.Get()) == "-") {
					topEle, _ := stack.Pull()
					PostfixExpression = append(PostfixExpression,  byte(topEle))
				}

				//最后将减号(码值)入栈
				stack.Push(int(strSlice[i]))

				//fmt.Println("-")
			default:
				//默认是数字的话直接加入后缀表达式
				PostfixExpression = append(PostfixExpression,strSlice[i])

				//表达式扫描完后需要检查栈中是否还有运算符,有运算符的话需要依次弹出,并加入后缀表达式
				if i == len(strSlice) - 1 {
					for !stack.IsEmpty() {
						topEle, _ := stack.Pull()
						PostfixExpression = append(PostfixExpression,  byte(topEle))
					}
				}

				//fmt.Println("默认")
			}

	}

	//返回后缀表达式的字节切片
	return PostfixExpression
}

//用栈实现后缀表达式的计算
func Calculate(str string) int {

	//1.从左往右扫描后缀表达式,直到处理完所有元素
	//2.若扫描到操作数,则压入栈,并回到步骤 1 ,否则执行步骤 3
	//3.若扫描到运算符,则弹出两个栈顶元素,执行相应计算,预算结果重新入栈,再返回步骤 1

	//将中缀表达式转成后缀表达式 (以码值的形式存储)
	PostfixExpression := Transition(str)

	//将码值转换成数值
	numSlice := []string{}

	for i := 0; i < len(PostfixExpression); i++ {
		numSlice = append(numSlice,string(PostfixExpression[i]))
	}

	//fmt.Println(PostfixExpression)
	//fmt.Println(numSlice)

	stack := InitSequentialStack()

	num_a := 0
	num_b := 0
	num_c := 0

	//依次扫描后缀表达式
	for i := 0; i < len(PostfixExpression); i++ {
		//fmt.Println(i)
		switch string(PostfixExpression[i]) {
			case "+":
				//弹出栈顶两个操作数
				num_a, _ = stack.Pull()
				num_b, _ = stack.Pull()

				//计算 num_b + num_a
				num_c = num_b + num_a

				//将计算结果压入栈中
				stack.Push(num_c)

			case "-":
				//弹出栈顶两个操作数
				num_a, _ = stack.Pull()
				num_b, _ = stack.Pull()

				//计算 num_b - num_a
				num_c = num_b - num_a

				//将计算结果压入栈中
				stack.Push(num_c)

			case "*":
				//弹出栈顶两个操作数
				num_a, _ = stack.Pull()
				num_b, _ = stack.Pull()

				//计算 num_b * num_a
				num_c = num_b * num_a

				//将计算结果压入栈中
				stack.Push(num_c)

			case "/":
				//弹出栈顶两个操作数
				num_a, _ = stack.Pull()
				num_b, _ = stack.Pull()

				//计算 num_b / num_a
				num_c = num_b / num_a

				//将计算结果压入栈中
				stack.Push(num_c)

		default:
			num_a, _ = strconv.Atoi(string(PostfixExpression[i]))
			stack.Push(num_a)

		}
	}

	//fmt.Println("num_c的值:",num_c)
	//fmt.Println("后缀表达式:",string(PostfixExpression))

	return num_c
}