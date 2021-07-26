package goSnek

func moveSnek(pos Body, dir Body) Body {
	b := Body{
		x: pos.x + dir.x,
		y: pos.y + dir.y,
	}

	return b
}

func removeAt(arr []Body, index int) []Body {
	return append(arr[:index], arr[index+1:]...)
}

func bodyEquals(a Body, b Body) bool {
	if a.x == b.x && a.y == b.y {
		return true
	}
	return false
}
