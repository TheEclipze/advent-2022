package day9

type Move int

const (
	Up Move = iota
	Down
	Left
	Right
)

type RopeGame struct {
	Head        Position
	Tail        Position
	TailVisited []Position
}

func MakeRopeGame() *RopeGame {
	rg := RopeGame{TailVisited: make([]Position, 0)}

	return &rg
}

type Position struct {
	X int
	Y int
}

// func parseMoves(input []string) []Move {
// 	// match any move (UDRL) and number of them (0 to inf)
// 	r := regexp.MustCompile(`([UDRL])\s([0-9]*)`)

// 	for _, e := range input {
// 		// res := r.FindAllStringSubmatch(e, -1)
// 	// }
// }

func Day9PartOne(input []string) int {
	MakeRopeGame()
	return 0
}

func Day9PartTwo(input []string) int {
	return 0
}
