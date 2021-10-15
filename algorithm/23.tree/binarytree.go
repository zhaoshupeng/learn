package main

import (
	"fmt"
)

/**
1、想要存储一棵二叉树，我们有两种方法，一种是基于指针或者引用的二叉链式存储法，一种是基于数组的顺序存储法。
	顺序存储法(用数组存储)：把根节点存储在下标 i = 1 的位置，那左子节点存储在下标 2 * i = 2 的位置，右子节点存储在 2 * i + 1 = 3，类推
2、遍历方式：前序遍历、中序遍历、后续遍历：时间复杂度都是O(n)
	其中，前、中、后序，表示的是节点与它的左右子树节点遍历打印的先后顺序。
3、按层遍历：
	层次遍历需要借助队列这样一个辅助数据结构。（其实也可以不用，这样就要自己手动去处理节点的关系，代码不太好理解，好处就是
	空间复杂度是o(1)。不过用队列比较好理解，缺点就是空间复杂度是o(n)）。根节点先入队列，然后队列不空，取出对头元素，
	如果左孩子存在就入列队，否则什么也不做，右孩子同理。直到队列为空，则表示树层次遍历结束。树的层次遍历，其实也是
	一个广度优先的遍历算法。
*/

// 参考极客时间 https://time.geekbang.org/column/article/67856
var geek BinaryTreeNode = BinaryTreeNode{
	Data: "A",
	Left: &BinaryTreeNode{
		Data: "B",
		Left: &BinaryTreeNode{
			Data: "D",
			Left: nil,
		},
		Right: &BinaryTreeNode{
			Data: "E",
		},
	},
	Right: &BinaryTreeNode{
		Data: "C",
		Left: &BinaryTreeNode{
			Data: "F",
		},
		Right: &BinaryTreeNode{
			Data: "G",
		},
	},
}

// https://blog.csdn.net/u014754841/article/details/79361772
// origin [0, 1, 2, 3, 4, 5, 0, 6, 0, 7, 8]
// 图中“0”表示不存在此结点。由此可见顺序存储结构仅适用于完全二叉树。
var origin = BinaryTreeNode{
	Data: 1,
	Left: &BinaryTreeNode{
		Data: 2,
		Left: &BinaryTreeNode{
			Data: 4,
			Left: nil,
			Right: &BinaryTreeNode{
				Data: 7,
			},
		},
		Right: &BinaryTreeNode{
			Data: 5,
			Left: &BinaryTreeNode{
				Data: 8,
			},
		},
	},
	Right: &BinaryTreeNode{
		Data: 3,
		Left: nil,
		Right: &BinaryTreeNode{
			Data: 6,
		},
	},
}

func main() {

	preOrder := preOrderTraversal(&geek, []interface{}{})
	fmt.Println("前序遍历：", preOrder, preOrderTraversalOptimize(&geek), preOrderTraversalBak(&geek)) // [A B D E C F G]

	inOrder := inOrderTraversal(&geek, []interface{}{})
	fmt.Println("中序遍历：", inOrder, inOrderTraversalBak(&geek)) // [D B E A F C G]

	postOrder := postOrderTraversal(&geek, []interface{}{})
	fmt.Println("后序遍历：", postOrder, postOrderTraversalBak(&geek)) // [D E B F G C A]
	fmt.Println("树的高度：", binaryTreeHeight(&geek))

	//fmt.Println("按层遍历：", floorLevelTree(&geek))  ????

	//var testCreate *BinaryTreeNode
	//test := preOrderCreate(testCreate)
	//testOrder := preOrderTraversal(testCreate, []interface{}{})
	//fmt.Println("前序创建并遍历：", testCreate, test, testOrder, preOrderTraversalOptimize(testCreate))
}

