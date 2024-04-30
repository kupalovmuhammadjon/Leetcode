package main

import "fmt"

type LinkedList struct{
	Val int
	Next *LinkedList
}

type Queue struct{
	Head *LinkedList
	Length int
	Tail *LinkedList
}

func main(){
	queue := new(Queue)
	queue.EnQueue(5)
	queue.EnQueue(7)
	fmt.Println(queue.Head.Val)
	queue.DeQueue()
	fmt.Println(queue.Head)
	fmt.Println(queue.Tail)
	fmt.Println(queue.Length)
}

func (this *Queue) EnQueue(val int){
	newNode := &LinkedList{Val: val, Next: nil}
	this.Length++
	if this.Tail == nil{
		this.Head = newNode
		this.Tail = newNode
		return
	}
	this.Tail.Next = &LinkedList{Val: val, Next: nil}
	this.Tail = this.Tail.Next
}
func (this *Queue) DeQueue() *LinkedList{
	if this.Length == 0{
		return nil
	}
	currentHead := this.Head
	this.Head = this.Head.Next
	if this.Head == nil{
		this.Tail = nil
	}
	this.Length--

	return currentHead
}