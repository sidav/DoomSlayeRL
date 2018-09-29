package main

func movePawn(p *pawn, d *dungeon, x, y int) {
	// px, py := p.x, p.y
	nx, ny := p.x+x, p.y+y
	if d.isTilePassable(nx, ny) {
		p.x += x
		p.y += y
	}
}
