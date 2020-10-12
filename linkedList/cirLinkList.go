package linkedList

import (
	"errors"
	"fmt"
)

//环形链表:Circular linked Ljst

const (
	MaxSize = 15
)


//环形单链表
type CircularSingleLL struct {
	data interface{}			//节点数据
	next *CircularSingleLL		//下一个接节点
	// 其他字段
}



//初始化一个环形单链表
func InitCircularSingleLL() *CircularSingleLL {
	CircularSingleLL := &CircularSingleLL{
		data: nil,
		next: nil,
	}
	CircularSingleLL.next = CircularSingleLL
	return CircularSingleLL
}

//新建一个环形单链表节点
func (CL *CircularSingleLL) New(data interface{}) *CircularSingleLL {
	return &CircularSingleLL{
		data: data,
		next: nil,
	}
}

//添加节点到环形单链表首部
func (CL *CircularSingleLL) Push_head(data interface{}) error {
	//判断环形单链表是否已满
	if CL.IsFull() {
		return errors.New("当前环形单链表已满,请Pull后重试!")
	}

	//新入链节点
	node := CL.New(data)

	//将新节点指向头节点后面的那个节点
	node.next = CL.next

	//入链
	CL.next = node
	return nil
}

//添加节点到环形单链表第 N 个位置
func (CL *CircularSingleLL) Insert(data interface{}, N int) error {
	//判断N是否合法
	if N > MaxSize || N < 0 {
		return errors.New("您输入的位置有误,请重试!")
	}

	//判断环形单链表是否已满
	if CL.IsFull() {
		return errors.New("当前环形单链表已满,请Pull后重试!")
	}

	//新插入的节点
	node := CL.New(data)

	//中间变量,用于找到第 N 个位置前的一个节点
	temp := CL
	//技计数器
	count := 0

	//寻找第 N 个位置前的一个节点
	for count < N - 1 {
		count++
		temp = temp.next
	}

	//将新插入的节点的下一个指向即第 N 个节点
	node.next = temp.next

	//新节点入链
	temp.next = node
	return nil
}

//添加节点到环形单链表尾部
func (CL *CircularSingleLL) Push(data interface{}) error {
	//判断环形单链表是否已满
	if CL.IsFull() {
		return errors.New("当前环形单链表已满,请Pull后重试!")
	}

	//新建节点
	node := CL.New(data)

	//中间节点,用于寻找尾部节点
	temp := CL
	for temp.next != CL {
		temp = temp.next
	}

	//将新插入的节点的下一个指向头节点
	node.next = CL

	//新节点入链
	temp.next = node
	return nil
}

//获取环形单链表首部节点数据
func (CL *CircularSingleLL) Get_head() (interface{}, error) {
	if CL.next == CL {
		return nil, errors.New("当前环形单链表为空,请Push后重试!")
	}
	return CL.next.data, nil
}

//获取环形单链表第 N 个位置处的节点数据
func (CL *CircularSingleLL) Get_N(N int) (interface{}, error) {
	//判断N是否合法
	if N > MaxSize || N < 0 {
		return nil, errors.New("您输入的位置有误,请重试!")
	}

	//中间变量,用于找到第 N 个位置前的一个节点
	temp := CL
	//技计数器
	count := 0

	//寻找第 N 个位置处的节点
	for count < N {
		count++
		temp = temp.next
	}


	return temp.data, nil
}

//获取环形单链表尾部节点数据
func (CL *CircularSingleLL) Get() (interface{}, error) {
	//判断链表是否为空
	if CL.next == CL {
		return nil, errors.New("当前环形单链表为空,请Push后重试!")
	}

	//中间节点,用于寻找尾部节点
	temp := CL
	for temp.next != CL {
		temp = temp.next
	}

	return temp.data, nil
}

//删除环形单链表首部节点
func (CL *CircularSingleLL) Pull_head() (interface{}, error) {
	//判断环形单链表是否为空
	if CL.next == CL {
		return nil, errors.New("当前环形单链表为空,请Push后重试!")
	}

	//要删除的节点数据
	data := CL.next.data

	//删除节点
	CL.next = CL.next.next

	return data, nil
}

//删除环形单链表第 N 个位置处的节点
func (CL *CircularSingleLL) Pull_N(N int) (interface{}, error) {
	//判断N是否合法
	if N > MaxSize || N < 0 {
		return nil, errors.New("您输入的位置有误,请重试!")
	}

	if CL.next == CL {
		return nil, errors.New("当前环形单链表为空,请Push后重试!")
	}

	//中间变量,用于找到第 N 个位置前的一个节点
	temp := CL
	//技计数器
	count := 0

	//寻找第 N 个位置前的一个节点
	for count < N - 1 {
		count++
		temp = temp.next
	}

	data := temp.next.data

	//删除第N个节点
	temp.next = temp.next.next

	return data, nil
}

