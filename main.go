package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	SliceDelete(arr, 0)
}

func SliceDelete[T any](slice []T, idx int) []T {
	if idx < 0 && idx >= len(slice) {
		fmt.Printf("id为%v, 切片长度为%v,请输入正确的下标", idx, arr)
		return nil
	}
	slice = append(slice[:idx], slice[idx+1:]...)
	ReduceCap(slice)
	return slice
}
func ReduceCap[T any](slice []T) []T {
	sliceCap := cap(slice)
	sliceLen := len(slice)
	processCap := sliceCap * 3 / 4
	halfCap := sliceCap / 2
	var newSlice []T

	if sliceCap > 256 && sliceLen < processCap {
		newSlice = make([]T, sliceLen, processCap)
	} else if sliceCap <= 256 && sliceLen < halfCap {
		newSlice = make([]T, sliceLen, halfCap)
	} else {
		return slice
	}

	copy(newSlice, slice)
	return newSlice
}
