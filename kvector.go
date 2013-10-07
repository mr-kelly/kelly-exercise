/*

C-like Vector data structure in GoLang implement

Author: Mrkelly

*/
package klib

const (
	SPARE_CAPACITY = 16 // Init Space
)

type KVector struct {
	items    []_T // container
	size     int  // real size
	capacity int  // real in memory
}

// args[0] -> initSize
func NewKVector(args ...int) *KVector {
	var initSize int = 0 // default 0 size
	if cap(args) > 0 {   // with arguments
		initSize = args[0]
	}

	var newVec *KVector = new(KVector)
	newVec.size = initSize
	newVec.capacity = initSize + SPARE_CAPACITY // default 2,   1 times increase
	newVec.items = make([]_T, newVec.capacity)

	return newVec
}

//    with * star,  will be a instance function
// New Capacity of items Slice
func (this *KVector) reverse(newCapacity int) {
	oldItems := this.items
	newItems := make([]_T, newCapacity)
	copy(newItems, oldItems)
	this.items = newItems

	this.capacity = newCapacity
	// del oldItems    auto release memory
}

func (this *KVector) Resize(newSize int) {
	if newSize > this.capacity { // Larger than the capacity,  create more cap
		this.reverse(newSize*2 + 1) // if newSize = 0, capacity = 1
	}
}

// vector push_back
func (this *KVector) PushBack(obj _T) {
	if this.size == this.capacity {
		this.reverse(this.capacity*2 + 1)
	}

	this.items[this.size] = obj
	this.size++
}

func (this *KVector) Size() int {
	return this.size
}

func (this *KVector) Capacity() int {
	return this.capacity
}
func (this *KVector) IsEmpty() bool {
	return this.size == 0
}
func (this *KVector) PopBack() {
	this.items[this.size] = nil
	this.size--

	// size too small~   shorten it for save memory
	if this.size < (this.capacity-2)/4 {
		this.reverse(this.size*2 + 1) // if newSize = 0, capacity = 1
	}
}

func (this *KVector) Get(index int) _T {
	return this.items[index]
}

func (this *KVector) Set(index int, value _T) {
	this.items[index] = value
}
