package q_sandbox_lib

import (
	"math"
)

type ShieldType int

const (
	NONE_SHIELD_TYPE ShieldType = iota
	WOODEN_BUCKLER_SHIELD_TYPE
)

type Shield struct {
	ToBlock           float64
	HP                float64
	ClassRestrictions ClassType
}

func (self *Shield) Block(damage float64) float64 {
	if self.HP == 0 || damage < self.ToBlock {
		return damage
	}
	damage = math.Min(self.HP, damage)
	self.HP -= damage
	return damage
}

var SHIELD_LOOKUP map[ShieldType]Shield = map[ShieldType]Shield{
	NONE_SHIELD_TYPE: Shield{
		ClassRestrictions: FIGHTER_CLASS_TYPE | THIEF_CLASS_TYPE,
	},
	WOODEN_BUCKLER_SHIELD_TYPE: Shield{
		ToBlock:           10,
		HP:                10,
		ClassRestrictions: FIGHTER_CLASS_TYPE | THIEF_CLASS_TYPE,
	},
}
