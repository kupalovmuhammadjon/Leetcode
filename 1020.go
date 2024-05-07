type UnionFind struct{
    Parents []int
    Size []int
}

func numEnclaves(grid [][]int) int {
    uf := constructor(len(grid) * len(grid[0]))
    for i := 0; i < len(grid); i++{
        for j := 0; j < len(grid[i]); j++{
            if grid[i][j] == 1{
                makeConnections(grid, uf, i, j)
            }
        }
    }

    // getting borders
    border := make(map[int]bool)

    // Top and bottom borders
    for i := 0; i < len(grid[0]); i++ {
        cur := i
        cur = uf.find(cur)
        border[cur] = true

        cur2 := (len(grid)-1) * len(grid[0]) + i
        cur2 = uf.find(cur2)
        border[cur2] = true
    }

    // Left and right borders
    for i := 0; i < len(grid); i++ {
        cur := i * len(grid[0])
        cur = uf.find(cur)
        border[cur] = true

        cur2 := i * len(grid[0]) + len(grid[0])-1
        cur2 = uf.find(cur2)
        border[cur2] = true
    }

    sizeOfRemoteIslands := 0
    removed := map[int]bool{}
    for i := 1; i < len(grid)-1; i++{
        for j := 1; j < len(grid[i])-1; j++{
            if grid[i][j] == 1{
                cur := i * len(grid[0]) + j
                cur = uf.find(cur)
                if !border[cur] && !removed[cur]{
                    sizeOfRemoteIslands += uf.Size[cur]
                    removed[cur] = true
                }
            }
        }
    }

    return sizeOfRemoteIslands
}

func constructor(size int) *UnionFind{
    Parents := make([]int, size)
    Size := make([]int, size)

    for i := 0; i < size; i++{
        Parents[i] = i
        Size[i] = 1
    }

    return &UnionFind{Parents, Size}
}

func makeConnections(grid [][]int, uf *UnionFind, i, j int){
    directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
    cur := i * len(grid[0]) + j
    for _, drc := range directions{
        newI := i + drc[0]
        newJ := j + drc[1]
        if newI >= 0 && newI < len(grid) && newJ >= 0 && newJ < len(grid[0]) && grid[newI][newJ] == 1{
            cur2 := newI * len(grid[0]) + newJ
            uf.union(cur, cur2)
        }
    }
}

func (uf *UnionFind) find(node int) int{
    if uf.Parents[node] == node{
        return node
    }
    uf.Parents[node] = uf.find(uf.Parents[node])
    return uf.Parents[node]
}

func (uf *UnionFind) union(node1, node2 int) {
    root1 := uf.find(node1)
    root2 := uf.find(node2)

    if root1 == root2 { return }

    if uf.Size[root1] > uf.Size[root2]{
        uf.Parents[root2] = root1
        uf.Size[root1] += uf.Size[root2]
    }else{
        uf.Parents[root1] = root2
        uf.Size[root2] += uf.Size[root1]
    }
}