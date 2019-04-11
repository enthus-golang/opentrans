package opentrans

type multiLocaleString00050 struct {
	multiLocale
	Value string `xml:",chardata" validate:"min=1,max=50"`
}
