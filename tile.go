package main

type d_doorData struct {
	isOpened     bool
	chrForOpened rune
}

type d_tile struct {
	IsPassable              bool
	opaque, wasSeenByPlayer bool
	cCell                   *consoleCell // is for CLOSED when the tile is a door
	doorData                *d_doorData
}

func (t *d_tile) getAppearance() *consoleCell {
	if t.doorData != nil && t.doorData.isOpened {
		return &consoleCell{t.doorData.chrForOpened, t.cCell.color}
	}
	return t.cCell
}

func (t *d_tile) isPassable() bool {
	if t.doorData != nil && t.doorData.isOpened {
		return true
	}
	return t.IsPassable
}

func (t *d_tile) isOpaque() bool {
	if t.doorData != nil && t.doorData.isOpened {
		return false
	}
	return t.opaque
}
