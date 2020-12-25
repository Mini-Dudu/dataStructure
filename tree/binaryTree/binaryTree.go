package binaryTree

import (
	"DataStruct/queue"
	"DataStruct/stack"
	"fmt"
)

type dataType int

const (
	L = 0
	R = 1
)

//满二叉树:高度为h,且含有2的h次方-1个节点的二叉树为满二叉树。
//BinaryTree:二叉树
type BiTree struct {
	//……
	left, right *BiTree  //左孩子与右孩子
	data        dataType //数据
}

////构造空二叉树T   空二叉树:一个结点都没有的二叉树
//func InitBITree() *BiTree {
//	//暂时不知道如何构建没有节点的二叉树
//	return nil
//}

//新建一个二叉树节点
func New(data dataType) *BiTree {
	return &BiTree{data: data}
}

//销毁二叉树-----利用清空模拟
func (t *BiTree) DestroyBiTree() bool {
	//利用清空模拟销毁操作
	t.data = dataType(0)
	t.left = nil
	t.right = nil
	return true
}

////按definition构造二叉树T
//func (t BiTree) CreateBiTree(T BiTree) () {
//
//}

//将二叉树t清空
func (t *BiTree) CliearBiTree() () {
	t.data = dataType(0)
	t.left = nil
	t.right = nil
}

//判断二叉树T是否为空   空二叉树:一个结点都没有的二叉树
func (t *BiTree) IsBiTreeEmpty() bool {
	//能调用此方法就说明树不为空
	return false
}

//返回二叉树T的深度
func (t *BiTree) BiTreeDepth() (depth dataType) {
	//初始化一个队列
	queue := queue.InitCircleQueue_sequential()
	queue.Push(t)
	for !queue.IsEmpty() {
		//深度加一
		depth++
		//获取当前层的宽度
		width := queue.Len()
		for i := 0; i < width; i++ {
			//出队一个元素
			data, _ := queue.Pull()
			value, _ := data.(*BiTree)
			if value.left != nil { //左孩子入队
				queue.Push(value.left)
			}
			if value.right != nil { //右孩子入队
				queue.Push(value.right)
			}
		}
	}
	return
}

//返回二叉树T的根节点
func (t *BiTree) Root() *BiTree {
	//默认调用者t就是根节点
	return t
}

//返回节点e的值
func (t *BiTree) Value(e *BiTree) *dataType {
	//遍历一遍树t,找到e节点,然后返回e.Data
	//中序非递归遍历
	temp := t
	stack := stack.InitChainStack() //初始化一个栈
	for temp != nil {               //初始状态时将根节点及左边所有节点入栈
		stack.Push(temp)
		temp = temp.left
	}
	temp_e := stack.Pull() //出栈一个节点,并访问它
	tempValue, ok := temp_e.(*BiTree)
	for ok {

		if tempValue == e { //在树t中找到e节点
			return &tempValue.data
		}
		//扫描该节点的右子节点,并将它入栈;
		if tempValue.right != nil {
			stack.Push(tempValue.right)
			temp = tempValue.right
			for temp.left != nil { //依次扫描该右子节点的所有左侧节点并一一入栈;
				stack.Push(temp.left)
				temp = temp.left
			}
		}
		temp_e = stack.Pull()
		tempValue, ok = temp_e.(*BiTree)
	}
	return nil
}

//为节点e赋值
func (t *BiTree) Assign(e *BiTree, value dataType) bool {
	//遍历一遍树t,找到e节点,然后为e节点赋值,并返回赋值是否成功
	//中序非递归遍历
	temp := t
	stack := stack.InitChainStack() //1.初始化一个栈
	for temp != nil {               //2.初始状态时将根节点及左边所有节点入栈
		stack.Push(temp)
		temp = temp.left
	}
	temp_e := stack.Pull() //3.出栈一个节点,并访问它
	tempValue, ok := temp_e.(*BiTree)
	for ok {
		if tempValue == e { //找到e节点,并未e节点赋值
			tempValue.data = value
			return true
		}
		//4.扫描该节点的右子节点,并将它入栈;
		if tempValue.right != nil {
			stack.Push(tempValue.right)
			temp = tempValue.right
			for temp.left != nil { //5.依次扫描该右子节点的所有左侧节点并一一入栈;
				stack.Push(temp.left)
				temp = temp.left
			}
		}
		temp_e = stack.Pull()
		tempValue, ok = temp_e.(*BiTree)
	}
	return false
}

//若e是T的非根节点,返回e的双亲节点,否则返回nil
func (t *BiTree) Parent(e *BiTree) *BiTree {
	if e == t {
		return nil
	}
	return t.SearchParent(e)
}

