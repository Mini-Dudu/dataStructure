//
// Created by 嘟嘟 on 2020/10/4.
//



//队列:queue
//顺序表:sequence list
package queue

import "errors"

//begin:开始实现顺序存储结构的一次性队列及其相关核心操作

//缺点:此队列为一次性队列,即满后就算全部取出也不能在使用当前队列,可重新New一个队列使用或者使用优化队列或者链式存储结构队列
//优化:使用循环对列,链式存储结构队列
type Queue_sequential struct {		//sequential:顺序结构
	queue [MaxSize]interface{}		//队列数据
	head int						//队头
	tail int    					//队尾		Tail:尾,尾巴,圣少女,尾部
}

const (
	MaxSize = 5		//实际上只能存MaxSize - 1 个数据
)

//新建一个队列
func NewQueue_sequential() *Queue_sequential {
	return &Queue_sequential{
		queue: [MaxSize]interface{}{},
		head:  -1,
		tail:  0,
	}
}

//入队列
func (Q *Queue_sequential) Push(data interface{}) error {
	if Q.IsFull() {
		return errors.New("队列已满!")
	}
	Q.queue[Q.tail] = data
	Q.tail++
	return nil
}

//获取队尾数据
func (Q *Queue_sequential) Get() (interface{},error) {
	if Q.IsEmpty() {
		return nil,errors.New("当前队列为空,请Push后重试!")
	}
	return Q.queue[Q.tail - 1], nil
}

//出队列
func (Q *Queue_sequential) Pull() (interface{},error) {
	if Q.IsEmpty() {
		return nil,errors.New("队列为空!")
	}
	Q.head++
	return Q.queue[Q.head],nil
}

//判断队列是否为空
func (Q *Queue_sequential) IsEmpty() bool {
	//仅仅判断q.tatil 是否等于 -1 是不可去取的,因为如果存入一个后马上取出队列也为空
	//return q.tail == -1
	return Q.head + 1 == Q.tail
}

//判断队列是否已满
func (Q *Queue_sequential) IsFull() bool {
	//因为不是循环队列所以只需要判断q.head是否等于最大容量减一即可
	return Q.tail == MaxSize - 1
}

//查看队列当前长度
func (Q *Queue_sequential) Len() int {
	return Q.tail - Q.head - 1
}

//end:顺序存储结构的一次性队列及其相关核心操作的实现已完成






//begin:开始实现链式存储结构队列及其相关核心操作


//说明:此链式结构队列带有头节点且以头节点那端为队头,以链表尾部为队尾
type Queue_chainStructure struct {				//chain structure: 链式结构
	data interface{}				//链表数据
	next *Queue_chainStructure		//下一个元素
	tail *Queue_chainStructure		//队尾
}

//初始化一个链式队列
func InitQueue_chainStructure() *Queue_chainStructure {
	return &Queue_chainStructure{
		data: nil,
		next: nil,
		tail: nil,
	}
}

//新建队列节点
func (Q *Queue_chainStructure) New(data interface{}) *Queue_chainStructure {

	return &Queue_chainStructure{
		data: data,
		next: nil,
		tail: nil,
	}
}

//入队列
func (Q *Queue_chainStructure) Push(data interface{}) {
	newNode := Q.New(data)
	if Q.IsEmpty() {
		Q.tail = Q
		Q.next = newNode
		return
	}
	Q.tail.next.next = newNode
	Q.tail = Q.tail.next
}

//出队列,将新节点添加到队列尾部,即链表尾部
func (Q *Queue_chainStructure) Pull() (interface{}, error) {
	if Q.IsEmpty() {
		return nil, errors.New("当前队列为空,请Push后重试!")
	}
	data := Q.next.data
	Q.next = Q.next.next
	return data,nil
}

//获得队头元素数据
func (Q *Queue_chainStructure) Get() (interface{},error) {
	if Q.IsEmpty() {
		return nil, errors.New("当前队列为空,请Push后重试!")
	}
	return Q.next.data, nil
}

//判断是否为空
func (Q *Queue_chainStructure) IsEmpty() bool {
	return Q.next == nil
}

//由于链式队列的链式存储结构,可以没有最大容量,这里就不写判断是否满了的方法

//获取链式队列的长度
func (Q *Queue_chainStructure) Len() int {
	var count int
	temp := Q
	for temp.next != nil {
		count++
		temp = temp.next
	}
	return count
}


//end:链式存储结构队列及其核心相关操作的实现已完成