package axis

// addDistance computes the new position from the given position and the distance
func addDistance(position Position, distance Distance) Position {
	return Position(int64(position) + int64(distance))
}
