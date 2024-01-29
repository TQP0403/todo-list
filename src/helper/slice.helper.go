package helper

func SliceCut(index int, slice []any) (any, []any) {
	res := slice[index]
	nSli := append(slice[:index], slice[index+1:]...)
	return res, nSli
}
