package main

import "fmt"

// It doesn't matter if the innerT1 struct unexported or exported, the result is the same.
// It also doesn't matter if the R type parameter is infer to a pointer or not, the result is the same.
type innerT1[T any, R *T1[T]] struct {
	Ref   *R
	origT T
}

type T1[T any] struct {
	T innerT1[T, *T1[T]]
}

func main() {

	iT := innerT1[int, *T1[int]]{
		origT: 2,
	}

	t1 := T1[int]{T: iT}

	fmt.Println(t1)
	fmt.Println(iT)

	fmt.Println("didn't panic")
}
