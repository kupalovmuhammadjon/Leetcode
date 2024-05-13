type KthLargest struct {
    Heap []int
    K int
}


func Constructor(k int, nums []int) KthLargest {
    h := KthLargest{K: k}

    sort.Slice(nums, func(x, y int) bool{
        return nums[x] > nums[y]
    })
    h.Push(0)
    for i := 0; i < k && i < len(nums); i++{
        h.Push(nums[i])
    }

    return h 
}   


func (h *KthLargest) Add(val int) int {
    if len(h.Heap) <= h.K{
        h.Push(val)
        return h.Heap[1]
    }
    if h.Heap[1] >= val{
        return h.Heap[1]
    }
    h.Heap[1] = val
    h.Reorder()
    return h.Heap[1] 

}

func (h *KthLargest) Push(val int) {
    if len(h.Heap) <= 1{
        h.Heap = append(h.Heap, val)
        return
    }
    h.Heap = append(h.Heap, val)
    i := len(h.Heap)-1
    
    for i/2 >= 1 && h.Heap[i] < h.Heap[i/2]{
        h.Heap[i], h.Heap[i/2] = h.Heap[i/2], h.Heap[i]
        i /= 2
    }
}

func (h *KthLargest) Reorder() {

    i := 1
    for 2 * i  < len(h.Heap){
        if 2 * i + 1 < len(h.Heap) && h.Heap[2 * i + 1] < h.Heap[i] &&
        h.Heap[2 * i + 1] < h.Heap[2*i]{
            h.Heap[2 * i + 1], h.Heap[i] = h.Heap[i], h.Heap[2 * i + 1]
            i = i * 2 + 1
        }else if h.Heap[2*i] < h.Heap[i] {
            h.Heap[2*i], h.Heap[i] = h.Heap[i], h.Heap[2*i]
            i = 2 * i
        }else{
            break
        }
    }
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */