package tis620

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"testing"
	"unicode/utf8"
)

const TEST_FILENAME = "tis620-testcases.txt"

func TestConvert_ToUTF8_AllChars(t *testing.T) {
	testCases := readTestFile()

	for _, line := range testCases {
		tis620byte := byte(line.In)
		utf8bytes := make([]byte, 3)
		l := utf8.EncodeRune(utf8bytes, line.Out)
		utf8bytes = utf8bytes[:l]

		tis620bytes := make([]byte, 1)
		tis620bytes[0] = tis620byte

		if !bytes.Equal(ToUTF8(tis620bytes), utf8bytes) {
			t.Error(fmt.Sprintf("Unexpected output - %d, %d, %s", tis620byte, line.Out, ToUTF8(tis620bytes)))
			return
		}
	}
}

func TestConvert_ToUTF8_Sentence(t *testing.T) {
	utf8 := []byte("ทดสอบภาษาไทย and English")
	tis620 := []byte{183, 180, 202, 205, 186, 192, 210, 201, 210, 228, 183, 194, 32, 97, 110, 100, 32, 69, 110, 103, 108, 105, 115, 104}

	if !bytes.Equal(utf8, ToUTF8(tis620)) {
		t.Error("Error converting tis620 sentence")
	}
}

func readTestFile() []InOut {
	file, err := os.Open(TEST_FILENAME)
	panicIf(err)
	defer file.Close()

	reader := csv.NewReader(file)

	data, err := reader.ReadAll()
	panicIf(err)

	var result []InOut
	for _, line := range data {
		in, _ := strconv.ParseInt(line[0], 0, 32)
		out, _ := strconv.ParseInt(line[1], 0, 32)
		result = append(result, InOut{int32(in), int32(out)})
	}

	return result
}

type InOut struct {
	In, Out int32
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
