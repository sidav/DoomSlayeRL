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
		instantlyPickupable bool
	}
)

func (i *i_item) getType() string {
	if i.weaponData != nil {
		return "weapon"
	}
	if i.ammoData != nil {
		return "ammo"
	}
	if i.medicalData != nil {
		return "medical"
	}
	return "item"
}
