package PathFinder

// Creates The Board
func Board(slice []string) [][]string {
	empty := [][]string{}
	str := []string{}
	for i := 0; i < len(slice); i++ {
		for _, v := range slice[i] {
			str = append(str, string(v))
		}
		empty = append(empty, str)
		str = []string{}
	}
	return empty
}