//删除环形单链表尾部节点
func (CL *CircularSingleLL) Pull() (interface{}, error) {
	//判断链表是否为空
	if CL.IsFull() {
		return nil, errors.New("当前环形单链表为空,请Push后重试!")
	}

	//中间节点,用于寻找尾部前一个节点
	temp := CL

	//因为要删除的是最后那个节点,所以找到倒数第二个节点才能进行删除
	for temp.next.next != CL {
		temp = temp.next
	}

	//要删除的节点的数据
	data := temp.next.data

	//删除节点
	temp.next = CL
	return data, nil
}

//判断环形单链表是否为空
func (CL *CircularSingleLL) IsEmpty() bool {
	return CL.next == CL
}

//判断环形单链表是否已满
func (CL *CircularSingleLL) IsFull() bool {
	count := 0
	temp := CL
	for temp.next != CL {
		temp = temp.next
		count++
	}
	return count == MaxSize
}

//获取当前环形单链表的长度
func (CL * CircularSingleLL) Len() int {
	//计数器
	count := 0

	//中间变量,用于遍历单链表
	temp := CL

	for temp.next != CL {
		count++
		temp = temp.next
	}

	return count
}

//显示环形单链表所有节点的数据
func (CL *CircularSingleLL) SHow() error {
	//判断环形单链表是否为空
	if CL.IsEmpty() {
		return errors.New("当前环形单链表为空,请Push后重试!")
	}

	//辅助变量,用于遍历链表
	temp := CL

	//只要temp的下一个节点不是头节点,就进行循环
	for temp.next != CL {
		if temp.next.next != CL {
			fmt.Printf("%d-->",temp.next.data)
		}else {
			fmt.Printf("%d\n",temp.next.data)
		}
		temp = temp.next
	}

	return nil

}





//环形双链表
type CircularDoubleLL struct {
	prior *CircularDoubleLL		//前一个节点
	data interface{}			//节点数据
	next *CircularDoubleLL		//下一个节点
}

//初始化一个环形双链表
func InitCircularDoubleLL() *CircularDoubleLL {
	CircularDoubleLL := &CircularDoubleLL{
		prior: nil,
		data:  nil,
		next:  nil,
	}
	CircularDoubleLL.next = CircularDoubleLL
	CircularDoubleLL.prior = CircularDoubleLL
	return CircularDoubleLL
}

//新建一个环形双链表节点
func (CDL *CircularDoubleLL) New(data interface{}) *CircularDoubleLL {
	return &CircularDoubleLL{
		prior: nil,
		data:  data,
		next:  nil,
	}
}

//添加节点到环形双链表首部
func (CDL *CircularDoubleLL) Push_head(data interface{}) error {
	//判断环形双链表是否已满
	if CDL.IsFull() {
		return errors.New("当前环形双链表已满,请Pull后重试!")
	}

	//新入链节点
	node := CDL.New(data)

	//新节点入链
	//第一步,将新节点的next指向首节点
	node.next = CDL.next

	//第二步,将首节点的prior指向新节点
	CDL.next.prior = node

	//第三步,将新节点的prior指向头节点
	node.prior = CDL

	//第四步,将头节点的next指向新节点
	CDL.next = node

	return nil
}

//添加节点到环形双链表第 N 个位置
func (CDL *CircularDoubleLL) Insert(data interface{}, N int) error {
	//判断N是否合法
	if N > MaxSize || N < 0 {
		return errors.New("您输入的位置有误,请重试!")
	}

	//判断环形双链表是否已满
	if CDL.IsFull() {
		return errors.New("当前环形双链表已满,请Pull后重试!")
	}

	//新插入的节点
	node := CDL.New(data)

	//中间变量,用于找到第 N 个位置前的一个节点
	temp := CDL
	//技计数器
	count := 0

	//寻找第 N 个位置前的一个节点
	for count < N - 1 {
		count++
		temp = temp.next
	}


	//新节点入链
	//第一步,将新节点的next指向第N个节点
	node.next = temp.next

	//第二步,将第N个节点的prior指向新节点
	temp.next.prior = node

	//第三步,将新节点的prior指向 第N个节点的前一个节点
	node.prior = temp

	//第四步,将新节点入链,即将第N个节点的前一个节点的next指向新节点
	temp.next = node

	return nil
}

