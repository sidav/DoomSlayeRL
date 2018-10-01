package procedurals

import "GoRoguelike/routines"

func randInRange(from, to int) int { //should be inclusive
	if to < from {
		t := from
		from = to
		to = t
	}
	if from == to {
		return from
	}
	return routines.Random(to-from) + from // TODO: replace routines.random usage with package own implementation
}

const (
	MAX_SPLITS        = 5
	MAP_W             = 80
	MAP_H             = 20
	SPLIT_MIN_PERCENT = 40
	SPLIT_MAX_PERCENT = 100 - SPLIT_MIN_PERCENT
	MIN_ROOM_W        = 3
	MIN_ROOM_H        = 3
	TRIES_FOR_SPLITTING
)

type container struct {
	x, y, w, h int
}

type treeNode struct {
	parent, left, right *treeNode
	room                *container
}

func getSplitRangeForPercent(wh int, percent int) (int, int) {
	min := wh * percent / 100
	return min, wh-min
}

func (t *treeNode) splitHoriz() { // splits node into "lower" and "upper"
	current_w := t.room.w
	current_h := t.room.h
	current_x := t.room.x
	current_y := t.room.y
	minSplSize, maxSplSize := getSplitRangeForPercent(current_h, SPLIT_MIN_PERCENT)
	// Let's try to split the node without breaking min room size constraints
	for try := 0; try < TRIES_FOR_SPLITTING; try ++ {
		upper_h := randInRange(minSplSize, maxSplSize)
		lower_h := current_h - upper_h
		if upper_h < MIN_ROOM_H || lower_h < MIN_ROOM_H {
			continue
		} else { // Okay, sizes are acceptable. Let's do the split
			upperNode := treeNode{parent: t, room: &container{x: current_x, y: current_y, w: current_w, h: upper_h}}
			// Most error-probable place:
			lowerNode := treeNode{parent: t, room: &container{x: current_x, y: current_y+upper_h, w: current_w, h: lower_h}}
			// hm... Left is upper and right is lower. Everything is obvious.
			t.left = &upperNode
			t.right = &lowerNode
			return
		}
	}
}

func (t *treeNode) splitVert() { // splits node into left and right
	current_w := t.room.w
	current_h := t.room.h
	current_x := t.room.x
	current_y := t.room.y
	minSplSize, maxSplSize := getSplitRangeForPercent(current_w, SPLIT_MIN_PERCENT)
	// Let's try to split the node without breaking min room size constraints
	for try := 0; try < TRIES_FOR_SPLITTING; try ++ {
		left_w := randInRange(minSplSize, maxSplSize)
		right_w := current_w - left_w
		if left_w < MIN_ROOM_H || right_w < MIN_ROOM_H {
			continue
		} else { // Okay, sizes are acceptable. Let's do the split
			leftNode := treeNode{parent: t, room: &container{x: current_x, y: current_y, w: left_w, h: current_h}}
			// Most error-probable place:
			rightNode := treeNode{parent: t, room: &container{x: current_x+left_w, y: current_y, w: right_w, h: current_h}}
			t.left = &leftNode
			t.right = &rightNode
			return
		}
	}
}

/////////////////////////////////////////

var (
	dmap     [][]rune
	treeRoot *treeNode
)

func generateDungeon() {
	// generate parent node
	treeRoot = &treeNode{room: &container{x: 0, y: 0, w: MAP_W, h: MAP_H}}
}