// 二叉树一个节点结构
type BinaryTreeNode struct {
	Data  interface{}
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

//-------二叉树生成(链式存储)
//1、前序创建二叉树
func preOrderCreate(root *BinaryTreeNode) []interface{} {
	// 从标准输入流中接收输入数据
	fmt.Println("Please input node value:")
	nodeValue := 0
	fmt.Scanln(&nodeValue)
	var arr []interface{}
	if nodeValue == 0 {
		return arr
	}

	if nodeValue != 0 { // 标记结束创建
		root = &BinaryTreeNode{}

		root.Data = nodeValue
		fmt.Println(root)
		arr = append(arr, nodeValue)
		arr = append(arr, preOrderCreate(root.Left)...)
		arr = append(arr, preOrderCreate(root.Right)...)

	}
	return arr
}

// 前序遍历, res 存储打印顺序
func preOrderTraversal(root *BinaryTreeNode, res []interface{}) []interface{} {
	if root == nil {
		return res
	}
	res = append(res, root.Data) // 先存储本节点，类似于打印
	//fmt.Println("遍历的节点值:", root.Data)
	if root.Left != nil { //处理左子树
		res = preOrderTraversal(root.Left, res)
	}
	if root.Right != nil {
		res = preOrderTraversal(root.Right, res)
	}
	return res
}

func preOrderTraversalOptimize(root *BinaryTreeNode) []interface{} {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []interface{}{root.Data}
	}
	res := []interface{}{root.Data}
	if root.Left != nil { //处理左子树
		res = append(res, preOrderTraversalOptimize(root.Left)...)
	}
	if root.Right != nil {
		res = append(res, preOrderTraversalOptimize(root.Right)...)
	}
	return res
}

func preOrderTraversalBak(root *BinaryTreeNode) []interface{} {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []interface{}{root.Data}
	}
	var stack []*BinaryTreeNode
	var res []interface{}
	stack = append(stack, root)
	for len(stack) != 0 {
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, e.Data)
		if e.Right != nil {
			stack = append(stack, e.Right)
		}
		if e.Left != nil {
			stack = append(stack, e.Left)
		}
	}
	return res
}

// 中序遍历
func inOrderTraversal(root *BinaryTreeNode, res []interface{}) []interface{} {
	if root == nil {
		return res
	}
	if root.Left != nil { //处理左子树
		res = inOrderTraversal(root.Left, res)
	}
	res = append(res, root.Data) // 先存储本节点，类似于打印
	//fmt.Println("遍历的节点值:", root.Data)
	if root.Right != nil {
		res = inOrderTraversal(root.Right, res)
	}

	return res
}

func inOrderTraversalBak(root *BinaryTreeNode) []interface{} {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []interface{}{root.Data}
	}
	res := inOrderTraversalBak(root.Left)
	res = append(res, root.Data)
	res = append(res, inOrderTraversalBak(root.Right)...)

	return res
}

// 后续遍历
func postOrderTraversal(root *BinaryTreeNode, res []interface{}) []interface{} {
	if root == nil {
		return res
	}
	if root.Left != nil { //处理左子树
		res = postOrderTraversal(root.Left, res)
	}
	if root.Right != nil {
		res = postOrderTraversal(root.Right, res)
	}
	res = append(res, root.Data) // 存储本节点，类似于打印

	return res
}
func postOrderTraversalBak(root *BinaryTreeNode) []interface{} {
	if root == nil {
		return nil
	}
	var res []interface{}
	if root.Left != nil {
		lres := postOrderTraversalBak(root.Left)
		if len(lres) > 0 {
			res = append(res, lres...)
		}
	}
	if root.Right != nil {
		rres := postOrderTraversalBak(root.Right)
		if len(rres) > 0 {
			res = append(res, rres...)
		}
	}
	res = append(res, root.Data)
	return res
}

func binaryTreeHeight(root *BinaryTreeNode) int {
	var height int = 0
	if root != nil {
		leftHeight := binaryTreeHeight(root.Left)
		rightHeight := binaryTreeHeight(root.Right)
		if leftHeight > rightHeight {
			height = leftHeight
		} else {
			height = rightHeight
		}
		height++
	}
	return height
}

// 按层遍历
func floorLevelTree(root *BinaryTreeNode) []interface{} {
	var res []interface{}
	var quene []*BinaryTreeNode
	quene = append(quene, root)
	for len(quene) > 0 {
		node := quene[0]
		res = append(res, node.Data)
		quene = quene[0 : len(quene)-1]
		if node.Left != nil {
			quene = append(quene, node.Left)
		}

		if node.Right != nil {
			quene = append(quene, node.Right)
		}
	}
	return res
}
