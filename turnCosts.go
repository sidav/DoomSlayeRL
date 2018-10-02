package main

func turnCostFor(actionName string) int {
	switch actionName {
	case "shoot":
		return 15
	case "step":
		return 10
	default:
		return 10 	
	}
}
