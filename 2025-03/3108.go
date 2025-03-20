// Runtime : 39 ms
// Memory  : 35.58 MB
package mar25

type DisjointSet struct {
	parent []int
	size   []int
	weight map[int]int
}

func Construct(n int) *DisjointSet {
	parent, size := make([]int, n), make([]int, n)
	weight := make(map[int]int)

	for i := range parent {
		parent[i] = i
		size[i] = 1
		weight[i] = -1
	}

	return &DisjointSet{
		parent: parent,
		size:   size,
		weight: weight,
	}
}

func (ds *DisjointSet) Find(p int) int {
	for p != ds.parent[p] {
		ds.parent[p] = ds.parent[ds.parent[p]]
		p = ds.parent[p]
	}
	return p
}

func (ds *DisjointSet) Union(p, q, w int) {
	pRoot, qRoot := ds.Find(p), ds.Find(q)
	if pRoot == qRoot {
		ds.weight[pRoot] &= w
		return
	}
	if ds.size[pRoot] > ds.size[qRoot] {
		ds.parent[qRoot] = pRoot
		ds.size[pRoot] += ds.size[qRoot]
		ds.weight[pRoot] &= w & ds.weight[qRoot]
	} else {
		ds.parent[pRoot] = qRoot
		ds.size[qRoot] += ds.size[pRoot]
		ds.weight[qRoot] &= w & ds.weight[pRoot]
	}
}

func minimumCost(n int, edges [][]int, query [][]int) []int {
	ds := Construct(n)
	ans := make([]int, len(query))

	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		ds.Union(u, v, w)
	}

	for i, q := range query {
		s, t := q[0], q[1]
		if sRoot, tRoot := ds.Find(s), ds.Find(t); sRoot != tRoot {
			ans[i] = -1
		} else {
			ans[i] = ds.weight[sRoot]
		}
	}

	return ans
}
