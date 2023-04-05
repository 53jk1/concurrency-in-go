package queuing

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func BenchmarkUnbufferedWrite(b *testing.B) {
	performWrite(b, tmpFileOrFatal())
}

func BenchmarkBufferedWrite(b *testing.B) {
	bufferedFile := bufio.NewWriter(tmpFileOrFatal())
	performWrite(b, bufio.NewWriter(bufferedFile))
}

func tmpFileOrFatal() *os.File {
	tmpFile, err := ioutil.TempFile("", "test")
	if err != nil {
		log.Fatalf("error creating temp file: %v", err)
	}
	return tmpFile
}

func performWrite(b *testing.B, writer io.Writer) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for bt := range take(done, b.N, repeat(done, byte(0))) {
		writer.Write([]byte{bt.(byte)})
	}
}
