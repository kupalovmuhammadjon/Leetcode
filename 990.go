type UnionFind struct{
    Parents []int
    Size []int
}

func equationsPossible(equations []string) bool {
    uf := constructor()
    restricted := map[byte][]byte{}
    for _, eqs := range equations{
        if eqs[1:3] == "=="{
            for _, r := range restricted[eqs[0]]{
                r1 := uf.find(int(eqs[3] - 'a'))
                if uf.find(int(r - 'a')) == r1 || r == eqs[3]{
                    return false
                }
            }
            for _, r := range restricted[eqs[3]]{
                r1 := uf.find(int(eqs[0] - 'a'))
                if uf.find(int(r - 'a')) == r1 || r == eqs[0]{
                    return false
                }
            }
            uf.union(int(eqs[0]-'a'), int(eqs[3] - 'a'))
        }else{
            root1 := uf.find(int(eqs[0]-'a'))
            root2 := uf.find(int(eqs[3]-'a'))
            if root1 == root2{
                return false
            }
            restricted[eqs[0]] = append(restricted[eqs[0]], eqs[3])
            restricted[eqs[3]] = append(restricted[eqs[3]], eqs[0])
        }
    }
    return true
}

func constructor() *UnionFind{
    p := make([]int, 26)
    s := make([]int, 26)
    for i := 0; i < 26; i++{
        p[i] = i
        s[i] = 1
    }

    return &UnionFind{p, s}
}

func (uf *UnionFind) find(node int) int{
    if uf.Parents[node] == node{
        return node
    }
    uf.Parents[node] = uf.find(uf.Parents[node])
    return  uf.Parents[node]
}

func (uf *UnionFind) union(node1, node2 int) {
    root1 := uf.find(node1)
    root2 := uf.find(node2)

    if root1 == root2 { return }

    if uf.Size[root1] > uf.Size[root2]{
        uf.Parents[root2] = root1
        uf.Size[root1] += uf.Size[root2] 
    }else{
        uf.Parents[root2] = root1
        uf.Size[root1] += uf.Size[root2] 
    }
    
}