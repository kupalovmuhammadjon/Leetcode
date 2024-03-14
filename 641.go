type MyCircularDeque struct {
	deque []int
	ln    int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{[]int{}, k}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.ln <= len(this.deque) {
		return false
	}
	this.deque = append([]int{value}, this.deque...)
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.ln <= len(this.deque) {
		return false
	}
	this.deque = append(this.deque, value)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if 0 == len(this.deque) {
		return false
	}
	this.deque = this.deque[1:]
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if 0 == len(this.deque) {
		return false
	}
	this.deque = this.deque[:len(this.deque)-1]
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if 0 == len(this.deque) {
		return -1
	}
	return this.deque[0]
}

func (this *MyCircularDeque) GetRear() int {
	if 0 == len(this.deque) {
		return -1
	}
	return this.deque[len(this.deque)-1]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return len(this.deque) == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.ln == len(this.deque)
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */