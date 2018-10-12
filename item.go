package main

type (
	i_medicalData struct {
		healAmount     int
		ignoresMaximum bool
	}

	i_item struct {
		x, y                int
		ccell               *consoleCell
		name                string
		weaponData          *i_weaponData
		ammoData            *i_ammoData
		medicalData         *i_medicalData
		armorData           *i_armorData
		instantlyPickupable bool
	}
)

func (i *i_item) getType() string {
	switch {

	case i.weaponData != nil:
		return "weapon"

	case i.ammoData != nil:
		return "ammo"

	case i.medicalData != nil:
		return "medical"

	case i.armorData != nil:
		return "armor"

	default:
		return "item"

	}
}
