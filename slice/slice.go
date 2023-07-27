package myslice

import "errors"

func DeleteSliceV1(slice []int, idx int) (res []int, err error) {
	if idx < 0 || idx > len(slice) {
		return res, errors.New("下标越界")
	}
	for i := 0; i < len(slice); i++ {
		if idx != i {
			res = append(res, slice[i])
		}
	}
	return res, nil
}

func DeleteSliceV2(slice []int, idx int) (res []int, err error) {
	if idx < 0 || idx > len(slice) {
		return res, errors.New("下标越界")
	}
	res = append(res, slice[:idx]...)
	res = append(res, slice[idx+1:]...)
	return res, nil
}

func DeleteSliceV3[T Number](slice []T, idx int) (res []T, err error) {
	if idx < 0 || idx > len(slice) {
		return res, errors.New("下标越界")
	}
	if idx == len(slice)-1 {
		res = append(res, slice[:idx]...)
		return res, nil
	}
	res = append(res, slice[:idx]...)
	res = append(res, slice[idx+1:]...)
	return res, nil
}

type Number interface {
	~int | ~int16 | ~int8 | ~int32 | ~int64 |
		~float32 | ~float64
}
