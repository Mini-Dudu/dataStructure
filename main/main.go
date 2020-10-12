package main

import (
	"DataStruct/linkedList"
	"DataStruct/queue"
	"DataStruct/stack"
	"fmt"
)
func main() {
	//测试单链表各个功能
	fmt.Println("单链表测试信息如下:")
	linkedList.Test_sgleLinkList()
	fmt.Println()

	//测试双链表的各个功能
	fmt.Println("双链表测试信息如下:")
	linkedList.Tese_doubleLinkList()
	fmt.Println()

	//测试顺序存储结构的队列的各个功能
	fmt.Println("顺序结构队列测试信息如下:	MaxSize为5的测试数据")
	queue.Test_Queue_sequential()
	fmt.Println()

	//测试链式存储结构的队列的各个功能
	fmt.Println("链式结构队列测试信息如下:")
	queue.Test_Queue_chainStructure()
	fmt.Println()

	//测试顺序存储结构的环形队列的各个功能
	fmt.Println("顺序存储结构的环形队列测试信息如下:	MaxSize为5的测试数据")
	queue.Test_CircleQueue_sequential()
	fmt.Println()

	//测试链式结构的环形队列的各个功能
	fmt.Println("链式结构的环形队列测试信息如下:")
	queue.Test_CircleQueue_chainStructure()
	fmt.Println()

	//测试环形单链表的各个功能
	fmt.Println("环形单链表测试信息如下:")
	linkedList.Tese_CircularSgleLinkList()
	fmt.Println()

	//测试环形双链表的各个功能
	fmt.Println("环形双链表测试信息如下:")
	linkedList.Tese_CircularDoubleLinkList()
	fmt.Println()

	//测试顺序存储结构的栈的相关功能
	fmt.Println("顺序存储结构的栈测试信息如下:")
	stack.Test_SequentialStack()
	fmt.Println()


	//测试链式存储结构的栈的相关功能
	fmt.Println("链式存储结构的栈测试信息如下:")
	stack.Test_chainStructure()
	fmt.Println()
}