//添加节点到环形双链表尾部
func (CDL *CircularDoubleLL) Push(data interface{}) error {
	//判断环形双链表是否已满
	if CDL.IsFull() {
		return errors.New("当前环形双链表已满,请Pull后重试!")
	}

	//新建节点
	node := CDL.New(data)

	//新节点入链
	//第一步,将新节点的next指向头节点
	node.next = CDL

	//第二步,将为节点的next指向新节点
	CDL.prior.next = node

	//第三步,将新节点的prior指向尾节点
	node.prior = CDL.prior

	//第四步,将头节点的prior指向新节点
	CDL.prior = node

	return nil
}

//获取环形双链表首部节点数据
func (CDL *CircularDoubleLL) Get_head() (interface{}, error) {
	if CDL.next == CDL {
		return nil, errors.New("当前环形双链表为空,请Push后重试!")
	}
	return CDL.next.data, nil
}

//获取环形双链表第 N 个位置处的节点数据
func (CDL *CircularDoubleLL) Get_N(N int) (interface{}, error) {
	//判断N是否合法
	if N > MaxSize || N < 0 {
		return nil, errors.New("您输入的位置有误,请重试!")
	}

	//中间变量,用于找到第 N 个位置前的一个节点
	temp := CDL
	//技计数器
	count := 0

	//寻找第 N 个位置处的节点
	for count < N {
		count++
		temp = temp.next
	}

	return temp.data, nil
}

//获取环形双链表尾部节点数据
func (CDL *CircularDoubleLL) Get() (interface{}, error) {
	//判断链表是否为空
	if CDL.next == CDL {
		return nil, errors.New("当前环形双链表为空,请Push后重试!")
	}

	return CDL.prior.data, nil
}

//删除环形双链表首部节点
func (CDL *CircularDoubleLL) Pull_head() (interface{}, error) {
	//判断链表是否为空
	if CDL.next == CDL {
		return nil, errors.New("当前环形双链表为空,请Push后重试!")
	}

	//要删除的节点数据
	data := CDL.next.data

	//删除节点
	CDL.next.next.prior = CDL
	CDL.next = CDL.next.next

	return data, nil
}

//删除环形双链表第 N 个位置处的节点
func (CDL *CircularDoubleLL) Pull_N(N int) (interface{}, error) {
	//判断N是否合法
	if N > MaxSize || N < 0 {
		return nil, errors.New("您输入的位置有误,请重试!")
	}

	if CDL.next == CDL {
		return nil, errors.New("当前环形双链表为空,请Push后重试!")
	}

	//中间变量,用于找到第 N 个位置的前的一个节点
	temp := CDL
	//技计数器
	count := 0

	//寻找第 N 个位置前的一个节点
	for count < N - 1 {
		count++
		temp = temp.next
	}

	//要删除的节点中的数据
	data := temp.next.data

	//删除第N个节点
	CDL.next.next.prior = CDL
	CDL.next = CDL.next.next

	return data, nil
}

//删除环形双链表尾部节点
func (CDL *CircularDoubleLL) Pull() (interface{}, error) {
	//判断双链表是否为空
	if CDL.IsFull() {
		return nil, errors.New("当前环形双链表为空,请Push后重试!")
	}

	//要删除的节点的数据
	data := CDL.prior.data

	//删除节点
	CDL.prior = CDL.prior.prior
	CDL.prior.next = CDL


	return data, nil
}

//判断环形双链表是否为空
func (CDL *CircularDoubleLL) IsEmpty() bool {
	return CDL.next == CDL || CDL.prior == CDL
}

//判断环形双链表是否已满
func (CDL *CircularDoubleLL) IsFull() bool {
	//计数器
	count := 0

	//中间变量,用于遍历环形双链表
	temp := CDL

	for temp.next != CDL {
		temp = temp.next
		count++
	}

	return count == MaxSize
}

//获取当前环形双链表的长度
func (CDL * CircularDoubleLL) Len() int {
	//计数器
	count := 0

	//中间变量,用于遍历环形双链表
	temp := CDL

	for temp.next != CDL {
		count++
		temp = temp.next
	}

	return count
}

//显示环形双链表所有节点中的data数据
func (CDL *CircularDoubleLL) SHow() error {
	//判断链表是否为空
	if CDL.IsEmpty() {
		return errors.New("当前环形单链表为空,请Push后重试!")
	}

	//辅助变量,用于遍历链表
	temp := CDL

	//只要temp的下一个节点不是头节点,就进行循环
	for temp.next != CDL {
		if temp.next.next != CDL {
			fmt.Printf("%d-->",temp.next.data)
		}else {
			fmt.Printf("%d\n",temp.next.data)
		}
		temp = temp.next
	}

	return nil
}