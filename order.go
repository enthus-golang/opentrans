package opentrans

import (
	"encoding/xml"

	"gitlab.com/mclgmbh/gomod/bmecat"
)

type OrderType string

const (
	OrderTypeStandard    = "standard"
	OrderTypeExpress     = "express"
	OrderTypeRelease     = "release"
	OrderTypeConsignment = "consignment"
)

type Order struct {
	XMLName       xml.Name  `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER" json:"-" yaml:"-"`
	Version       string    `xml:"version,attr" validate:"required,eq=2.1"`
	Type          OrderType `xml:"type,attr" validate:"required,min=1,max=20,oneof=standard express release consignment"`
	OrderHeader   OrderHeader
	OrderItemList OrderItemList
	OrderSummary  OrderSummary
}

func NewOrder(typ OrderType) *Order {
	return &Order{
		Version: "2.1",
		Type:    typ,
	}
}

type OrderHeader struct {
	XMLName      xml.Name      `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_HEADER"`
	ControlInfo  *ControlInfo  `xml:",omitempty"`
	SourcingInfo *SourcingInfo `xml:",omitempty"`
	OrderInfo    *OrderInfo    `validate:"required"`
}

type OrderItemList struct {
	XMLName xml.Name `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_ITEM_LIST"`
}
type OrderSummary struct {
	XMLName xml.Name `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_SUMMARY"`
}

type SourcingInfo struct {
	XMLName xml.Name `xml:"http://www.opentrans.org/XMLSchema/2.1 SOURCING_INFO"`
}

type OrderInfo struct {
	XMLName                xml.Name                `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_INFO"`
	OrderID                string                  `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_ID" validate:"min=1,max=250"`
	OrderDate              Datetime                `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_DATE"`
	DeliveryDate           *Datetime               `xml:"http://www.opentrans.org/XMLSchema/2.1 DELIVERY_DATE"`
	Language               []bmecat.Language       `xml:"http://www.bmecat.org/bmecat/2005 LANGUAGE"`
	MIMERoot               string                  `xml:"http://www.bmecat.org/bmecat/2005 MIME_ROOT,omitempty" validate:"max=250"`
	Parties                Parties                 `xml:"http://www.opentrans.org/XMLSchema/2.1 PARTIES"`
	CustomerOrderReference *CustomerOrderReference `xml:"http://www.opentrans.org/XMLSchema/2.1 CUSTOMER_ORDER_REFERENCE,omitempty"`
	OrderPartiesReference  OrderPartiesReference   `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_PARTIES_REFERENCE"`
}

type CustomerOrderReference struct {
	XMLName xml.Name `xml:"http://www.opentrans.org/XMLSchema/2.1 CUSTOMER_ORDER_REFERENCE"`
}

type OrderPartiesReference struct {
	XMLName       xml.Name              `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_PARTIES_REFERENCE"`
	BuyerIDRef    *bmecat.BuyerIDRef    `xml:"http://www.bmecat.org/bmecat/2005 BUYER_IDREF,omitempty"`
	SupplierIDRef *bmecat.SupplierIDRef `xml:"http://www.bmecat.org/bmecat/2005 SUPPLIER_IDREF,omitempty"`
}
