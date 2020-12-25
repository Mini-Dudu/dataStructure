package binaryTree

import "fmt"

func TestBiTree() {
	B := New(0)

	node1 := New(1)
	node2 := New(2)
	node3 := New(3)
	node4 := New(4)
	node5 := New(5)
	node6 := New(6)
	node7 := New(7)
	node8 := New(8)
	node9 := New(9)
	node10 := New(10)
	node11 := New(11)
	node12 := New(12)

	B.InsertChild(B, node1, L)
	B.InsertChild(B, node2, R)

	B.InsertChild(node1, node3, L)
	B.InsertChild(node1, node4, R)

	B.InsertChild(node2, node5, L)
	B.InsertChild(node2, node6, R)

	B.InsertChild(node4, node7, L)
	B.InsertChild(node4, node8, R)

	B.InsertChild(node8, node9, L)
	B.InsertChild(node8, node10, R)

	B.InsertChild(node10, node11, L)
	B.InsertChild(node11, node12, R)

	//B.left = node1
	//B.right = node2
	//
	//node1.left = node3
	//node1.right = node4
	//
	//node2.left = node5
	//node2.right = node6
	//
	////node3.left = nil
	////node3.right = nil
	//
	//node4.left = node7
	//node4.right = node8
	//
	////node5.left = nil
	////node5.right = nil
	//
	////node6.left = nil
	////node6.right = nil
	//
	////node7.left = nil
	////node7.right = nil
	//
	//node8.left = node9
	//node8.right = node10
	//
	//node10.left = node11
	//node11.right = node12

	fmt.Println(B.Parent(node1) == B)

	//fmt.Println(B.RightChild(B).data)
	//B.LevelOrderTraverse()
	//fmt.Println(B.RightSibling(node1) == node2)

	//fmt.Printf("节点11的双亲节点是否为节点2:%v\n",B.SearchParent(node11) == node2)
	//fmt.Println("非递归中序遍历:")//3 1 7 4 9 8 10 0 5 2 6
	//fmt.Println(B.SearchParent(node11))

	//fmt.Printf("递归先序遍历:\n")//0 1 3 4 7 8 9 10 2 5 6
	//B.PreOrderTraverse()
	//
	//fmt.Println("\n递归中序遍历:")//3 1 7 4 9 8 10 0 5 2 6
	//B.InOrderTraverse()
	//
	//fmt.Println("\n递归后序遍历:")//3 7 9 10 8 4 1 5 6 2 0
	//B.PostOrderTraverse()
	//fmt.Printf("\n树的高度:%v",B.BiTreeDepth())

	//l, r := B.SearchChild(node6)
	//fmt.Printf("\n左孩子:%v\t右孩子:%v\n", l, r)

}
