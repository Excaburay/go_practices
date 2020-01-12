package learntest

func RemoveElement(slice []int, n int) []int{
	copy(slice[n:],slice[n+1:])
	return slice[:len(slice)-1]
}

func NoEmpty(strings []string) []string{
	i :=0
	for _,s := range strings {
		if s != "" {
			strings[i]= s
			i++
		}
	}
	return strings[:i]
}

func ReverseSlice(slice []int) []int{
	for i,j :=0,len(slice)-1 ;i<=j ; i,j = i+1,j-1 {
		slice[i] , slice[j] = slice[j], slice[i]
	}
	return slice
}