package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Flags struct {
	L          *bool
	W          *bool
	M          *bool
	CountFlags int
}

func GetCountWord(filename string) string {
	fileOpen, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fileOpen.Close()
	input, err := ioutil.ReadAll(fileOpen)
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	return strconv.Itoa(count) + " " + filename
}

func GetCountLine(filename string) string {
	fileOpen, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fileOpen.Close()
	input, err := ioutil.ReadAll(fileOpen)
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	return strconv.Itoa(count) + " " + filename
}

func GetCountChar(filename string) string {
	fileOpen, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fileOpen.Close()
	bytes, err := ioutil.ReadAll(fileOpen)
	return strconv.Itoa(len(bytes)) + " " + filename // \n also sum in count
}

func GetCountFiles(args []string) int {
	var count int
	for i := range args {
		if args[i] != "-w" && args[i] != "-m" && args[i] != "-l" {
			count++
		}
	}
	return count
}

func main() {
	var (
		flags Flags
		args  []string
		fl    string
	)
	flags.M = flag.Bool("m", false, "This flag for counting characters")
	flags.L = flag.Bool("l", false, "This flag for counting lines")
	flags.W = flag.Bool("w", false, "This flag for counting words")
	args = os.Args[1:]
	countFiles := GetCountFiles(args)
	chanSlice := make(chan string, countFiles)
	for i := 0; i < len(args); i++ {
		if args[i] == "-w" || args[i] == "-m" || args[i] == "-l" {
			fl = args[i]
			i++
		}
		go func(i int, fl string, args []string) {
			switch fl {
			case "-w":
				chanSlice <- GetCountWord(args[i])
			case "-m":
				chanSlice <- GetCountChar(args[i])
			case "-l":
				chanSlice <- GetCountLine(args[i])
			}
		}(i, fl, args)
	}
	for i := 0; i < countFiles; i++ {
		val := <-chanSlice
		fmt.Println(val)
	}
}
