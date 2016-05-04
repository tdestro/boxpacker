package boxpacker

import (
	"container/heap"
	"errors"
	//"log"
)

var (
	HEAP_ERR_NOT_PACKEDBOX = "The type of the object being returned is not PackedBox."
	HEAP_ERR_NOT_BOX = "The type of the object being returned is not Box."
	HEAP_ERR_NOT_ITEM = "The type of the object being returned is not Item."
	HEAP_ERR_NOT_HEAPABLE = "Proper interface is not implemented on this object."
	HEAP_ERR_INDEX_OUT_OF_BOUNDS = "Given index is out of bounds."
)

type MinHeap []MinHeapable

//
// This identities that an object has the capibility to be added to the heap.
//

type MinHeapable interface {
	Priority() int
	// packed box only
	ItemCount() int
	BoxInnerVolume() int
	BoxWeight() int
}

//
// This identifies those additional functions besides those native to the built in heap.
//

type PackerMinHeapInterface interface {
	heap.Interface
	Copy() *MinHeap
	AtIndex(i int)
	RemoveAt(i int)
	PopItem() (*Item)
	PopBox() (*Box)
	PeekItem() *Item
	ItemAtIndex(i int)
	PeekBox() *Box
	PushItem(item Item)
	PushBox(box Box)
}

//
// Creates a new heap.
//

func NewMinHeap() *MinHeap {

	mh := &MinHeap{}
	heap.Init(mh)
	return mh
}

//
// Functions required to satisfy heap.interface
// they are called from functions within heap.interface and those are the functions that are called by the wrappers.
// calling these directly results in the contents not being kept as a heap.
//

func (mh MinHeap) Len() int {
	return len(mh)
}

func (mh MinHeap) Less(i, j int) bool {

	aI, okaI := mh[i].(Item)
	bI, okbI := mh[j].(Item)

	if okaI && okbI {
		return aI.Priority() < bI.Priority()
	}

	aB, okaB := mh[i].(Box)
	bB, okbB := mh[j].(Box)

	if okaB && okbB {
		return bB.Priority() > aB.Priority()
	}

	// Dealing with a PackedBox
	// are we equal on count ?
	choice := mh[i].ItemCount() - mh[j].ItemCount()

	// if we are equal on count, check box volume.
	if (choice == 0) {
		choice = mh[i].BoxInnerVolume() - mh[j].BoxInnerVolume()

	}
	if (choice == 0) {
		choice = mh[i].BoxWeight() - mh[j].BoxWeight()
	}

	if (choice < 0) {
		return false
	} else {
		return true
	}

}

func (mh *MinHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	x := old[n - 1]
	*mh = old[0 : n - 1]
	return x
}

func (mh *MinHeap) Push(item interface{}) {
	v, ok := item.(MinHeapable)
	if !ok {
		panic(HEAP_ERR_NOT_HEAPABLE)
	}
	*mh = append(*mh, v)
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

//
// Additional functions; satisfies PackerMinHeapInterface
//

func (mh *MinHeap) Peek() interface{} {
	items := *mh
	return items[0]
}

func (mh *MinHeap) RemoveAt(i int) (err error) {
	items := *mh
	if i < items.Len() && i > -1 {
		heap.Remove(mh, i)
	} else {
		err = errors.New(HEAP_ERR_INDEX_OUT_OF_BOUNDS)
	}
	return
}

// Returns the item at index i
func (mh *MinHeap) AtIndex(i int) (x MinHeapable, err error) {
	items := *mh
	if i < items.Len() && i > -1 {
		x = items[i]
	} else {
		err = errors.New(HEAP_ERR_INDEX_OUT_OF_BOUNDS)
	}
	return
}

func (mh *MinHeap) Copy() *MinHeap {
	newHeap := make(MinHeap, mh.Len())
	copy(newHeap, *mh)
	return &newHeap
}


//
// Additional functions; convenience unboxing wrappers; satisfies PackerMinHeapInterface
//

func (mh *MinHeap) PeekItem() *Item {
	if v, ok := mh.Peek().(Item); ok {
		return &v
	} else {
		panic(HEAP_ERR_NOT_ITEM)
	}
}
func (mh *MinHeap) PeekBox() *Box {
	if v, ok := mh.Peek().(Box); ok {
		return &v
	} else {
		panic(HEAP_ERR_NOT_BOX)
	}
}
func (mh *MinHeap) PeekPackedBox() *PackedBox {
	if v, ok := mh.Peek().(PackedBox); ok {
		return &v
	} else {
		panic(HEAP_ERR_NOT_PACKEDBOX)
	}
}

func (mh *MinHeap) PopItem() *Item {
	if v, ok := heap.Pop(mh).(Item); ok {
		return &v
	} else {
		panic(HEAP_ERR_NOT_ITEM)
	}
}
func (mh *MinHeap) PopBox() *Box {
	if v, ok := heap.Pop(mh).(Box); ok {
		return &v
	} else {
		panic(HEAP_ERR_NOT_BOX)
	}
}

func (mh *MinHeap) PopPackedBox() *PackedBox {
	if v, ok := heap.Pop(mh).(PackedBox); ok {
		return &v
	} else {
		panic(HEAP_ERR_NOT_PACKEDBOX)
	}
}
func (mh *MinHeap) ItemAtIndex(i int) (x Item, err error) {

	v, err := mh.AtIndex(i)

	if err != nil {
		return
	}

	if val, ok := v.(Item); ok {
		x = val
		return
	} else {
		panic(HEAP_ERR_NOT_ITEM)
	}

	return
}

//
// Additional functions; wrappers to make sure the indended type is what is put in; satisfies PackerMinHeapInterface
// Could be replaced by Push itself.
//
func (mh *MinHeap) PushItem(item Item) {
	heap.Push(mh, item)
}

func (mh *MinHeap) PushBox(box Box) {
	heap.Push(mh, box)
}

func (mh *MinHeap) PushPackedBox(packedbox PackedBox) {
	mh.Push(packedbox)
}
