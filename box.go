package boxpacker

import (
	"github.com/tdestro/NarwhalAPI/utils"
)

type Box struct {
	Reference   string
	OuterWidth  int
	OuterLength int
	OuterDepth  int
	EmptyWeight int
	InnerWidth  int
	InnerLength int
	InnerDepth  int
	InnerVolume int
	MaxWeight   int
}

func (i Box) Priority() int {
	return i.InnerVolume
}


// dummy funcs, only packed box uses these
func (i Box) ItemCount() int {
	return 0
}

func (i Box) BoxInnerVolume() int {
	return 0
}

func (i Box) BoxWeight() int {
	return 0
}

func CalcVolBox(b *Box) *Box {
	b.InnerVolume = b.InnerWidth * b.InnerLength * b.InnerDepth
	return b
}

func CalcOuterDims(b *Box) *Box {
	b.EmptyWeight = 0
	b.OuterLength = b.InnerLength + int(utils.RoundFloat64(3 / 8))
	b.OuterWidth = b.InnerWidth + int(utils.RoundFloat64(3 / 8))
	b.OuterDepth = b.InnerDepth + int(utils.RoundFloat64(5 / 8))
	return b
}

