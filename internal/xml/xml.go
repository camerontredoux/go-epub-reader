package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
	// "strings"
)

func UnmarshalXml(filename string) Html {
	file, err := os.Open(filename)
	if err != nil {
		panic("cannot find file")
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		panic("cannot read file content")
	}

	var html Html
	err = xml.Unmarshal(content, &html)
	if err != nil {
		panic("cannot unmashal content")
	}

	return html
}

func ParseXml(html Html) {
}

func DecodeXml(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic("cannot find file")
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)

	var text []string
	for {
		ct, err := decoder.Token()
		if err != nil {
			fmt.Println("all tokens read")
			break
		}

		switch t := ct.(type) {
		case xml.StartElement:
			// fmt.Println("start", t.Name.Local)
		case xml.EndElement:
			// fmt.Println("end", t.Name.Local)
		case xml.CharData:
			data := string(t)
			if len(data) > 0 {
				// fmt.Print("<", strings.TrimSpace(data), ">")
				text = append(text, strings.TrimSpace(data))
			}
		}

	}

	for _, data := range text {
		fmt.Println(data)
	}
}

func ReaderFromXml(filename string) *bytes.Reader {
	file, err := os.Open(filename)
	if err != nil {
		panic("cannot find file")
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		panic("cannot read file content")
	}

	var html Html
	err = xml.Unmarshal(content, &html)
	if err != nil {
		panic("cannot unmashal content")
	}

  body := html.Body.Content
  reader := bytes.NewReader(body)
  return reader
}
