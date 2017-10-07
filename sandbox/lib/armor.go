package q_sandbox_lib

import (
	"math"
)

type ArmorType int

const (
	NONE_ARMOR_TYPE ArmorType = iota
	SHIRT_ARMOR_TYPE
)

type Armor struct {
	AbsorptionRatio   float64
	ToBlock           float64
	HP                float64
	ClassRestrictions ClassType
}

func (self *Armor) Block(damage float64) float64 {
	if damage < self.ToBlock {
		return damage
	} else {
		absorbed := math.Min((damage * self.AbsorptionRatio), self.HP)
		self.HP -= absorbed
		return damage - absorbed
	}
}

var ARMOR_LOOKUP map[ArmorType]Armor = map[ArmorType]Armor{
	NONE_ARMOR_TYPE: Armor{
		ClassRestrictions: FIGHTER_CLASS_TYPE | THIEF_CLASS_TYPE,
	},
	SHIRT_ARMOR_TYPE: Armor{
		AbsorptionRatio:   0.1,
		ToBlock:           10,
		HP:                10,
		ClassRestrictions: FIGHTER_CLASS_TYPE | THIEF_CLASS_TYPE,
	},
}
