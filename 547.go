type UnionFind struct{
    Parent []int
    Size []int
}

func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    uf := constructor(n)

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if isConnected[i][j] == 1 {
                uf.unionBySize(i, j)
            }
        }
    }
    
    components := map[int]bool{}
    for _, p := range uf.Parent {
        components[uf.find(p)] = true
    }
    
    return len(components)
}

func constructor(size int) *UnionFind {
    p := make([]int, size)
    s := make([]int, size)

    for i := 0; i < size; i++ {
        p[i] = i
        s[i] = 1
    }

    return &UnionFind{p, s}
}

func (uf *UnionFind) find(node int) int {
    if uf.Parent[node] == node {
        return node
    }
    uf.Parent[node] = uf.find(uf.Parent[node])
    return uf.Parent[node]
}

func (uf *UnionFind) unionBySize(node1, node2 int) {
    root1 := uf.find(node1)
    root2 := uf.find(node2)

    if root1 != root2 {
        if uf.Size[root1] > uf.Size[root2] {
            uf.Parent[root2] = root1
            uf.Size[root1] += uf.Size[root2]
        } else {
            uf.Parent[root1] = root2
            uf.Size[root2] += uf.Size[root1]
        }
    }
}
