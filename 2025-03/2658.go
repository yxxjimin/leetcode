// Runtime : 6 ms
// Memory  : 9.73 MB
package mar25

type DisjointSet2 struct {
	parent []int
	size   []int
	edges  []int
}

func Construct2(n int) *DisjointSet2 {
	parent, size, edges := make([]int, n), make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	return &DisjointSet2{
		parent: parent,
		size:   size,
		edges:  edges,
	}
}

func (this *DisjointSet2) Find(p int) int {
	if this.parent[p] != p {
		this.parent[p] = this.Find(this.parent[p])
	}
	return this.parent[p]
}

func (this *DisjointSet2) Union(p, q int) {
	this.edges[p]++
	this.edges[q]++

	pr, qr := this.Find(p), this.Find(q)

	if pr == qr {
		return
	} else if this.size[qr] > this.size[pr] {
		pr, qr = qr, pr
	}

	this.parent[qr] = pr
	this.size[pr] += this.size[qr]
}

func (this *DisjointSet2) IsComplete(p int) bool {
	root := this.Find(p)
	nodes := this.size[root]
	return this.edges[p] == nodes-1
}

func countCompleteComponents(n int, edges [][]int) int {
	ds := Construct2(n)
	isComplete := make(map[int]bool)
	ans := 0

	for _, edge := range edges {
		p, q := edge[0], edge[1]
		ds.Union(p, q)
	}

	for i := 0; i < n; i++ {
		root := ds.Find(i)

		if _, ok := isComplete[root]; !ok {
			isComplete[root] = true
		}

		if !ds.IsComplete(i) {
			isComplete[root] = false
		}
	}

	for _, comp := range isComplete {
		if comp {
			ans++
		}
	}

	return ans
}
