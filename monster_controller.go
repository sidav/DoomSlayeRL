package main

// AI if I can say so

func ai_decideMove(monster *p_pawn, dung *dungeon) {
	ex, ey := dung.player.getCoords()
	vx, vy := ai_getVectorToTarget(monster, ex, ey)
	m_moveOrMeleeAttackPawn(monster, dung, vx, vy)
}

func ai_getVectorToTarget(monster *p_pawn, ex, ey int) (int, int) { // should be later replaced with pathfinding algorithm
	x, y := monster.getCoords()
	var resx, resy int
	if x == ex {
		resx = 0
	}
	if x > ex {
		resx = -1
	}
	if x < ex {
		resx = 1
	}
	if y == ey {
		resy = 0
	}
	if y > ey {
		resy = -1
	}
	if y < ey {
		resy = 1
	}
	return resx, resy
}
