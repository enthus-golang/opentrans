package opentrans

import (
	"encoding/xml"
	"fmt"
	"io"
	"sort"
	"strings"
)

func CleanXMLNamespaces(x string) (string, error) {
	replacer := map[string]string{
		"http://www.opentrans.org/XMLSchema/2.1": "",
		"http://www.bmecat.org/bmecat/2005":      "bmecat",
	}

	d := xml.NewDecoder(strings.NewReader(x))
	fx := ""
	depth := 0
	for {
		t, err := d.Token()
		if err == io.EOF {
			return fx, nil
		}
		if err != nil {
			return "", err
		}
		switch tt := t.(type) {
		case xml.StartElement:
			//fmt.Println(">", tt.Name.Local, tt.Name.Space)
			fx += "<"
			if v := replacer[tt.Name.Space]; v != "" {
				fx += v + ":"
			}
			fx += tt.Name.Local
			if depth == 0 {
				attrs := make([]string, len(replacer))
				keys := make([]string, len(replacer))
				i := 0
				for k := range replacer {
					keys[i] = k
					i++
				}
				i = 0
				sort.Strings(keys)
				for _, k := range keys {
					attrs[i] = "xmlns"
					if replacer[k] != "" {
						attrs[i] += ":" + replacer[k]
					}
					attrs[i] += `="` + k + `"`
					i++
				}
				fx += " " + strings.Join(attrs, " ")
			}
			for _, v := range tt.Attr {
				if v.Name.Local == "xmlns" {
					continue
				}
				fx += " "
				fx += fmt.Sprintf(`%s="%s"`, v.Name.Local, v.Value)
			}
			fx += ">"
			depth++
		case xml.EndElement:
			//fmt.Println("<", tt.Name.Local, tt.Name.Space)
			fx += "</"
			if v := replacer[tt.Name.Space]; v != "" {
				fx += v + ":"
			}
			fx += tt.Name.Local + ">"
			depth--
		case xml.CharData:
			fx += string(tt.Copy())
		}
	}
}
