package opentrans

import (
	"encoding/xml"

	"gitlab.com/mclgmbh/golang-pkg/bmecat"
)

type OrderResponse struct {
	XMLName             xml.Name            `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDERRESPONSE"`
	Version             string              `xml:"version,attr" validate:"required,eq=2.1"`
	OrderResponseHeader OrderResponseHeader `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDERRESPONSE_HEADER"`
}

type OrderResponseHeader struct {
	XMLName           xml.Name           `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDERRESPONSE_HEADER"`
	ControlInfo       *ControlInfo       `xml:",omitempty"`
	OrderResponseInfo *OrderResponseInfo `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDERRESPONSE_INFO" validate:"required"`
}

type OrderResponseInfo struct {
	XMLName                    xml.Name                    `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDERRESPONSE_INFO"`
	OrderID                    string                      `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_ID" validate:"min=1,max=250"`
	OrderResponseDate          bmecat.Datetime             `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDERRESPONSE_DATE"`
	OrderDate                  *bmecat.Datetime            `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_DATE"`
	AltCustomerOrderID         []string                    `xml:"http://www.opentrans.org/XMLSchema/2.1 ALT_CUSTOMER_ORDER_ID" validate:"dive,max=250"`
	SupplierOrderID            string                      `xml:"http://www.opentrans.org/XMLSchema/2.1 SUPPLIER_ORDER_ID" validate:"max=250"`
	DeliveryDate               *bmecat.Datetime            `xml:"http://www.opentrans.org/XMLSchema/2.1 DELIVERY_DATE"`
	Language                   []bmecat.Language           `xml:"http://www.bmecat.org/bmecat/2005 LANGUAGE"`
	MIMERoot                   string                      `xml:"http://www.bmecat.org/bmecat/2005 MIME_ROOT,omitempty" validate:"max=250"`
	Parties                    Parties                     `xml:"http://www.opentrans.org/XMLSchema/2.1 PARTIES"`
	OrderPartiesReference      OrderPartiesReference       `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_PARTIES_REFERENCE"`
	Currency                   bmecat.Currency             `xml:"http://www.bmecat.org/bmecat/2005 CURRENCY" validate:"required"`
	Remarks                    []Remarks                   `xml:"http://www.opentrans.org/XMLSchema/2.1 REMARKS"`
	HeaderUserDefinedExtension *HeaderUserDefinedExtension `xml:"http://www.opentrans.org/XMLSchema/2.1 HEADER_UDX"`
}
