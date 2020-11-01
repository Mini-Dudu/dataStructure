//
// Created by 嘟嘟 on 2020/10/4.
//



package queue

import "fmt"

//测试顺序存储结构的队列的各个功能
func Test_Queue_sequential() {
	//以下测试数据将MaxSize设为5

	q := NewQueue_sequential()
	a := [5]int{1,2,3,4,5}
	err := q.Push(a)
	err = q.Push("嘟嘟")
	err = q.Push("冰糖")
	err = q.Push("葫芦")
	//满了后再次尝试入队
	err = q.Push("你是不是傻!")
	fmt.Println("队列满了后再次尝试入队:",err)
	data,err := q.Pull()
	fmt.Println("第一次取出队列数据:",data)

	len := q.Len()
	fmt.Println("第一次取出队列数据后队列长度:",len)
	data,err = q.Pull()
	data,err = q.Pull()
	data,err = q.Pull()

	data,err = q.Pull()
	fmt.Println("全部数据取出后再次尝试取数据:",data)


	fmt.Println("取出全部数据后再次尝试出队:",err)

	fmt.Println("全部数据取出后队列的长度:",q.Len(),"是否为空:",q.IsEmpty())


}


//测试链式存储结构的队列的各个功能
func Test_Queue_chainStructure() {
	q := InitQueue_chainStructure()
	q.Push("第一个入队")
	q.Push("第二个入队")
	q.Push("第三个入队")
	q.Push("第四个入队")
	q.Push("第五个入队")
	fmt.Println("Push五次后队列的长度:",q.Len())

	data, _ := q.Pull()
	fmt.Println("第一个出队元素:",data)

	data, _ = q.Pull()
	fmt.Println("第二个出队元素:",data)

	data, _ = q.Pull()
	data, _ = q.Pull()
	data, _ = q.Pull()
	fmt.Println("第五个出队元素:",data)

	data, err := q.Pull()
	fmt.Println("第六个出队元素:",data,"错误信息:",err)
}

//测试顺序存储结构的环形队列的各个功能
func Test_CircleQueue_sequential() {
	cq := InitCircleQueue_sequential()

	cq.Push("hahahhaha")
	cq.Push("hahahhaha")
	//fmt.Println(cq.Pull())
	fmt.Println(cq.Len())
	cq.Pull()
	cq.Pull()
	cq.Pull()
	fmt.Println(cq.Len())
	fmt.Println(cq.Push("hahahhaha"))
	fmt.Println(cq.Len())

}

//测试链式结构的环形队列的各个功能
func Test_CircleQueue_chainStructure() {
	q := InitCircleQueue_chainStructure()
	for i := 1; i <= 20; i++ {
		q.Push(i)
	}
	q.Show()

	fmt.Println(q.Push(12))
	q.Show()

	for i := 0; i < 20; i++ {
		q.Pull()
	}

	fmt.Println(q.Get())
	q.Show()

	fmt.Println(q.Push(12))
	q.Show()

}