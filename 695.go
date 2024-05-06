type UnionFind struct{
    Parent []int
    Size []int
    maxSize int
}

func constructor(ln int) *UnionFind{
    p := make([]int, ln)
    s := make([]int, ln)

    for i := 0; i < ln; i++{
        p[i] = i
        s[i] = 1
    }
    return &UnionFind{p, s, 0}
}

func maxAreaOfIsland(grid [][]int) int {
    
    uf := constructor(len(grid) * len(grid[0]))
    oneExists := false
    for i := 0; i < len(grid); i++{
        for j := 0; j < len(grid[0]); j++{
            if grid[i][j] == 1{
                oneExists = true
                findConnections(grid, uf, i, j)
            }
        }
    }
    if uf.maxSize == 0 && oneExists{
        return 1
    }
    return uf.maxSize
}

func findConnections(grid [][]int, uf *UnionFind, i, j int){
    cur := i * len(grid[0]) + j
    directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
    for _, drc := range directions{
        ci := i + drc[0]
        cj := j + drc[1]
        if ci < 0 || cj < 0 || ci >= len(grid) || cj >= len(grid[0]){
            continue
        }
        if grid[ci][cj] == 1{
            ind := ci * len(grid[0]) + cj
            uf.unionBySize(cur, ind)
        }
    }
}

func (uf *UnionFind) find(node int) int{
    if uf.Parent[node] == node{
        return node
    }
    uf.Parent[node] = uf.find(uf.Parent[node])
    return uf.Parent[node]
}

func (uf *UnionFind) unionBySize(node1, node2 int){
    ulpNode1 := uf.find(node1)
    ulpNode2 := uf.find(node2)
    if ulpNode1 == ulpNode2{ return }
    if uf.Size[ulpNode1] > uf.Size[ulpNode2]{
        uf.Parent[ulpNode2] = ulpNode1
        uf.Size[ulpNode1] += uf.Size[ulpNode2]
        uf.maxSize = max(uf.maxSize, uf.Size[ulpNode1])
    }else{
        uf.Parent[ulpNode1] = ulpNode2
        uf.Size[ulpNode2] += uf.Size[ulpNode1]
        uf.maxSize = max(uf.maxSize, uf.Size[ulpNode2])
    }
}