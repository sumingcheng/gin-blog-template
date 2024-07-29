package util

import "errors"

// ForEach - 遍历数组，对每个元素执行提供的回调函数，没有返回值。
func ForEach[T any](arr []T, callback func(T, int, []T)) {
	for i := 0; i < len(arr); i++ {
		callback(arr[i], i, arr)
	}
}

// Map - 创建一个新切片，包含原切片每个元素经过回调函数处理后的结果。
func Map[T any](arr []T, callback func(T, int, []T) T) []T {
	_arr := make([]T, 0, len(arr)) // 预分配足够的空间
	for i, value := range arr {
		item := callback(value, i, arr)
		_arr = append(_arr, item)
	}
	return _arr
}

// Filter - 返回一个新的数组，包含通过 callback 测试的所有元素。如果没有任何元素通过测试，则返回一个空数组。
func Filter[T any](arr []T, callback func(T, int, []T) bool) []T {
	var result []T
	for i, value := range arr {
		if callback(value, i, arr) {
			result = append(result, value)
		}
	}
	return result
}

// Reduce - 将数组元素通过回调函数累积处理成一个单一的结果，从左到右处理元素。
func Reduce[T any, U any](slice []T, initial U, reducer func(U, T) U) U {
	accumulator := initial
	for _, value := range slice {
		accumulator = reducer(accumulator, value)
	}
	return accumulator
}

// ReduceRight - 从右到左处理数组元素。
func ReduceRight[T any, U any](slice []T, initial U, reducer func(U, T) U) U {
	accumulator := initial
	// 从后向前遍历切片
	for i := len(slice) - 1; i >= 0; i-- {
		accumulator = reducer(accumulator, slice[i])
	}
	return accumulator
}

// Every - 检查数组中的所有元素是否都满足指定的测试条件，如果全部满足返回 true，否则返回 false。
func Every[T any](arr []T, callback func(T, int, []T) bool) bool {
	for i, value := range arr {
		if !callback(value, i, arr) {
			return false
		}
	}
	return true
}

// Some - 检查数组中是否至少有一个元素满足指定的测试条件，如果有则返回 true，否则返回 false。
func Some[T any](arr []T, callback func(T, int, []T) bool) bool {
	for i, value := range arr {
		if callback(value, i, arr) {
			return true
		}
	}
	return false
}

// Push 函数向切片的末尾添加一个元素，并返回新的切片以及新的长度
func Push[T any](arr []T, elements ...T) ([]T, int) {
	newArr := append(arr, elements...) // 接收任意数量的 element
	return newArr, len(newArr)         // 返回新的切片和长度
}

// Unshift - 向数组的开头添加一个元素，并返回新的切片以及新的长度。
func Unshift[T any](slice []T, elements ...T) ([]T, int) {
	newSlice := append(elements, slice...)
	return newSlice, len(newSlice)
}

// Pop 从切片末尾移除一个元素
// 返回被移除的元素，更新后的切片，以及可能发生的错误
func Pop[T any](slice []T) (T, []T, error) {
	if len(slice) == 0 {
		var zeroValue T
		return zeroValue, nil, errors.New("cannot pop from an empty slice")
	}

	// 获取切片最后一个元素
	element := slice[len(slice)-1]
	// 移除切片的最后一个元素
	slice = slice[:len(slice)-1]

	return element, slice, nil
}

// Shift - 删除数组的第一个元素
func Shift[T any](slice []T) (T, []T, error) {
	if len(slice) == 0 {
		var zeroValue T
		return zeroValue, nil, errors.New("cannot shift from an empty slice")
	}

	// 获取切片的第一个元素
	element := slice[0]
	// 移除切片的第一个元素
	slice = slice[1:]
	return element, slice, nil
}

// Reverse - 将数组中的元素顺序反转。
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Splice - 从数组中添加/删除元素，然后返回被删除的元素的切片。
func Splice[T any](slice []T, start, deleteCount int, items ...T) ([]T, []T) {
	if start < 0 {
		start = len(slice) + start
	}
	if start > len(slice) {
		start = len(slice)
	}
	if deleteCount < 0 {
		deleteCount = 0
	}
	if start+deleteCount > len(slice) {
		deleteCount = len(slice) - start
	}

	// 要删除的元素
	deleted := make([]T, deleteCount)
	copy(deleted, slice[start:start+deleteCount])

	// 拼接后的新切片
	newSlice := append(slice[:start], items...)
	newSlice = append(newSlice, slice[start+deleteCount:]...)

	return newSlice, deleted
}

// Slice - 返回数组的一个部分，不改变原数组。
func Slice[T any](slice []T, start int, end ...int) []T {
	if start < 0 {
		start = len(slice) + start
		if start < 0 {
			start = 0
		}
	}
	if start > len(slice) {
		start = len(slice)
	}

	sliceEnd := len(slice)
	if len(end) > 0 && end[0] >= 0 {
		sliceEnd = end[0]
		if sliceEnd > len(slice) {
			sliceEnd = len(slice)
		}
	}

	if start > sliceEnd {
		return []T{}
	}

	return slice[start:sliceEnd]
}

// Find 查找切片中第一个满足条件的元素。
// 返回找到的元素和一个布尔值，指示是否找到元素。
func Find[T any](slice []T, test func(T) bool) (T, bool) {
	var zeroValue T
	for _, elem := range slice {
		if test(elem) {
			return elem, true
		}
	}
	return zeroValue, false
}

// FindIndex 返回切片中第一个满足条件的元素的索引。
// 如果没有元素满足条件，返回 -1。
func FindIndex[T any](slice []T, test func(T) bool) int {
	for index, elem := range slice {
		if test(elem) {
			return index
		}
	}
	return -1 // 如果没有元素满足条件，返回 -1
}
