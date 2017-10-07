package q_sandbox_lib

type ClassType int

const (
	FIGHTER_CLASS_TYPE ClassType = 1 << iota
	THIEF_CLASS_TYPE             = 1 << iota
)

var CHARACTER_CONFIG_LOOKUP map[ClassType]CharacterConfig = map[ClassType]CharacterConfig{
	FIGHTER_CLASS_TYPE: CharacterConfig{
		Class: FIGHTER_CLASS_TYPE,
		HP:    25,
		Level: 1,
	},
	THIEF_CLASS_TYPE: CharacterConfig{
		Class: THIEF_CLASS_TYPE,
		HP:    25,
		Level: 1,
	},
}
