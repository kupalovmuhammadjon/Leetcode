type Heap struct{
    Heap []int
}

func (h *Heap) Heapify(){
    h.Heap = append(h.Heap, h.Heap[0])

    for cur := len(h.Heap) / 2; cur > 0; cur--{
        i := cur
        for 2 * i < len(h.Heap){
            if 2 * i + 1 < len(h.Heap) && h.Heap[2 * i + 1] > h.Heap[i] &&
            h.Heap[2 * i + 1] > h.Heap[2 * i]{
                h.Heap[2 * i + 1], h.Heap[i] = h.Heap[i], h.Heap[2 * i + 1]
                i = 2 * i + 1
            }else if h.Heap[i] < h.Heap[2 * i]{
                h.Heap[2 * i], h.Heap[i] = h.Heap[i], h.Heap[2 * i]
                i = 2 * i
            }else{
                break
            }
        }
    }
}

func (h *Heap) Push(val int){
    h.Heap = append(h.Heap, val)
    i := len(h.Heap) - 1

    for i / 2 >= 1 && h.Heap[i] > h.Heap[i/2]{
        h.Heap[i], h.Heap[i/2] = h.Heap[i/2], h.Heap[i]
        i = i / 2
    }
}

func (h *Heap) Pop() int{
    maxValue := h.Heap[1]
    h.Heap[1] = h.Heap[len(h.Heap)-1]
    h.Heap = h.Heap[:len(h.Heap)-1]

    i := 1
    for 2 * i < len(h.Heap) {
        if 2 * i + 1 < len(h.Heap) && h.Heap[2 * i + 1] > h.Heap[i] &&
        h.Heap[2 * i + 1] > h.Heap[2*i]{
            h.Heap[2 * i + 1], h.Heap[i] = h.Heap[i], h.Heap[2 * i + 1]
            i = i * 2 + 1
        }else if h.Heap[2 * i] > h.Heap[i]{
            h.Heap[2 * i], h.Heap[i] = h.Heap[i], h.Heap[2 * i]
            i = i * 2
        }else{
            break
        }
    }
    return maxValue
}

func lastStoneWeight(stones []int) int {
    h := Heap{stones}
    h.Heapify()

    for len(h.Heap) > 2{
        stone1 := h.Pop()
        stone2 := h.Pop()
        if stone1 - stone2 > 0{
            h.Push(stone1 - stone2)
        }
    }
    if len(h.Heap) == 1{
        return 0
    }
    return h.Pop()
}