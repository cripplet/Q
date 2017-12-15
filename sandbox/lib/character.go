package q_sandbox_lib

import (
	"errors"
	"math"
	"math/rand"
)

type MissileWeapon struct{}
type Spell struct{}

type Character struct {
	name         string
	stats        CharacterConfig
	meleeWeapons map[string]MeleeWeapon
	// missile []MissileWeapon
	// spells []Spell
	techniques map[string]Technique
	armor      Armor
	shield     Shield
	// mp int
}

func NewCharacter(name string) *Character {
	c := Character{
		name:         name,
		meleeWeapons: make(map[string]MeleeWeapon),
		techniques: make(map[string]Technique),
	}
	return &c
}

func (self *Character) Load(
	config CharacterConfig,
	armor Armor,
	shield Shield,
	meleeWeapons []MeleeWeapon,
	techniques []Technique) error {
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
	println("loading weapons")
	for _, meleeWeapon := range meleeWeapons {
		e = self.EquipMeleeWeapon(meleeWeapon)
		if e != nil {
			panic(e)
			return e
		}
	}

	println("loading techniques")
	for _, technique := range techniques {
		e = self.EquipTechnique(technique)
		println(technique.Name)
		if e != nil {
			return e
		}
	}
	return e
}

func (self *Character) LoadConfig(config CharacterConfig) {
	self.stats = config
}

func (self *Character) GetMeleeWeapons() map[string]MeleeWeapon {
	return self.meleeWeapons
}

func (self *Character) GetTechniques() map[string]Technique {
	return self.techniques
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
		self.meleeWeapons[String(5)] = weapon
	}
	return errors.New("Melee weapon class restriction mismatch.")
}

func (self *Character) EquipTechnique(technique Technique) error {
	if self.stats.Class & technique.ClassRestrictions == self.stats.Class {
		self.techniques[String(5)] = technique
	}
	return errors.New("Technique class restriction mismatch.")
}

func (self *Character) GetHP() float64 {
	return self.stats.HP
}

func (self *Character) GetName() string {
	return self.name
}

func (self *Character) MeleeAttack(
	other *Character,
	meleeWeaponKey string,
	techniqueKey string) (bool, error) {
	w, isPresent := self.meleeWeapons[meleeWeaponKey]
	if !isPresent {
		return false, errors.New("Invalid meleeWeaponKey.")
	}

	/*
	_, isPresent = self.techniques[techniqueKey]
	if !isPresent {
		return false, errors.New("Invalid techniqueKey.")
	}*/

	// Calculate hit percentage.
	damage := (rand.Float64() * float64(w.BaseDamage)) * float64(self.stats.Level) // TODO(cripplet): Use dice notation, see github.com/justinian/dice.
	win := other.Block(damage)
	if win {
		self.stats.XP += float64(other.stats.Level)
	}
	return win, nil
}

func (self *Character) Block(damage float64) bool {
	damage = self.shield.Block(damage)
	damage = math.Min(self.armor.Block(damage), self.GetHP())
	self.stats.HP -= damage
	return self.GetHP() == 0
}
