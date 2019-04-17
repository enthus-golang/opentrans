package opentrans

import (
	"encoding/xml"

	"gitlab.com/mclgmbh/gomod/bmecat"
)

type OrderType string

const (
	OrderTypeStandard    OrderType = "standard"
	OrderTypeExpress               = "express"
	OrderTypeRelease               = "release"
	OrderTypeConsignment           = "consignment"
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
	XMLName   xml.Name    `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_ITEM_LIST"`
	OrderItem []OrderItem `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_ITEM" validate:"required,gt=0,dive,required"`
}

type OrderItem struct {
	XMLName         xml.Name         `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_ITEM"`
	LineItemID      string           `xml:"http://www.opentrans.org/XMLSchema/2.1 LINE_ITEM_ID" validate:"required,max=50"`
	ProductID       ProductID        `xml:"http://www.opentrans.org/XMLSchema/2.1 PRODUCT_ID" validate:"required"`
	Quantity        float64          `xml:"http://www.opentrans.org/XMLSchema/2.1 QUANTITY" validate:"required"`
	OrderUnit       bmecat.OrderUnit `xml:"http://www.bmecat.org/bmecat/2005 ORDER_UNIT" validate:"required,max=3"`
	PriceLineAmount float64          `xml:"http://www.opentrans.org/XMLSchema/2.1 PRICE_LINE_AMOUNT,omitempty"`
	ProductPriceFix *ProductPriceFix `xml:"http://www.opentrans.org/XMLSchema/2.1 PRODUCT_PRICE_FIX"`
}

type ProductID struct {
	XMLName          xml.Name                 `xml:"http://www.opentrans.org/XMLSchema/2.1 PRODUCT_ID"`
	SupplierPID      *bmecat.SupplierPID      `xml:"http://www.bmecat.org/bmecat/2005 SUPPLIER_PID"`
	SupplierIDRef    *bmecat.SupplierIDRef    `xml:"http://www.bmecat.org/bmecat/2005 SUPPLIER_IDREF"`
	ConfigCodeFix    string                   `xml:"http://www.opentrans.org/XMLSchema/2.1 CONFIG_CODE_FIX,omitempty" validate:"max=6000"`
	LOTNumber        []string                 `xml:"http://www.opentrans.org/XMLSchema/2.1 LOT_NUMBER" validate:"dive,max=80"`
	SerialNumber     []string                 `xml:"http://www.opentrans.org/XMLSchema/2.1 SERIAL_NUMBER" validate:"dive,max=80"`
	InternationalPID *bmecat.InternationalPID `xml:"http://www.bmecat.org/bmecat/2005 INTERNATIONAL_PID"`
	BuyerPID         *bmecat.BuyerPID         `xml:"http://www.bmecat.org/bmecat/2005 BUYER_PID"`
	DescriptionShort *DescriptionShort        `xml:"http://www.opentrans.org/XMLSchema/2.1 DESCRIPTION_SHORT"`
	DescriptionLong  *DescriptionLong         `xml:"http://www.opentrans.org/XMLSchema/2.1 DESCRIPTION_LONG"`
}

type DescriptionShort struct {
	multiLocaleString00150
	XMLName xml.Name `xml:"http://www.opentrans.org/XMLSchema/2.1 DESCRIPTION_SHORT"`
}

type DescriptionLong struct {
	multiLocaleString64000
	XMLName xml.Name `xml:"http://www.opentrans.org/XMLSchema/2.1 DESCRIPTION_LONG"`
}

type ProductPriceFix struct {
	XMLName     xml.Name           `xml:"http://www.opentrans.org/XMLSchema/2.1 PRODUCT_PRICE_FIX"`
	PriceAmount bmecat.PriceAmount `xml:"http://www.bmecat.org/bmecat/2005 PRICE_AMOUNT" validate:"required"`
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
	OrderDate              bmecat.Datetime         `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_DATE"`
	DeliveryDate           *bmecat.Datetime        `xml:"http://www.opentrans.org/XMLSchema/2.1 DELIVERY_DATE"`
	Language               []bmecat.Language       `xml:"http://www.bmecat.org/bmecat/2005 LANGUAGE"`
	MIMERoot               string                  `xml:"http://www.bmecat.org/bmecat/2005 MIME_ROOT,omitempty" validate:"max=250"`
	Parties                Parties                 `xml:"http://www.opentrans.org/XMLSchema/2.1 PARTIES"`
	CustomerOrderReference *CustomerOrderReference `xml:"http://www.opentrans.org/XMLSchema/2.1 CUSTOMER_ORDER_REFERENCE,omitempty"`
	OrderPartiesReference  OrderPartiesReference   `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_PARTIES_REFERENCE"`
	Currency               bmecat.EMail            `xml:"http://www.bmecat.org/bmecat/2005 CURRENCY" validate:"required"`
	PartialShipmentAllowed bool                    `xml:"http://www.opentrans.org/XMLSchema/2.1 PARTIAL_SHIPMENT_ALLOWED"`
}

type CustomerOrderReference struct {
	XMLName xml.Name `xml:"http://www.opentrans.org/XMLSchema/2.1 CUSTOMER_ORDER_REFERENCE"`
}

type OrderPartiesReference struct {
	XMLName       xml.Name              `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_PARTIES_REFERENCE"`
	BuyerIDRef    *bmecat.BuyerIDRef    `xml:"http://www.bmecat.org/bmecat/2005 BUYER_IDREF,omitempty"`
	SupplierIDRef *bmecat.SupplierIDRef `xml:"http://www.bmecat.org/bmecat/2005 SUPPLIER_IDREF,omitempty"`
}
