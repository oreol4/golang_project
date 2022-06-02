package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Flags struct {
	dir  *bool
	file *bool
	sl   *bool
	ext  *bool
	ex   string
}

func GetPath(args []string) string {
	var filename string
	for i := 0; i < len(args); i++ {
		if strings.Contains(args[i], "/") {
			filename = args[i]
		} else if args[i] == "." {
			filename = "./"
		}
	}
	return filename
}

func dirThree(path string, flags *Flags) {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	data, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	for i := range data {
		if data[i].IsDir() {
			if *flags.dir {
				fmt.Println(path + data[i].Name())
			}
			dirThree(path+data[i].Name(), flags)
		} else {
			link, err := os.Readlink(path + data[i].Name())
			if err == nil {
				_, errl := os.Open(path + link) // выдаст ошибку, если файла нет и зайдет в след условие(42)
				if *flags.sl && errl != nil {
					fmt.Printf("%s -> %s\n", path+data[i].Name(), "[broken]")
				} else if *flags.sl {
					fmt.Println(path+data[i].Name(), "->", link)
				}
				continue
			}
			_, err2 := os.Open(path + data[i].Name()) // выдаст ошибку, если это ссылка или директория и зайдет в условие
			if err2 != nil {
				continue // не следует печать путь+название ссылки(директории)(это сделано выше)
			}
			if *flags.file && err2 == nil && !*flags.ext {
				fmt.Println(path + data[i].Name())
			} else if strings.Contains(data[i].Name(), flags.ex) && *flags.ext && *flags.file {
				fmt.Println(path + data[i].Name())
			}
		}
	}
}

func searchEx(flags *Flags, args []string) string {
	var exName string
	if *flags.ext {
		exName = args[len(args)-2]
	}
	return "." + exName
}
func main() {
	var (
		workPath string
		flags    Flags
	)
	flags.dir = flag.Bool("d", false, "Use this flag for directory")
	flags.file = flag.Bool("f", false, "Use this flag for files")
	flags.sl = flag.Bool("sl", false, "Use this flag for symbolic links")
	flags.ext = flag.Bool("ext", false, "Use this flag for ex files")
	flag.Parse()
	args := os.Args[1:]
	workPath = GetPath(args)
	flags.ex = searchEx(&flags, args)
	dirThree(workPath, &flags)
}
