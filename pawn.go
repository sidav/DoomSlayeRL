package main

type (
	p_meleeAttackData struct {
		meleeAttackString string
		// 3d6 + 1 == dnum d dval + dmod
		dnum, dval, dmod int
	}
	p_playerData struct {
	}
	p_pawn struct {
		appearance      rune
		hp, maxhp, x, y int
		name            string
		melee           *p_meleeAttackData
		playerData      *p_playerData
	}
)

func (m *p_meleeAttackData) rollForDamage() int {
	return rollDice(m.dnum, m.dval, m.dmod)
}

func (p *p_pawn) isPlayer() bool {
	return p.playerData != nil
}

func (p *p_pawn) getCoords() (int, int) {
	return p.x, p.y
}
