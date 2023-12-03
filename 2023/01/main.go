package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func getCalibration(str string) int {
	builder := strings.Builder{}

	for i := 0; i < len(str); i++ {
		if '0' <= str[i] && str[i] <= '9' {
			builder.WriteByte(str[i])
		} else {
			digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

			for j, digit := range digits {
				if strings.HasPrefix(str[i:], digit) {
					builder.WriteByte(byte('0' + j + 1))
				}
			}

		}
	}

	resStr := builder.String()

	res := 10*int(resStr[0]-'0') + int(resStr[len(resStr)-1]-'0')
	return res
}

func sum(reader io.ByteReader) int {
	builder := strings.Builder{}

	res := 0

	for {
		curByte, err := reader.ReadByte()

		if err == io.EOF {
			break
		}

		if curByte == '\n' {
			res += getCalibration(builder.String())

			builder.Reset()
		} else {
			builder.WriteByte(curByte)
		}
	}

	res += getCalibration(builder.String())

	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := sum(reader)

	fmt.Printf("%d\n", res)
}
