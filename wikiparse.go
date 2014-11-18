package main

// An example streaming XML  to json parser fro use against wikepdia xml extracts.
//  Version 1: extracts text from wikipedia articles on chemical elements then writes each element's data into a seprate text file in /tmp/elements
// phase 1 just get  extractingt text <-  that works nicely . Fork from https://github.com/dps/go-xml-parse

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
)

// flag( name, default value, description)
var inputFile = flag.String("infile", "Wikipedia-20141117154753.xml", "Input file path")
var indexFile = flag.String("indexfile", "/tmp/article_list.txt", "article list output file")

var filter, _ = regexp.Compile("^file:.*|^talk:.*|^special:.*|^wikipedia:.*|^wiktionary:.*|^user:.*|^user_talk:.*")

// Here is an example article from the Wikipedia XML dump
//
// <page>
// 	<title>Apollo 11</title>
//      <redirect title="Foo bar" />
// 	...
// 	<revision>
// 	...
// 	  <text xml:space="preserve">
// 	  {{Infobox Space mission
// 	  |mission_name=&lt;!--See above--&gt;
// 	  |insignia=Apollo_11_insignia.png
// 	...
// 	  </text>
// 	</revision>
// </page>
//
// Note how the tags on the fields of Page and Redirect below
// describe the XML schema structure.

type Redirect struct {
	Title string `xml:"title,attr" json:"title"`
}

type Page struct {
	Title          string   `xml:"title" json:"title"`
	CanonicalTitle string   `xml:"ctitle" json:"ctitle"`
	Redir          Redirect `xml:"redirect" json:"redirect"`
	Text           string   `xml:"revision>text" json:"text"`
}

func CanonicalizeTitle(title string) string {
	can := strings.ToLower(title)
	can = strings.Replace(can, " ", "_", -1)
	can = url.QueryEscape(can)
	return can
}

func WriteJsonPage(title string, text string) {
	outFile, err := os.Create("/tmp/elements/" + title + ".json")
	if err == nil {
		writer := bufio.NewWriter(outFile)
		defer outFile.Close()
		writer.WriteString(text)
		writer.Flush()
	}
}

func main() {
	flag.Parse()

	xmlFile, err := os.Open(*inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// automatically call Close() at the end of current method
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	total := 0
	var inElement string
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		// Inspect the type of the token just read.
		switch se := t.(type) {
		case xml.StartElement:
			// If we just read a StartElement token
			inElement = se.Name.Local
			// ...and its name is "page"
			if inElement == "page" {
				var p Page
				// decode a whole chunk of following XML into the
				// variable p which is a Page (see above)
				decoder.DecodeElement(&p, &se)

				// Do some stuff with the page.
				p.Title = CanonicalizeTitle(p.Title)
				m := filter.MatchString(p.Title)
				if !m && p.Redir.Title == "" {

					// MarshalIndent  makes json output prettier !

					//  js_string, err := json.Marshal(p)
					js_string, err := json.MarshalIndent(p, "", " ")
					if err == nil {
						jsonString := string(js_string)
						WriteJsonPage(p.Title, jsonString)
					} else {
						fmt.Println(err)
					}
					total++
				}
			}
		default:
		}

	}

	fmt.Printf("Total articles: %d \n", total)
}
