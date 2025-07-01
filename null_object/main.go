package main

import (
	"fmt"
	"io"
)

// NullWriter は io.Writer の Null Object 実装
type NullWriter struct{}

// Write は何もせず、len(p) のバイト数を書き込んだことにする
func (n *NullWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

func main() {
	var w io.Writer = &NullWriter{}

	n, err := w.Write([]byte("this will be discarded"))
	if err != nil {
		fmt.Println("write error:", err)
		return
	}
	fmt.Printf("wrote %d bytes (but discarded)\n", n)
}
