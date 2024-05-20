type SmallestInfiniteSet struct {
    Heap []int
    Numbers map[int]bool
    Last int
}


func Constructor() SmallestInfiniteSet {
    return SmallestInfiniteSet{[]int{0}, map[int]bool{}, 1}
}


func (h *SmallestInfiniteSet) PopSmallest() int {
    if len(h.Heap) <= 1{
        res := h.Last
        h.Last++
        return res
    }
    mn := h.Heap[1]
    h.Heap[1] = h.Heap[len(h.Heap)-1]
    h.Heap = h.Heap[:len(h.Heap)-1]
    if h.Last <= mn{
        h.Last++
    }
    delete(h.Numbers, mn)
    i := 1
    for i * 2 < len(h.Heap){
        if i * 2 + 1 < len(h.Heap) && h.Heap[i] > h.Heap[i * 2 + 1] && h.Heap[i*2] > h.Heap[i * 2 + 1]{
            h.Heap[i], h.Heap[i * 2 + 1] = h.Heap[i * 2 + 1], h.Heap[i]
            i = i * 2 + 1
        }else if h.Heap[i] > h.Heap[i*2]{
            h.Heap[i], h.Heap[i*2] = h.Heap[i*2], h.Heap[i]
        }else{
            break
        }
    }

    return mn
}

func (h *SmallestInfiniteSet) AddBack(num int) {
    if h.Last <= num || h.Numbers[num]{
        return
    }
    h.Numbers[num] = true
    h.Heap = append(h.Heap, num)
    i := len(h.Heap) - 1
    for i / 2 >= 1 && h.Heap[i] < h.Heap[i/2]{
        h.Heap[i], h.Heap[i/2] = h.Heap[i/2], h.Heap[i]
        i /= 2
    }
}



/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */