package comression

type Encoder interface {
	Encode(str string) []byte
}

type Decoder interface {
	Decode(encodedData []byte) (string, error)
}
