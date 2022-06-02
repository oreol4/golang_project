package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Flags struct {
	L *bool
	W *bool
	M *bool
}

func GetFilesName(args []string) []string {
	var files []string
	for i := range args {
		if args[i] != "-w" && args[i] != "-m" && args[i] != "-l" {
			files = append(files, args[i])
		}
	}
	return files
}

func GetCountWord(byteSlice []byte) int {
	tmp := string(byteSlice)
	splitWord := strings.Split(tmp, " ")
	return len(splitWord)
}

func GetCountLine(byteSlice []byte) int {
	tmp := string(byteSlice)
	splitWord := strings.Split(tmp, "\n")
	return len(splitWord)
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
	chanSlice := make(chan int)
	for i := 0; i < len(args); i++ {
		if args[i] == "-w" || args[i] == "-m" || args[i] == "-l" {
			fl = args[i]
			i++
		}
		go func(i int, fl string, args []string) {
			fileName, _ := os.Open(args[i])
			byteInput, _ := ioutil.ReadAll(fileName)
			switch fl {
			case "-w":
				fmt.Print(fileName.Name())
				chanSlice <- GetCountWord(byteInput)
			case "-m":
				fmt.Print(fileName.Name())
				chanSlice <- len(byteInput)
			case "-l":
				fmt.Print(fileName.Name())
				chanSlice <- GetCountLine(byteInput)
			}
		}(i, fl, args)
	}
	for i := 0; i < 3; i++ {
		val := <-chanSlice
		fmt.Println(val)
	}
}
