package main

import "fmt"

var PC1 = []int{
	57, 49, 41, 33, 25, 17, 9,
	1, 58, 50, 42, 34, 26, 18,
	10, 2, 59, 51, 43, 35, 27,
	19, 11, 3, 60, 52, 44, 36,
	63, 55, 47, 39, 31, 23, 15,
	7, 62, 54, 46, 38, 30, 22,
	14, 6, 61, 53, 45, 37, 29,
	21, 13, 5, 28, 20, 12, 4,
}

func pc1(m uint64) uint64 {

	result := uint64(0)

	for index, v := range PC1 {

		num := (m >> (64 - v)) & 0x1

		result |= num << (56 - (index + 1))

	}

	return result
}

func permutation(data uint64, size int, n []int) uint64 {
	outSize := len(n)
	result := uint64(0)
	for i := 0; i < outSize; i++ {
		j := n[i]
		bit := (data >> (size - j)) & 1
		result |= bit << (outSize - i - 1)
	}

	return result
}

func main() {

	m := uint64(0x1122334455667788)

	output1 := pc1(m)
	output2 := permutation(m, 64, PC1)

	fmt.Printf("cout result: %016x\n", output1)
	fmt.Printf("check result: %016x\n", output2)
}
