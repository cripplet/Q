package q_sandbox_lib

import (
	"errors"
	"math"
	"math/rand"
)

type MissileWeapon struct{}
type Spell struct{}
type Technique struct{}

type Character struct {
	name         string
	stats        CharacterConfig
	meleeWeapons []MeleeWeapon
	// missile []MissileWeapon
	// spells []Spell
	// techniques []Technique
	armor  Armor
	shield Shield
	// mp int
}

func NewCharacter(name string) *Character {
	c := Character{
		name: name,
	}
	return &c
}

func (self *Character) Load(config CharacterConfig, armor Armor, shield Shield, meleeWeapons []MeleeWeapon) error {
	self.LoadConfig(config)
	var e error = nil
	e = self.EquipArmor(armor)
	if e != nil {
		return e
	}
	e = self.EquipShield(shield)
	if e != nil {
		return e
	}
	for _, meleeWeapon := range meleeWeapons {
		e = self.EquipMeleeWeapon(meleeWeapon)
		if e != nil {
			return e
		}
	}
	return e
}

func (self *Character) LoadConfig(config CharacterConfig) {
	self.stats = config
}

func (self *Character) EquipArmor(armor Armor) error {
	if self.stats.Class&armor.ClassRestrictions == self.stats.Class {
		self.armor = armor
		return nil
	}
	return errors.New("Armor class restriction mismatch.")
}

func (self *Character) EquipShield(shield Shield) error {
	if self.stats.Class&shield.ClassRestrictions == self.stats.Class {
		self.shield = shield
		return nil
	}
	return errors.New("Shield class restriction mismatch.")
}

func (self *Character) EquipMeleeWeapon(weapon MeleeWeapon) error {
	if self.stats.Class&weapon.ClassRestrictions == self.stats.Class {
		self.meleeWeapons = append(self.meleeWeapons, weapon)
	}
	return errors.New("Melee weapon class restriction mismatch.")
}

func (self *Character) GetHP() float64 {
	return self.stats.HP
}

func (self *Character) GetName() string {
	return self.name
}

func (self *Character) MeleeAttack(other *Character, w MeleeWeapon) bool { // TODO(cripplet): Find a way to use a Character.meleeWeapons reference instead of an instance as an arg.
	// Calculate hit percentage.
	damage := (rand.Float64() * float64(w.BaseDamage)) * float64(self.stats.Level) // TODO(cripplet): Use dice notation, see github.com/justinian/dice.
	win := other.Block(damage)
	if win {
		self.stats.XP += float64(other.stats.Level)
	}
	return win
}

func (self *Character) Block(damage float64) bool {
	damage = self.shield.Block(damage)
	damage = math.Min(self.armor.Block(damage), self.GetHP())
	self.stats.HP -= damage
	return self.GetHP() == 0
}
