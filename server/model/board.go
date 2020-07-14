package model

//GetTileColor figures out what color a tile is!
func GetTileColor(x, y int) bool {
	return (x+y)%2 == 0
}
