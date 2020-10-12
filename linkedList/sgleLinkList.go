package linkedList

import (
	"errors"
	"fmt"
)

//单链表: single linked list
type SingleLL struct {
	//id int				//链表节点编号
	data interface{}	//节点数据
	//其他字段
	next *SingleLL		//后继节点的指针
}

//初始化一个单链表头节点
func InitSingleLL() *SingleLL {
	return &SingleLL{
		data: nil,
		next: nil,
	}
}

//新建一个单链表节点
func (S *SingleLL) NewSingleLL(data interface{}) *SingleLL {
	return &SingleLL{
		//id:   0,
		data: data,
		next: nil,
	}
}

//往单链表头节点后面添加一个新的节点
func (S *SingleLL) Insert_H_Rear(data interface{}) error {
	//要入链的新节点
	newNode := S.NewSingleLL(data)

	//先将新节点的next指向头节点的下一个节点
	newNode.next = S.next

	//再将新节点入链
	S.next = newNode

	return nil
}

//往单链表中第 N 个位置添加节点
func (S *SingleLL) Insert(N int,data interface{}) error{
	//先判断N是否合法
	if N > S.Len() || N < 0{
		return errors.New("不合法的操作,请重试!")
	}

	//中间变量,用于寻找第N - 1 个节点
	temp := S

	//计数器,结合中间变量,寻找第 N - 1 个节点
	count := 0

	//寻找第 N - 1 个节点
	for count < N - 1{
		temp = temp.next
		count++
	}

	//新入链节点
	newNode := S.NewSingleLL(data)

	//入链
	newNode.next = temp.next
	temp.next = newNode

	return nil
}

//往单链表尾部添加节点
func (S *SingleLL) Push(data interface{}) error {
	//辅助节点
	temp := S

	//寻找链表末尾
	for temp.next != nil {
		temp = temp.next
	}

	//新入链节点
	newNode := S.NewSingleLL(data)

	//将新节点添加到头节点为head的链表上
	temp.next = newNode

	return nil
}

//删除单链表头节点后的一个节点
func (S *SingleLL) Delete_H_Rear() (interface{}, error) {
	//判断单链表是否为空
	if S.IsEmpty() {
		return nil, errors.New("当前单链表为空,请Push后重试!")
	}

	//要删除的节点的data数据,用于返回
	data := S.data

	//删除节点
	S.next = S.next.next

	return data, nil
}

//删除单链表尾部节点
func (S *SingleLL) Delete() (interface{}, error) {
	//判断单链表是否为空
	if S.IsEmpty() {
		return nil, errors.New("当前单链表为空,请Push后重试!")
	}

	//辅助节点,用于寻找倒数第二个节点
	temp := S

	//寻找倒数第二个节点
	for temp.next.next != nil {
		temp = temp.next
	}

	//尾节点的data数据
	data := temp.next.data

	//删除尾节点
	temp.next = nil

	return data, nil
}

//删除单链表中的第 N 个节点
func (S *SingleLL) Delete_N(N int) (interface{},error) {
	//先判断N是否合法
	if N > S.Len() || N < 0 {
		return nil,errors.New("不合法的操作,请重试!")
	}

	//辅助节点,用于寻找第 N- 1 个节点
	temp := S

	//计数器,结合辅助节点,寻找第 N- 1 个节点
	count := 0

	//寻找第 N- 1 个节点
	for count < N - 1{
		temp = temp.next
		count++
	}

	//第N个节点的data数据
	data := temp.next.data

	//删除第N个节点
	temp.next = temp.next.next

	return data,nil
}

//获取单链表的长度
func (S *SingleLL) Len() int {
	//辅助变量,结合count计数器计算单链表的长度
	temp := S
	count := 0

	//计算单链表的长度
	for temp.next != nil {
		temp = temp.next
		count++
	}

	return count
}

//判断单链表是否为空
func (S *SingleLL) IsEmpty() bool {
	return S.next == nil
}

//显示链表所有节点
func (S *SingleLL) Show() {

	if S.IsEmpty() {
		fmt.Println("当前链表为空,请添加节点后再进行此操作!")
		return
	}
	temp := S
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