package coordinate

type Position struct {
	X int
	Y int
}

type Direction Position

var UP Direction = Direction{X: 0, Y: -1}
var DOWN Direction = Direction{X: 0, Y: 1}
var RIGHT Direction = Direction{X: 1, Y: 0}
var LEFT Direction = Direction{X: -1, Y: 0}

func (a Position) Move(dir Direction) Position {
	return Position{X: a.X + dir.X, Y: a.Y + dir.Y}
}

func (a Direction) Turn90() Direction {
	switch a {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	}

	panic("Unexpected direction, verify dirction variable")
}
