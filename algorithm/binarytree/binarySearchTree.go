package main

import "fmt"

/**
1.既然有了这么高效的散列表，使用二叉树的地方是不是都可以替换成散列表呢？有没有哪些地方是散列表做不了，必须要用二叉树来做的呢？
	查找不光是查找某一个值，还会查找一个特定的范围，这在散列表里面就不一定适用了。类似B+树之类的，只在叶子节点保存数据，并且
将其用链表连起来。散列表在扩缩容的时候，性能不大稳定，同时由于散列冲突的存在，虽然散列表的时间复杂度是常数级别的，但实际应用中，
由于其不稳定，性能也不一定会比平衡二叉搜索树好。

2.二叉查找树（Binary Search Tree）:
	二叉查找树要求，在树中的任意一个节点，其左子树中的每个节点的值，都要小于这个节点的值，而右子树节点的值都大于这个节点的值。
3.二叉查找树的查找、插入操作都比较简单易懂，但是它的删除操作就比较复杂了 。针对要删除节点的子节点个数的不同，我们需要分三种情况来处理。
	1)要删除的节点没有子节点
	2)要删除的节点只有一个子节点
	3)要删除的节点有两个子节点:要先找到右子树中最小的节点，把他替换到要删除的节点上。然后再删掉这个最小节点(最小节点必然没有左子节点)
*/

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

var binarySearch = Node{
	Data: 33,
	Left: &Node{
		Data: 17,
		Left: &Node{
			Data: 13,
			Left: nil,
			Right: &Node{
				Data: 16,
			},
		},
		Right: &Node{
			Data: 18,
			Right: &Node{
				Data: 25,
			},
		},
	},
	Right: &Node{
		Data: 50,
		Left: nil,
		Right: &Node{
			Data: 58,
		},
	},
}

func main() {
	fmt.Println(find(50), find(51))
}

//--------------------默认节点没有重复数据

func find(data int) *Node {
	p := &binarySearch

	for p != nil {
		if data < p.Data {
			p = p.Left
		} else if data > p.Data {
			p = p.Right
		} else {
			return p
		}
	}
	return nil
}

func insert(data int) {
	if &binarySearch == nil {
		binarySearch = Node{
			Data: data,
		}
		return
	}

	p := &binarySearch

	for p != nil {
		if data < p.Data {
			if p.Left == nil {
				p.Left = &Node{
					Data: data,
				}
				return
			}
			p = p.Left
		} else {
			if p.Right == nil {
				p.Right = &Node{
					Data: data,
				}
				return
			}
			p = p.Right
		}
	}
	return
}

// 二叉查找树的查找、插入操作都比较简单易懂，但是它的删除操作就比较复杂了 。
func delete(data int) {
	p := &binarySearch // p指向要删除的节点，初始化指向根节点
	var pp *Node       //pp记录的是p的父节点
	for p != nil && p.Data != data {
		pp = p
		if p.Data < data {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	if p == nil {
		return // 没找到
	}
	// 要删除的节点有两个节点
	if p.Left != nil && p.Right != nil {
		minP := p.Right
		minPP := minP
		for minP.Left != nil {
			minPP = minP
			minP = minP.Left
		}
		p.Data = minP.Data // 将minP的数据替换到p中
		p = minP           // 下面就变成删除minP了
		pp = minPP
	}

	// 要删除的节点是叶子节点或仅有一个子节点
	var child *Node
	if p.Left != nil {
		child = p.Left
	} else if p.Right != nil {
		child = p.Right
	} else {
		child = nil
	}

	// 要删除的是根节点
	if pp == nil {
		binarySearch = *child
	} else if pp.Left == p {
		pp.Left = child
	} else {
		pp.Right = child
	}
}
