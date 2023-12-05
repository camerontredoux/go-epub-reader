package xml

import "encoding/xml"

type Html struct {
	XMLName xml.Name `xml:"html"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Epub    string   `xml:"epub,attr"`
	Head    Head     `xml:"head"`
	Body    Body     `xml:"body"`
}

type Head struct {
	Text  string `xml:",chardata"`
	Title string `xml:"title"`
	Meta  []Meta `xml:"meta"`
	Link  []Link `xml:"link"`
}

type Meta struct {
	Text      string `xml:",chardata"`
	Content   string `xml:"content,attr"`
	Name      string `xml:"name,attr"`
	HTTPEquiv string `xml:"http-equiv,attr"`
}

type Link struct {
	Text string `xml:",chardata"`
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
}

type Heading struct {
	Text string `xml:",chardata"`
}

type Body struct {
	Content any `xml:",any"`
}

type Section struct {
	Text    string     `xml:",chardata"`
	Heading []Heading  `xml:",any"`
	A       []Anchor   `xml:"a"`
	Span    []Modifier `xml:"span"`
	Italic  []Modifier `xml:"i"`
	Strong  []Modifier `xml:"strong"`
	Small   []Modifier `xml:"small"`
}

type Anchor struct {
	Text string `xml:",chardata"`
	Href string `xml:"href,attr"`
}

type Modifier struct {
	Text string `xml:",chardata"`
}

type HeadingAlt struct {
	Heading1 []Modifier `xml:"h1"`
	Heading2 []Modifier `xml:"h2"`
	Heading3 []Modifier `xml:"h3"`
	Heading4 []Modifier `xml:"h4"`
	Heading5 []Modifier `xml:"h5"`
	Heading6 []Modifier `xml:"h6"`
}
