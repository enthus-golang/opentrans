package opentrans

import (
	"encoding/xml"

	"gitlab.com/mclgmbh/golang-pkg/bmecat"
)

type Invoice struct {
	XMLName         xml.Name `xml:"http://www.opentrans.org/XMLSchema/2.1 INVOICE" json:"-" yaml:"-"`
	Namespace       string   `xml:"xmlns,attr"`
	NamespaceBMEcat string   `xml:"xmlns:bmecat,attr"`
	Version         string   `xml:"version,attr" validate:"required,eq=2.1"`
	InvoiceHeader   InvoiceHeader
	InvoiceItemList InvoiceItemList
	InvoiceSummary  InvoiceSummary
}

type InvoiceHeader struct {
	XMLName      xml.Name      `xml:"http://www.opentrans.org/XMLSchema/2.1 INVOICE_HEADER"`
	ControlInfo  *ControlInfo  `xml:",omitempty"`
	SourcingInfo *SourcingInfo `xml:",omitempty"`
	InvoiceInfo  *InvoiceInfo  `validate:"required"`
}

type InvoiceInfo struct {
	XMLName       xml.Name              `xml:"http://www.opentrans.org/XMLSchema/2.1 INVOICE_INFO"`
	InvoiceID     string                `xml:"http://www.opentrans.org/XMLSchema/2.1 INVOICE_ID" validate:"min=1,max=250"`
	InvoiceDate   string                `xml:"http://www.opentrans.org/XMLSchema/2.1 INVOICE_DATE"`
	DeliveryDate  *string               `xml:"http://www.opentrans.org/XMLSchema/2.1 DELIVERY_DATE"`
	Language      []bmecat.Language     `xml:"http://www.bmecat.org/bmecat/2005 LANGUAGE"`
	MIMERoot      string                `xml:"http://www.bmecat.org/bmecat/2005 MIME_ROOT,omitempty" validate:"max=250"`
	Parties       Parties               `xml:"http://www.opentrans.org/XMLSchema/2.1 PARTIES"`
	BuyerIDRef    *bmecat.BuyerIDRef    `xml:"http://www.bmecat.org/bmecat/2005 BUYER_IDREF"`
	SupplierIDRef *bmecat.SupplierIDRef `xml:"http://www.bmecat.org/bmecat/2005 SUPPLIER_IDREF"`
	Currency      bmecat.Currency       `xml:"http://www.bmecat.org/bmecat/2005 CURRENCY" validate:"required"`
}

type InvoiceItemList struct {
	XMLName xml.Name `xml:"http://www.opentrans.org/XMLSchema/2.1 INVOICE_ITEM_LIST"`
}
type InvoiceSummary struct {
	XMLName xml.Name `xml:"http://www.opentrans.org/XMLSchema/2.1 INVOICE_SUMMARY"`
}
