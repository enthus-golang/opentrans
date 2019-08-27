package opentrans

import (
	"encoding/xml"
	"strings"

	"gitlab.com/mclgmbh/golang-pkg/bmecat"
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
	XMLName                  xml.Name                  `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_ITEM"`
	LineItemID               string                    `xml:"http://www.opentrans.org/XMLSchema/2.1 LINE_ITEM_ID" validate:"required,max=50"`
	ProductID                ProductID                 `xml:"http://www.opentrans.org/XMLSchema/2.1 PRODUCT_ID" validate:"required"`
	Quantity                 float64                   `xml:"http://www.opentrans.org/XMLSchema/2.1 QUANTITY" validate:"required"`
	OrderUnit                bmecat.OrderUnit          `xml:"http://www.bmecat.org/bmecat/2005 ORDER_UNIT" validate:"required,max=3"`
	ProductPriceFix          *ProductPriceFix          `xml:"http://www.opentrans.org/XMLSchema/2.1 PRODUCT_PRICE_FIX"`
	PriceLineAmount          float64                   `xml:"http://www.opentrans.org/XMLSchema/2.1 PRICE_LINE_AMOUNT,omitempty"`
	CustomerOrderReference   *CustomerOrderReference   `xml:"http://www.opentrans.org/XMLSchema/2.1 CUSTOMER_ORDER_REFERENCE,omitempty"`
	Remarks                  []Remarks                 `xml:"http://www.opentrans.org/XMLSchema/2.1 REMARKS"`
	ItemUserDefinedExtension *ItemUserDefinedExtension `xml:"http://www.opentrans.org/XMLSchema/2.1 ITEM_UDX"`
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
	XMLName      xml.Name `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_SUMMARY"`
	TotalItemNum int      `xml:"http://www.opentrans.org/XMLSchema/2.1 TOTAL_ITEM_NUM" validate:"gte=0"`
	TotalAmount  *float64 `xml:"http://www.opentrans.org/XMLSchema/2.1 TOTAL_AMOUNT"`
}

type SourcingInfo struct {
	XMLName     xml.Name   `xml:"http://www.opentrans.org/XMLSchema/2.1 SOURCING_INFO"`
	QuotationID string     `xml:"http://www.opentrans.org/XMLSchema/2.1 QUOTATION_ID,omitempty" validate:"max=250"`
	Agreement   *Agreement `xml:"http://www.opentrans.org/XMLSchema/2.1 AGREEMENT"`
}

type AgreementDescription struct {
	multiLocaleString00250
	XMLName xml.Name `xml:"http://www.opentrans.org/XMLSchema/2.1 AGREEMENT_DESCR"`
}

type Agreement struct {
	XMLName              xml.Name              `xml:"http://www.opentrans.org/XMLSchema/2.1 AGREEMENT"`
	Type                 string                `xml:"type,omitempty" validate:"max=50"`
	Default              string                `xml:"default,omitempty" validate:"omitempty,oneof=true false"`
	AgreementID          string                `xml:"http://www.bmecat.org/bmecat/2005 AGREEMENT_ID,omitempty" validate:"min=1,max=50"`
	AgreementStartDate   *bmecat.Datetime      `xml:"http://www.bmecat.org/bmecat/2005 AGREEMENT_START_DATE"`
	AgreementEndDate     *bmecat.Datetime      `xml:"http://www.bmecat.org/bmecat/2005 AGREEMENT_END_DATE" validate:"required"`
	AgreementDescription *AgreementDescription `xml:"http://www.opentrans.org/XMLSchema/2.1 AGREEMENT_DESCR,omitempty"`
}

