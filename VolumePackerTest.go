package boxpacker

import (
	"golang.org/x/net/context"
	"testing"
	"github.com/stretchr/testify/assert"
	aelog "google.golang.org/appengine/log"
)

func NewTestItem(Description string, Width int, Length int, Depth int, Weight int) Item {
	return *CalcVolItem(&Item{Description:"Baseball", Width: Width, Length:Length, Depth: Depth, Weight: Weight})
}
func NewTestBox(Reference string, OuterWidth int, OuterLength int, OuterDepth int, EmptyWeight int, InnerWidth int, InnerLength int, InnerDepth int, MaxWeight int) Box {
	return *CalcVolBox(&Box{Reference: Reference, InnerLength: InnerLength, InnerWidth: InnerWidth, InnerDepth: InnerDepth, MaxWeight: MaxWeight})
}

func testPackBoxThreeItemsFitEasily(ctx context.Context, t *testing.T) bool {
	aelog.Infof(ctx, "test");
	assert := assert.New(t)
	box := NewTestBox("Le box", 300, 300, 10, 10, 296, 296, 8, 1000)

	items := NewMinHeap()

	items.PushItem(NewTestItem("Item 1", 250, 250, 2, 200))
	items.PushItem(NewTestItem("Item 2", 250, 250, 2, 200))
	items.PushItem(NewTestItem("Item 3", 250, 250, 2, 200))

	packedBox := VolumePackerPack(box, items.Copy());

	return assert.Equal(3, packedBox.ItemCount(), "they should be equal")
}

func testPackBoxThreeItemsFitExactly(ctx context.Context, t *testing.T) bool {
	assert := assert.New(t)
	box := NewTestBox("Le box", 300, 300, 10, 10, 296, 296, 8, 1000)
	items := NewMinHeap()

	items.PushItem(NewTestItem("Item 1", 296, 296, 2, 200))
	items.PushItem(NewTestItem("Item 2", 296, 296, 2, 500))
	items.PushItem(NewTestItem("Item 3", 296, 296, 4, 290))

	packedBox := VolumePackerPack(box, items.Copy());
	return assert.Equal(3, packedBox.ItemCount(), "they should be equal")
}

func testPackBoxThreeItemsFitExactlyNoRotation(ctx context.Context, t *testing.T) bool{
	assert := assert.New(t)
	box := NewTestBox("Le box", 300, 300, 10, 10, 296, 296, 8, 1000)
	items := NewMinHeap()

	items.PushItem(NewTestItem("Item 1", 296, 148, 2, 200))
	items.PushItem(NewTestItem("Item 2", 296, 148, 2, 500))

	packedBox := VolumePackerPack(box, items.Copy());
	return assert.Equal(2, packedBox.ItemCount(), "they should be equal")
}

func testPackBoxThreeItemsFitSizeButOverweight(ctx context.Context, t *testing.T) bool {
	assert := assert.New(t)
	box := NewTestBox("Le box", 300, 300, 10, 10, 296, 296, 8, 1000)
	items := NewMinHeap()

	items.PushItem(NewTestItem("Item 1", 250, 250, 2, 400))
	items.PushItem(NewTestItem("Item 2", 250, 250, 2, 500))
	items.PushItem(NewTestItem("Item 3", 250, 250, 2, 200))

	packedBox := VolumePackerPack(box, items.Copy());
	return assert.Equal(2, packedBox.ItemCount(), "they should be equal")
}

func testPackBoxThreeItemsFitWeightBut2Oversize(ctx context.Context, t *testing.T) bool {
	assert := assert.New(t)
	box := NewTestBox("Le box", 300, 300, 10, 10, 296, 296, 8, 1000)
	items := NewMinHeap()

	items.PushItem(NewTestItem("Item 1", 297, 296, 2, 200))
	items.PushItem(NewTestItem("Item 2", 297, 296, 2, 500))
	items.PushItem(NewTestItem("Item 3", 296, 296, 4, 290))

	packedBox := VolumePackerPack(box, items.Copy());
	return assert.Equal(1, packedBox.ItemCount(), "they should be equal")
}

func testPackTwoItemsFitExactlySideBySide(ctx context.Context, t *testing.T) bool {
	assert := assert.New(t)
	box := NewTestBox("Le box", 300, 400, 10, 10, 296, 496, 8, 1000)
	items := NewMinHeap()

	items.PushItem(NewTestItem("Item 1", 296, 248, 8, 200))
	items.PushItem(NewTestItem("Item 2", 248, 296, 8, 200))

	packedBox := VolumePackerPack(box, items.Copy());
	return assert.Equal(2, packedBox.ItemCount(), "they should be equal")
}

func testPackThreeItemsBottom2FitSideBySideOneExactlyOnTop(ctx context.Context, t *testing.T) bool {
	assert := assert.New(t)
	box := NewTestBox("Le box", 300, 300, 10, 10, 296, 296, 8, 1000)
	items := NewMinHeap()

	items.PushItem(NewTestItem("Item 1", 248, 148, 4, 200))
	items.PushItem(NewTestItem("Item 2", 148, 248, 4, 200))
	items.PushItem(NewTestItem("Item 3", 296, 296, 4, 200))

	packedBox := VolumePackerPack(box, items.Copy());
	return assert.Equal(3, packedBox.ItemCount(), "they should be equal")
}

func testPackThreeItemsBottom2FitSideBySideWithSpareSpaceOneOverhangSlightlyOnTop(ctx context.Context, t *testing.T) bool {
	assert := assert.New(t)
	box := NewTestBox("Le box", 250, 250, 10, 10, 248, 248, 8, 1000)
	items := NewMinHeap()

	items.PushItem(NewTestItem("Item 1", 200, 200, 4, 200))
	items.PushItem(NewTestItem("Item 2", 110, 110, 4, 200))
	items.PushItem(NewTestItem("Item 3", 110, 110, 4, 200))

	packedBox := VolumePackerPack(box, items.Copy());
	return assert.Equal(3, packedBox.ItemCount(), "they should be equal")
}

func testPackSingleItemFitsBetterRotated(ctx context.Context, t *testing.T) bool {
	assert := assert.New(t)
	box := NewTestBox("Le box", 400, 300, 10, 10, 396, 296, 8, 1000)
	items := NewMinHeap()

	items.PushItem(NewTestItem("Item 1", 250, 290, 2, 200))

	packedBox := VolumePackerPack(box, items.Copy());
	return assert.Equal(1, packedBox.ItemCount(), "they should be equal")
}
