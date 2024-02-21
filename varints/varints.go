package varints

import (
	"encoding/binary"
	"fmt"
)

func encodeVarint(num uint64) []byte {

	var buf []byte

	for {
		b := byte(num & 0x7f)
		num = num >> 7
		if num != 0 {
			b |= 0x80
		}

		buf = append(buf, b)
		if num == 0 {
			break
		}
	}

	return buf

}

func decodeVarint(buf []byte) uint64 {

	length := len(buf)
	b := uint64(0)

	for i, v := range buf {
		if i < length {
			b |= uint64(v&0x7f) << (7 * i)
		}
	}
	return b
}

func varintsTest(a uint64) []byte {

	b := make([]byte, 10)

	binary.PutUvarint(b, uint64(a))

	return b
}

func Vints() {

	a1 := uint64(9876543216549876513)

	enput := encodeVarint(a1)
	fmt.Printf("%x\n", enput)

	deput := decodeVarint(enput)
	fmt.Printf("%v\n", deput)

	tput := varintsTest(a1)
	fmt.Printf("B: %x\n", tput)
}
