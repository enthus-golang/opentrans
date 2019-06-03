package opentrans

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"sort"
	"strings"
)

func CleanXMLNamespaces(x []byte) ([]byte, error) {
	replacer := map[string]string{
		"http://www.opentrans.org/XMLSchema/2.1": "",
		"http://www.bmecat.org/bmecat/2005":      "bmecat",
	}

	d := xml.NewDecoder(bytes.NewReader(x))
	var fx []byte
	depth := 0
	for {
		t, err := d.Token()
		if err == io.EOF {
			return fx, nil
		}
		if err != nil {
			return nil, err
		}
		switch tt := t.(type) {
		case xml.StartElement:
			//fmt.Println(">", tt.Name.Local, tt.Name.Space)
			fx = append(fx, '<')
			if v := replacer[tt.Name.Space]; v != "" {
				fx = append(fx, []byte(v+":")...)
			}
			fx = append(fx, []byte(tt.Name.Local)...)
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
				fx = append(fx, []byte(" "+strings.Join(attrs, " "))...)
			}
			for _, v := range tt.Attr {
				if v.Name.Local == "xmlns" {
					continue
				}
				fx = append(fx, []byte(" "+fmt.Sprintf(`%s="%s"`, v.Name.Local, v.Value))...)
			}
			fx = append(fx, '>')
			depth++
		case xml.EndElement:
			//fmt.Println("<", tt.Name.Local, tt.Name.Space)
			fx = append(fx, '<', '/')
			if v := replacer[tt.Name.Space]; v != "" {
				fx = append(fx, []byte(v+":")...)
			}
			fx = append(fx, []byte(tt.Name.Local+">")...)
			depth--
		case xml.CharData:
			fx = append(fx, tt.Copy()...)
		}
	}
}
