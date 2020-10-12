package linkedList

import (
	"fmt"
)

//测试单链表的各个功能
func Test_sgleLinkList() {
	//初始化一个头节点
	head := InitSingleLL()

	//新建一个节点并添加到头节点为 head 的链表中
	head.Push("嘟嘟一号")

	//同上
	head.Push("嘟嘟二号")

	//同上
	head.Push("嘟嘟三号")

	//同上
	head.Push("嘟嘟四号")

	//新建一个节点并添加到头节点后面
	//newNode.Push(head,newNode)		//注: 同一节点不能同时添加到链表尾部和首部,否则将变成环状单链表
	head.Insert_H_Rear("嘟嘟五号")

	//同上
	head.Insert(3,"嘟嘟六号")
	//fmt.Println(newNode.Insert(head,9,newNode))

	//嘟嘟五号-->嘟嘟一号-->嘟嘟六号-->嘟嘟二号-->嘟嘟三号-->嘟嘟四号

	head.Delete_H_Rear()
	head.Delete()
	head.Delete_N(2)

	//以上步骤:
	//先将一号至四号依次添加到单链表尾部,然后将五号插入头节点后面,(也就是首部)
	//再将六号插入第三个位置,最后删除链表尾部节点和首节点以及第二个位置的节点
	//最后链表为: 嘟嘟一号-->嘟嘟二号-->嘟嘟三号

	fmt.Printf("当前单链表长度:%d\n",head.Len())
	head.Show()
}

//测试双链表的各个功能
func Tese_doubleLinkList() {
	head := InitDoubleLL()

	head.Push("冰糖一号")

	head.Push("冰糖二号")

	head.Push("冰糖三号")

	head.Push("冰糖四号")

	head.Insert_H_Rear("冰糖五号")

	err := head.Insert(2,"冰糖六号")
	err = head.Insert(20,"冰糖六号")
	//冰糖五号-->冰糖六号-->冰糖一号-->冰糖二号-->冰糖三号-->冰糖四号
	fmt.Println(err)
	head.Delete_H_Rear()
	head.Delete()
	head.Delete_N(1)

	fmt.Printf("当前单链表长度:%d\n",head.Len())
	head.Show()

	//以上步骤:
	//先将一号至四号依次添加到单链表尾部,然后将五号插入头节点后面,(也就是首部)
	//再将六号插入第二个位置,最后删除链表尾部节点和首节点以及第一个位置的节点
	//最后链表为: 冰糖一号-->冰糖二号-->冰糖三号
}


//测试环形单链表的各个功能
func Tese_CircularSgleLinkList() {
	csll := InitCircularSingleLL()

	for i := 3; i < MaxSize; i++ {
		csll.Push(i + 1)
	}

	csll.Push_head(2)
	csll.Insert(3,2)
	csll.Push(16)

	fmt.Println(csll.Get_head())
	fmt.Println(csll.Get_N(3))
	fmt.Println(csll.Get())
	//fmt.Println(csll.Push("c"))
	csll.SHow()
	fmt.Println(csll.Len())


	fmt.Println(csll.Pull_head())
	fmt.Println(csll.Pull_N(3))
	fmt.Println(csll.Pull())
	//fmt.Println(csll.Push("c"))
	csll.SHow()
	fmt.Println(csll.Len())
}

//测试环形双链表的各个功能
func Tese_CircularDoubleLinkList() {
	csll := InitCircularDoubleLL()

	for i := 3; i < MaxSize; i++ {
		csll.Push(i + 1)
	}

	csll.Push_head(2)
	csll.Insert(3,2)
	csll.Push(16)

	fmt.Println(csll.Get_head())
	fmt.Println(csll.Get_N(3))
	fmt.Println(csll.Get())
	//fmt.Println(csll.Push("c"))
	csll.SHow()
	fmt.Println(csll.Len())


	fmt.Println(csll.Pull_head())
	fmt.Println(csll.Pull_N(3))
	fmt.Println(csll.Pull())
	//fmt.Println(csll.Push("c"))
	csll.SHow()
	fmt.Println(csll.Len())}