type OrderInfo struct {
	XMLName                    xml.Name                    `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_INFO"`
	OrderID                    string                      `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_ID" validate:"min=1,max=250"`
	OrderDate                  bmecat.Datetime             `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_DATE"`
	DeliveryDate               *bmecat.Datetime            `xml:"http://www.opentrans.org/XMLSchema/2.1 DELIVERY_DATE"`
	Language                   []bmecat.Language           `xml:"http://www.bmecat.org/bmecat/2005 LANGUAGE"`
	MIMERoot                   string                      `xml:"http://www.bmecat.org/bmecat/2005 MIME_ROOT,omitempty" validate:"max=250"`
	Parties                    Parties                     `xml:"http://www.opentrans.org/XMLSchema/2.1 PARTIES"`
	CustomerOrderReference     *CustomerOrderReference     `xml:"http://www.opentrans.org/XMLSchema/2.1 CUSTOMER_ORDER_REFERENCE,omitempty"`
	OrderPartiesReference      OrderPartiesReference       `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_PARTIES_REFERENCE"`
	Currency                   bmecat.Currency             `xml:"http://www.bmecat.org/bmecat/2005 CURRENCY" validate:"required"`
	PartialShipmentAllowed     bool                        `xml:"http://www.opentrans.org/XMLSchema/2.1 PARTIAL_SHIPMENT_ALLOWED"`
	Remarks                    []Remarks                   `xml:"http://www.opentrans.org/XMLSchema/2.1 REMARKS"`
	HeaderUserDefinedExtension *HeaderUserDefinedExtension `xml:"http://www.opentrans.org/XMLSchema/2.1 HEADER_UDX"`
}

type CustomerOrderReference struct {
	XMLName    xml.Name         `xml:"http://www.opentrans.org/XMLSchema/2.1 CUSTOMER_ORDER_REFERENCE"`
	OrderID    string           `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_ID,omitempty"`
	LineItemID string           `xml:"http://www.opentrans.org/XMLSchema/2.1 LINE_ITEM_ID,omitempty"`
	OrderDate  *bmecat.Datetime `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_DATE"`
	//OrderDescription string          `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_DESCR"`
}

type OrderPartiesReference struct {
	XMLName                  xml.Name                  `xml:"http://www.opentrans.org/XMLSchema/2.1 ORDER_PARTIES_REFERENCE"`
	BuyerIDRef               bmecat.BuyerIDRef         `xml:"http://www.bmecat.org/bmecat/2005 BUYER_IDREF,omitempty"`
	SupplierIDRef            bmecat.SupplierIDRef      `xml:"http://www.bmecat.org/bmecat/2005 SUPPLIER_IDREF,omitempty"`
	InvoiceRecipientIDRef    *InvoiceRecipientIDRef    `xml:"http://www.opentrans.org/XMLSchema/2.1 INVOICE_RECIPIENT_IDREF"`
	ShipmentPartiesReference *ShipmentPartiesReference `xml:"http://www.opentrans.org/XMLSchema/2.1 SHIPMENT_PARTIES_REFERENCE"`
}

type InvoiceRecipientIDRef struct {
	XMLName xml.Name `xml:"http://www.opentrans.org/XMLSchema/2.1 INVOICE_RECIPIENT_IDREF"`
	Type    string   `xml:"type,attr,omitempty" validate:"min=1,max=250"`
	Value   string   `xml:",chardata" validate:"min=1,max=250"`
}

type ShipmentPartiesReference struct {
	XMLName            xml.Name            `xml:"http://www.opentrans.org/XMLSchema/2.1 SHIPMENT_PARTIES_REFERENCE"`
	DeliveryIDRef      DeliveryIDRef       `xml:"http://www.opentrans.org/XMLSchema/2.1 DELIVERY_IDREF"`
	FinalDeliveryIDRef *FinalDeliveryIDRef `xml:"http://www.opentrans.org/XMLSchema/2.1 FINAL_DELIVERY_IDREF"`
	DelivererIDRef     *DelivererIDRef     `xml:"http://www.opentrans.org/XMLSchema/2.1 DELIVERER_IDREF"`
}

type DeliveryIDRef struct {
	XMLName xml.Name `xml:"http://www.opentrans.org/XMLSchema/2.1 DELIVERY_IDREF"`
	Type    string   `xml:"type,attr,omitempty" validate:"min=1,max=250"`
	Value   string   `xml:",chardata" validate:"min=1,max=250"`
}

