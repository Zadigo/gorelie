package internal

import "io"

type Entry struct {
	Key   []byte
	Value []byte
}

func (e *Entry) Encode() []byte {
	return nil
}

func (e *Entry) Decode(buffer io.Reader) error {
	return nil
}
