package boxpacker

type PackedBox struct {
	Box             Box
	Items           MinHeap
	Weight          int
	RemainingWidth  int
	RemainingLength int
	RemainingDepth  int
	RemainingWeight int
}




func (pb PackedBox) ItemCount() int {
	return pb.Items.Len()
}

func (pb PackedBox) BoxInnerVolume() int{
	return pb.Box.InnerVolume
}

func (pb PackedBox) BoxWeight() int {

	if pb.Weight > 0 {
		return pb.Weight
	}

	pb.Weight = pb.Box.EmptyWeight

	for i:=0; i < pb.Items.Len(); i++ {
		obj, _ := pb.Items.ItemAtIndex(i)
		pb.Weight += obj.Weight
	}

	return pb.Weight;
}
// dummy func, never used, only to satisfy interface

func (i PackedBox) Priority() int {
	return 0
}


func NewPackedBox(box Box, itemList * MinHeap, remainingWidth int, remainingLength int, remainingDepth int, remainingWeight int) * PackedBox {
	pBox := &PackedBox{}
	pBox.Box = box
	pBox.Items = *itemList
	pBox.RemainingWidth = remainingWidth
	pBox.RemainingLength = remainingLength
	pBox.RemainingDepth = remainingDepth
	pBox.RemainingWeight = remainingWeight
	return pBox
}