//若存在,则返回树t中节点e的左孩子,否则返回nil
func (t *BiTree) LeftChild(e *BiTree) *BiTree {
	//遍历一遍树t,找到e节点,并返回e的左孩子
	//中序非递归遍历
	temp := t
	stack := stack.InitChainStack() //1.初始化一个栈
	for temp != nil {               //2.初始状态时将根节点及左边所有节点入栈
		stack.Push(temp)
		temp = temp.left
	}
	temp_e := stack.Pull() //3.出栈一个节点,并访问它
	tempValue, ok := temp_e.(*BiTree)
	for ok {
		if tempValue == e { //找到e节点,并返回其左孩子
			return tempValue.left
		}
		//4.扫描该节点的右子节点,并将它入栈;
		if tempValue.right != nil {
			stack.Push(tempValue.right)
			temp = tempValue.right
			for temp.left != nil { //5.依次扫描该右子节点的所有左侧节点并一一入栈;
				stack.Push(temp.left)
				temp = temp.left
			}
		}
		temp_e = stack.Pull()
		tempValue, ok = temp_e.(*BiTree)
	}
	return nil
}

//若存在,则返回树t中节点e的右孩子,否则返回nil
func (t *BiTree) RightChild(e *BiTree) *BiTree {
	//遍历一遍树t,找到e节点,并返回e的左孩子
	//中序非递归遍历
	temp := t
	stack := stack.InitChainStack() //1.初始化一个栈
	for temp != nil {               //2.初始状态时将根节点及左边所有节点入栈
		stack.Push(temp)
		temp = temp.left
	}
	temp_e := stack.Pull() //3.出栈一个节点,并访问它
	tempValue, ok := temp_e.(*BiTree)
	for ok {
		if tempValue == e { //找到e节点,并返回其右孩子
			return tempValue.right
		}
		//4.扫描该节点的右子节点,并将它入栈;
		if tempValue.right != nil {
			stack.Push(tempValue.right)
			temp = tempValue.right
			for temp.left != nil { //5.依次扫描该右子节点的所有左侧节点并一一入栈;
				stack.Push(temp.left)
				temp = temp.left
			}
		}
		temp_e = stack.Pull()
		tempValue, ok = temp_e.(*BiTree)
	}
	return nil
}

//若存在,则返回树t中节点e的左兄弟,否则返回nil
func (t *BiTree) LeftSibling(e *BiTree) *BiTree {
	if e == t {
		return nil
	}

	//找到e节点的双亲节点
	parent := t.SearchParent(e)
	if parent != nil { //e节点有可能不在树t中,此时parent为空
		l := parent.left
		//r := parent.right
		if e != l { //判断e节点是否为parent的左孩子
			return l
		}
	}

	return nil
}

//若存在,则返回树t中节点e的右兄弟,否则返回nil
func (t *BiTree) RightSibling(e *BiTree) *BiTree {
	if e == t {
		return nil
	}

	//找到e节点的双亲节点
	parent := t.SearchParent(e)
	if parent != nil { //e节点有可能不在树t中,此时parent为空
		//l := parent.left
		r := parent.right
		if e != r {
			return r
		} //判断e节点是否为parent的左孩子
	}

	return nil
}

//根据lr为0或1,,将c插入为e所指节点的左或右子树,并返回是否插入成功
func (t *BiTree) InsertChild(e, c *BiTree, lr int) bool {
	//fmt.Println("执行插入孩子节点")
	if lr != L && lr != R {
		return !(lr != L || lr != R)
	}
	//思路:利用非递归遍历算法在树t中找到e节点后,进行相关操作
	queue := queue.InitCircleQueue_sequential()
	queue.Push(t)
	data, _ := queue.Pull()
	value, ok := data.(*BiTree)

	for ok {
		if value == e {
			if lr == L {
				//log.Println("插入左孩子成功")
				value.left = c
				return true
			} else {
				//log.Println("插入右孩子成功")
				value.right = c
				return true
			}
		}
		if value.left != nil {
			queue.Push(value.left)
		}
		if value.right != nil {
			queue.Push(value.right)
		}
		data, _ = queue.Pull()
		value, ok = data.(*BiTree)

	}
	return false
}

//根据lr为0或1,删除e节点的左或右子树,并返回是否删除成功
func (t *BiTree) DeleteChild(e *BiTree, lr int) bool {
	if lr != L && lr != R {
		return !(lr != L || lr != R)
	}

	//思路:利用非递归遍历算法在树t中找到e节点后,进行相关操作
	queue := queue.InitCircleQueue_sequential()
	queue.Push(t)
	data, _ := queue.Pull()
	value, ok := data.(*BiTree)

	for ok {
		if value == e {
			if lr == L {
				value.left = nil
				return true
			} else {
				value.right = nil
				return true
			}
		}
		if value.left != nil {
			queue.Push(value.left)
		}
		if value.right != nil {
			queue.Push(value.right)
		}
		data, _ = queue.Pull()
		value, ok = data.(*BiTree)

	}
	return false
}

