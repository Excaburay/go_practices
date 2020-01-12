package learntest

func IsExistElement(ages map[string]int) bool {
	if _, ok := ages["bob"]; !ok{
		return false
	}
	return true
}

func MapEqual(x, y map[string]int) bool {
	if len(x) != len(y){
		return false
	}
	for key, xv := range x {
		if yv, ok := y[key]; !ok || xv != yv{
			return false
		}
	}
	return true
}