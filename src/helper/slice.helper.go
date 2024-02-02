package helper

func SliceCut(index int, slice []any) (any, []any) {
	return slice[index], append(slice[:index], slice[index+1:]...)
}
