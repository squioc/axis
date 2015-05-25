package axis

func addDistance(position Position, distance Distance) Position {
    return Position(int64(position) + int64(distance))
}
