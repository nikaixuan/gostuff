package algorithms

type Uf struct {
	count  int
	parent []int
	size   []int
}

func (u *Uf) Uf(n int) {
	u.count = n
	u.parent = make([]int, n)
	u.size = make([]int, n)
	for i := 0; i < n; i++ {
		u.parent[i] = i
		u.size[i] = i
	}
}

func (u *Uf) Union(p, q int) {
	rootp := u.Find(p)
	rootq := u.Find(q)
	if rootp == rootq {
		return
	}
	if u.size[rootq] > u.size[rootp] {
		u.parent[rootq] = rootp
		u.size[rootp] += u.size[rootq]
	} else {
		u.parent[rootp] = rootq
		u.size[rootq] += u.size[rootp]
	}
	u.count = u.count - 1
}

func (u *Uf) Find(x int) int {
	for u.parent[x] != x {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}

func (u *Uf) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}
