//
// Created by 嘟嘟 on 2020/10/2.
//


package linkedList

import (
	"errors"
	"fmt"
)

//自定义数据类型,方便拓展
type DataType string

//双链表:double linked list
type DoubleLL struct {
	prior *DoubleLL		//指向前一个节点的指针
	data interface{}	//节点数据
	//其他字段
	next *DoubleLL		//指向后一个节点的指针
}

//初始化一条双链表(头节点)
func InitDoubleLL() *DoubleLL {
	return &DoubleLL{
		prior: nil,
		data:  nil,
		next:  nil,
	}
}

//新建一个双链表节点
func (D *DoubleLL) NewDbleLL(data interface{}) *DoubleLL {
	return &DoubleLL{
		prior: nil,
		data:  data,
		next:  nil,
	}
}

//往双链表头节点后面添加一个新的节点
func (D *DoubleLL) Insert_H_Rear(data interface{}) error {
	//新入链节点
	newNode := D.NewDbleLL(data)

	//中间变量,用于操作新节点的入链
	temp := D

	//当环形双链表为空,第一个节点入链是的操作
	if temp.next == nil {
		temp.next = newNode
		newNode.prior = temp
		return nil
	}

	//入链
	//第一步,将新节点的next指向原来环形双链表的第一个节点
	newNode.next = temp.next

	//第二步,将原来环形双链表的第一个节点的prior指向新节点
	temp.next.prior = newNode

	//第三步,将新节点的prior指向头节点
	newNode.prior = temp

	//第四步,将新节点入链
	temp.next = newNode

	return nil
}

//往双链表中第 N 个位置添加节点
func (D *DoubleLL) Insert(N int,data interface{}) error{
	//先判断N是否比链表的总长还大
	if N > D.Len() || N < 0{
		return errors.New("不合法的操作,请重试!")
	}

	//新节点
	newNode := D.NewDbleLL(data)

	//中间变量,用于遍历环形双链表
	temp := D

	//计数器,结合中间变量寻找第N个节点的前一个节点
	count := 0
	for count < N - 1{
		temp = temp.next
		count++
	}

	//入链
	//第一步,将新节点的next指向第N个节点
	newNode.next = temp.next

	//第二步,进第N个节点的prior指向新节点
	temp.next.prior = newNode

	//第三步,将新节点入链
	temp.next = newNode
	newNode.prior = temp

	return nil
}

//往双链表尾部添加节点
func (D *DoubleLL) Push(data interface{}) {

	newNode := D.NewDbleLL(data)

	//辅助节点,用于寻找链表末尾节点
	temp := D

	//寻找链表末尾
	for temp.next != nil {
		temp = temp.next
	}

	//将新节点添加到头节点为head的链表上
	temp.next = newNode
	newNode.prior = temp
}

//删除双链表头节点后的一个节点
func (D *DoubleLL) Delete_H_Rear() (interface{},error) {
	//判断环形双链表是否为空
	if D.next == nil {
		return nil, errors.New("当前双链表为空,请Push后重试!")
	}

	//中间变量,指向头节点,用于环形双链表首节点
	temp := D

	//环形双链表只有一个节点的请况
	if temp.next.next == nil {
		data := temp.next.data
		temp.next.prior = nil
		temp.next = nil
		return data, nil
	}

	//环形双链表节点大于一的请况
	data  := temp.next.data
	temp.next.next.prior = temp
	temp.next = temp.next.next
	return data, nil
}

//删除双链表尾部节点
func (D *DoubleLL) Delete() (interface{},error) {
	//判断环形双链表是否为空链表
	if D.next == nil {
		return nil, errors.New("当前双链表为空,请Push后重试!")
	}

	//中间变量,指向倒数第二个节点,用于删除环形双链表尾节点
	temp := D

	//寻找倒数第二个节点
	for temp.next.next != nil {
		temp = temp.next
	}

	//删除尾节点
	data := temp.next.data
	temp.next.prior = nil
	temp.next = nil

	return data, nil
}

//删除双链表中的第 N 个节点
func (D *DoubleLL) Delete_N(N int) (interface{},error) {
	//先判断N是否合法
	if N > D.Len() || N < 0 {
		return nil, errors.New("不合法的操作,请重试!")
	}

	//中间变量,指向倒数第N个节点的前一个节点,用于删除环形双链表中的第N个节点
	temp := D

	//计数器,结合中间变量,用于找第N - 1 个节点
	count := 0

	//寻找第N - 1 个节点
	for count < N - 1{		//找到第 N 个的前一个节点
		temp = temp.next
		count++
	}

	//将要删除的节点中的data数据保存,用于返回数据
	data := temp.next.data

	if temp.next.next == nil {		//说明删除的是最后的尾节点
		temp.next.prior = nil
		temp.next = nil
	}else {							//其他节点
		temp.next.next.prior = temp
		temp.next = temp.next.next
	}

	return data, nil
}

//获取双链表的长度
func (D *DoubleLL) Len() int {
	//中间变量,用于遍历环形双链表
	temp := D

	//计数器,结合中间变量,计算环形双量表长度
	count := 0

	//计算长度
	for temp.next != nil {
		temp = temp.next
		count++
	}

	return count
}

//判断双链表是否为空
func (D *DoubleLL) IsEmpty() bool {
	return D.next == nil
}

//显示双链表所有节点
func (D *DoubleLL) Show() {

	//判断环形双链表是否为空
	if D.IsEmpty() {
		fmt.Println("当前链表为空,请添加节点后再进行此操作!")
		return
	}

	//中间变量,用于遍历环形双链表,打印节点data数据
	temp := D

	fmt.Printf("链表信息如下:\n\t")
	for temp.next != nil {
		temp = temp.next
		if temp.next != nil {
			fmt.Printf("%s-->",temp.data)
		}else {
			fmt.Printf("%s\n",temp.data)
		}
	}

}