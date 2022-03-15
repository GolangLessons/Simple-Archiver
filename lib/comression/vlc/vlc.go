package vlc

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"strings"

	"archiver/lib/comression/vlc/table"
)

type EncoderDecoder struct {
	tblGenerator table.Generator
}

func New(tblGenerator table.Generator) EncoderDecoder {
	return EncoderDecoder{tblGenerator: tblGenerator}
}

func (ed EncoderDecoder) Encode(text string) []byte {
	tbl := ed.tblGenerator.NewTable(text)

	encoded := encodeBin(text, tbl)

	return buildEncodedFile(tbl, encoded)
}

func (ed EncoderDecoder) Decode(encodedData []byte) (string, error) {
	tbl, data, err := parseFile(encodedData)
	if err != nil {
		return "", err
	}

	return tbl.Decode(data), nil
}

func buildEncodedFile(tbl table.EncodingTable, data string) []byte {
	encodedTbl := encodeTable(tbl)

	var res bytes.Buffer

	res.Write(encodeInt(len(encodedTbl)))
	res.Write(encodeInt(len(data)))
	res.Write(encodedTbl)
	res.Write(splitByChunks(data, chunksSize).Bytes())

	return res.Bytes()
}

func parseFile(data []byte) (table.EncodingTable, string, error) {
	const (
		tableSizeBytesCount = 4
		dataSizeBytesCount  = 4
	)

	tableSizeBinary, data := data[:tableSizeBytesCount], data[tableSizeBytesCount:]
	dataSizeBinary, data := data[:dataSizeBytesCount], data[dataSizeBytesCount:]

	tableSize := binary.BigEndian.Uint32(tableSizeBinary)
	dataSize := binary.BigEndian.Uint32(dataSizeBinary)

	tblBinary, data := data[:tableSize], data[tableSize:]

	tbl, err := decodeTable(tblBinary)
	if err != nil {
		return nil, "", err
	}

	body := NewBinChunks(data).Join()

	return tbl, body[:dataSize], nil
}

func encodeInt(num int) []byte {
	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, uint32(num))

	return res
}

func decodeTable(tblBinary []byte) (table.EncodingTable, error) {
	var tbl table.EncodingTable

	r := bytes.NewReader(tblBinary)
	if err := gob.NewDecoder(r).Decode(&tbl); err != nil {
		return nil, err
	}

	return tbl, nil
}

func encodeTable(tbl table.EncodingTable) []byte {
	var tableBuf bytes.Buffer

	if err := gob.NewEncoder(&tableBuf).Encode(tbl); err != nil {
		panic(err)
	}

	return tableBuf.Bytes()
}

// encodeBin encodes str into binary codes string without spaces.
func encodeBin(str string, table table.EncodingTable) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(bin(ch, table))
	}

	return buf.String()
}

func bin(ch rune, table table.EncodingTable) string {
	res, ok := table[ch]
	if !ok {
		panic("unknown character: " + string(ch))
	}

	return res
}
