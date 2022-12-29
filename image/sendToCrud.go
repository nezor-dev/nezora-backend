package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func ConvertImageToPDF_BASE64() string {
	// Read the entire file into a byte slice
	bytes, err := ioutil.ReadFile("./scan.pdf")
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	case "application/pdf":
		base64Encoding += "data:application/pdf;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	// Print the full base64 representation of the image
	return base64Encoding
}

func main() {

	var args = os.Args[1:]

	name := args[0]
	date := args[1]
	sender := args[2]

	content, err := ioutil.ReadFile("./output.txt")
	if err != nil {
		log.Fatal(err)
	}

	base64 := getBase64ofFolderImages_PDF()

	values := map[string]string{
		"name":    name,
		"image":   base64,
		"content": string(content),
		"date":    date,
		"sender":  sender,
	}

	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://127.0.0.1:8001/api/v1/mail", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(resp)

	cleanUpFiles()
}

func getBase64ofFolderImages_PDF() string {
	var base64 = ConvertImageToPDF_BASE64()
	return base64
}

func cleanUpFiles() {
	e := os.Remove("output.txt")
	if e != nil {
		log.Fatal(e)
	}
	e = os.Remove("scan.pdf")
	if e != nil {
		log.Fatal(e)
	}
	e = os.Remove("*.tiff")
	if e != nil {
		log.Fatal(e)
	}
}
