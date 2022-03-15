package table

import (
	"strings"
)

type EncodingTable map[rune]string

type decodingTree struct {
	Value string
	Zero  *decodingTree
	One   *decodingTree
}

type Generator interface {
	NewTable(text string) EncodingTable
}

func (et EncodingTable) Decode(str string) string {
	dt := et.decodingTree()

	return dt.Decode(str)
}

func (et EncodingTable) decodingTree() decodingTree {
	res := decodingTree{}

	for ch, code := range et {
		res.add(code, ch)
	}

	return res
}

func (dt *decodingTree) Decode(str string) string {
	var buf strings.Builder

	currNode := dt

	for _, ch := range str {
		if currNode.Value != "" {
			buf.WriteString(currNode.Value)
			currNode = dt
		}

		switch ch {
		case '0':
			currNode = currNode.Zero
		case '1':
			currNode = currNode.One
		}
	}

	if currNode.Value != "" {
		buf.WriteString(currNode.Value)
		currNode = dt
	}

	return buf.String()
}

func (dt *decodingTree) add(code string, value rune) {
	currNode := dt

	for _, ch := range code {
		switch ch {
		case '0':
			if currNode.Zero == nil {
				currNode.Zero = &decodingTree{}
			}

			currNode = currNode.Zero
		case '1':
			if currNode.One == nil {
				currNode.One = &decodingTree{}
			}

			currNode = currNode.One
		}
	}

	currNode.Value = string(value)
}
