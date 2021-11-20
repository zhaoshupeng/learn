package graph

import "container/list"

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
