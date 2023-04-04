package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

type ServerConfigs struct {
	Servers []struct {
		ServerId int    `yaml:"serverId"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	} `yaml:"servers"`
}

func readServerConfigs(configPath string) ServerConfigs {
	f, err := ioutil.ReadFile(configPath)

	if err != nil {
		log.Fatalf("could not read config file %s : %v", configPath, err)
	}

	scs := ServerConfigs{}
	err = yaml.Unmarshal(f, &scs)

	return scs
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) != 5 {
		log.Fatal("Usage : ./netsort {serverId} {inputFilePath} {outputFilePath} {configFilePath}")
	}

	// What is my serverId
	serverId, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid serverId, must be an int %v", err)
	}
	fmt.Println("My server Id:", serverId)

	// Read server configs from file
	scs := readServerConfigs(os.Args[4])
	fmt.Println("Got the following server configs:", scs)

	/*
		Implement Distributed Sort
	*/
	// First Step: Read the Input file
	bytes_, err := os.ReadFile(os.Args[2])
	if os.IsNotExist(err) {
		log.Fatalf("File not exist\n")
	} else if err != nil {
		log.Fatalf("File access error\n")
	}
	fmt.Printf("%s open success, length: %d\n", os.Args[2], len(bytes_))
	// Second Step: Appropriately partition the data
	pair_num := len(bytes_) / 100
	var mask byte = 0b11000000
	for idx := 0; idx < pair_num; idx++ {
		key := bytes_[idx*100 : (idx+1)*100][:10]
		//val := bytes_[idx*100 : (idx+1)*100][10:]
		send_idx := int((key[9] & mask) >> 6)
		fmt.Printf("%08b %08b %d\n", mask, key[9], send_idx)
	}
}
