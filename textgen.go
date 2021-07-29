package main

import (
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	words, err := ioutil.ReadFile("dictionary.txt")
	if err != nil {
		panic(fmt.Errorf("Missing dictionary.txt - how can I proceed???"))
	}
	var chunkSize int
	var totalSize int
	var intSeed int64
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s: Generate random chunks of data\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.IntVar(&chunkSize, "chunk-size", 10000, "Size of a single file (chunk)")
	flag.IntVar(&totalSize, "total-size", 1000000, "Total size to generate")
	flag.Int64Var(&intSeed, "random-seed", -1, "Random seed")
	flag.Parse()
	if intSeed == -1 {
		intSeed = time.Now().Unix()
	}
	rand.Seed(intSeed)
	wordary := strings.Split(string(words), "\n")
	wordary_len := len(wordary)
	count := totalSize / chunkSize
	var wg sync.WaitGroup
	for count > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			genchunk(wordary, wordary_len, chunkSize)
		}()
		count -= 1
	}
	wg.Wait()
}

func genchunk(source []string, sourcelen, size int) {
	chunk := strings.Builder{}
	chunk.Grow(size)
	f := 0
	for size > f {
		chunk.WriteString(source[rand.Intn(sourcelen)])
		chunk.WriteRune(' ')
		f = chunk.Len()
	}
	chunk.WriteRune('\n')
	s := chunk.String()
	h := sha1.New()
	h.Write([]byte(s))
	fn := hex.EncodeToString(h.Sum(nil))
	dpath := fmt.Sprintf("%s/%s", fn[0:2], fn[1:3])
	fpath := fmt.Sprintf("%s/%s", dpath, fn)
	os.MkdirAll(dpath, 0750)
	ioutil.WriteFile(fpath, []byte(s), 0640)
}
