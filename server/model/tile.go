package model

import "github.com/team142/angrychess/util"

//A tile struct for piece on the board
type Tile struct {
	X, Y int
}

func (t *Tile) Equal(other *Tile) bool {
	return t.X == other.X && t.Y == other.Y
}

//get the tiles by calculating the tiles left between 2 struct. Then closes the channel and returns the channel
//only checks diagonal distance and not other possiblities.
func (t *Tile) GetTilesUntil(end *Tile) chan Tile {
	c := make(chan Tile)
	go func() {
		xd := util.GetDirection(t.X, end.X)
		yd := util.GetDirection(t.Y, end.Y)
		current := Tile{X: t.X, Y: t.Y}
		for {
			current.X += xd
			current.Y += yd
			c <- current //adds the value to the channel
			if current.Equal(end) {
				close(c)
				return
			}
		}
	}()
	return c
}
