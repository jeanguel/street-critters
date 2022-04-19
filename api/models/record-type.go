package models

type RecordType string

const (
	RT_POOP   RecordType = "poop"
	RT_ANIMAL RecordType = "dead_animal"
)

func (rt RecordType) IsValid() bool {
	switch rt {
	case RT_POOP,
		RT_ANIMAL:
		return true

	default:
		return false
	}
}
