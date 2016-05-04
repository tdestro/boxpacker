package boxpacker

import "log"

func Report(packedBoxes *MinHeap) {

	log.Printf("These items fitted into %d box(es).\n", packedBoxes.Len())
	for _, pb := range *packedBoxes {
		packedBox := pb.(PackedBox)
		boxType := packedBox.Box
		log.Printf("This box is a %s, it is %d in wide, %d in long and %d in high.\n", boxType.Reference, boxType.OuterWidth, boxType.OuterLength, boxType.OuterDepth)
		log.Printf("The combined weight of this box and the items inside it is %d lbs.", packedBox.BoxWeight());
		log.Printf("The items in this box are:\n");
		itemsInTheBox := packedBox.Items

		for _, i := range itemsInTheBox {
			item := i.(Item)
			log.Printf("%s", item.Description)
		}

	}
}


func Debugf(format string, args ...interface{}) {
	if *debug {
		log.Printf("DEBUG "+format, args)
	}
}

func Errorf(format string, args ...interface{}) {
	log.Printf("ERROR "+format, args)
}
