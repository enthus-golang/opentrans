package opentrans

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"sort"
	"strings"
	"unicode/utf8"
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
			var buff bytes.Buffer
			err = escapeText(&buff, tt.Copy(), false)
			if err != nil {
				return nil, err
			}
			fx = append(fx, buff.Bytes()...)
		}
	}
}

var (
	escQuot = []byte("&#34;") // shorter than "&quot;"
	escApos = []byte("&#39;") // shorter than "&apos;"
	escAmp  = []byte("&amp;")
	escLT   = []byte("&lt;")
	escGT   = []byte("&gt;")
	escTab  = []byte("&#x9;")
	escNL   = []byte("&#xA;")
	escCR   = []byte("&#xD;")
	escFFFD = []byte("\uFFFD") // Unicode replacement character
)

func escapeText(w io.Writer, s []byte, escapeNewline bool) error {
	var esc []byte
	last := 0
	for i := 0; i < len(s); {
		r, width := utf8.DecodeRune(s[i:])
		i += width
		switch r {
		case '"':
			esc = escQuot
		case '\'':
			esc = escApos
		case '&':
			esc = escAmp
		case '<':
			esc = escLT
		case '>':
			esc = escGT
		case '\t':
			esc = escTab
		case '\n':
			if !escapeNewline {
				continue
			}
			esc = escNL
		case '\r':
			esc = escCR
		default:
			if !isInCharacterRange(r) || (r == 0xFFFD && width == 1) {
				esc = escFFFD
				break
			}
			continue
		}
		if _, err := w.Write(s[last : i-width]); err != nil {
			return err
		}
		if _, err := w.Write(esc); err != nil {
			return err
		}
		last = i
	}
	_, err := w.Write(s[last:])
	return err
}

func isInCharacterRange(r rune) (inrange bool) {
	return r == 0x09 ||
		r == 0x0A ||
		r == 0x0D ||
		r >= 0x20 && r <= 0xD7FF ||
		r >= 0xE000 && r <= 0xFFFD ||
		r >= 0x10000 && r <= 0x10FFFF
}
