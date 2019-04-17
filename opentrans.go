package opentrans

import (
	"encoding/xml"

	"gitlab.com/mclgmbh/gomod/bmecat"
)

type ControlInfo struct {
	XMLName                 xml.Name `xml:"http://www.opentrans.org/XMLSchema/2.1 CONTROL_INFO"`
	StopAutomaticProcessing string   `xml:"http://www.opentrans.org/XMLSchema/2.1 STOP_AUTOMATIC_PROCESSING,omitempty" validate:"max=250"`
	GeneratorInfo           string   `xml:"http://www.opentrans.org/XMLSchema/2.1 GENERATOR_INFO,omitempty" validate:"max=250"`
	GenerationDate          string   `xml:"http://www.opentrans.org/XMLSchema/2.1 GENERATION_DATE"`
}

type Parties struct {
	XMLName xml.Name `xml:"http://www.opentrans.org/XMLSchema/2.1 PARTIES"`
	Party   []Party  `xml:"http://www.opentrans.org/XMLSchema/2.1 PARTY" validate:"required,min=1,dive"`
}

type PartyRole string

const (
	PartyRoleBuyer               PartyRole = "buyer"
	PartyRoleCentralRegulator              = "central_regulator"
	PartyRoleCustomer                      = "customer"
	PartyRoleDeliverer                     = "deliverer"
	PartyRoleDelivery                      = "delivery"
	PartyRoleDocumentCreator               = "document_creator"
	PartyRoleFinalDelivery                 = "final_delivery"
	PartyRoleIntermediary                  = "intermediary"
	PartyRoleInvoiceIssuer                 = "invoice_issuer"
	PartyRoleInvoiceRecipient              = "invoice_recipient"
	PartyRoleIPPOperator                   = "ipp_operator"
	PartyRoleManufacturer                  = "manufacturer"
	PartyRoleMarketplace                   = "marketplace"
	PartyRolePayer                         = "payer"
	PartyRoleRemittee                      = "remittee"
	PartyRoleStandardizationBody           = "standardization_body"
	PartyRoleSupplier                      = "supplier"
	PartyRoleThrustedThirdParty            = "trustedthirdparty"
	PartyRoleOther                         = "other"
)

type Party struct {
	XMLName   xml.Name         `xml:"http://www.opentrans.org/XMLSchema/2.1 PARTY"`
	PartyID   []bmecat.PartyID `xml:"http://www.bmecat.org/bmecat/2005 PARTY_ID"`
	PartyRole []PartyRole      `xml:"http://www.opentrans.org/XMLSchema/2.1 PARTY_ROLE" validate:"dive,oneof=buyer central_regulator customer deliverer delivery document_creator final_delivery intermediary invoice_issuer invoice_recipient ipp_operator manufacturer marketplace payer remittee standardization_body supplier trustedthirdparty other"`
	Address   []Address        `xml:"http://www.opentrans.org/XMLSchema/2.1 ADDRESS"`
	Account   string           `xml:"http://www.opentrans.org/XMLSchema/2.1 ACCOUNT"`
	MIMEInfo  string           `xml:"http://www.opentrans.org/XMLSchema/2.1 MIME_INFO"`
}

//func (p *Party) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
//	for {
//		t, _ := d.Token()
//		switch tt := t.(type) {
//		case xml.StartElement:
//			fmt.Println(">", tt.Name.Local, tt.Name.Space)
//		case xml.EndElement:
//			fmt.Println("<", tt.Name.Local, tt.Name.Space)
//		}
//	}
//}

