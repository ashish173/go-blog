package main

func Add(xs []int) int {
	total := int(0)
	for _, x := range xs {
		total += x
	}
	return total
}
