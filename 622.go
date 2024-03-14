type MyCircularQueue struct {
	queue []int
	ln    int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{[]int{}, k}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if len(this.queue) == this.ln {
		return false
	}
	this.queue = append(this.queue, value)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if len(this.queue) == 0 {
		return false
	}
	this.queue = this.queue[1:]
	return true
}

func (this *MyCircularQueue) Front() int {
	if len(this.queue) > 0 {
		return this.queue[0]
	}
	return -1
}

func (this *MyCircularQueue) Rear() int {
	if len(this.queue) > 0 {
		return this.queue[len(this.queue)-1]
	}
	return -1
}

func (this *MyCircularQueue) IsEmpty() bool {
	return len(this.queue) == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return len(this.queue) == this.ln
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */