package main

func turnCostFor(actionName string) int { // TODO: make this accept pawn as an argument
	switch actionName {
	case "melee_attack":
		return 12
	case "ranged_attack":
		return 10
	case "glory_kill":
		return 17
	case "shoot":
		return 15
	case "step":
		return 10
	case "step_diag":
		return 14
	default:
		log.appendMessagef("Turn cost for unknown action %s requested!", actionName)
		return 10
	}
}
