package opentrans

import (
	"encoding/xml"

	"gitlab.com/mclgmbh/gomod/bmecat"
)

type Parties struct {
	XMLName xml.Name `xml:"PARTIES"`
	Party   []Party  `xml:"PARTY" validate:"required,min=1,dive"`
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
	XMLName   xml.Name       `xml:"PARTY"`
	PartyID   bmecat.PartyID `xml:"bmecat:PARTY_ID"`
	PartyRole PartyRole      `xml:"PARTY_ROLE" validate:"oneof=buyer central_regulator customer deliverer delivery document_creator final_delivery intermediary invoice_issuer invoice_recipient ipp_operator manufacturer marketplace payer remittee standardization_body supplier trustedthirdparty other"`
	Address   []Address      `xml:"ADDRESS"`
	Account   string         `xml:"ACCOUNT"`
	MIMEInfo  string         `xml:"MIME_INFO"`
}

type Address struct {
	XMLName        xml.Name            `xml:"ADDRESS"`
	Name           []bmecat.Name       `xml:"bmecat:NAME"`
	Name2          []bmecat.Name2      `xml:"bmecat:NAME2"`
	Name3          []bmecat.Name3      `xml:"bmecat:NAME3"`
	Department     []bmecat.Department `xml:"bmecat:DEPARTMENT"`
	ContactDetails []ContactDetails    `xml:"CONTACT_DETAILS"`
	Street         []bmecat.Street     `xml:"bmecat:STREET"`
	Zip            []bmecat.Zip        `xml:"bmecat:ZIP"`
	BoxNo          []bmecat.BoxNo      `xml:"bmecat:BOXNO"`
	ZipBox         []bmecat.ZipBox     `xml:"bmecat:ZIPBOX"`
	City           []bmecat.City       `xml:"bmecat:CITY"`
	State          []bmecat.State      `xml:"bmecat:STATE"`
	Country        []bmecat.Country    `xml:"bmecat:COUNTRY"`
	CountryCoded   string              `xml:"bmecat:COUNTRY_CODED,omitempty"`
	VatID          string              `xml:"bmecat:VAT_ID,omitempty"`
	Phone          []bmecat.Phone      `xml:"bmecat:PHONE"`
	Fax            []bmecat.Fax        `xml:"bmecat:FAX"`
	EMail          []string            `xml:"bmecat:EMAIL" validate:"required,dive,required,max=255"`
	URL            bmecat.URL          `xml:"bmecat:URL,omitempty" validate:"max=255"`
}

type ContactDetails struct {
	XMLName            xml.Name                    `xml:"CONTACT_DETAILS"`
	ContactID          string                      `xml:"bmecat:CONTACT_ID,omitempty" validate:"max=60"`
	ContactName        []bmecat.ContactName        `xml:"bmecat:CONTACT_NAME"`
	FirstName          []bmecat.FirstName          `xml:"bmecat:FIRST_NAME"`
	Title              []bmecat.Title              `xml:"bmecat:TITLE"`
	AcademicTitle      []bmecat.AcademicTitle      `xml:"bmecat:ACADEMIC_TITLE"`
	ContractRole       []ContractRole              `xml:"CONTACT_ROLE"`
	ContactDescription []bmecat.ContactDescription `xml:"bmecat:CONTACT_DESCR"`
	Phone              []bmecat.Phone              `xml:"bmecat:PHONE"`
	Fax                []bmecat.Fax                `xml:"bmecat:FAX"`
	URL                bmecat.URL                  `xml:"bmecat:URL,omitempty" validate:"max=255"`
	EMails             *bmecat.EMails              `xml:"bmecat:EMAILS"`
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
	XMLName xml.Name         `xml:"CONTACT_ROLE"`
	Type    ContractRoleType `xml:"type,attr" validate:"omitempty,oneof=administrativ commercial document_issuer special_treatment technical others"`
}
