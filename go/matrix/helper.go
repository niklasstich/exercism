package matrix

// Little helper function in the spirit of LINQ's 'where' function
func filter(slice []string, filter func(string) bool) (ret []string) {
	for _, val := range slice {
		if filter(val) {
			ret = append(ret, val)
		}
	}
	return
}
