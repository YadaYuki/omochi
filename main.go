package main

import (
	"bytes"
	"compress/zlib"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/YadaYuki/omochi/app/domain/entities"
)

func main() {
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.

	postings := make([]entities.Posting, 0)
	for i := 0; i < 100; i++ {
		postings = append(postings, *entities.NewPosting(1, []int{1, 2, 3}))
	}
	enc.Encode(&postings)
	var compressedDataBuffer bytes.Buffer
	w := zlib.NewWriter(&compressedDataBuffer)
	w.Write(network.Bytes())
	w.Close()

	fmt.Printf("size before compress:%v, after compress:%v \n", len(network.Bytes()), len(compressedDataBuffer.Bytes()))
	compressedDataBuffer.Reset()
	fmt.Println(len(compressedDataBuffer.Bytes()))
	w.Flush()
	fmt.Println(len(compressedDataBuffer.Bytes()))

	//
	var decompressedDataBuffer bytes.Buffer
	r, _ := zlib.NewReader(&compressedDataBuffer)
	io.Copy(&decompressedDataBuffer, r)
	var postingDecoded []entities.Posting
	dec := gob.NewDecoder(&decompressedDataBuffer) // Will read from network.
	dec.Decode(&postingDecoded)
	fmt.Println("decompressed: ", postingDecoded)

}
