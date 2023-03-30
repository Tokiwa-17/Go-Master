package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"os"
	"sort"
)

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
	mp := make(map[uint16]string, pairNum)
	idx := 0
	for ; idx < pairNum; idx++ {
		key := bytes_[idx*100 : (idx+1)*100][:10]
		val := bytes_[idx*100 : (idx+1)*100][10:]
		var j uint16
		buf := bytes.NewReader(key)
		binary.Read(buf, binary.LittleEndian, &j)
		mp[j] = string(val)
	}

	// for key, val := range mp {
	// 	log.Println(key, val)
	// }
	var key_arr []uint16
	for k := range mp {
		key_arr = append(key_arr, k)
	}
	sort.Slice(key_arr, func(i, j int) bool {
		return key_arr[i] < key_arr[j]
	})
	for _, k := range key_arr {
		log.Println(k, mp[k])
	}
	log.Printf("Sorting %s to %s\n", os.Args[1], os.Args[2])
}
