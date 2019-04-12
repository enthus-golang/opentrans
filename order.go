package opentrans

import (
	"encoding/xml"
	"time"

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
	XMLName         xml.Name  `xml:"ORDER" json:"-" yaml:"-"`
	Namespace       string    `xml:"xmlns,attr"`
	NamespaceBMEcat string    `xml:"xmlns:bmecat,attr"`
	Version         string    `xml:"version,attr" validate:"required,eq=2.1"`
	Type            OrderType `xml:"type,attr" validate:"required,min=1,max=20,oneof=standard express release consignment"`
	OrderHeader     OrderHeader
	OrderItemList   OrderItemList
	OrderSummary    OrderSummary
}

func NewOrder(typ OrderType) *Order {
	return &Order{
		Namespace:       "http://www.opentrans.org/XMLSchema/2.1",
		NamespaceBMEcat: bmecat.Namespace,
		Version:         "2.1",
		Type:            typ,
	}
}

type OrderHeader struct {
	XMLName      xml.Name          `xml:"ORDER_HEADER"`
	ControlInfo  *OrderControlInfo `xml:",omitempty"`
	SourcingInfo *SourcingInfo     `xml:",omitempty"`
	OrderInfo    *OrderInfo        `validate:"required"`
}

type OrderItemList struct {
	XMLName xml.Name `xml:"ORDER_ITEM_LIST"`
}
type OrderSummary struct {
	XMLName xml.Name `xml:"ORDER_SUMMARY"`
}

type OrderControlInfo struct {
	XMLName                 xml.Name   `xml:"CONTROL_INFO"`
	StopAutomaticProcessing string     `xml:"STOP_AUTOMATIC_PROCESSING,omitempty" validate:"max=250"`
	GeneratorInfo           string     `xml:"GENERATOR_INFO,omitempty" validate:"max=250"`
	GenerationDate          *time.Time `xml:"GENERATION_DATE"`
}

type SourcingInfo struct {
	XMLName xml.Name `xml:"SOURCING_INFO"`
}

type OrderInfo struct {
	XMLName                xml.Name               `xml:"ORDER_INFO"`
	OrderID                string                 `xml:"ORDER_ID" validate:"min=1,max=250"`
	OrderDate              time.Time              `xml:"ORDER_DATE"`
	DeliveryDate           *time.Time             `xml:"DELIVERY_DATE"`
	Language               []bmecat.Language      `xml:"bmecat:LANGUAGE"`
	MIMERoot               string                 `xml:"bmecat:MIME_ROOT,omitempty" validate:"max=250"`
	Parties                Parties                `xml:"PARTIES"`
	CustomerOrderReference CustomerOrderReference `xml:"CUSTOMER_ORDER_REFERENCE"`
}

type CustomerOrderReference struct {
	XMLName xml.Name `xml:"CUSTOMER_ORDER_REFERENCE"`
}
