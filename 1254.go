type UnionFind struct{
    Parents []int
    Sizes []int
}

func (uf *UnionFind) Find(node int) int{
    if uf.Parents[node] == node{
        return node
    }
    uf.Parents[node] = uf.Find(uf.Parents[node])
    return uf.Parents[node]
}

func (uf *UnionFind) Union(node1, node2 int){
    root1 := uf.Find(node1)
    root2 := uf.Find(node2)

    if root1 == root2 { return }

    if uf.Sizes[root1] > uf.Sizes[root2]{
        uf.Parents[root2] = root1
        uf.Sizes[root1] += uf.Sizes[root2]
    }else{
        uf.Parents[root1] = root2
        uf.Sizes[root2] += uf.Sizes[root1]
    }
}

func constructor(ln int) *UnionFind{
    p := make([]int, ln)
    s := make([]int, ln)

    for i := 0; i < ln; i++{
        p[i] = i
        s[i] = 1
    }

    return &UnionFind{p, s}
}

func closedIsland(grid [][]int) int {
    uf := constructor(len(grid)*len(grid[0]))
    buildConnections(uf, grid) 

    border := getBorders(uf, grid)
    
    used := map[int]bool{}
    count := 0
    for i := 1; i < len(grid)-1; i++{
        for j := 1; j < len(grid[i])-1; j++{
            if grid[i][j] == 0{
                parent := uf.Find(i * len(grid[0]) + j)
                if !used[parent]  && !border[parent]{
                    count++
                    used[parent] = true
                }
            }
        }
    }

    return count
}

func buildConnections(uf *UnionFind, grid [][]int){

    for i := 0; i < len(grid); i++{
        for j := 0; j < len(grid[i]); j++{
            if grid[i][j] == 0{
                connectNeighbours(uf, grid, i, j)
            }
        }
    }
}

func connectNeighbours(uf *UnionFind, grid [][]int, i, j int){
    directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

    for _, dir := range directions{
        r := i + dir[0]
        c := j + dir[1]
        if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]){
            continue
        }
        if grid[r][c] == 0{
            uf.Union(r * len(grid[0]) + c, i * len(grid[0]) + j)
        }
    }
}

func getBorders(uf *UnionFind, grid [][]int) map[int]bool {
    rows := len(grid)
    cols := len(grid[0])
    borders := map[int]bool{}

    for i := 0; i < cols; i++ {
        borders[uf.Find(i)] = true                   // First row
        borders[uf.Find((rows-1)*cols + i)] = true   // Last row
    }

    for i := 0; i < rows; i++ {
        borders[uf.Find(i*cols)] = true              // First column
        borders[uf.Find(i*cols + (cols-1))] = true   // Last column
    }

    return borders
}
