package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Cake struct {
	Name        string        `json:"name" xml:"name"`
	Time        string        `json:"time" xml:"stovetime"`
	Ingredients []Ingredients `json:"ingredients" xml:"ingredients>item"`
}

type Ingredients struct {
	XMLName xml.Name `json:"-" xml:"item"`
	Name    string   `json:"ingredient_name" xml:"itemname"`
	Count   string   `json:"ingredient_count" xml:"itemcount"`
	Unit    string   `json:"ingredient_unit,omitempty" xml:"itemunit"`
}

type Reciept struct {
	XMLName xml.Name `json:"-" xml:"recipes"`
	Cakes   []Cake   `json:"cake" xml:"cake"`
}

type BDReader interface {
	Read(input []byte) ([]byte, error)
}

type bdJSON struct {
	fileName *os.File
}

type bdXML struct {
	fileName *os.File
}

func (js *bdJSON) Read(input []byte) ([]byte, error) {
	var cakes Reciept
	err := json.Unmarshal(input, &cakes)
	Xml, err := xml.MarshalIndent(cakes, "", "    ")
	return Xml, err
}

func (t *bdXML) Read(input []byte) ([]byte, error) {
	var cakes Reciept
	err := xml.Unmarshal(input, &cakes)
	Json, err := json.MarshalIndent(cakes, "", "  ")
	return Json, err
}

func worker(filename string) error {
	var (
		JsonName *bdJSON
		XmlName  *bdXML
		RealExt  string
		reciept  []byte
	)
	RealExt = filepath.Ext(filename)
	fileOpen, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileOpen.Close()

	bytesRead, err := ioutil.ReadAll(fileOpen)
	if err != nil {
		return err
	}
	switch RealExt {
	case ".xml":
		reciept, err = XmlName.Read(bytesRead)
	case ".json":
		reciept, err = JsonName.Read(bytesRead)
	default:
		fmt.Println("wrong extension files")
		os.Exit(-1)
	}
	fmt.Println(string(reciept))
	return err
}

func main() {
	flags := flag.Bool("f", false, "You must use -f flag")
	flag.Parse()
	if !(*flags) {
		fmt.Println("Error input argument")
		return
	}
	fileName := os.Args[2]
	err := worker(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
}
