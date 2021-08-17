package main

import (
	"fmt"
	"math"
	"math/rand"
)

/**
https://juejin.cn/post/6844903446475177998  可着重看
https://cloud.tencent.com/developer/article/1353762 Redis源码学习之跳表

时间复杂度：O(logN)
空间复杂度: O(N)


*/

const (
	// 最高层数
	MAX_LEVEL = 16
)

type skipListNode struct {
	// 跳表保存的值
	v interface{}
	// 用于排序的分值
	score int
	// 层高
	level int
	// 每层的前进指针
	forwards []*skipListNode
	//每个元素为一个节点，每个forwards实际是当前节点值的后续节点，而不是像图中所画的一个数据分好多个节点 ,与层级实际关系不大。
	//再说明白一点，p.forwards[i]表示节点p的第i层的下一个节点。(即每个节点中的forwards存储的是该节点存在的那些层的对应该节点的下一个节点)
}

func newSkipListNode(v interface{}, score, level int) *skipListNode {
	return &skipListNode{v: v, score: score, level: level, forwards: make([]*skipListNode, level, level)}
}

// 跳表结构体
type SkipList struct {
	// 跳表头结点
	head *skipListNode
	// 跳表当前层数
	level int
	// 跳表长度，不包含头结点
	length int
}

// 实例化跳表对象
func NewSkipList() *SkipList {
	// 头结点，便于操作
	head := newSkipListNode(0, math.MinInt32, MAX_LEVEL)
	//
	return &SkipList{head: head, level: 1, length: 0}
}

// 获取跳表长度
func (sl *SkipList) Length() int {
	return sl.length
}

// 获取跳表层级
func (sl *SkipList) Level() int {
	return sl.level
}

// 插入节点到跳表,返回0代表成功,不支持插入相同值
func (sl *SkipList) Insert(v interface{}, score int) int {
	if v == nil {
		return 1
	}

	// update 更名为 preNodes ，即 待插新节点在每一层的pre 结点数组；
	// forwards 更名为 nextNodes，即 当前节点在每一层的next节点数组”

	// 查找插入位置
	cur := sl.head
	// 记录每层的路径
	update := [MAX_LEVEL]*skipListNode{}
	i := MAX_LEVEL - 1
	// 从最上层索引开始查找，主要目的是为了记录待插入节点的前置节点
	for ; i >= 0; i-- {
		// 在同一层索引中查找,i值相同的代表同一层
		for nil != cur.forwards[i] { // 后面存在有效节点
			if cur.forwards[i].v == v { // 后面节点的值与要插入的值相等
				return 2
			}
			if cur.forwards[i].score > score { // 后面节点的分值大于要插入的节点分值
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
		if nil == cur.forwards[i] {
			update[i] = cur
		}
	}

	//通过随机算法获取该节点层数,redis里的随机算法获取层数是一个核心点,确保了跳表不会退化
	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}

	//创建一个新的跳表节点
	newNode := newSkipListNode(v, score, level)

	//原有节点连接
	for i := 0; i <= level-1; i++ {
		next := update[i].forwards[i]   // 在第i层待插入节点的后续节点(注意，不是每个节点值冗余的存在于多个层中)
		update[i].forwards[i] = newNode // 类似于将要插入节点分别保存到对应各层
		newNode.forwards[i] = next
	}

	//如果当前节点的层数大于之前跳表的层数
	//更新当前跳表层数
	if level > sl.level {
		sl.level = level
	}

	//更新跳表长度
	sl.length++

	return 0

}

// 查找
func (sl *SkipList) Find(v interface{}, score int) *skipListNode {
	if nil == v || sl.length == 0 {
		return nil
	}
	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				return cur.forwards[i]
			} else if cur.forwards[i].score > score {
				break // 终止本层查找
			}
			cur = cur.forwards[i] // 同层向后查找
		}
	}
	return nil
}

func (sl *SkipList) Delete(v interface{}, score int) int {
	if nil == v {
		return 1
	}
	// 查找前驱节点
	cur := sl.head
	// 记录前驱路径
	update := [MAX_LEVEL]*skipListNode{}
	for i := sl.level - 1; i >= 0; i-- {
		update[i] = sl.head // 每一层的前驱节点开始都是头结点
		for nil != cur.forwards[i] {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				update[i] = cur
				break // 找到该层的前驱节点，终止本层的查找
			}
			cur = cur.forwards[i]
		}
	}

	cur = update[0].forwards[0] // 确定待删除节点，获取节点层高
	for i := cur.level - 1; i >= 0; i-- {
		if update[i] == sl.head && cur.forwards[i] == nil { // 待删除节点是最上层唯一一个存在数据的节点
			sl.level = i
		}
		if nil == update[i].forwards[i] {
			update[i].forwards[i] = nil // 重复
		} else {
			update[i].forwards[i] = update[i].forwards[i].forwards[i]
		}
	}
	sl.length--
	return 0
}

func (sl *SkipList) String() string {
	return fmt.Sprintf("level:%+v, length:%+v", sl.level, sl.length)
}
