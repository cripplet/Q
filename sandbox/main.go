package main

import (
	"fmt"
	ql "github.com/cripplet/Q/sandbox/lib"
	"math/rand"
	"time"
)

func main() {
	a := ql.NewCharacter("A")
	tech := []ql.Technique{ql.TECHNIQUE_LOOKUP[ql.NORMAL_TECHNIQUE_TYPE]}
	a.Load(
		ql.CHARACTER_CONFIG_LOOKUP[ql.FIGHTER_CLASS_TYPE],
		ql.ARMOR_LOOKUP[ql.NONE_ARMOR_TYPE], ql.SHIELD_LOOKUP[ql.NONE_SHIELD_TYPE],
		[]ql.MeleeWeapon{ql.MELEE_WEAPON_LOOKUP[ql.NONE_MELEE_WEAPON_TYPE]},
		tech,
	)

	b := ql.NewCharacter("B")
	b.Load(
		ql.CHARACTER_CONFIG_LOOKUP[ql.FIGHTER_CLASS_TYPE],
		ql.ARMOR_LOOKUP[ql.NONE_ARMOR_TYPE], ql.SHIELD_LOOKUP[ql.NONE_SHIELD_TYPE],
		[]ql.MeleeWeapon{ql.MELEE_WEAPON_LOOKUP[ql.NONE_MELEE_WEAPON_TYPE]},
		tech,
	)

	var default_a_weapon_key string
	var default_b_weapon_key string
	for default_a_weapon_key, _ = range a.GetMeleeWeapons() {
		break
	}
	for default_b_weapon_key, _ = range b.GetMeleeWeapons() {
		break
	}

	var default_a_technique_key string
	var default_b_technique_key string
	for default_a_technique_key, _ = range a.GetTechniques() {
		break
	}
	for default_b_technique_key, _ = range b.GetTechniques() {
		break
	}

	fmt.Printf("%s: %.2f\n", a.GetName(), a.GetHP())
	fmt.Printf("%s: %.2f\n", b.GetName(), b.GetHP())

	for {
		if succ, _ := a.MeleeAttack(b, default_a_weapon_key, default_a_technique_key); succ {
			fmt.Printf("%s wins\n", a.GetName())
			break
		} else {
			if succ, _ := b.MeleeAttack(a, default_b_weapon_key, default_b_technique_key); succ {
				fmt.Printf("%s wins\n", b.GetName())
				break
			}
		}
	}

	fmt.Printf("%s: %.2f\n", a.GetName(), a.GetHP())
	fmt.Printf("%s: %.2f\n", b.GetName(), b.GetHP())
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
