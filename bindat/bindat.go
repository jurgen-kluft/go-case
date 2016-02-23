package bindat

import ()

func NewReader(filepath string) (reader Reader, err error) {
	return NewFileReader(filepath)
}

func NewWriter(filepath string) (writer Writer, err error) {
	return NewFileAppendWriter(filepath)
}


