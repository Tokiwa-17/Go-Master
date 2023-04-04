package main

import (
	"log"
	"math/big"
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
	mp := make(map[string]string, pairNum)
	idx := 0
	for ; idx < pairNum; idx++ {
		key := bytes_[idx*100 : (idx+1)*100][:10]
		val := bytes_[idx*100 : (idx+1)*100][10:]
		j := new(big.Int).SetBytes(key)
		mp[j.String()] = string(val)
	}

	// for key, val := range mp {
	// 	log.Println(key, val)
	// }
	var key_arr []string
	for k := range mp {
		key_arr = append(key_arr, k)
	}
	sort.Strings(key_arr)
	file, err := os.OpenFile(os.Args[2], os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("file access error\n")
	} else {
		buf := make([]byte, pairNum*100)
		for i := 0; i < pairNum; i++ {
			key_buf := []byte(key_arr[i])[:10]
			val_buf := []byte(mp[key_arr[i]])
			pair_buf := append(key_buf, val_buf...)
			buf = append(buf[:i*100], append(pair_buf, buf[(i+1)*100:]...)...)
		}
		if len(buf)%100 != 0 {
			log.Fatalf("file size not divisible by 100\n")
		}
		file.Write(buf)
	}
	log.Printf("Sorting %s to %s\n", os.Args[1], os.Args[2])
}
