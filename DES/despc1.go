package main

import "fmt"

var (
	PC1 = []int{
		57, 49, 41, 33, 25, 17, 9,
		1, 58, 50, 42, 34, 26, 18,
		10, 2, 59, 51, 43, 35, 27,
		19, 11, 3, 60, 52, 44, 36,
		63, 55, 47, 39, 31, 23, 15,
		7, 62, 54, 46, 38, 30, 22,
		14, 6, 61, 53, 45, 37, 29,
		21, 13, 5, 28, 20, 12, 4,
	}

	PC2 = []int{
		14, 17, 11, 24, 1, 5,
		3, 28, 15, 6, 21, 10,
		23, 19, 12, 4, 26, 8,
		16, 7, 27, 20, 13, 2,
		41, 52, 31, 37, 47, 55,
		30, 40, 51, 45, 33, 48,
		44, 49, 39, 56, 34, 53,
		46, 42, 50, 36, 29, 32,
	}

	IterateShiftTable = []int{
		1, 1, 2, 2, 2, 2, 2, 2,
		1, 2, 2, 2, 2, 2, 2, 1,
	}
)

func pc(m uint64, size int, PC []int) uint64 {

	result := uint64(0)
	length := len(PC)

	for index, v := range PC {

		num := (m >> (size - v)) & 0x1

		result |= num << (length - (index + 1))

	}

	return result
}

func partRemove(n uint64, x int) uint64 {
	if x == 1 {
		num := (n&0x7ffffff)<<1 | (n&0x8000000)>>27
		return num
	}

	num := (n&0x3ffffff)<<2 | (n&0xc000000)>>26
	return num

}

func generateKeys(num uint64, keys []uint64) []uint64 {

	output1 := pc(num, 64, PC1)
	l, r := (output1>>28)&0xfffffff, (output1 & 0xfffffff)

	for i := 0; i < 16; i++ {

		lm := partRemove(l, IterateShiftTable[i])
		rm := partRemove(r, IterateShiftTable[i])

		m := lm<<28 | rm

		key := pc(m, 56, PC2)
		keys = append(keys, key)

		l = lm
		r = rm
	}

	return keys

}

//check
func permutation_check(data uint64, size int, n []int) uint64 {
	outSize := len(n)
	result := uint64(0)
	for i := 0; i < outSize; i++ {
		j := n[i]
		bit := (data >> (size - j)) & 1
		result |= bit << (outSize - i - 1)
	}

	return result
}

func leftShift28(data uint64, count int) uint64 {
	if count == 1 {
		return ((data & 0x800_0000) >> 27) | ((data & 0x7ff_ffff) << 1)
	}

	return ((data & 0xc00_0000) >> 26) | ((data & 0x3ff_ffff) << 2)
}

func makeKeys(key64 uint64, subKeys48 []uint64) []uint64 {
	pcKey56 := permutation_check(key64, 64, PC1)
	c28, d28 := (pcKey56>>28)&0x0fffffff, (pcKey56>>0)&0x0fffffff
	for i := 0; i < 16; i++ {
		c28 = leftShift28(c28, IterateShiftTable[i])
		d28 = leftShift28(d28, IterateShiftTable[i])
		cd56 := (c28 << 28) | d28
		subKey48 := permutation_check(cd56, 56, PC2)
		subKeys48 = append(subKeys48, subKey48)
	}

	return subKeys48
}

func main() {

	num := uint64(0x1a2b3c4d5a6b7c8d)
	keys := make([]uint64, 0)

	output1 := generateKeys(num, keys)
	output2 := makeKeys(num, keys)

	fmt.Println("generate keys:", output1)
	//fmt.Printf("cout result: %016x\n", output1)
	fmt.Println("check keys:", output2)
	//fmt.Printf("check result: %016x\n", output2)
}
