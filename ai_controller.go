package main

import "GoRoguelike/routines"

// AI if I can say so

type (
	ai_aiState byte

	p_aiData struct {
		state            ai_aiState
		stateTimeoutTurn int
		currentTarget    *p_pawn
		targetx, targety int
	}
)

const (
	AI_SILENT    ai_aiState = 0
	AI_ROAMING   ai_aiState = 1
	AI_ENGAGING  ai_aiState = 2
	AI_STAGGERED ai_aiState = 3
)

func ai_decideMove(monster *p_pawn, dung *dungeon) {
	ai_reactToSurroundings(monster, dung)
	aiData := monster.aiData
	switch aiData.state {
	case AI_ENGAGING:
		ex, ey := dung.player.getCoords()
		vx, vy := ai_getVectorToTarget(monster, ex, ey)
		if monster.canShoot() {
			m_rangedAttack(monster, monster.aiData.currentTarget, dung)
		}
		m_moveOrMeleeAttackPawn(monster, dung, vx, vy)
		return
	case AI_ROAMING:
		stepx, stepy := routines.RandomUnitVectorInt()
		m_movePawn(monster, dung, stepx, stepy)
	}
}

func ai_reactToSurroundings(monster *p_pawn, dung *dungeon) { //change state if sees something or whatever
	mx, my := monster.getCoords()
	ex, ey := dung.player.getCoords()
	aiData := monster.aiData
	if monster.aiData.state == AI_STAGGERED {
		if monster.aiData.stateTimeoutTurn <= CURRENT_TURN {
			monster.aiData.state = AI_ROAMING
		} else {
			return
		}
	}
	if dung.visibleLineExists(mx, my, ex, ey) {
		aiData.state = AI_ENGAGING
		aiData.currentTarget = dung.player
		aiData.targetx, aiData.targety = ex, ey
	} else {
		switch aiData.state {
		case AI_ENGAGING:
			aiData.state = AI_ROAMING
			return
		}
	}
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
