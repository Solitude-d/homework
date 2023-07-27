package main

import (
	"fmt"
	myslice "homework/slice"
)

func main() {
	res, _ := myslice.DeleteSliceV3([]float32{1.1, 2.1, 3.3, 4.4}, 2)
	fmt.Printf("%v", res)

}
