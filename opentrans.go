package opentrans

import "encoding/xml"

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
	XMLName   xml.Name  `xml:"PARTY"`
	PartyID   PartyID   `xml:"bmecat:PARTY_ID"`
	PartyRole PartyRole `xml:"PARTY_ROLE" validate:"oneof=buyer central_regulator customer deliverer delivery document_creator final_delivery intermediary invoice_issuer invoice_recipient ipp_operator manufacturer marketplace payer remittee standardization_body supplier trustedthirdparty other"`
	Address   []Address `xml:"ADDRESS"`
	Account   string    `xml:"ACCOUNT"`
	MIMEInfo  string    `xml:"MIME_INFO"`
}

type Address struct {
	XMLName        xml.Name         `xml:"ADDRESS"`
	Name           []Name           `xml:"bmecat:NAME"`
	Name2          []Name2          `xml:"bmecat:NAME2"`
	Name3          []Name3          `xml:"bmecat:NAME3"`
	Department     []Department     `xml:"bmecat:DEPARTMENT"`
	ContactDetails []ContactDetails `xml:"CONTACT_DETAILS"`
	Street         []Street         `xml:"bmecat:STREET"`
	Zip            []Zip            `xml:"bmecat:ZIP"`
	BoxNo          []BoxNo          `xml:"bmecat:BOXNO"`
	ZipBox         []ZipBox         `xml:"bmecat:ZIPBOX"`
	City           []City           `xml:"bmecat:CITY"`
	State          []State          `xml:"bmecat:STATE"`
	Country        []Country        `xml:"bmecat:COUNTRY"`
	CountryCoded   string           `xml:"bmecat:COUNTRY_CODED,omitempty"`
	VatID          string           `xml:"bmecat:VAT_ID,omitempty"`
	Phone          []Phone          `xml:"bmecat:PHONE"`
	Fax            []Fax            `xml:"bmecat:FAX"`
	EMail          []string         `xml:"bmecat:EMAIL" validate:"required,dive,required,max=255"`
	URL            URL              `xml:"bmecat:URL,omitempty" validate:"max=255"`
}

type ContactDetails struct {
	XMLName            xml.Name             `xml:"CONTACT_DETAILS"`
	ContactID          string               `xml:"bmecat:CONTACT_ID,omitempty" validate:"max=60"`
	ContactName        []ContactName        `xml:"bmecat:CONTACT_NAME"`
	FirstName          []FirstName          `xml:"bmecat:FIRST_NAME"`
	Title              []Title              `xml:"bmecat:TITLE"`
	AcademicTitle      []AcademicTitle      `xml:"bmecat:ACADEMIC_TITLE"`
	ContractRole       []ContractRole       `xml:"CONTACT_ROLE"`
	ContactDescription []ContactDescription `xml:"bmecat:CONTACT_DESCR"`
	Phone              []Phone              `xml:"bmecat:PHONE"`
	Fax                []Fax                `xml:"bmecat:FAX"`
	URL                URL                  `xml:"bmecat:URL,omitempty" validate:"max=255"`
	EMails             *EMails              `xml:"bmecat:EMAILS"`
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