type Address struct {
	XMLName        xml.Name            `xml:"http://www.opentrans.org/XMLSchema/2.1 ADDRESS"`
	Name           []bmecat.Name       `xml:"http://www.bmecat.org/bmecat/2005 NAME"`
	Name2          []bmecat.Name2      `xml:"http://www.bmecat.org/bmecat/2005 NAME2"`
	Name3          []bmecat.Name3      `xml:"http://www.bmecat.org/bmecat/2005 NAME3"`
	Department     []bmecat.Department `xml:"http://www.bmecat.org/bmecat/2005 DEPARTMENT"`
	ContactDetails []ContactDetails    `xml:"http://www.opentrans.org/XMLSchema/2.1 CONTACT_DETAILS"`
	Street         []bmecat.Street     `xml:"http://www.bmecat.org/bmecat/2005 STREET"`
	Zip            []bmecat.Zip        `xml:"http://www.bmecat.org/bmecat/2005 ZIP"`
	BoxNo          []bmecat.BoxNo      `xml:"http://www.bmecat.org/bmecat/2005 BOXNO"`
	ZipBox         []bmecat.ZipBox     `xml:"http://www.bmecat.org/bmecat/2005 ZIPBOX"`
	City           []bmecat.City       `xml:"http://www.bmecat.org/bmecat/2005 CITY"`
	State          []bmecat.State      `xml:"http://www.bmecat.org/bmecat/2005 STATE"`
	Country        []bmecat.Country    `xml:"http://www.bmecat.org/bmecat/2005 COUNTRY"`
	CountryCoded   bmecat.CountryCoded `xml:"http://www.bmecat.org/bmecat/2005 COUNTRY_CODED,omitempty"`
	VatID          bmecat.VatID        `xml:"http://www.bmecat.org/bmecat/2005 VAT_ID,omitempty"`
	Phone          []bmecat.Phone      `xml:"http://www.bmecat.org/bmecat/2005 PHONE"`
	Fax            []bmecat.Fax        `xml:"http://www.bmecat.org/bmecat/2005 FAX"`
	EMail          []bmecat.EMail      `xml:"http://www.bmecat.org/bmecat/2005 EMAIL" validate:"required,dive,required,max=255"`
	URL            bmecat.URL          `xml:"http://www.bmecat.org/bmecat/2005 URL,omitempty" validate:"max=255"`
}

type ContactDetails struct {
	XMLName            xml.Name                    `xml:"http://www.opentrans.org/XMLSchema/2.1 CONTACT_DETAILS"`
	ContactID          string                      `xml:"bmecat:CONTACT_ID,omitempty" validate:"max=60"`
	ContactName        []bmecat.ContactName        `xml:"http://www.bmecat.org/bmecat/2005 CONTACT_NAME"`
	FirstName          []bmecat.FirstName          `xml:"http://www.bmecat.org/bmecat/2005 FIRST_NAME"`
	Title              []bmecat.Title              `xml:"http://www.bmecat.org/bmecat/2005 TITLE"`
	AcademicTitle      []bmecat.AcademicTitle      `xml:"http://www.bmecat.org/bmecat/2005 ACADEMIC_TITLE"`
	ContractRole       []ContractRole              `xml:"http://www.opentrans.org/XMLSchema/2.1 CONTACT_ROLE"`
	ContactDescription []bmecat.ContactDescription `xml:"http://www.bmecat.org/bmecat/2005 CONTACT_DESCR"`
	Phone              []bmecat.Phone              `xml:"http://www.bmecat.org/bmecat/2005 PHONE"`
	Fax                []bmecat.Fax                `xml:"http://www.bmecat.org/bmecat/2005 FAX"`
	URL                bmecat.URL                  `xml:"http://www.bmecat.org/bmecat/2005 URL,omitempty" validate:"max=255"`
	EMails             *bmecat.EMails              `xml:"http://www.bmecat.org/bmecat/2005 EMAILS"`
}

type ContractRoleType string

const (
	ContactRoleAdministrativ    ContractRoleType = "administrativ"
	ContactRoleCommercial                        = "commercial"
	ContactRoleDocumentIssuer                    = "document_issuer"
	ContactRoleSpecialTreatment                  = "special_treatment"
	ContactRoleTechnical                         = "technical"
	ContactRoleOthers                            = "others"
)

type ContractRole struct {
	multiLocaleString00050
	XMLName xml.Name         `xml:"http://www.opentrans.org/XMLSchema/2.1 CONTACT_ROLE"`
	Type    ContractRoleType `xml:"type,attr" validate:"omitempty,oneof=administrativ commercial document_issuer special_treatment technical others"`
}
