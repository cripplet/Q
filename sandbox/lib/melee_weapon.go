package q_sandbox_lib

type MeleeWeaponType int

const (
	NONE_MELEE_WEAPON_TYPE MeleeWeaponType = iota
	DAGGER_MELEE_WEAPON_TYPE
)

type MeleeWeapon struct {
	Name              string
	BaseDamage        float64
	ClassRestrictions ClassType
}

var MELEE_WEAPON_LOOKUP map[MeleeWeaponType]MeleeWeapon = map[MeleeWeaponType]MeleeWeapon{
	NONE_MELEE_WEAPON_TYPE: MeleeWeapon{
		Name:              "None",
		BaseDamage:        1,
		ClassRestrictions: FIGHTER_CLASS_TYPE | THIEF_CLASS_TYPE,
	},
	DAGGER_MELEE_WEAPON_TYPE: MeleeWeapon{
		Name:              "Dagger",
		BaseDamage:        5,
		ClassRestrictions: FIGHTER_CLASS_TYPE | THIEF_CLASS_TYPE,
	},
}
