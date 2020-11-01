//
// Created by 嘟嘟 on 2020/10/4.
//


//环形队列: circle queue
package queue

import (
	"errors"
	"fmt"
)

//实现顺序存储结构的环形队列,极其相关操作
type CircleQueue_sequential struct {
	head int			//队头
	queue [size]interface{}
	tail int			//队尾
}

const (
	size = 15	//顺序存储结构的环形队列中存储数据的数组大小,,实际能存储的个数为: size - 1
)

//初始化一个顺序结构的环形队列
func InitCircleQueue_sequential() *CircleQueue_sequential {
	return &CircleQueue_sequential{
		head:   0,
		queue: [size]interface{}{},
		tail:  0,
	}
}

//入队操作
func (Q *CircleQueue_sequential) Push(data interface{}) error {
	//判断队列是否已满
	if Q.IsFull() {
		return errors.New("队列已满,请Pull后重试!")
	}

	//将数据入队
	Q.queue[Q.tail] = data
	Q.tail = (Q.tail + 1) % size
	return nil
}

//获得队尾元素
func (Q *CircleQueue_sequential) Get() (interface{},error) {
	if Q.IsEmpty() {
		return nil, errors.New("当前队列为空,请Push后重试!")
	}
	return Q.queue[Q.tail], nil
}
//出队列操作
func (Q *CircleQueue_sequential) Pull() (interface{},error) {
	//判断队列是否为空
	if Q.IsEmpty() {
		return nil, errors.New("当前顺序结构环形队列为空,请Push后重试!")
	}
	//数据出队列
	Q.head = (Q.head + 1) % size
	data := Q.queue[Q.head]

	return data, nil
}

//判断队列是否已满
func (Q *CircleQueue_sequential) IsFull() bool {
	return (Q.tail + 1) % size == Q.head
}

//判断队列是否为空
func (Q *CircleQueue_sequential) IsEmpty() bool {
	return Q.head == Q.tail
}

//获取队列中当前的长度
func (Q *CircleQueue_sequential) Len() int {
	return (Q.tail + size - Q.head) % size
}



//实现链式(单链表)存储结构的环形队列,极其相关操作
type CircleQueue_chainStructure struct {
	data interface{}					//入队数据
	next *CircleQueue_chainStructure	//下一个元素
	tail *CircleQueue_chainStructure	//队尾元素
	//用头节点端作为队头
}

//初始化一个链式结构的环形队列
func InitCircleQueue_chainStructure() *CircleQueue_chainStructure {
	CircleQueue_chainStructure := &CircleQueue_chainStructure{
		data: nil,
		next: nil,
		tail: nil,
	}
	CircleQueue_chainStructure.next = CircleQueue_chainStructure
	CircleQueue_chainStructure.tail = CircleQueue_chainStructure

	return CircleQueue_chainStructure
}

//新建一个单链表节点
func (Q *CircleQueue_chainStructure) New(data interface{}) *CircleQueue_chainStructure {
	return &CircleQueue_chainStructure{
		data: data,
		next: nil,
		tail: nil,
	}
}

//入队操作
func (Q *CircleQueue_chainStructure) Push(data interface{}) error {
	//判断队列是否已满
	if Q.IsFull() {
		return errors.New("队列已满,请Pull后重试!")
	}

	//将数据入队
	newNode := Q.New(data)

	Q.tail.next = newNode
	newNode.next = Q
	Q.tail = Q.tail.next
	return nil
}

//获得队尾元素
func (Q *CircleQueue_chainStructure) Get() (interface{},error) {
	//判断队列是否为空
	if Q.IsEmpty() {
		return nil, errors.New("当前队列为空,请Push后重试!")
	}

	//返回数据
	return Q.next.data, nil
}
//出队列操作
func (Q *CircleQueue_chainStructure) Pull() (interface{},error) {
	//判断队列是否为空
	if Q.IsEmpty() {
		return nil, errors.New("当前顺序结构环形队列为空,请Push后重试!")
	}
	//数据出队列
	data := Q.next.data
	Q.next = Q.next.next
	return data, nil
}

//判断队列是否已满
func (Q *CircleQueue_chainStructure) IsFull() bool {
	//初始化后没有添加元素进队列中的时候Q.next为空
	if Q.next == nil {
		return false
	}

	//计数器
	count := 0

	//定义一个中间变量,辅助计数
	temp := Q

	//添加数据后最后一个队列节点的next会指向头节点
	for temp.next != Q {
		count++
		temp = temp.next
	}

	//其实链式结构不用定义最大存储容量,这里为了测试定义为20
	return count == 20
}

//判断队列是否为空
func (Q *CircleQueue_chainStructure) IsEmpty() bool {
	//有可能没有进行添加数据也判断是否为空
	return Q.next == Q || Q.tail == Q
}

//获取队列中当前的长度
func (Q *CircleQueue_chainStructure) Len() int {
	//初始化后没有添加元素进队列中的时候Q.next为空
	if Q.next == nil {
		return 0
	}
	//计数器
	count := 0

	//定义一个中间变量,辅助计数
	temp := Q

	//添加数据后最后一个队列节点的next会指向头节点
	for temp.next != Q {
		count++
		temp = temp.next
	}
	return count
}

//显示队列
func (Q *CircleQueue_chainStructure) Show() {
	//长度为零说明队列中没有元素
	if Q.Len() == 0 {
		return
	}

	//定义一个中间变量,帮忙遍历队列
	temp := Q

	for temp.next != Q {
		if temp.next.next != Q {
			fmt.Printf("%d-->",temp.next.data)
		}else {
			fmt.Printf("%d\n",temp.next.data)
		}
		temp = temp.next
	}

}