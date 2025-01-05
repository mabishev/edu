func minimumFromFour() int {
	var w, x, y, z int
	fmt.Scan(&w, &x, &y, &z)
	if w < x && w < y && w < z {
		return w
	}
	if x < w && x < y && x < z {
		return x
	}
	if y < w && y < x && y < z {
		return y
	}
	return z
}
