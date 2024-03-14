type MyLinkedList struct {
	ls []int
}

func Constructor() MyLinkedList {
	return MyLinkedList{[]int{}}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= len(this.ls) {
		return -1
	}
	return this.ls[index]
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.ls = append([]int{val}, this.ls...)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.ls = append(this.ls, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

	if index > len(this.ls) {
		return
	}

	this.ls = append(this.ls[:index], append([]int{val}, this.ls[index:]...)...)
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= 0 && index < len(this.ls) {
		this.ls = append(this.ls[:index], this.ls[index+1:]...)
	}
}
