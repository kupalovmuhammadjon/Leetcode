/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *       
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

 type PeekingIterator struct {
    iter *Iterator
    peeks []int
}

func Constructor(iter *Iterator) *PeekingIterator {
    return &PeekingIterator{iter, []int{}}
}

func (this *PeekingIterator) hasNext() bool {
    if len(this.peeks) > 0{
        return true
    }
    return this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
    if len(this.peeks) > 0{
        val := this.peeks[len(this.peeks)-1]
        this.peeks = this.peeks[:len(this.peeks)-1]
        return val
    }
    return this.iter.next()
}

func (this *PeekingIterator) peek() int {
    if len(this.peeks) > 0{
        return this.peeks[len(this.peeks)-1]
    }
    this.peeks = append(this.peeks, this.iter.next())
    return this.peeks[len(this.peeks)-1]
}