package graph

import (
	"container/list"
	"fmt"
)

/**
（1）定义
	顶点的度（degree）: 就是跟顶点相连接的边的条数。
	顶点的入度: 表示有多少条边指向这个顶点；
	顶点的出度: 表示有多少条边是以这个顶点为起点指向其他顶点。
	带权图（weighted graph）: 在带权图中，每条边都有一个权重（weight），我们可以通过这个权重来表示 QQ 好友间的亲密度。
	稀疏图（Sparse Matrix）: 也就是说，顶点很多，但每个顶点的边并不多。

	我们需要一个逆邻接表。邻接表中存储了用户的关注关系，逆邻接表中存储的是用户的被关注关系。
*/
/**
如何在内存中存储图这种数据结构呢？
	（1）邻接矩阵存储方法：邻接矩阵的底层依赖一个二维数组。对于无向图来说，如果顶点 i 与顶点 j 之间有边，我们就将 A[i][j] 和 A[j][i] 标记为 1；
	（2）邻接表存储方法（Adjacency List）：
		邻接表是不是有点像散列表？每个顶点对应一条链表，链表中存储的是与这个顶点相连接的其他顶点。一个有向图的邻接表存储方式，每个顶点对应的链表里面，存储的是指向的顶点。。
对于无向图来说，也是类似的，不过，每个顶点的链表中存储的，是跟这个顶点有边相连的顶点。邻接表存储起来比较节省空间，但是使用起来就比较耗时间。
（我们可以将邻接表中的链表改成平衡二叉查找树。实际开发中，我们可以选择用红黑树。这样，我们就可以更加快速地查找两个顶点之间是否存在边了。
当然，这里的二叉查找树可以换成其他动态数据结构，比如跳表、散列表等。除此之外，我们还可以将链表改成有序动态数组，可以通过二分查找的方
法来快速定位两个顶点之间否是存在边。）
*/

/**
需要说明一下，深度优先搜索算法和广度优先搜索算法，既可以用在无向图，也可以用在有向图上。
	(1) 广度优先算法(BFS)
	(2) 深度优先算法(DFS)
*/

/*
问题：
	（1）给你一个用户，如何找出这个用户的所有三度（其中包含一度、二度和三度）好友关系？
*/

//adjacency table, 以无向图作为示例
type Graph struct {
	adj []*list.List // 邻接表 // list 包实现了一个双链表（doubly linked list）
	v   int          // 顶点的个数
}

// init graphh according to capacity
func NewGraph(v int) *Graph {
	graphh := &Graph{}
	graphh.v = v
	graphh.adj = make([]*list.List, v)
	for i := range graphh.adj {
		graphh.adj[i] = list.New() //  // 创建一个新的链表，头节点是空(哨兵)
	}
	return graphh
}

//insert as add edge，一条边存2次
func (graph *Graph) addEdge(s int, t int) { // 无向图一条边存两次
	graph.adj[s].PushBack(t)
	graph.adj[t].PushBack(s)
}

// 广度优先搜索（Breadth-First-Search），我们平常都简称 BFS。直观地讲，它其实就是一种“地毯式”层层推进的搜索策略，
//即先查找离起始顶点最近的，然后是次近的，依次往外搜索。
// s 表示起始顶点，t 表示终止顶点；
// 实际上，这样求得的路径就是从 s 到 t 的最短路径。
// 时间复杂度：O(V+E)可以简写为 O(E)。，其中，V 表示顶点的个数，E 表示边的个数； （当然，对于一个连通图来说，也就是说一个图中的所有顶点都是连通的，E 肯定要大于等于 V-1）
// 空间复杂度： O(V)。（空间消耗主要在几个辅助变量 visited 数组、queue 队列、prev 数组上。这三个存储空间的大小都不会超过顶点的个数）
func (graph *Graph) BFS(s int, t int) {
	if s == t {
		return
	}

	// 三个核心变量

	//init prev
	// prev 用来记录搜索路径。
	// 当我们从顶点 s 开始，广度优先搜索到顶点 t 后，prev 数组中存储的就是搜索的路径。不过，这个路径是反向存储的。
	// prev[w]存储的是，顶点 w 是从哪个前驱顶点遍历过来的。
	// 比如，我们通过顶点 2 的邻接表访问到顶点 3，那 prev[3]就等于 2。
	prev := make([]int, graph.v)
	for index := range prev {
		prev[index] = -1
	}

	//search by queue
	// visited 是用来记录已经被访问的顶点，用来避免顶点被重复访问。如果顶点 q 被访问，那相应的 visited[q]会被设置为 true。
	visited := make([]bool, graph.v)
	visited[s] = true

	// queue 是一个队列，用来存储已经被访问、但相连的顶点还没有被访问的顶点。
	// 因为广度优先搜索是逐层访问的，也就是说，我们只有把第 k 层的顶点都访问完成之后，才能访问第 k+1 层的顶点。
	// 当我们访问到第 k 层的顶点的时候，我们需要把第 k 层的顶点记录下来，稍后才能通过第 k 层的顶点来找第 k+1 层的顶点。所以，我们用这个队列来实现记录的功能。
	var queue []int
	queue = append(queue, s)

	isFound := false

	// 都访问完成或者发现目标顶点则停止循环
	for len(queue) > 0 && !isFound {
		// 获取队首
		top := queue[0]
		queue = queue[1:]
		linkedlist := graph.adj[top] // 获取对应顶点的链表

		for e := linkedlist.Front(); e != nil; e = e.Next() {
			k := e.Value.(int) // 类型断言
			if !visited[k] {   // 如果该顶点没被访问
				prev[k] = top // 标记从top可以到达k
				if k == t {
					// todo 也可以在这里直接打印路径，这样就不需要  isFound
					// printPrev(prev, s, t)
					isFound = true
					break
				}

				queue = append(queue, k)

				visited[k] = true // 标记 k 顶点被访问
			}
		}
	}

	if isFound { // 找到则，打印路径
		printPrev(prev, s, t)
	} else {
		fmt.Printf("no path found from %d to %d\n", s, t)
	}
}

//print path recursively
func printPrev(prev []int, s int, t int) { // 递归打印s->t的路径
	// 递归截止条件
	if t != s && prev[t] != -1 {
		printPrev(prev, s, prev[t])
	}
	fmt.Printf("%d ", t)

}

// 深度优先搜索（Depth-First-Search），简称 DFS。最直观的例子就是“走迷宫”。
// 实际上，深度优先搜索用的是一种比较著名的算法思想，回溯思想。
// 深度优先搜索代码实现也用到了 prev、visited 变量以及 print() 函数，它们跟广度优先搜索代码实现里的作用是一样的
// 不过，深度优先搜索代码实现里，有个比较特殊的变量 found，它的作用是，当我们已经找到终止顶点 t 之后，我们就不再递归地继续查找了。
//search by DFS
// 时间复杂度：深度优先搜索算法的时间复杂度是 O(E)，E 表示边的个数。
// 空间复杂度就是 O(V)。
func (graph *Graph) DFS(s int, t int) {
	prev := make([]int, graph.v)
	for index := range prev {
		prev[index] = -1
	}

	visited := make([]bool, graph.v)
	visited[s] = true

	isFound := false

	graph.recurse(s, t, prev, visited, isFound)

	printPrev(prev, s, t)
}

//recursivly find path
func (graph *Graph) recurse(s int, t int, prev []int, visited []bool, isFound bool) {
	if isFound {
		return
	}

	visited[s] = true

	if s == t {
		isFound = true
		return
	}

	linkedlist := graph.adj[s]
	for e := linkedlist.Front(); e != nil; e = e.Next() {
		k := e.Value.(int)
		if !visited[k] {
			prev[k] = s
			graph.recurse(k, t, prev, visited, false)
		}
	}
}

//func main() {
//	graph := newGraph(8)
//	graph.addEdge(0, 1)
//	graph.addEdge(0, 3)
//	graph.addEdge(1, 2)
//	graph.addEdge(1, 4)
//	graph.addEdge(2, 5)
//	graph.addEdge(3, 4)
//	graph.addEdge(4, 5)
//	graph.addEdge(4, 6)
//	graph.addEdge(5, 7)
//	graph.addEdge(6, 7)
//
//	graph.BFS(0, 7)
//	fmt.Println()
//	graph.BFS(1, 3)
//	fmt.Println()
//	graph.DFS(0, 7)
//	fmt.Println()
//	graph.DFS(1, 3)
//	fmt.Println()
//}
