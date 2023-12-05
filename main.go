// package main
//
// import (
//
//	"nori/internal/xml"
//
// )
//
//	func main() {
//		// xml.ParseXml(xml.UnmarshalXml("test.xhtml"))
//	  xml.DecodeXml("test.xhtml")
//	}
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
	"golang.org/x/net/html"
)

func Text(s []*html.Node) string {
	var buf bytes.Buffer

	// Slightly optimized vs calling Each: no single selection object created
	var f func(n *html.Node)
	f = func(n *html.Node) {
		// Keep newlines and spaces, like jQuery
		if n.Data == "p" {
			buf.WriteString("\t")
		}
		if n.Type == html.TextNode {
			if n.Parent.Data == "i" {
				buf.WriteString("\033[3m" + n.Data + "\033[0m")
			} else {
				buf.WriteString(n.Data)
			}
		}
		if n.FirstChild != nil {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
	}

	for _, n := range s {
		f(n)
	}

	return strings.Trim(buf.String(), "\n")
}

func main() {

	data, _ := os.Open("chapter2.xhtml")

	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		log.Fatal(err)
	}
	chapter := Text(doc.Find("body").Nodes)
	s := wordwrap.String(chapter, 80)
  padding := lipgloss.NewStyle().Padding(0,35)
  
  fmt.Println(padding.Render(s))
  fmt.Println(padding.GetFrameSize())


	// fmt.Println(strings.TrimSpace(Text(doc.Find("body"))))

	/* var text strings.Builder
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		text.WriteString(s.Text())
	})

	result := text.String()
	fmt.Println(result) */
}
