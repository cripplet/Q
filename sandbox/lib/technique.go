package q_sandbox_lib

type TechniqueType int

const (
	NORMAL_TECHNIQUE_TYPE TechniqueType = iota
	DOUBLE_STRIKE_TECHNIQUE_TYPE
)

// TODO(cripplet): Add strike penalties (e.g. accuracy debuff).
// TODO(cripplet): Add special attacks, a la SolarBeam.
type StrikeWeight int

type Technique struct {
	Name              string
	Weights           []StrikeWeight
	ClassRestrictions ClassType
	// TODO(cripplet): Add AttackType (MELEE|MISSILE|SPELL) restriction whitelist.
}

var TECHNIQUE_LOOKUP map[TechniqueType]Technique = map[TechniqueType]Technique{
	NORMAL_TECHNIQUE_TYPE: Technique{
		Name:              "Normal Strike",
		Weights:           []StrikeWeight{100},
		ClassRestrictions: FIGHTER_CLASS_TYPE | THIEF_CLASS_TYPE,
	},
}