//二叉树的递归遍历算法：
//二叉树的先序遍历、中序遍历、后续遍历，都是相对于根节点而言的，在再行左优先原则
//先序遍历：先访问根节点，再访问左子树，最后访问右子树，其余两个同理

//先序遍历T
func (t *BiTree) PreOrderTraverse() {
	temp := t
	Visit(temp)
	if temp.left != nil {
		temp.left.PreOrderTraverse()
	}
	if temp.right != nil {
		temp.right.PreOrderTraverse()
	}
}

//中序遍历T----递归算法
func (t *BiTree) InOrderTraverse() {
	temp := t

	if temp.left != nil {
		temp.left.InOrderTraverse()
	}
	Visit(temp)
	if temp.right != nil {
		temp.right.InOrderTraverse()
	}
}

//后序遍历T----递归算法
func (t *BiTree) PostOrderTraverse() {
	temp := t

	if temp.left != nil {
		temp.left.PostOrderTraverse()
	}
	if temp.right != nil {
		temp.right.PostOrderTraverse()
	}
	Visit(temp)
}

/*二叉树的层次遍历:
算法思想:借助队列
	1)初始将根入队列并访问,然后出队;
	2)若有左子树,则将左子树的根入队;
	3)若有右子树,则将右子树的根入队;
	4)然后出队;并访问该节点;
	5)重复以上步骤,直到队列为空;
*/
//层次遍历T----非递归算法
func (t *BiTree) LevelOrderTraverse() {
	queue := queue.InitCircleQueue_sequential()
	queue.Push(t)
	data, _ := queue.Pull()
	value, ok := data.(*BiTree)

	for ok {

		Visit(value)

		if value.left != nil {
			queue.Push(value.left)
		}

		if value.right != nil {
			queue.Push(value.right)
		}

		data, _ = queue.Pull()

		value, ok = data.(*BiTree)

	}
}

//访问e
func Visit(e *BiTree) {
	fmt.Printf("%v\t", e.data)
}

/*二叉树的中序非递归遍历:
算法思想:借助栈
	1)初始时依次扫描根节点所有左侧节点(树中所有最左边的节点)并将它们一一进栈;
	2)出栈一个节点,并访问它;
	3)扫描该节点的右子节点,并将它入栈;
	4)依次扫描右子节点的所有左侧节点并一一入栈;
	5)重复以上步骤,直到栈为空;
其他同理!
*/
//中序遍历非递归算法
//利用根节点t,寻找某个节点,返回其双亲节点
func (t *BiTree) SearchParent(e *BiTree) *BiTree {
	temp := t
	stack := stack.InitChainStack()
	for temp != nil {
		stack.Push(temp)
		temp = temp.left
	}
	temp_e := stack.Pull()
	tempValue, ok := temp_e.(*BiTree)
	for ok {
		if tempValue.left == e || tempValue.right == e {
			return tempValue
		}
		if tempValue.right != nil {
			stack.Push(tempValue.right)
			temp = tempValue.right
			for temp.left != nil {
				stack.Push(temp.left)
				temp = temp.left
			}
		}
		temp_e = stack.Pull()
		tempValue, ok = temp_e.(*BiTree)
	}
	return nil
}

//先序遍历非递归算法
//利用根节点t,寻找某个节点,返回其孩子节点
func (t *BiTree) SearchChild(e *BiTree) (l, r *BiTree) {
/*思路:借助栈
	1)初始时依次扫描根节点及所有左侧节点(树中所有最左边的节点)并将它们一一进栈,并进行局部访问(访问data数据);
	2)出栈一个节点,并进行局部访问(访问right字段);
	3)扫描该节点的右子节点,将它入栈并进行局部访问(访问data数据及left字段);
	4)依次扫描右子节点的所有左侧节点并一一入栈;
	5)重复以上步骤,直到栈为空;
 */
	temp := t
	stack := stack.InitChainStack()
	for temp != nil {
		stack.Push(temp)
		//访问data数据
		Visit(temp)
		temp = temp.left
	}
	temp_e := stack.Pull()
	tempValue, ok := temp_e.(*BiTree)
	for ok {
		if tempValue == e {
			return tempValue.left, tempValue.right
		}
		if tempValue.right != nil {
			stack.Push(tempValue.right)
			//访问data数据
			Visit(tempValue.right)
			temp = tempValue.right
			for temp.left != nil {
				stack.Push(temp.left)
				temp = temp.left
				//访问data数据
				Visit(temp)
			}
		}
		temp_e = stack.Pull()
		tempValue, ok = temp_e.(*BiTree)
	}
	return nil, nil
}


//测试git分支提交
