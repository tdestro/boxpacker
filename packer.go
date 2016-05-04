package boxpacker

import (
	"errors"
)

var (
	Items * MinHeap
	Boxes * MinHeap
	PACKER_ERR_ITEM_TOO_BIG = "The Item is too large for any box."
)

func  NewPacker(){
	Items = NewMinHeap()
	Boxes = NewMinHeap()
}

func AddItem(item *Item, qty int) {
	itemVol := CalcVolItem(item)
	for i := 0; i < qty; i++ {
		Items.PushItem(*itemVol)
	}

	Debugf("added item %d x %s", qty, item.Description);
}

func AddBox(box *Box) {
	Boxes.PushBox(*CalcVolBox(CalcOuterDims(box)))
	Debugf("added box %s %v", box.Reference, box)
}

// Pack items into boxes
func Pack() *MinHeap {

	packedBoxes, _ := doVolumePacking()
	if packedBoxes.Len() > 1 {
		// packedBoxes = RedistributeWeight(boxes, packedBoxes)
	}
	Debugf("packing completed %d", packedBoxes.Len());
	return packedBoxes
}

// Pack items into boxes using the principle of largest volume item first

func doVolumePacking() (packedBoxes *MinHeap, err error) {
	packedBoxes = NewMinHeap()
	//Keep going until everything packed
	for (Items.Len() > 0) {
		boxesToEvaluate := Boxes.Copy()
		packedBoxesIteration := NewMinHeap()

		for (boxesToEvaluate.Len() > 0) {
			box := boxesToEvaluate.PopBox()
			packedBox := VolumePackerPack(*box, Items.Copy());

			if packedBox.Items.Len() > 0 {
				packedBoxesIteration.PushPackedBox(*packedBox)
				//Have we found a single box that contains everything?
				if packedBox.Items.Len() == Items.Len() {
					break;
				}
			}
		}

		//Check iteration was productive
		if packedBoxesIteration.Len() == 0 {
			Debugf("%s is too large to fit into any box.", Items.PeekItem().Description);
			err = errors.New(PACKER_ERR_ITEM_TOO_BIG)
			return
		}

		//Find best box of iteration, and remove packed items from unpacked list

		unPackedItems := Items.Copy()
		bestBox := packedBoxesIteration.PeekPackedBox()
		bestBoxItems := bestBox.Items.Copy()
		for _, packedItem := range * bestBoxItems {
			for unpackedKey, unpackedItem := range *unPackedItems {
				if ItemCompare(packedItem, unpackedItem) {
					err = unPackedItems.RemoveAt(unpackedKey)
					if err != nil {
						Debugf("unPackedItems.RemoveAt %s", err);
						return nil, err
					}
					break
				}
			}
		}

		unpackedItemList := NewMinHeap()

		for i := 0; i < unPackedItems.Len(); i++ {
			unpackedItem, err := unPackedItems.ItemAtIndex(i)
			if err != nil {
				Debugf("unPackedItems.ItemAtIndex %s", err);
				return nil, err
			}
			unpackedItemList.PushItem(unpackedItem)
		}

		Items = unpackedItemList
		packedBoxes.PushPackedBox(*bestBox);
	}
	return
}

