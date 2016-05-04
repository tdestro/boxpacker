package boxpacker

import (
	"math"
	"log"
)

func VolumePackerPack(box Box, items *MinHeap) *PackedBox {

	log.Printf("[EVALUATING BOX] %s", box.Reference)

	packedItems := NewMinHeap()
	remainingDepth := box.InnerDepth
	remainingWeight := box.MaxWeight - box.EmptyWeight
	remainingWidth := box.InnerWidth
	remainingLength := box.InnerLength

	layerWidth := 0
	layerLength := 0
	layerDepth := 0

	for (items.Len() > 0) {

		itemToPack := items.PeekItem()
		// skip items that are simply too large

		if isItemTooLargeForBox(itemToPack, remainingDepth, remainingWeight) {
			items.PopItem()
			log.Printf("%s depth or weight too much.", itemToPack.Description)
			continue
		}

		log.Printf("evaluating item %s %d x %d x %d", itemToPack.Description, itemToPack.Length, itemToPack.Width, itemToPack.Depth)
		log.Printf("remaining width: %d  length: %d  depth: %d", remainingWidth, remainingLength, remainingDepth)
		log.Printf("remaining layerWidth: %d  layerLength: %d  layerDepth: %d", layerWidth, layerLength, layerDepth)

		itemWidth := itemToPack.Width
		itemLength := itemToPack.Length

		if (fitsGap(itemToPack, remainingWidth, remainingLength)) {


			packedItems.PushItem(*items.PopItem())

			remainingWeight -= itemToPack.Weight

			var nextItem *Item
			if (items.Len() > 0) {
				nextItem = items.PeekItem()
			}

			if (fitsBetterRotated(itemToPack, nextItem, remainingWidth, remainingLength)) {
				log.Printf("Fits (better) unrotated")
				remainingLength -= itemLength
				layerLength += itemLength
				layerWidth = max(itemWidth, layerWidth)
			} else {
				log.Printf("Fits (better) rotated")
				remainingLength -= itemWidth
				layerLength += itemWidth
				layerWidth = max(itemLength, layerWidth)
			}
			layerDepth = max(layerDepth, itemToPack.Depth) //greater than 0, items will always be less deep
			//allow items to be stacked in place within the same footprint up to current layerdepth
			maxStackDepth := layerDepth - itemToPack.Depth





			for items.Len() > 0  && canStackItemInLayer(itemToPack, items.PeekItem(), maxStackDepth, remainingWeight) {
				remainingWeight -= items.PeekItem().Weight
				maxStackDepth -= items.PeekItem().Depth
				packedItems.PushItem(*items.PopItem())
			}


		} else {
			if (remainingWidth >= min(itemWidth, itemLength) && isLayerStarted(layerWidth, layerLength, layerDepth)) {
				log.Printf("No more fit in lengthwise, resetting for new row")
				remainingLength += layerLength
				remainingWidth -= layerWidth
				layerWidth = 0
				layerLength = 0
				continue
			} else if (remainingLength < min(itemWidth, itemLength) || layerDepth == 0) {
				log.Printf("Doesn't fit on layer even when empty")
				items.PopItem()
				continue
			}

			remainingWidth = box.InnerWidth
			if (layerWidth > 0) {
				remainingWidth = min(int(math.Floor(float64(layerWidth) * 1.1)), box.InnerWidth)
			}

			remainingLength = box.InnerLength
			if (layerLength > 0) {
				remainingLength = min(int(math.Floor(float64(layerLength) * 1.1)), box.InnerLength)
			}

			remainingDepth -= layerDepth
			layerWidth = 0
			layerLength = 0
			layerDepth = 0

			log.Printf("Doesn't fit, so starting next vertical layer")
		}
	}


	log.Printf("Done with this box.")
	packedBox := NewPackedBox(box, packedItems, remainingWidth, remainingLength, remainingDepth, remainingWeight)
	return packedBox
}
// max only used with two params ever, and only int, full php-style implementation not needed.
func max(n0 int, n1 int) int {
	// if multiple values of different types evaluate as equal the first provided to the function will be returned.
	if n1 > n0 {
		return n1
	}
	return n0
}

// min only used with two params ever, and only int, full php-style implementation not needed.
func min(n0 int, n1 int) int {
	// if multiple values of different types evaluate as equal the first provided to the function will be returned.
	if n1 < n0 {
		return n1
	}
	return n0
}

func isItemTooLargeForBox(item *Item, remainingDepth int, remainingWeight int) bool {
	return item.Depth > remainingDepth || item.Weight > remainingWeight
}

// Figure out space left for next item if we pack this one in it's regular orientation
func fitsSameGap(item *Item, remainingWidth int, remainingLength int) int {
	return min(remainingWidth - item.Width, remainingLength - item.Length)
}

// Figure out space left for next item if we pack this one rotated by 90deg
func fitsRotatedGap(item *Item, remainingWidth int, remainingLength int) int {
	return min(remainingWidth - item.Length, remainingLength - item.Width)
}

func fitsBetterRotated(item *Item, nextItem *Item, remainingWidth int, remainingLength int) bool {
	fitsSameGap := fitsSameGap(item, remainingWidth, remainingLength)
	fitsRotatedGap := fitsRotatedGap(item, remainingWidth, remainingLength)

	return bool(fitsRotatedGap < 0 || (fitsSameGap >= 0 && fitsSameGap <= fitsRotatedGap) || (item.Width <= remainingWidth && nextItem == item && remainingLength >= 2 * item.Length))
}
// Does item fit in specified gap
func fitsGap(item *Item, remainingWidth int, remainingLength int) bool {
	return fitsSameGap(item, remainingWidth, remainingLength) >= 0 || fitsRotatedGap(item, remainingWidth, remainingLength) >= 0
}

// Figure out if we can stack the next item vertically on top of this rather than side by side
// Used when we've packed a tall item, and have just put a shorter one next to it
func canStackItemInLayer(item *Item, nextItem *Item, maxStackDepth int, remainingWeight int) bool {
	return nextItem.Depth <= maxStackDepth && nextItem.Weight <= remainingWeight && nextItem.Width <= item.Width && nextItem.Length <= item.Length
}

func isLayerStarted(layerWidth int, layerLength int, layerDepth int) bool {
	return layerWidth > 0 && layerLength > 0 && layerDepth > 0
}