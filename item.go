package boxpacker

type Item struct {
	Description string
	Width       int
	Length      int
	Depth       int
	Volume      int
	Weight      int
}

func (i Item) Priority() int { return i.Volume }


// dummy funcs, only packed box uses these
func (i Item) ItemCount() int {
	return 0
}

func (i Item) BoxInnerVolume() int {
	return 0
}

func (i Item) BoxWeight() int {
	return 0
}

func CalcVolItem(i * Item) *Item {
	i.Volume = i.Width * i.Length * i.Depth
	return i
}

func ItemCompare(i, j MinHeapable) bool {

	var a Item
	if v, ok := i.(Item); ok {
		a = v
	} else {
		panic(HEAP_ERR_NOT_ITEM)
	}

	var b Item
	if v, ok := j.(Item); ok {
		b = v
	} else {
		panic(HEAP_ERR_NOT_ITEM)
	}

	if a.Description != b.Description {
		return false
	}
	if a.Width != b.Width {
		return false
	}
	if a.Length != b.Length {
		return false
	}
	if a.Depth != b.Depth {
		return false
	}
	if a.Volume != b.Volume {
		return false
	}
	if a.Weight != b.Weight {
		return false
	}

	return true
}