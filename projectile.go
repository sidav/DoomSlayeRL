package main

type projectile struct {
	x, y, targetX, targetY, nextTurnToMove, turnsForOneTile int
	damageDice                                              *dice
	cCell                                                   *consoleCell
}