type FinalDeliveryIDRef struct {
	XMLName xml.Name `xml:"http://www.opentrans.org/XMLSchema/2.1 FINAL_DELIVERY_IDREF"`
	Type    string   `xml:"type,attr,omitempty" validate:"min=1,max=250"`
	Value   string   `xml:",chardata" validate:"min=1,max=250"`
}

type DelivererIDRef struct {
	XMLName xml.Name `xml:"http://www.opentrans.org/XMLSchema/2.1 DELIVERER_IDREF"`
	Type    string   `xml:"type,attr,omitempty" validate:"min=1,max=250"`
	Value   string   `xml:",chardata" validate:"min=1,max=250"`
}

type RemarkType string

const (
	RemarkDeliveryNote           RemarkType = "deliverynote"
	RemarkDispatchNotification              = "dispatchnotification"
	RemarkGeneral                           = "general"
	RemarkInvoice                           = "invoice"
	RemarkOrder                             = "order"
	RemarkOrderChange                       = "orderchange"
	RemarkOrderResponse                     = "orderresponse"
	RemarkQuotation                         = "quotation"
	RemarkReceiptacknowledgement            = "receiptacknowledgement"
	RemarkRemittanceAdvice                  = "remittanceadvice"
	RemarkInvoiceList                       = "invoicelist"
	RemarkRFQ                               = "rfq"
	RemarkTransport                         = "transport"
)

type Remarks struct {
	XMLName xml.Name   `xml:"http://www.opentrans.org/XMLSchema/2.1 REMARKS"`
	Type    RemarkType `xml:"type,attr,omitempty" validate:"min=1,max=250"`
	Value   string     `xml:",chardata" validate:"min=1,max=64000"`
}

const UserDefinedExtensionPrefix = "UDX."

type HeaderUserDefinedExtension struct {
	XMLName  xml.Name `xml:"http://www.opentrans.org/XMLSchema/2.1 HEADER_UDX" json:"-" yaml:"-"`
	Elements map[string][]byte
}

func (h HeaderUserDefinedExtension) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range h.Elements {
		err = e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: UserDefinedExtensionPrefix + strings.ToUpper(k)}})
		if err != nil {
			return err
		}
	}

	err = e.EncodeToken(xml.EndElement{Name: start.Name})
	if err != nil {
		return err
	}

	return nil
}

func (h *HeaderUserDefinedExtension) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	h.Elements = map[string][]byte{}

	key := ""

	for {
		t, _ := d.Token()
		switch tt := t.(type) {

		case xml.StartElement:
			if strings.HasPrefix(tt.Name.Local, UserDefinedExtensionPrefix) {
				key = tt.Name.Local[len(UserDefinedExtensionPrefix):]
			}
		case xml.EndElement:
			if tt.Name == start.Name {
				return nil
			}

			key = ""
		case xml.CharData:
			if key != "" {
				h.Elements[key] = tt.Copy()
			}
		}
	}
}

type ItemUserDefinedExtension struct {
	XMLName  xml.Name `xml:"http://www.opentrans.org/XMLSchema/2.1 ITEM_UDX"`
	Elements map[string][]byte
}

func (i ItemUserDefinedExtension) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range i.Elements {
		err = e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: UserDefinedExtensionPrefix + strings.ToUpper(k)}})
		if err != nil {
			return err
		}
	}

	err = e.EncodeToken(xml.EndElement{Name: start.Name})
	if err != nil {
		return err
	}

	return nil
}

func (i *ItemUserDefinedExtension) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	i.Elements = map[string][]byte{}

	key := ""

	for {
		t, _ := d.Token()
		switch tt := t.(type) {

		case xml.StartElement:
			if strings.HasPrefix(tt.Name.Local, UserDefinedExtensionPrefix) {
				key = tt.Name.Local[len(UserDefinedExtensionPrefix):]
			}
		case xml.EndElement:
			if tt.Name == start.Name {
				return nil
			}

			key = ""
		case xml.CharData:
			if key != "" {
				i.Elements[key] = tt.Copy()
			}
		}
	}
}
