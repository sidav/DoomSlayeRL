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
	AI_SHOOT_CHANCE                   = 15
	AI_STEP_BACK_CHANCE               = 80
	AI_STEP_BACK_THRESHOLD            = 3 * 3
	AI_SILENT              ai_aiState = 0
	AI_ROAMING             ai_aiState = 1
	AI_ENGAGING            ai_aiState = 2
	AI_STAGGERED           ai_aiState = 3
)

func ai_decideMove(monster *p_pawn, dung *dungeon) {
	ai_reactToSurroundings(monster, dung)
	aiData := monster.aiData
	switch aiData.state {
	case AI_ENGAGING:
		ex, ey := dung.player.getCoords()
		vx, vy := ai_getVectorToTarget(monster, ex, ey)
		if monster.canShoot() {
			// if the distance is less than threshold, then step back. Maybe.
			if getSqDistance(monster.x, monster.y, ex, ey) < AI_STEP_BACK_THRESHOLD && routines.RandomPercent() < AI_STEP_BACK_CHANCE {
				m_movePawn(monster, dung, -vx, -vy)
			}
			if routines.RandomPercent() < AI_SHOOT_CHANCE {
				m_rangedAttack(monster, monster.aiData.currentTarget, dung)
				return
			}
		}
		m_moveOrMeleeAttackPawn(monster, dung, vx, vy)
		return
	case AI_ROAMING:
		ai_roam(monster, dung)
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

func ai_roam(monster *p_pawn, dung *dungeon) {
	stepx, stepy := routines.RandomUnitVectorInt()
	m_movePawn(monster, dung, stepx, stepy)
}

func getSqDistance(fx, fy, tx, ty int) int {
	return (fx-tx)*(fx-tx) + (fy-ty)*(fy-ty)
}
