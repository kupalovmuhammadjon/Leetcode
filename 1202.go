type UnionFind struct{
    Parents []int
    Sizes []int
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

func smallestStringWithSwaps(s string, pairs [][]int) string {
    uf := constructor(len(s))
    uf.buildConnection(pairs)

    sortedPairs := map[int][]int{}
    for i := 0; i < len(s); i++{
        p := uf.Find(i)
        sortedPairs[p] = append(sortedPairs[p], i)
    }
    
result := make([]byte, len(s))

	for _, indices := range sortedPairs {
		characters := make([]byte, len(indices))
		for i, index := range indices {
			characters[i] = s[index]
		}

		sort.Slice(characters, func(i, j int) bool {
			return characters[i] < characters[j]
		})

		sort.Ints(indices)
		for i, index := range indices {
			result[index] = characters[i]
		}
	}

	return string(result)
}

func (uf *UnionFind) buildConnection(pairs [][]int){
    for _, pair := range pairs{
        uf.Union(pair[0], pair[1])
    }
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