package main

type (
	p_meleeAttackData struct {
		meleeAttackString string
		// 3d6 + 1 == dnum d dval + dmod
		damageDice *dice
	}
	p_playerData struct {
	}
	p_aiData struct {
		state ai_aiState
		stateTimeoutTurn int
		currentTarget *p_pawn
		targetx, targety int
	}
	p_pawn struct {
		appearance      rune
		hp, maxhp, x, y int
		name            string
		meleeData       *p_meleeAttackData
		playerData      *p_playerData
		weaponInHands   *i_item
		inventory       *inventory
		aiData *p_aiData
	}
)

func (p *p_pawn) canMelee() bool {
	return p.meleeData != nil
}

func (p *p_pawn) isPlayer() bool {
	return p.playerData != nil
}

func (p *p_pawn) getCoords() (int, int) {
	return p.x, p.y
}

func (p *p_pawn) getHpPercent() int {
	return p.hp * 100 / p.maxhp
}
