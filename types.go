package opentrans

import (
	"gitlab.com/mclgmbh/golang-pkg/bmecat"
)

type multiLocaleString00050 struct {
	bmecat.MultiLocale
	Value string `xml:",chardata" validate:"min=1,max=50"`
}
type multiLocaleString00150 struct {
	bmecat.MultiLocale
	Value string `xml:",chardata" validate:"min=1,max=150"`
}
type multiLocaleString00250 struct {
	bmecat.MultiLocale
	Value string `xml:",chardata" validate:"min=1,max=250"`
}
type multiLocaleString64000 struct {
	bmecat.MultiLocale
	Value string `xml:",chardata" validate:"min=1,max=150"`
}
