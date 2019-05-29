package opentrans

import (
	"regexp"
	"sort"
	"strings"
)

func CleanXMLNamespaces(xml string) (string, error) {
	replacer := map[string]string{
		"http://www.opentrans.org/XMLSchema/2.1": "",
		"http://www.bmecat.org/bmecat/2005":      "bmecat",
	}

	submatchReplace := func(re *regexp.Regexp, str string, repl func(int, []string) string) string {
		result := ""
		lastIndex := 0

		for i, v := range re.FindAllSubmatchIndex([]byte(str), -1) {
			var groups []string
			for j := 0; j < len(v); j += 2 {
				groups = append(groups, str[v[j]:v[j+1]])
			}

			result += str[lastIndex:v[0]] + repl(i, groups)
			lastIndex = v[1]
		}

		return result + str[lastIndex:]
	}

	re := regexp.MustCompile(`<([^\s]+) xmlns="([^"]+)"`)
	result := submatchReplace(re, xml, func(i int, groups []string) string {
		if i == 0 {
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
			return "<" + groups[1] + " " + strings.Join(attrs, " ")
		}
		prefix := replacer[groups[2]]
		if prefix == "" {
			return "<" + groups[1]
		}

		return "<" + prefix + ":" + groups[1]
	})

	return result, nil
}
