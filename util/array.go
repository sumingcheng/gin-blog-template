package util

// ForEach - 遍历数组，对每个元素执行提供的回调函数，没有返回值。
func ForEach[T any](arr []T, callback func(T, int, []T)) {
	for i := 0; i < len(arr); i++ {
		callback(arr[i], i, arr)
	}
}

// Map - 创建一个新切片，包含原切片每个元素经过回调函数处理后的结果。
func Map[T any](arr []T, callback func(T, int, []T) T) []T {
	var _arr []T
	for i := 0; i < len(arr); i++ {
		item := callback(arr[i], i, arr)
		_arr = append(_arr, item)
	}
	return _arr
}

// Filter - 创建一个新切片，只包含符合回调函数条件（即回调函数返回 true）的原切片元素。
// Reduce - 将数组元素通过回调函数累积处理成一个单一的结果，从左到右处理元素。
// ReduceRight - 类似于 Reduce，但是从右到左处理数组元素。
// Every - 检查数组中的所有元素是否都满足指定的测试条件，如果全部满足返回 true，否则返回 false。
// Some - 检查数组中是否至少有一个元素满足指定的测试条件，如果有则返回 true，否则返回 false。
// Push - 向数组的末尾添加一个或多个元素，并返回新的数组长度。
// Unshift - 向数组的开头添加一个元素，并返回新的数组长度。
// Pop - 删除数组的最后一个元素，并返回这个元素。
// Shift - 删除数组的第一个元素，并返回这个元素。
// Reverse - 将数组中的元素顺序反转。
// Splice - 从数组中添加/删除元素，然后返回被删除的元素的切片。
// Slice - 返回数组的一个部分，不改变原数组。
// Find - 返回数组中满足提供的测试函数的第一个元素的指针，如果没有找到符合条件的元素，返回 nil。
// FindIndex - 返回数组中满足提供的测试函数的第一个元素的索引，如果没有找到符合条件的元素，返回 -1。
