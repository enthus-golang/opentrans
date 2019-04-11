package opentrans

import (
	"encoding/xml"
	"regexp"
)

type Language struct {
	XMLName xml.Name `xml:"bmecat:LANGUAGE"`
	Default bool     `xml:"default,attr,omitempty"`
	Value   string   `xml:",chardata" validate:"oneof=aar abk ace ach ada afa afh afr aka akk alb ale alg amh ang apa ara arc arm arn arp art arw asm ath aus ava ave awa aym aze bad bai bak bal bam ban baq bas bat bej bel bem ben ber bho bih bik bin bis bla bnt bod bos bra bre btk bua bug bul bur cad cai car cat cau ceb cel ces cha chb che chg chi chk chm chn cho chp chr chu chv chy cmc cop cor cos cpe cpf cpp cre crp cus cym cze dak dan day del den deu dgr din div doi dra dua dum dut dyu dzo efi egy eka ell elx eng enm epo est eus ewe ewo fan fao fas fat fij fin fiu fon fra fre frm fro fry ful fur gaa gay gba gem geo ger gez gil gla gle glg glv gmh goh gon gor got grb grc gre grn guj gwi hai hau haw heb her hil him hin hit hmn hmo hrv hun hup hye iba ibo ice ijo iku ile ilo ina inc ind ine ipk ira iro isl ita jav jpn jpr jrb kaa kab kac kal kam kan kar kas kat kau kaw kaz kha khi khm kho kik kin kir kmb kok kom kon kor kos kpe kro kru kua kum kur kut lad lah lam lao lat lav lez lin lit lol loz ltz lua lub lug lui lun luo lus mac mad mag mah mai mak mal man mao map mar mas may mdr men mga mic min mis mkd mkh mlg mlt mnc mni mno moh mol mon mos mri msa mul mun mus mwr mya myn nah nai nau nav nbl nde ndo nds nep new nia nic niu nld nno nob non nor nso nub nya nym nyn nyo nzi oci oji ori orm osa oss ota oto paa pag pal pam pan pap pau peo per phi phn pli pol pon por pra pro pus qaa que raj rap rar roa roh rom ron rum run rus sad sag sah sai sal sam san sas sat scc sco scr sel sem sga sgn shn sid sin sio sit sla slk slo slv sme smi smo sna snd snk sog som son sot spa sqi srd srp srr ssa ssw suk sun sus sux swa swe syr tah tai tam tat tel tem ter tet tgk tgl tha tib tig tir tiv tkl tli tmh tog ton tpi tsi tsn tso tuk tum tur tut tvl twi tyv uga uig ukr umb und urd uzb vai ven vie vol vot wak wal war was wel wen wol xho yao yap yid yor ypk zap zen zha zho znd zul"`
}

const (
	PartyTypeBuyerSpecific    = "buyer_specific"
	PartyTypeCustomerSpecific = "customer_specific"
	PartyTypeDUNS             = "duns"
	PartyTypeILN              = "iln"
	PartyTypeGLN              = "gln"
	PartyTypeSupplierSpecific = "supplier_specific"
)

type PartyID struct {
	XMLName xml.Name `xml:"bmecat:PARTY_ID"`
	Type    string   `xml:"type,attr" validate:"min=1,max=250"`
	Value   string   `xml:",chardata" validate:"min=1,max=250"`
}

type multiLocale struct {
	Lang   string `xml:"lang,attr,omitempty" validate:"oneof=aar abk ace ach ada afa afh afr aka akk alb ale alg amh ang apa ara arc arm arn arp art arw asm ath aus ava ave awa aym aze bad bai bak bal bam ban baq bas bat bej bel bem ben ber bho bih bik bin bis bla bnt bod bos bra bre btk bua bug bul bur cad cai car cat cau ceb cel ces cha chb che chg chi chk chm chn cho chp chr chu chv chy cmc cop cor cos cpe cpf cpp cre crp cus cym cze dak dan day del den deu dgr din div doi dra dua dum dut dyu dzo efi egy eka ell elx eng enm epo est eus ewe ewo fan fao fas fat fij fin fiu fon fra fre frm fro fry ful fur gaa gay gba gem geo ger gez gil gla gle glg glv gmh goh gon gor got grb grc gre grn guj gwi hai hau haw heb her hil him hin hit hmn hmo hrv hun hup hye iba ibo ice ijo iku ile ilo ina inc ind ine ipk ira iro isl ita jav jpn jpr jrb kaa kab kac kal kam kan kar kas kat kau kaw kaz kha khi khm kho kik kin kir kmb kok kom kon kor kos kpe kro kru kua kum kur kut lad lah lam lao lat lav lez lin lit lol loz ltz lua lub lug lui lun luo lus mac mad mag mah mai mak mal man mao map mar mas may mdr men mga mic min mis mkd mkh mlg mlt mnc mni mno moh mol mon mos mri msa mul mun mus mwr mya myn nah nai nau nav nbl nde ndo nds nep new nia nic niu nld nno nob non nor nso nub nya nym nyn nyo nzi oci oji ori orm osa oss ota oto paa pag pal pam pan pap pau peo per phi phn pli pol pon por pra pro pus qaa que raj rap rar roa roh rom ron rum run rus sad sag sah sai sal sam san sas sat scc sco scr sel sem sga sgn shn sid sin sio sit sla slk slo slv sme smi smo sna snd snk sog som son sot spa sqi srd srp srr ssa ssw suk sun sus sux swa swe syr tah tai tam tat tel tem ter tet tgk tgl tha tib tig tir tiv tkl tli tmh tog ton tpi tsi tsn tso tuk tum tur tut tvl twi tyv uga uig ukr umb und urd uzb vai ven vie vol vot wak wal war was wel wen wol xho yao yap yid yor ypk zap zen zha zho znd zul"`
	Locale string `xml:"locale,attr,omitempty"`
}

