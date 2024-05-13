package main

type Heap struct{
	Heap []int
}

func (h *Heap) Push(val int){
	h.Heap = append(h.Heap, val)
	i := len(h.Heap)-1

	for i > 
}

func (h *Heap) Pop(){
	
func (h *KthLargest) Pop() int{
    if len(h.Heap) == 1{
        return 0
    }

    var minValue int
    minValue = h.Heap[1]
    h.Heap[1] = h.Heap[len(h.Heap)-1]
    h.Heap = h.Heap[:len(h.Heap)-1]

    if len(h.Heap) == 1{
        return minValue
    }

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
    return minValue
}
}

func (h *Heap) Heapify(){
	
}

func main(){

}