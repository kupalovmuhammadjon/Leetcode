type UnionFind struct {
    Parents map[byte]byte
    Sizes   map[byte]int
}

func constructor() *UnionFind {
    p := make(map[byte]byte)
    s := make(map[byte]int)
    for ch := byte('a'); ch <= byte('z'); ch++ {
        p[ch] = ch
        s[ch] = 1
    }
    return &UnionFind{p, s}
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
    uf := constructor()
    for i := 0; i < len(s1); i++ {
        uf.Union(s1[i], s2[i])
    }

    var newBase strings.Builder
    for i := 0; i < len(baseStr); i++{
        p := uf.Find(baseStr[i])
        newBase.WriteByte(p)
    }

    return newBase.String()
}

func (uf *UnionFind) Find(node byte) byte {
    if uf.Parents[node] == node {
        return node
    }
    uf.Parents[node] = uf.Find(uf.Parents[node])
    return uf.Parents[node]
}

func (uf *UnionFind) Union(node1, node2 byte) {
    root1 := uf.Find(node1)
    root2 := uf.Find(node2)

    if root1 == root2 {
        return
    }

    if root1 < root2 {
        uf.Parents[root2] = root1
    } else {
        uf.Parents[root1] = root2
    }
}

