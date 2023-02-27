package main

import (
	"bytes"
	"fmt"
)

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {
}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func t1() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{buffer: bytes.NewBuffer([]byte{})}
}

func t2() {
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello Youtube listeners, this is a test"))
	wc.Close()
}

func t3() {
	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("default")
	}
}

func main() {
	t3()
}