type Name struct {
	multiLocale
	XMLName xml.Name `xml:"bmecat:NAME"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type Name2 struct {
	multiLocale
	XMLName xml.Name `xml:"bmecat:NAME2"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type Name3 struct {
	multiLocale
	XMLName xml.Name `xml:"bmecat:NAME3"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type Department struct {
	multiLocale
	XMLName xml.Name `xml:"bmecat:DEPARTMENT"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type Street struct {
	multiLocale
	XMLName xml.Name `xml:"bmecat:STREET"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type Zip struct {
	multiLocale
	XMLName xml.Name `xml:"bmecat:ZIP"`
	Value   string   `xml:",chardata" validate:"min=1,max=20"`
}

type BoxNo struct {
	multiLocale
	XMLName xml.Name `xml:"bmecat:BOXNO"`
	Value   string   `xml:",chardata" validate:"min=1,max=20"`
}

type ZipBox struct {
	multiLocale
	XMLName xml.Name `xml:"bmecat:ZIPBOX"`
	Value   string   `xml:",chardata" validate:"min=1,max=20"`
}

type City struct {
	multiLocale
	XMLName xml.Name `xml:"bmecat:CITY"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type State struct {
	multiLocale
	XMLName xml.Name `xml:"bmecat:STATE"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type Country struct {
	multiLocale
	XMLName xml.Name `xml:"bmecat:COUNTRY"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

var countryCodedPattern = regexp.MustCompile(`^(AD|AE|AF|AG|AI|AL|AM|AN|AO|AQ|AR|AS|AT|AU|AW|AZ|BA|BB|BD|BE|BF|BG|BH|BI|BJ|BM|BN|BO|BR|BS|BT|BV|BW|BY|BZ|CA|CC|CD|CF|CG|CH|CI|CK|CL|CM|CN|CO|CR|CU|CV|CX|CY|CZ|DE|DJ|DK|DM|DO|DZ|EC|EE|EG|EH|ER|ES|ET|FI|FJ|FK|FM|FO|FR|GA|GB|GD|GE|GF|GH|GI|GL|GM|GN|GP|GQ|GR|GS|GT|GU|GW|GY|HK|HM|HN|HR|HT|HU|ID|IE|IL|IN|IO|IQ|IR|IS|IT|JM|JO|JP|KE|KG|KH|KI|KM|KN|KP|KR|KW|KY|KZ|LA|LB|LC|LI|LK|LR|LS|LT|LU|LV|LY|MA|MC|MD|MG|MK|ML|MM|MN|MO|MP|MQ|MR|MS|MT|MU|MV|MW|MX|MY|MZ|NA|NC|NE|NF|NG|NI|NL|NO|NP|NR|NU|NZ|OM|PA|PE|PF|PG|PH|PK|PL|PM|PN|PR|PS|PT|PW|PY|QA|RE|RO|RU|RW|SA|SB|SC|SD|SE|SG|SH|SI|SJ|SK|SL|SM|SN|SO|SR|ST|SV|SY|SZ|TC|TD|TF|TG|TH|TJ|TK|TM|TN|TO|TP|TR|TT|TV|TW|TZ|UA|UG|US|UY|UZ|VA|VC|VE|VG|VI|VN|VU|WF|WS|YE|YT|YU|ZA|ZM|ZW){1,1}(-[A-Z|0-9]{1,3}){0,1}$`)

//func ValidateCountryCoded(field reflect.Value) interface{} {
//	if valuer, ok := field.Interface().(CountryCoded); ok {
//		countryCodedPattern.Match([]byte(valuer.Value))
//
//	}
//}

type PhoneType string

const (
	PhoneMobile  PhoneType = "mobile"
	PhoneOffice            = "office"
	PhonePrivate           = "private"
)

type Phone struct {
	multiLocale
	XMLName xml.Name  `xml:"bmecat:PHONE"`
	Type    PhoneType `xml:"type,attr,omitempty" validate:"min=1,max=50"`
	Value   string    `xml:",chardata" validate:"min=1,max=50"`
}

type FaxType string

const (
	FaxOffice  FaxType = "office"
	FaxPrivate         = "private"
)

type Fax struct {
	multiLocale
	XMLName xml.Name `xml:"bmecat:FAX"`
	Type    FaxType  `xml:"type,attr,omitempty" validate:"min=1,max=50"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type ContactName struct {
	multiLocale
	XMLName xml.Name `xml:"bmecat:CONTACT_NAME"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type FirstName struct {
	multiLocale
	XMLName xml.Name `xml:"bmecat:FIRST_NAME"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type Title struct {
	multiLocale
	XMLName xml.Name `xml:"bmecat:TITLE"`
	Value   string   `xml:",chardata" validate:"min=1,max=20"`
}

type AcademicTitle struct {
	multiLocale
	XMLName xml.Name `xml:"bmecat:ACADEMIC_TITLE"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type ContactDescription struct {
	multiLocale
	XMLName xml.Name `xml:"bmecat:CONTACT_DESCR"`
	Value   string   `xml:",chardata" validate:"min=1,max=250"`
}

// 1-255
type URL string

type EMails struct {
	XMLName xml.Name `xml:"bmecat:EMAILS"`
	EMail   []string `xml:"bmecat:EMAIL" validate:"required,dive,required,max=255"`
}

type Authentification struct {
	XMLName  xml.Name `xml:"bmecat:AUTHENTIFICATION"`
	Login    string   `xml:"bmecat:LOGIN" validate:"required,max=60"`
	Password string   `xml:"bmecat:PASSWORD,omitempty" validate:"max=20"`
}
