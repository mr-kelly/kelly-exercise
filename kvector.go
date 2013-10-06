/*

C-like Vector data structure in GoLang implement

Author: Mrkelly

*/
package klib

const (
	SPARE_CAPACITY = 16 // Init Space
)

// Similar to  C-Vector data structure
type KVector_T interface{}

type KVector struct {
	items    []*KVector_T // container
	size     int          // real size
	capacity int          // real in memory
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
	newVec.items = make([]*KVector_T, newVec.capacity)

	return newVec
}

func (this *KVector) reverse() {

}

// vector push_back
func (this *KVector) PushBack(obj *KVector_T) {
	this.items[this.size] = obj
	this.size++

	// return true
}

func (this *KVector) Size() int {
	return this.size
}

func (this *KVector) Capacity() int {
	return this.capacity
}
