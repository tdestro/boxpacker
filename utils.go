package boxpacker

import (
	aelog "google.golang.org/appengine/log"
	"golang.org/x/net/context"
)

func Report(ctx context.Context, packedBoxes *MinHeap) {

	aelog.Infof(ctx, "These items fitted into %d box(es).\n", packedBoxes.Len())
	for _, pb := range *packedBoxes {
		packedBox := pb.(PackedBox)
		boxType := packedBox.Box
		aelog.Infof(ctx, "This box is a %s, it is %d in wide, %d in long and %d in high.\n", boxType.Reference, boxType.OuterWidth, boxType.OuterLength, boxType.OuterDepth)
		aelog.Infof(ctx, "The combined weight of this box and the items inside it is %d lbs.", packedBox.BoxWeight());
		aelog.Infof(ctx, "The items in this box are:\n");
		itemsInTheBox := packedBox.Items

		for _, i := range itemsInTheBox {
			item := i.(Item)
			aelog.Infof(ctx, "%s", item.Description)
		}

	}
}

