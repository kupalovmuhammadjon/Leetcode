type UnionFind struct{
    Parent []int
    Size []int
    Max int
}
func longestConsecutive(nums []int) int {
    existingNumbers := map[int]int{}
    for i := 0; i < len(nums); i++{
        existingNumbers[nums[i]] = i
    }
    if len(existingNumbers) == 1{
        return 1
    }else if len(existingNumbers) == 0{
        return 0
    }

    uf := constructor(nums)
    for num, i := range existingNumbers{
        if _, ok := existingNumbers[num+1]; ok{
            uf.unionBySize(i, existingNumbers[num+1])
        }
        if _, ok := existingNumbers[num-1]; ok{
            uf.unionBySize(i, existingNumbers[num-1])
        }
    }
    if uf.Max == 0{
        return 1
    }
    return uf.Max
}

func constructor(nums []int) *UnionFind{
    p := make([]int, len(nums))
    s := make([]int, len(nums))
    for i := range nums{
        p[i] = i
        s[i] = 1
    }
    return &UnionFind{p, s, 0}
}

func (uf *UnionFind) find(ind int) int{
    if uf.Parent[ind] == ind{
        return ind
    }
    uf.Parent[ind] = uf.find(uf.Parent[ind])
    return uf.Parent[ind]
}

func (uf *UnionFind) unionBySize(node1, node2 int){
    ulpNode1 := uf.find(node1)
    ulpNode2 := uf.find(node2)
    if ulpNode1 == ulpNode2 { return }
    if uf.Size[ulpNode1] >= uf.Size[ulpNode2]{
        uf.Parent[ulpNode2] = ulpNode1
        uf.Size[ulpNode1] += uf.Size[ulpNode2]
        uf.Max = max(uf.Max, uf.Size[ulpNode1])
    }else{
        uf.Parent[ulpNode1] = ulpNode2
        uf.Size[ulpNode2] += uf.Size[ulpNode1]
        uf.Max = max(uf.Max, uf.Size[ulpNode2])
    }
}