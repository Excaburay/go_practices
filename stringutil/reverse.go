package stringutil

func Reverse(s string) string{

	//int32的别名，几乎在所有方面等同于int32
	//它用来区分字符值和整数值
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}
