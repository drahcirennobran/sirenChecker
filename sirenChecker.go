package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	inseeFileName := os.Args[1]
	inseeFile, err := os.Open(inseeFileName)
	if err != nil {
		log.Fatalf("Couldn't open %v\n", err)
	}
	inseeReader := csv.NewReader(bufio.NewReader(inseeFile))
	inseeReader.Comma = ','
	inseeMap := make(map[int]bool)
	//for i := 0; i < 10000; i++ {
	for i := 0; ; i++ {
		record, err := inseeReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		key, _ := strconv.Atoi(record[0])
		/*
			if error != nil {
				fmt.Printf("error on record %v\n", record[0])
			}
		*/
		inseeMap[key] = true
		/*
			if i%1000000 == 0 {
				fmt.Println(len(inseeMap))
			}
		*/
	}
	//fmt.Println(len(inseeMap))

	OBSFileName := os.Args[2]
	OBSFile, err := os.Open(OBSFileName)
	if err != nil {
		log.Fatalf("Couldn't open %v\n", err)
	}
	OBSReader := csv.NewReader(bufio.NewReader(OBSFile))
	OBSReader.Comma = ';'
	for i := 0; ; i++ {
		record, err := OBSReader.Read()
		if err == io.EOF || len(record[0]) == 0 {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		key, error := strconv.Atoi(record[0])
		if error == nil {
			_, exist := inseeMap[key]
			if exist {
				fmt.Printf("%s;%s;%s;%s;%s\n", record[0], "Ok", record[2], record[3], record[4])
			} else {
				fmt.Printf("%s;%s;%s;%s;%s\n", record[0], "X", record[2], record[3], record[4])
			}
		}
		//fmt.Printf("%s --- %s\n", record[0], record[20])
		//key, error := strconv.ParseUint(record[0], 0, 64)
		//fmt.Printf("\"%v\" : %v\n", record[0], error)
		/*
			_, exist := inseeMap[record[0]]
			if exist {
				fmt.Println(record[0])
			}
		*/
	}
}
