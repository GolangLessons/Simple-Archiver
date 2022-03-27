package compression

type Encoder interface {
	Encode(str string) []byte
}

type Decoder interface {
	Decode(data []byte) string
}
