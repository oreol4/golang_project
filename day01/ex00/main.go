package main

import (
	"encoding/json"
	"encoding/xml"
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
	Read(input []byte) (Reciept, error)
}

//func ()
type bdJSON struct {
	fileName *os.File
}

type bdXML struct {
	fileName *os.File
}

func (js *bdJSON) Read(input []byte) (Reciept, error) {
	var cakes Reciept
	err := json.Unmarshal(input, &cakes)
	Xml, err := xml.MarshalIndent(cakes, "", "    ")
	fmt.Println(string(Xml))
	return cakes, err
}

func (t *bdXML) Read(input []byte) (Reciept, error) {
	var cakes Reciept
	err := xml.Unmarshal(input, &cakes)
	Json, err := json.MarshalIndent(cakes, "", "  ")
	fmt.Println(string(Json))
	return cakes, err
}

func main() {
	var (
		JsonName *bdJSON
		XmlName  *bdXML
		//rec      Reciept ?????
		err     error
		RealExt string
	)
	//var cakes Reciept
	//if len(os.Args) != 2 {
	//	return
	//}
	fileName := os.Args[1]
	RealExt = filepath.Ext(fileName)
	fileContent, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fileContent.Close()

	fmt.Println("The File is opened successfully...")

	byteResult, err := ioutil.ReadAll(fileContent)
	if err != nil {
		log.Fatal(err)
		return
	}
	if RealExt == ".xml" {
		_, err = XmlName.Read(byteResult)
	} else {
		_, err = JsonName.Read(byteResult)
	}
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(rec)
	//err = json.Unmarshal(byteResult, &cakes)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//
	//XMLdata, err := xml.MarshalIndent(cakes, "", "    ")
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//fmt.Printf("%v\n", string(XMLdata))
}
