package shannon_fano

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"archiver/lib/comression/vlc/table"
)

type Table map[rune]Code

type Code struct {
	Char     rune
	Quantity int
	Bits     uint32
	Size     int
}

type charStat map[rune]int

type Generator struct{}

func NewGenerator() Generator {
	return Generator{}
}

func (g Generator) NewTable(text string) table.EncodingTable {
	return build(newCharStat(text)).Export()
}

func (t Table) Export() map[rune]string {
	res := make(map[rune]string)

	for k, v := range t {
		byteStr := fmt.Sprintf("%b", v.Bits)

		if lenDiff := v.Size - len(byteStr); lenDiff > 0 {
			byteStr = strings.Repeat("0", lenDiff) + byteStr
		}

		res[k] = byteStr
	}

	return res
}

func build(stat charStat) Table {
	codes := make([]Code, 0, len(stat))

	for ch, qty := range stat {
		codes = append(codes, Code{
			Char:     ch,
			Quantity: qty,
		})
	}

	sort.Slice(codes, func(i, j int) bool {
		if codes[i].Quantity != codes[j].Quantity {
			return codes[i].Quantity > codes[j].Quantity
		}

		return codes[i].Char < codes[j].Char
	})

	assignCodes(codes)

	res := make(Table)

	for _, code := range codes {
		res[code.Char] = code
	}

	return res
}

func assignCodes(codes []Code) {
	// TODO: fix case with one character
	if len(codes) < 2 {
		return
	}

	divider := bestDividePosition(codes)

	for i := 0; i < len(codes); i++ {
		codes[i].Bits <<= 1
		codes[i].Size++

		if i >= divider {
			codes[i].Bits |= 1
		}
	}

	assignCodes(codes[:divider])
	assignCodes(codes[divider:])
}

func bestDividePosition(codes []Code) int {
	total := 0
	for _, code := range codes {
		total += code.Quantity
	}

	left := codes[0].Quantity
	best := math.MaxInt
	bestPosition := 0

	for i := 0; i < len(codes)-1; i++ {
		right := total - left

		diff := abs(right - left)
		if diff >= best {
			break
		}

		best = diff
		left += codes[i].Quantity
		bestPosition = i + 1
	}

	return bestPosition
}

func newCharStat(str string) charStat {
	charStat := make(charStat)

	for _, ch := range str {
		charStat[ch]++
	}

	return charStat
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
