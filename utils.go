package main

func sumArray(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printList(list []string) {
	for _, v := range list {
		println(v)
	}
}

func reshape(items []string, size int) [][]string {
	var reshaped [][]string
	assertDivisibleArr(items, size)

	for i := 0; i < len(items); i += size {
		reshaped = append(reshaped, items[i:i+size])
	}

	return reshaped
}

type Queue[T any] struct {
	q []T
}

func (q *Queue[T]) Push(v T) {
	q.q = append(q.q, v)
}

func (q *Queue[T]) Pop() T {
	res := q.q[0]
	q.q = q.q[1:]
	return res
}
