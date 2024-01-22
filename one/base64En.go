package one

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var BaseIndex = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', //0-7
	'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', //8-15
	'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', //16-23
	'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', //24-31
	'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', //32-39
	'o', 'p', 'q', 'r', 's', 't', 'u', 'v', //40-47
	'w', 'x', 'y', 'z', '0', '1', '2', '3', //48-55
	'4', '5', '6', '7', '8', '9', '+', '/', //56-64
}

func CountBase64(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("open error: %s\n", err)
	}

	buf := make([]byte, 3)
	encoding := make([]byte, 0)
	for {

		length, err := file.Read(buf)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				fmt.Printf("read error: %s\n", err)
			}
			break
		}

		encoding = append(encoding, BaseIndex[(buf[0]&0xfc)>>2])

		if length >= 3 {
			encoding = append(encoding, BaseIndex[((buf[0]&0x03)<<4)|((buf[1]&0xf0)>>4)])
			encoding = append(encoding, BaseIndex[((buf[1]&0x0f)<<2)|((buf[2]&0xc0)>>6)])
			encoding = append(encoding, BaseIndex[(buf[2]&0x3f)])
			continue
		}

		if length == 2 {
			encoding = append(encoding, BaseIndex[((buf[0]&0x03)<<4)|((buf[1]&0xf0)>>4)])
			encoding = append(encoding, BaseIndex[((buf[1]&0x0f)<<2)])
			encoding = append(encoding, '=')
			continue
		}

		if length < 2 {
			encoding = append(encoding, BaseIndex[((buf[0]&0x03)<<4)])
			encoding = append(encoding, '=')
			encoding = append(encoding, '=')
		}

	}
	fmt.Println(string(encoding))
	//fmt.Println(encoding)
}

func Ba64En() {
	args := os.Args[1:]

	for _, arg := range args {

		CountBase64(arg)
	}

}
