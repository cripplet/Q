package main

import (
	"fmt"
	ql "github.com/cripplet/Q/sandbox/lib"
	"math/rand"
	"time"
)

func main() {
	a := ql.NewCharacter("A")
	a.Load(
		ql.CHARACTER_CONFIG_LOOKUP[ql.FIGHTER_CLASS_TYPE],
		ql.ARMOR_LOOKUP[ql.NONE_ARMOR_TYPE], ql.SHIELD_LOOKUP[ql.NONE_SHIELD_TYPE],
		[]ql.MeleeWeapon{ql.MELEE_WEAPON_LOOKUP[ql.NONE_MELEE_WEAPON_TYPE]},
	)

	b := ql.NewCharacter("B")
	b.Load(
		ql.CHARACTER_CONFIG_LOOKUP[ql.FIGHTER_CLASS_TYPE],
		ql.ARMOR_LOOKUP[ql.NONE_ARMOR_TYPE], ql.SHIELD_LOOKUP[ql.NONE_SHIELD_TYPE],
		[]ql.MeleeWeapon{ql.MELEE_WEAPON_LOOKUP[ql.NONE_MELEE_WEAPON_TYPE]},
	)

	fmt.Printf("%s: %.2f\n", a.GetName(), a.GetHP())
	fmt.Printf("%s: %.2f\n", b.GetName(), b.GetHP())

	for {
		if a.MeleeAttack(b, ql.MELEE_WEAPON_LOOKUP[ql.NONE_MELEE_WEAPON_TYPE]) {
			fmt.Printf("%s wins\n", a.GetName())
			break
		} else {
			if b.MeleeAttack(a, ql.MELEE_WEAPON_LOOKUP[ql.NONE_MELEE_WEAPON_TYPE]) {
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
