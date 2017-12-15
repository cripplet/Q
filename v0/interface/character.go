package q_lib

import (
	qi "github.com/cripplet/Q/lib/interface"
	"math/rand"
)

type MeleeWeapon struct {
	qi.WeaponInterface

	accuracy float64
	damage   float64
}

func (self *MeleeWeapon) GenerateDamageToken() float64 {
	if rand.Float64() < self.accuracy {
		return 0
	}
	return rand.Float64() * self.damage
}

type NStrikes struct {
	qi.AttackInterface
	nStrikes int
}

type AttackEvent struct {
	qi.AttackEventInterface
	target qi.CharacterInterface
	damage float64
}

func (self *NStrikes) GenerateAttack(char *qi.CharacterInterface, weapon *qi.WeaponInterface) []qi.AttackEventInterface {
	c := []qi.AttackEventInterface{}
	for i := 0; i < self.nStrikes; i++ {
		c = append(c, AttackEvent{
			damage: char.Dump().GetLevel() * weapon.GenerateDamageToken(),
		})
	}
	return c
}

type PlayerCharacter struct {
	qi.CharacterInterface

	spells  []qi.WeaponInterface
	melee   qi.WeaponInterface
	missile qi.WeaponInterface
}

/*
type CharacterStats interface {}

type AttackEventInterface interface {}

type AttackInterface interface{ // double, triple slash
	GenerateAttack(weapon WeaponInterface) chan AttackEventInterface
}

type WeaponInterface interface {}

type CharacterInterface interface {
	Load(stats CharacterStats)

	Dump() CharacterStats

	Equip(weapon WeaponInterface)

	Attack(other CharacterInterface, attack AttackInterface)

	Block(other CharacterInterface, attack AttackInterface)

	LevelUp()
}
*/
