package main

import (
	"log"
	"os"
	"strconv"
)

func sizeOfFile(fileName string) string {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		log.Printf("Error openning file to calculate size: %v", err)
		return "N/A"
	}
	st, _ := file.Stat()
	size := strconv.Itoa(int(st.Size()/1024)) + "KB"
	return size

}

func comment(flag string) int {
	switch flag {
	case "low":
		return 2
	case "middle":
		return 4
	case "big":
		return 8
	default:
		return 0
	}
}
