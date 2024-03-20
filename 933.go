type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{[]int{}}
}

func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	start := t - 3000
	if start < 0 {
		return len(this.queue)
	} else {
		for len(this.queue) > 0 {
			if this.queue[0] < start {
				this.queue = this.queue[1:]
			} else {
				break
			}
		}
	}
	return len(this.queue)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */