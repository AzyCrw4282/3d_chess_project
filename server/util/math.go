package util

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func OrderPoints(p1, p2 int) (int, int) {
	if p1 <= p2 {
		return p1, p2
	}
	return p2, p1
}

//TODO: iff untilTiles need to be checked for differnt direction
func GetDirection(p1, p2 int) int {
	if p1 == p2 {
		return 0
	} else if p1 <= p2 {
		return 1
	}
	return -1
}
