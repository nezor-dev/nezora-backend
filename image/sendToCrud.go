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
	"strings"

	"github.com/otiai10/gosseract/v2"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

var fileext = []string{"jpg", "png"}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func ConvertImageToPDF_BASE64(input []string) string {
	// Read the entire file into a byte slice
	imp, _ := api.Import("form:A3, pos:c, s:1.0", pdfcpu.POINTS)
	api.ImportImagesFile(input, "out.pdf", imp, nil)

	// Read the entire file into a byte slice
	bytes, err := ioutil.ReadFile("./out.pdf")
	if err != nil {
		log.Fatal(err)
	}
	os.Remove("./out.pdf")

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
	content := ""

	base64 := getBase64ofFolderImages_PDF()

	content = getContentFromCurrentFolder(content)

	values := map[string]string{
		"name":    name,
		"image":   base64,
		"content": content,
		"date":    date,
		"sender":  sender,
	}

	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://127.0.0.1:8001/documents", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(name)
	fmt.Println(content)
	fmt.Println(date)
	fmt.Println(sender)

}

func getBase64ofFolderImages_PDF() string {
	var strArray []string

	strArray = append(strArray, extractFileNamesFromCurrentFolder(fileext)...)

	var base64 = ConvertImageToPDF_BASE64(strArray)
	return base64
}

func getContentFromCurrentFolder(content string) string {
	client := gosseract.NewClient()
	defer client.Close()

	for _, file := range extractFileNamesFromCurrentFolder(fileext) {
		client.SetImage(file)
		image, _ := client.Text()
		content = content + image
	}
	return content
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func extractFileNamesFromCurrentFolder(file []string) []string {
	var strArray []string

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
		split := strings.Split(f.Name(), ".")
		if contains(file, split[len(split)-1]) {
			strArray = append(strArray, f.Name())
		}
	}
	return strArray
}
