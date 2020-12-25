//
// Created by 嘟嘟 on 2020/10/7.
//

package stack

import (
	"fmt"
)

//测试顺序存储结构的栈的各个功能
func Test_SequentialStack() {
	sStack := InitSequentialStack()
	fmt.Println("新建的顺序栈是否为空:", sStack.IsEmpty())
	sStack.Push(1)
	sStack.Push(2)
	sStack.Push(3)
	sStack.Push(4)
	sStack.Push(5)
	sStack.Pull()
	fmt.Println("进行五次入栈和一次出栈操作后顺序栈是否为空:", sStack.IsEmpty())
}

//测试链式存储结构的栈的各个功能
func Test_chainStructure() {
	stack := InitChainStack()
	stack.Push("嘟嘟1")
	stack.Push("嘟嘟2")
	stack.Push("嘟嘟3")
	stack.Push("嘟嘟4")

	fmt.Println(stack.Len())

	data := stack.Get()
	fmt.Println(data)

	data = stack.Pull()
	fmt.Println(data)

	fmt.Println()

	fmt.Println(stack.Len())
	stack.Show()

	//以上步骤:
	//1.先将嘟嘟1,2,3,4入栈
	//2.查看链式栈的元素的个数
	//3.获取栈顶元素并打印
	//4.弹出栈顶元素并打印
	//5.最后查看栈的元素个数,并显示栈内所有元素数据

}

//测试表达式的计算
func Test_Calculate() int {

	//中缀表达式
	str := "5+6-4+2+1" //10
	//str := "3+2*(5-3)/2"//
	//str := "((3+7*4)-3*3)/(3-1)"

	//经过转换,得到后缀表达式
	//PostfixExpre := Transition(str)

	//根据后缀表达式,计算结果
	count := Calculate(str)

	return count
}
