package repeatnum03

//the start search point can begain with the upper right corner or lower left
func Find(table [][]int, target int) bool {
	rlen := len(table)
	clen := len(table[0])

	//we look for the upper right corner first
	for r, c := 0, clen-1; r < rlen && c >= 0; {
		if table[r][c] == target {
			return true
		}
		if table[r][c] > target {
			c--
			continue
		} else {
			r++
		}
	}
	return false
}
