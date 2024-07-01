package util

func Unshift[T any](
	arr *[]T,
	item T,
) int {
	*arr = append([]T{item}, *arr...)
	return len(*arr)
}
