package hex
import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)
var hexConverted []string
var decimalorg uint
var mp = make(map[string]string)
func isMemory(s string) bool {
	memoryIR := []string{"LDA", "STA", "ADD", "AND", "BUN", "BSA"}
	for _, elem := range memoryIR {
		if elem == s {
			return true
		}
	}
	return false
}
func getDECofHEX(s string) uint {
	var dec float64
	for i := range s {
		if s[i] == 'F' {
			dec += 15 * (math.Pow(16, float64(len(s)-1-i)))
		} else if s[i] == 'E' {
			dec += 14 * (math.Pow(16, float64(len(s)-1-i)))
		} else if s[i] == 'D' {
			dec += 13 * (math.Pow(16, float64(len(s)-1-i)))
		} else if s[i] == 'C' {
			dec += 12 * (math.Pow(16, float64(len(s)-1-i)))
		} else if s[i] == 'B' {
			dec += 11 * (math.Pow(16, float64(len(s)-1-i)))
		} else if s[i] == 'A' {
			dec += 10 * (math.Pow(16, float64(len(s)-1-i)))
		} else if s[i] == '9' {
			dec += 9 * (math.Pow(16, float64(len(s)-1-i)))
		} else if s[i] == '8' {
			dec += 8 * (math.Pow(16, float64(len(s)-1-i)))
		} else if s[i] == '7' {
			dec += 7 * (math.Pow(16, float64(len(s)-1-i)))
		} else if s[i] == '6' {
			dec += 6 * (math.Pow(16, float64(len(s)-1-i)))
		} else if s[i] == '5' {
			dec += 5 * (math.Pow(16, float64(len(s)-1-i)))
		} else if s[i] == '4' {
			dec += 4 * (math.Pow(16, float64(len(s)-1-i)))
		} else if s[i] == '3' {
			dec += 3 * (math.Pow(16, float64(len(s)-1-i)))
		} else if s[i] == '2' {
			dec += 2 * (math.Pow(16, float64(len(s)-1-i)))
		} else if s[i] == '1' {
			dec += 1 * (math.Pow(16, float64(len(s)-1-i)))
		} else {
			break
		}
	}
	return uint(dec)
}
func getIRhex(s string) string {
	switch instruct := s; instruct {
	case "HLT":
		return "7001"
	case "CMA":
		return "7200"
	case "CLA":
		return "7800"
	case "CLE":
		return "7400"
	case "CME":
		return "7100"
	case "CIR":
		return "7080"
	case "CIL":
		return "7040"
	case "INC":
		return "7020"
	case "SPA":
		return "7010"
	case "SNA":
		return "7008"
	case "SZA":
		return "7004"
	case "SZE":
		return "7002"
	case "INP":
		return "F800"
	case "OUT":
		return "F400"
	case "SKI":
		return "F200"
	case "SKO":
		return "F100"
	case "ION":
		return "F080"
	case "IOF":
		return "F040"
	}
	return ""
}
func getmemoryIR(s string, b bool) byte {
	switch s {
	case "AND":
		if b {
			return '0'
		} else {
			return '8'
		}
	case "ADD":
		if b {
			return '1'
		} else {
			return '9'
		}
	case "LDA":
		if b {
			return '2'
		} else {
			return 'A'
		}
	case "STA":
		if b {
			return '3'
		} else {
			return 'B'
		}
	case "BUN":
		if b {
			return '4'
		} else {
			return 'C'
		}
	case "BSA":
		if b {
			return '5'
		} else {
			return 'D'
		}
	case "ISZ":
		if b {
			return '6'
		} else {
			return 'E'
		}
	}
	return ' '
}
func saveAddress(index int, line string, org uint) {
	pos := strings.Index(line, ",")
	str := line[:pos]
	mp[str] = toHEX(org + uint(index))
}
func addressing(line string, add uint) {
	pos := strings.Index(line, ",")
	if test := line[pos+2 : pos+5]; test == "DEC" {
		d, _ := strconv.ParseInt(line[pos+6:], 10, 64)
		dec := uint(d)
		hexConverted = append(hexConverted, toHEX(add)+" "+toHEX(dec))
	} else if test == "HEX" {
		hexConverted = append(hexConverted, toHEX(add)+" "+line[pos+6:])
	} else if !isMemory(test) {
		hexConverted = append(hexConverted, toHEX(add)+" "+getIRhex(line[pos+2:pos+5]))
	} else if check := strings.Contains(line, " I"); check && isMemory(test) {
		posI := strings.Index(line, "I")
		temp:= mp[line[5:posI-1]]
		hexConverted = append(hexConverted, toHEX(add)+" "+string(getmemoryIR(line[pos+2:pos+5], false))+temp)
	} else if check := strings.Contains(line, " I"); !check && isMemory(test) {
		temp:= mp[line[5:]]
		hexConverted = append(hexConverted, toHEX(add)+" "+string(getmemoryIR(line[pos+2:pos+5], true))+temp)
	}
}
func toHEX(num uint) string {
	var hex []byte
	for num > 0 {
		if k := num % 16; k < 10 {
			hex = append(hex, byte(48+k))
		} else {
			hex = append(hex, byte(55+k))
		}
		num /= 16
	}
	var hexReturn string
	for i := range hex {
		hexReturn += string(rune(hex[len(hex)-1-i]))
	}
	return hexReturn
}
func Run() {
	input := bufio.NewScanner(os.Stdin)
	var lines []string
	for {
		input.Scan()
		line := input.Text()
		if line == "~" {
			break
		}
		lines = append(lines, line)
	}
	for _, line := range lines {
		if check := strings.Contains(line, "ORG"); check {
			decimalorg = getDECofHEX(line[5:])
			for i, line := range lines {
				if check := strings.Contains(line, ","); check {
					saveAddress(i, line, decimalorg)
				}
			}
			continue
		} else if check := strings.Contains(line, "END"); check {
			break
		} else if check := strings.Contains(line, ","); check {
			decimalorg += 1
			addressing(line, decimalorg)
		} else {
			if check := strings.Contains(line, " I "); isMemory(line[1:4]) && check {
				decimalorg += 1
				pos := strings.Index(line, "I")
				temp := mp[line[5:pos-1]]
				hexConverted = append(hexConverted, toHEX(decimalorg)+" "+string(getmemoryIR(line[1:4], false))+temp)
			} else if check := strings.Contains(line, " I "); isMemory(line[1:4]) && !check {
				temp:= mp[line[5:]]
				decimalorg += 1
				hexConverted = append(hexConverted, toHEX(decimalorg)+" "+string(getmemoryIR(line[1:4], true))+temp)
			} else {
				decimalorg += 1
				hexConverted = append(hexConverted, toHEX(decimalorg)+" "+getIRhex(line[1:4]))
			}
		}
	}
	for _, elem := range hexConverted {
		println(elem)
	}
}
