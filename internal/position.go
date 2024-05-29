package internal

type Coordinate struct {
	X int
	Y int
}

type Position struct {
	Name string
}

const (
	POS_TL = "TL"
	POS_TR = "TR"
	POS_BL = "BL"
	POS_BR = "BR"
)

func (p *Position) GetCoordinate(appSize Coordinate, screenSize Coordinate) *Coordinate {
	switch p.Name {
	case POS_TL: //TOP LEFT
		return &Coordinate{X: 0, Y: 0}
	case POS_TR: //TOP RIGHT
		return &Coordinate{X: screenSize.X - appSize.X, Y: 0}
	case POS_BL: //BOTTOM LEFT
		return &Coordinate{X: 0, Y: screenSize.Y - appSize.Y}
	case POS_BR: //BOTTOM RIGHT
		return &Coordinate{X: screenSize.X - appSize.X, Y: screenSize.Y - appSize.Y}
	default:
		//DEFAULT POSITION: CENTER
		return nil
	}
}
