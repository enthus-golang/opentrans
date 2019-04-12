package opentrans

import "gitlab.com/mclgmbh/gomod/bmecat"

type multiLocaleString00050 struct {
	bmecat.MultiLocale
	Value string `xml:",chardata" validate:"min=1,max=50"`
}
