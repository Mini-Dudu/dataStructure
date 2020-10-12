//栈:stack
package stack

import (
	"errors"
	"fmt"
)

const (
	StackMaxSize = 30
)

//实现顺序存储结构的栈,极其相关操作
//顺序结构:sequential
type SequentialStack struct {
	top int				//栈顶
	data [StackMaxSize]interface{}	//栈内数据,,,,数组下标为零出为栈底
}

//初始化一个顺序存储结构的栈
func InitSequentialStack() *SequentialStack {
	return &SequentialStack{
		top:    -1,
		data:   [StackMaxSize]interface{}{},
	}
}

//入栈操作
func (S *SequentialStack) Push(data interface{}) error {
	if S.IsFull() {
		return errors.New("栈已满,请pull后重新操作!")
	}
	S.top++
	S.data[S.top] = data
	return nil
}

//出栈操作
func (S *SequentialStack) Pull() (interface{},error) {
	if S.IsEmpty() {
		return nil,errors.New("当前栈为空,请push后重新操作!")
	}
	data := S.data[S.top]
	S.top--
	return data,nil
}

//获得栈顶元素
func (S *SequentialStack) Get() interface{} {
	return S.data[S.top]
}

//判断顺序存储结构栈是否为空
func (S *SequentialStack) IsEmpty() bool {
	return S.top == -1
}

//判断顺序存储结构栈是否已满
func (S *SequentialStack) IsFull() bool {
	return S.top == StackMaxSize - 1
}





//实现链式存储结构的栈,极其相关操作
//链式结构 chain structure
type ChainStack struct {	//链式存储结构栈不需要栈顶及栈底,栈底用头节点表示
	data interface{}		//数据
	next *ChainStack		//下一个栈内数据
}

//说明:
	//链式存储结构栈不需要栈顶及栈底,可用头节点代替,例如:
		//1.栈底用头节点端表示
		//2.栈顶用头节点端表示

//以下链式栈采用第一种方式实现,效率低于第二种方式

//新建数据节点
func (C *ChainStack) New(data interface{}) *ChainStack {
	return &ChainStack{
		data: data,
		next: nil,
	}
}

//初始化一个链式存储结构的空栈
func InitChainStack() *ChainStack{
	return &ChainStack{
		data: nil,
		next: nil,
	}
}

//入栈操作
func (C *ChainStack) Push(data interface{}) {
	newNode := C.New(data)
	temp := C

	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

//出栈操作,栈顶元素出栈,并返回栈顶数据和错误
func (C *ChainStack) Pull() interface{} {
	if C.next == nil {
		return errors.New("当前栈为空,请入栈后重试!")
	}

	temp := C

	//找尾节点的前一个节点
	for temp.next.next != nil {
		temp = temp.next
	}
	data := temp.next.data
	temp.next = nil
	return data
}

//获得栈顶元素,并返回栈顶数据和错误
func (C *ChainStack) Get() interface{} {
	if C.next == nil {
		return errors.New("当前栈为空,请入栈后重试!")
	}

	temp := C

	for temp.next.next != nil {
		temp = temp.next
	}
	data := temp.next.data
	return data
}

//查看栈内有多少个数据
func (C *ChainStack) Len() int {
	count := 0
	temp := C
	for temp.next != nil {
		temp = temp.next
		count++
	}
	return count
}

//测试用于显示链式结构栈中的数据用
func (C *ChainStack) Show() {
	temp := C
	for temp.next != nil {
		if temp.next.next != nil {
			fmt.Printf("%s-->",temp.next.data)
		}else {
			fmt.Printf("%s\n",temp.next.data)
		}

		temp = temp.next
	}

}
