package bindex

import (
	"fmt"
	"testing"
)

type TestReader struct {
	count int
	dex   []byte
}

func (r *TestReader) Count() int {
	return r.count
}

func (r *TestReader) Read(i int) (dex Dex, err error) {
	offset := i * cSizeOfDex
	if offset <= (len(r.dex) - cSizeOfDex) {
		return Dex(r.dex[offset : offset+cSizeOfDex]), nil
	}
	return Dex([]byte{}), fmt.Errorf("TestReader::Read was asked to get a Dex entry at index %d while the maximum index is %d", i, r.Count()-1)
}

var TestContent = []byte{
	85, 31, 90, 218, 243, 11, 219, 124, 173, 122, 232, 157, 233, 43, 20, 109,
	0, 0, 0, 0,
	205, 98, 248, 61, 81, 93, 203, 68, 246, 23, 14, 21, 165, 221, 41, 217,
	0, 0, 0, 1,
	214, 93, 200, 212, 5, 193, 169, 14, 162, 12, 91, 229, 240, 134, 50, 107,
	0, 0, 0, 2,
	68, 179, 98, 161, 249, 242, 23, 221, 66, 102, 124, 48, 82, 163, 91, 38,
	0, 0, 0, 3,
}

func NewTestReader(content []byte) Reader {
	reader := &TestReader{dex: []byte{}}
	reader.count = 4
	reader.dex = content
	return reader
}

func TestTheReader(t *testing.T) {

	reader := NewTestReader(TestContent)
	for i := 0; i < reader.Count(); i++ {
		dex, err := reader.Read(i)
		if err != nil || (uint64(i*cDataAlignment) != dex.GetOffset()) {
			t.Fail()
		}
	}

}

type TestWriter struct {
	count int
	dex   []byte
}

func (r *TestWriter) Write(dex Dex) error {
	r.dex = append(r.dex, dex...)
	r.count++
	return nil
}

func TestTheWriter(t *testing.T) {
	writer := &TestWriter{dex: []byte{}}
	writer.count = 0

	for i := 0; i < 4; i++ {
		err := writer.Write(TestContent[i*cSizeOfDex : (i*cSizeOfDex)+cSizeOfDex])
		if err != nil {
			t.Fail()
		}
	}

	for i, b := range TestContent {
		if writer.dex[i] != b {
			t.Fail()
		}
	}
}
