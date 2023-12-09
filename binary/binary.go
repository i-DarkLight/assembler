package binary

import (
	"github.com/i-DarkLight/memari/hex"
	"strings"
)

var BinaryCode []string

func convert(s string) string {
	var binary string
	for i := range s {
		switch s[i] {
		case '0':
			binary += "0000"
		case '1':
			binary += "0001"
		case '2':
			binary += "0010"
		case '3':
			binary += "0011"
		case '4':
			binary += "0100"
		case '5':
			binary += "0101"
		case '6':
			binary += "0110"
		case '7':
			binary += "0111"
		case '8':
			binary += "1000"
		case '9':
			binary += "1001"
		case 'A':
			binary += "1010"
		case 'B':
			binary += "1011"
		case 'C':
			binary += "1100"
		case 'D':
			binary += "1101"
		case 'E':
			binary += "1110"
		case 'F':
			binary += "1111"
		}
	}
	return binary
}
func ToBinary() {
	defer println("===============================")
	for _, elem := range hex.HexConverted {
		temp := elem
		if leng := len(strings.TrimSpace(elem)); leng == 6 {
			temp = elem[:3] + "0" + elem[3:]
		}
		temp = strings.TrimSpace(temp)
		binaryOfLine := convert(temp)
		BinaryCode = append(BinaryCode, binaryOfLine[:12]+"		"+binaryOfLine[12:])
	}
	for _, elem := range BinaryCode {
		println(elem)
	}
}
