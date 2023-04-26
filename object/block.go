package object

type Block struct {
	Vitality Vital
}

type Vital int

const (
	ALIVE Vital = 0
	DEAD  Vital = 1
)
