package main

import (
	"bytes"
	"log"
	"math/big"
	"os"
	"sort"
)

type Pair struct {
	Key   *big.Int
	Value []byte
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) != 3 {
		log.Fatalf("Usage: %v inputfile outputfile\n", os.Args[0])
	}

	bytes_, err := os.ReadFile(os.Args[1])
	if os.IsNotExist(err) {
		log.Fatalf("File not exist\n")
	} else if err != nil {
		log.Fatalf("File access error\n")
	}
	log.Printf("%s open success, length: %d\n", os.Args[1], len(bytes_))
	var pairNum int = len(bytes_) / 100
	var pairs []Pair
	for idx := 0; idx < pairNum; idx++ {
		key := bytes_[idx*100 : (idx+1)*100][:10]
		val := bytes_[idx*100 : (idx+1)*100][10:]
		_key := new(big.Int).SetBytes(key)
		pairs = append(pairs, Pair{_key, val})
	}
	sort.Slice(pairs, func(i, j int) bool {
		x := *(pairs[i].Key)
		return x.Cmp(pairs[j].Key) > 0
	})

	file, err := os.OpenFile(os.Args[2], os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("file access error\n")
	} else {
		var buf []byte
		for i := 0; i < pairNum; i++ {
			key_buf := (*pairs[i].Key).Bytes()
			if len(key_buf) < 10 {
				key_buf = append(key_buf, bytes.Repeat([]byte{0}, 10-len(key_buf))...)
			} else {
				key_buf = key_buf[:10]
			}
			val_buf := []byte(pairs[i].Value)[:90]
			buf = append(buf, append(key_buf, val_buf...)...)
		}
		if len(buf)%100 != 0 {
			log.Fatalf("file size not divisible by 100\n")
		}
		file.Write(buf)
	}
	log.Printf("Sorting %s to %s\n", os.Args[1], os.Args[2])
}
