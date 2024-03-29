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

var BASE64_DECODE_MAP = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // 0-15
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //16-31
	//                                                                  + ,                     / ,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x3f, //32-47
	// 0,  1  ,  2  ,  3  ,  4  ,  5  ,  6  ,  7  ,  8  ,  9  ,  :  ,  ;  ,  <  ,  =  ,  >  ,  ?  ,
	0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //48-63
	//@ ,  A  ,  B,    C  ,  D  ,  E  ,   F ,  G  ,  H  ,  I  ,  J  ,   K ,   L ,  M  ,  N  ,  O  ,
	0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, //64-79
	//P ,  Q  ,  R  ,   S ,  T  ,  U  ,  V  ,  W  ,  X  ,  Y  ,  Z
	0x0F, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00, //80-95
	//  ,   a ,  b  ,   c ,   d ,   e ,   f ,  g  ,   h ,  i  ,   j ,   k ,   l ,   m ,   n ,   o ,
	0x00, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, //96-111
	// p,   q ,   r ,   s ,   t ,   u ,   v ,   w ,   x ,   y ,   z
	0x29, 0x2A, 0x2B, 0x2C, 0x2D, 0x2E, 0x2F, 0x30, 0x31, 0x32, 0x33, 0x00, 0x00, 0x00, 0x00, 0x00, //112-127
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //128-143
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //144-159
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //160-175
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //176-191
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //192-207
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //208-223
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //224-239
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //240-256
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

func CountBase64de(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("open error: %s\n", err)
		return
	}
	defer file.Close()

	buf := make([]byte, 4)

	for {
		length, err := file.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("read error: %s\n", err)
			return
		}

		a := make([]byte, 3)
		a[0] = (BASE64_DECODE_MAP[buf[0]])<<2 | (BASE64_DECODE_MAP[buf[1]])>>4
		a[1] = (BASE64_DECODE_MAP[buf[1]])<<4 | (BASE64_DECODE_MAP[buf[2]])>>2
		a[2] = (BASE64_DECODE_MAP[buf[2]])<<6 | BASE64_DECODE_MAP[buf[3]]

		if length > 1 {
			if buf[length-1] == '=' {
				a[2] = 0
				if buf[length-2] == '=' {
					a[1] = 0
				}
			}
		}

		fmt.Printf(string(a))
	}

}

func Ba64() {

	args := os.Args[1:]

	if args[0] == "-e" {
		for _, arg := range args[1:] {
			CountBase64(arg)
		}
	}

	if args[0] == "-d" {
		for _, arg := range args[1:] {
			CountBase64de(arg)
		}
	}

}
