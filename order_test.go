package opentrans

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.com/mclgmbh/golang-pkg/bmecat"
)

func TestNewOrder(t *testing.T) {
	d := bmecat.Datetime{Time: time.Date(2018, time.October, 1, 18, 54, 55, 0, time.UTC)}

	t.Run("Test1", func(t *testing.T) {
		o1 := NewOrder(OrderTypeStandard)
		o1.OrderHeader.ControlInfo = &ControlInfo{
			GeneratorInfo:  "test",
			GenerationDate: d,
		}
		o1.OrderHeader.OrderInfo = &OrderInfo{
			OrderID:   "ORDER_123456",
			OrderDate: d,
			Language: []bmecat.Language{
				{
					Default: true,
					Value:   "deu",
				},
			},
			Parties: Parties{
				Party: []Party{
					{
						PartyID: []bmecat.PartyID{
							{
								Type:  bmecat.PartyTypeBuyerSpecific,
								Value: "1234569",
							},
						},
						PartyRole: []PartyRole{
							PartyRoleSupplier,
						},
						Address: []Address{
							{
								Name: []bmecat.Name{
									{
										Value: "XYZ Name",
									},
								},
								Street: []bmecat.Street{
									{
										Value: "Some sort of street 26",
									},
								},
								Zip: []bmecat.Zip{
									{
										Value: "ZIP 123",
									},
								},
								City: []bmecat.City{
									{
										Value: "Awesomo",
									},
								},
								Country: []bmecat.Country{
									{
										Value: "Deutschland",
									},
								},
								CountryCoded: "DE",
							},
						},
					},
					{
						PartyID: []bmecat.PartyID{
							{
								Type:  bmecat.PartyTypeSupplierSpecific,
								Value: "25874852",
							},
						},
						PartyRole: []PartyRole{
							PartyRoleBuyer,
						},
						Address: []Address{
							{
								Name: []bmecat.Name{
									{
										Value: "ASD",
									},
								},
								ContactDetails: []ContactDetails{
									{
										ContactName: []bmecat.ContactName{
											{
												Value: "Bin",
											},
										},
										FirstName: []bmecat.FirstName{
											{
												Value: "Foo",
											},
										},
										Phone: []bmecat.Phone{
											{
												Type:  bmecat.PhoneOffice,
												Value: "+49 487 1474447448",
											},
										},
										EMails: &bmecat.EMails{
											EMail: []bmecat.EMail{
												"foo@bar.baz",
											},
										},
									},
								},
								Street: []bmecat.Street{
									{
										Value: "Another street 12",
									},
								},
								Zip: []bmecat.Zip{
									{
										Value: "45123",
									},
								},
								City: []bmecat.City{
									{
										Value: "Gotham",
									},
								},
								Country: []bmecat.Country{
									{
										Value: "Deutschland",
									},
								},
								CountryCoded: "DE",
								Phone: []bmecat.Phone{
									{
										Type:  bmecat.PhoneOffice,
										Value: "+49 148/147854788",
									},
								},
								Fax: []bmecat.Fax{
									{
										Type:  bmecat.FaxOffice,
										Value: "+49 148/147854789",
									},
								},
								URL: "www.example.tld",
							},
						},
					},
					{
						PartyID: []bmecat.PartyID{
							{
								Type:  bmecat.PartyTypeSupplierSpecific,
								Value: "25874852",
							},
						},
						PartyRole: []PartyRole{
							PartyRoleDelivery,
						},
						Address: []Address{
							{
								Name: []bmecat.Name{
									{
										Value: "ASD",
									},
								},
								ContactDetails: []ContactDetails{
									{
										ContactName: []bmecat.ContactName{
											{
												Value: "Bin",
											},
										},
										FirstName: []bmecat.FirstName{
											{
												Value: "Foo",
											},
										},
										Phone: []bmecat.Phone{
											{
												Type:  bmecat.PhoneOffice,
												Value: "+49 487 1474447448",
											},
										},
										EMails: &bmecat.EMails{
											EMail: []bmecat.EMail{
												"foo@bar.baz",
											},
										},
									},
								},
								Street: []bmecat.Street{
									{
										Value: "Another street 12",
									},
								},
								Zip: []bmecat.Zip{
									{
										Value: "45123",
									},
								},
								City: []bmecat.City{
									{
										Value: "Gotham",
									},
								},
								Country: []bmecat.Country{
									{
										Value: "Deutschland",
									},
								},
								CountryCoded: "DE",
								Phone: []bmecat.Phone{
									{
										Type:  bmecat.PhoneOffice,
										Value: "+49 148/147854788",
									},
								},
								Fax: []bmecat.Fax{
									{
										Type:  bmecat.FaxOffice,
										Value: "+49 148/147854789",
									},
								},
								URL:   "www.example.tld",
								EMail: "foo2@bar.baz",
							},
						},
					},
				},
			},
			OrderPartiesReference: OrderPartiesReference{
				BuyerIDRef: bmecat.BuyerIDRef{
					PartyID: bmecat.PartyID{
						Type:  bmecat.PartyTypeSupplierSpecific,
						Value: "123456",
					},
				},
				SupplierIDRef: bmecat.SupplierIDRef{
					PartyID: bmecat.PartyID{
						Type:  bmecat.PartyTypeBuyerSpecific,
						Value: "102315123",
					},
				},
			},
			Currency: "EUR",
			Remarks: []Remarks{
				{
					Type:  RemarkGeneral,
					Value: "RemarkTheFirst",
				},
			},
			HeaderUserDefinedExtension: &HeaderUserDefinedExtension{
				Elements: map[string][]byte{
					"DROPSHIPMENT": []byte("true"),
				},
			},
		}

		o1.OrderItemList.OrderItem = []OrderItem{
			{
				LineItemID: "123456",
				ProductID: ProductID{
					SupplierPID: &bmecat.SupplierPID{
						Type:  bmecat.PIDSupplierSpecific,
						Value: "5095055",
					},
					InternationalPID: &bmecat.InternationalPID{
						Value: "5702015867511",
					},
				},
				Quantity:        1.0,
				OrderUnit:       bmecat.UnitPiece,
				PriceLineAmount: 3.51,
				ProductPriceFix: &ProductPriceFix{
					PriceAmount: 3.51,
				},
			},
		}

		s, err := xml.MarshalIndent(o1, "", "  ")
		if err != nil {
			t.Error(err.Error())
			return
		}

		assert.Equalf(t, `<ORDER xmlns="http://www.opentrans.org/XMLSchema/2.1" version="2.1" type="standard">
  <ORDER_HEADER xmlns="http://www.opentrans.org/XMLSchema/2.1">
    <CONTROL_INFO xmlns="http://www.opentrans.org/XMLSchema/2.1">
      <GENERATOR_INFO xmlns="http://www.opentrans.org/XMLSchema/2.1">test</GENERATOR_INFO>
      <GENERATION_DATE xmlns="http://www.opentrans.org/XMLSchema/2.1">2018-10-01T18:54:55Z</GENERATION_DATE>
    </CONTROL_INFO>
    <ORDER_INFO xmlns="http://www.opentrans.org/XMLSchema/2.1">
      <ORDER_ID xmlns="http://www.opentrans.org/XMLSchema/2.1">ORDER_123456</ORDER_ID>
      <ORDER_DATE xmlns="http://www.opentrans.org/XMLSchema/2.1">2018-10-01T18:54:55Z</ORDER_DATE>
      <LANGUAGE xmlns="http://www.bmecat.org/bmecat/2005" default="true">deu</LANGUAGE>
      <PARTIES xmlns="http://www.opentrans.org/XMLSchema/2.1">
        <PARTY xmlns="http://www.opentrans.org/XMLSchema/2.1">
          <PARTY_ID xmlns="http://www.bmecat.org/bmecat/2005" type="buyer_specific">1234569</PARTY_ID>
          <PARTY_ROLE xmlns="http://www.opentrans.org/XMLSchema/2.1">supplier</PARTY_ROLE>
          <ADDRESS xmlns="http://www.opentrans.org/XMLSchema/2.1">
            <NAME xmlns="http://www.bmecat.org/bmecat/2005">XYZ Name</NAME>
            <STREET xmlns="http://www.bmecat.org/bmecat/2005">Some sort of street 26</STREET>
            <ZIP xmlns="http://www.bmecat.org/bmecat/2005">ZIP 123</ZIP>
            <CITY xmlns="http://www.bmecat.org/bmecat/2005">Awesomo</CITY>
            <COUNTRY xmlns="http://www.bmecat.org/bmecat/2005">Deutschland</COUNTRY>
            <COUNTRY_CODED xmlns="http://www.bmecat.org/bmecat/2005">DE</COUNTRY_CODED>
          </ADDRESS>
        </PARTY>
        <PARTY xmlns="http://www.opentrans.org/XMLSchema/2.1">
          <PARTY_ID xmlns="http://www.bmecat.org/bmecat/2005" type="supplier_specific">25874852</PARTY_ID>
          <PARTY_ROLE xmlns="http://www.opentrans.org/XMLSchema/2.1">buyer</PARTY_ROLE>
          <ADDRESS xmlns="http://www.opentrans.org/XMLSchema/2.1">
            <NAME xmlns="http://www.bmecat.org/bmecat/2005">ASD</NAME>
            <CONTACT_DETAILS xmlns="http://www.opentrans.org/XMLSchema/2.1">
              <CONTACT_NAME xmlns="http://www.bmecat.org/bmecat/2005">Bin</CONTACT_NAME>
              <FIRST_NAME xmlns="http://www.bmecat.org/bmecat/2005">Foo</FIRST_NAME>
              <PHONE xmlns="http://www.bmecat.org/bmecat/2005" type="office">+49 487 1474447448</PHONE>
              <EMAILS xmlns="http://www.bmecat.org/bmecat/2005">
                <EMAIL xmlns="http://www.bmecat.org/bmecat/2005">foo@bar.baz</EMAIL>
              </EMAILS>
            </CONTACT_DETAILS>
            <STREET xmlns="http://www.bmecat.org/bmecat/2005">Another street 12</STREET>
            <ZIP xmlns="http://www.bmecat.org/bmecat/2005">45123</ZIP>
            <CITY xmlns="http://www.bmecat.org/bmecat/2005">Gotham</CITY>
            <COUNTRY xmlns="http://www.bmecat.org/bmecat/2005">Deutschland</COUNTRY>
            <COUNTRY_CODED xmlns="http://www.bmecat.org/bmecat/2005">DE</COUNTRY_CODED>
            <PHONE xmlns="http://www.bmecat.org/bmecat/2005" type="office">+49 148/147854788</PHONE>
            <FAX xmlns="http://www.bmecat.org/bmecat/2005" type="office">+49 148/147854789</FAX>
            <URL xmlns="http://www.bmecat.org/bmecat/2005">www.example.tld</URL>
          </ADDRESS>
        </PARTY>
        <PARTY xmlns="http://www.opentrans.org/XMLSchema/2.1">
          <PARTY_ID xmlns="http://www.bmecat.org/bmecat/2005" type="supplier_specific">25874852</PARTY_ID>
          <PARTY_ROLE xmlns="http://www.opentrans.org/XMLSchema/2.1">delivery</PARTY_ROLE>
          <ADDRESS xmlns="http://www.opentrans.org/XMLSchema/2.1">
            <NAME xmlns="http://www.bmecat.org/bmecat/2005">ASD</NAME>
            <CONTACT_DETAILS xmlns="http://www.opentrans.org/XMLSchema/2.1">
              <CONTACT_NAME xmlns="http://www.bmecat.org/bmecat/2005">Bin</CONTACT_NAME>
              <FIRST_NAME xmlns="http://www.bmecat.org/bmecat/2005">Foo</FIRST_NAME>
              <PHONE xmlns="http://www.bmecat.org/bmecat/2005" type="office">+49 487 1474447448</PHONE>
              <EMAILS xmlns="http://www.bmecat.org/bmecat/2005">
                <EMAIL xmlns="http://www.bmecat.org/bmecat/2005">foo@bar.baz</EMAIL>
              </EMAILS>
            </CONTACT_DETAILS>
            <STREET xmlns="http://www.bmecat.org/bmecat/2005">Another street 12</STREET>
            <ZIP xmlns="http://www.bmecat.org/bmecat/2005">45123</ZIP>
            <CITY xmlns="http://www.bmecat.org/bmecat/2005">Gotham</CITY>
            <COUNTRY xmlns="http://www.bmecat.org/bmecat/2005">Deutschland</COUNTRY>
            <COUNTRY_CODED xmlns="http://www.bmecat.org/bmecat/2005">DE</COUNTRY_CODED>
            <PHONE xmlns="http://www.bmecat.org/bmecat/2005" type="office">+49 148/147854788</PHONE>
            <FAX xmlns="http://www.bmecat.org/bmecat/2005" type="office">+49 148/147854789</FAX>
            <EMAIL xmlns="http://www.bmecat.org/bmecat/2005">foo2@bar.baz</EMAIL>
            <URL xmlns="http://www.bmecat.org/bmecat/2005">www.example.tld</URL>
          </ADDRESS>
        </PARTY>
      </PARTIES>
      <ORDER_PARTIES_REFERENCE xmlns="http://www.opentrans.org/XMLSchema/2.1">
        <BUYER_IDREF xmlns="http://www.bmecat.org/bmecat/2005" type="supplier_specific">123456</BUYER_IDREF>
        <SUPPLIER_IDREF xmlns="http://www.bmecat.org/bmecat/2005" type="buyer_specific">102315123</SUPPLIER_IDREF>
      </ORDER_PARTIES_REFERENCE>
      <CURRENCY xmlns="http://www.bmecat.org/bmecat/2005">EUR</CURRENCY>
      <PARTIAL_SHIPMENT_ALLOWED xmlns="http://www.opentrans.org/XMLSchema/2.1">false</PARTIAL_SHIPMENT_ALLOWED>
      <REMARKS xmlns="http://www.opentrans.org/XMLSchema/2.1" type="general">RemarkTheFirst</REMARKS>
      <HEADER_UDX xmlns="http://www.opentrans.org/XMLSchema/2.1">
        <UDX.DROPSHIPMENT>true</UDX.DROPSHIPMENT>
      </HEADER_UDX>
    </ORDER_INFO>
  </ORDER_HEADER>
  <ORDER_ITEM_LIST xmlns="http://www.opentrans.org/XMLSchema/2.1">
    <ORDER_ITEM xmlns="http://www.opentrans.org/XMLSchema/2.1">
      <LINE_ITEM_ID xmlns="http://www.opentrans.org/XMLSchema/2.1">123456</LINE_ITEM_ID>
      <PRODUCT_ID xmlns="http://www.opentrans.org/XMLSchema/2.1">
        <SUPPLIER_PID xmlns="http://www.bmecat.org/bmecat/2005" type="supplier_specific">5095055</SUPPLIER_PID>
        <INTERNATIONAL_PID xmlns="http://www.bmecat.org/bmecat/2005">5702015867511</INTERNATIONAL_PID>
      </PRODUCT_ID>
      <QUANTITY xmlns="http://www.opentrans.org/XMLSchema/2.1">1</QUANTITY>
      <ORDER_UNIT xmlns="http://www.bmecat.org/bmecat/2005">C62</ORDER_UNIT>
      <PRODUCT_PRICE_FIX xmlns="http://www.opentrans.org/XMLSchema/2.1">
        <PRICE_AMOUNT xmlns="http://www.bmecat.org/bmecat/2005">3.51</PRICE_AMOUNT>
      </PRODUCT_PRICE_FIX>
      <PRICE_LINE_AMOUNT xmlns="http://www.opentrans.org/XMLSchema/2.1">3.51</PRICE_LINE_AMOUNT>
    </ORDER_ITEM>
  </ORDER_ITEM_LIST>
  <ORDER_SUMMARY xmlns="http://www.opentrans.org/XMLSchema/2.1">
    <TOTAL_ITEM_NUM xmlns="http://www.opentrans.org/XMLSchema/2.1">0</TOTAL_ITEM_NUM>
  </ORDER_SUMMARY>
</ORDER>`, string(s), "")
	})

	t.Run("Agreement", func(t *testing.T) {
		o1 := NewOrder(OrderTypeStandard)
		o1.OrderHeader.ControlInfo = &ControlInfo{
			GeneratorInfo:  "test",
			GenerationDate: d,
		}
		o1.OrderHeader.OrderInfo = &OrderInfo{
			OrderID:   "ORDER_123456",
			OrderDate: d,
			Currency:  "EUR",
		}
		o1.OrderHeader.SourcingInfo = &SourcingInfo{
			Agreement: &Agreement{
				Type:        "asd",
				AgreementID: "abced",
			},
		}

		o1.OrderItemList.OrderItem = []OrderItem{
			{
				LineItemID: "123456",
				ProductID: ProductID{
					SupplierPID: &bmecat.SupplierPID{
						Type:  bmecat.PIDSupplierSpecific,
						Value: "5095055",
					},
					InternationalPID: &bmecat.InternationalPID{
						Value: "5702015867511",
					},
				},
				Quantity:        1.0,
				OrderUnit:       bmecat.UnitPiece,
				PriceLineAmount: 3.51,
				ProductPriceFix: &ProductPriceFix{
					PriceAmount: 3.51,
				},
			},
		}

		s, err := xml.MarshalIndent(o1, "", "  ")
		if err != nil {
			t.Error(err.Error())
			return
		}

		assert.Equalf(t, `<ORDER xmlns="http://www.opentrans.org/XMLSchema/2.1" version="2.1" type="standard">
  <ORDER_HEADER xmlns="http://www.opentrans.org/XMLSchema/2.1">
    <CONTROL_INFO xmlns="http://www.opentrans.org/XMLSchema/2.1">
      <GENERATOR_INFO xmlns="http://www.opentrans.org/XMLSchema/2.1">test</GENERATOR_INFO>
      <GENERATION_DATE xmlns="http://www.opentrans.org/XMLSchema/2.1">2018-10-01T18:54:55Z</GENERATION_DATE>
    </CONTROL_INFO>
    <SOURCING_INFO xmlns="http://www.opentrans.org/XMLSchema/2.1">
      <AGREEMENT xmlns="http://www.opentrans.org/XMLSchema/2.1">
        <type>asd</type>
        <AGREEMENT_ID xmlns="http://www.bmecat.org/bmecat/2005">abced</AGREEMENT_ID>
      </AGREEMENT>
    </SOURCING_INFO>
    <ORDER_INFO xmlns="http://www.opentrans.org/XMLSchema/2.1">
      <ORDER_ID xmlns="http://www.opentrans.org/XMLSchema/2.1">ORDER_123456</ORDER_ID>
      <ORDER_DATE xmlns="http://www.opentrans.org/XMLSchema/2.1">2018-10-01T18:54:55Z</ORDER_DATE>
      <PARTIES xmlns="http://www.opentrans.org/XMLSchema/2.1"></PARTIES>
      <ORDER_PARTIES_REFERENCE xmlns="http://www.opentrans.org/XMLSchema/2.1">
        <BUYER_IDREF xmlns="http://www.bmecat.org/bmecat/2005" type=""></BUYER_IDREF>
        <SUPPLIER_IDREF xmlns="http://www.bmecat.org/bmecat/2005" type=""></SUPPLIER_IDREF>
      </ORDER_PARTIES_REFERENCE>
      <CURRENCY xmlns="http://www.bmecat.org/bmecat/2005">EUR</CURRENCY>
      <PARTIAL_SHIPMENT_ALLOWED xmlns="http://www.opentrans.org/XMLSchema/2.1">false</PARTIAL_SHIPMENT_ALLOWED>
    </ORDER_INFO>
  </ORDER_HEADER>
  <ORDER_ITEM_LIST xmlns="http://www.opentrans.org/XMLSchema/2.1">
    <ORDER_ITEM xmlns="http://www.opentrans.org/XMLSchema/2.1">
      <LINE_ITEM_ID xmlns="http://www.opentrans.org/XMLSchema/2.1">123456</LINE_ITEM_ID>
      <PRODUCT_ID xmlns="http://www.opentrans.org/XMLSchema/2.1">
        <SUPPLIER_PID xmlns="http://www.bmecat.org/bmecat/2005" type="supplier_specific">5095055</SUPPLIER_PID>
        <INTERNATIONAL_PID xmlns="http://www.bmecat.org/bmecat/2005">5702015867511</INTERNATIONAL_PID>
      </PRODUCT_ID>
      <QUANTITY xmlns="http://www.opentrans.org/XMLSchema/2.1">1</QUANTITY>
      <ORDER_UNIT xmlns="http://www.bmecat.org/bmecat/2005">C62</ORDER_UNIT>
      <PRODUCT_PRICE_FIX xmlns="http://www.opentrans.org/XMLSchema/2.1">
        <PRICE_AMOUNT xmlns="http://www.bmecat.org/bmecat/2005">3.51</PRICE_AMOUNT>
      </PRODUCT_PRICE_FIX>
      <PRICE_LINE_AMOUNT xmlns="http://www.opentrans.org/XMLSchema/2.1">3.51</PRICE_LINE_AMOUNT>
    </ORDER_ITEM>
  </ORDER_ITEM_LIST>
  <ORDER_SUMMARY xmlns="http://www.opentrans.org/XMLSchema/2.1">
    <TOTAL_ITEM_NUM xmlns="http://www.opentrans.org/XMLSchema/2.1">0</TOTAL_ITEM_NUM>
  </ORDER_SUMMARY>
</ORDER>`, string(s), "")
	})
}

func TestHeaderUserDefinedExtension(t *testing.T) {
	t.Run("MarshalXML", func(t *testing.T) {
		udx := struct {
			XMLName xml.Name `xml:"ASD"`
			A       HeaderUserDefinedExtension
		}{
			A: HeaderUserDefinedExtension{
				Elements: map[string][]byte{
					"DROPSHIPMENT": []byte("true"),
				},
			},
		}

		s, err := xml.MarshalIndent(udx, "", "  ")
		if err != nil {
			t.Error(err.Error())
			return
		}

		assert.Equalf(t, `<ASD>
  <HEADER_UDX xmlns="http://www.opentrans.org/XMLSchema/2.1">
    <UDX.DROPSHIPMENT>true</UDX.DROPSHIPMENT>
  </HEADER_UDX>
</ASD>`, string(s), "")
	})
	t.Run("UnmarshalXML", func(t *testing.T) {
		var udx struct {
			XMLName xml.Name `xml:"ASD"`
			A       HeaderUserDefinedExtension
		}

		txt := []byte(`<ASD>
  <HEADER_UDX xmlns="http://www.opentrans.org/XMLSchema/2.1">
    <UDX.DROPSHIPMENT>true</UDX.DROPSHIPMENT>
    <UDX.TEST>foo</UDX.TEST>
  </HEADER_UDX>
</ASD>`)

		err := xml.Unmarshal(txt, &udx)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(udx.A.Elements))
		val, ok := udx.A.Elements["DROPSHIPMENT"]
		assert.Equal(t, true, ok)
		assert.Equal(t, "true", string(val))
		val, ok = udx.A.Elements["TEST"]
		assert.Equal(t, true, ok)
		assert.Equal(t, "foo", string(val))
	})
}

func TestItemUserDefinedExtension(t *testing.T) {
	t.Run("MarshalXML", func(t *testing.T) {
		udx := struct {
			XMLName xml.Name `xml:"ASD"`
			A       ItemUserDefinedExtension
		}{
			A: ItemUserDefinedExtension{
				Elements: map[string][]byte{
					"DROPSHIPMENT": []byte("true"),
				},
			},
		}

		s, err := xml.MarshalIndent(udx, "", "  ")
		if err != nil {
			t.Error(err.Error())
			return
		}

		assert.Equalf(t, `<ASD>
  <ITEM_UDX xmlns="http://www.opentrans.org/XMLSchema/2.1">
    <UDX.DROPSHIPMENT>true</UDX.DROPSHIPMENT>
  </ITEM_UDX>
</ASD>`, string(s), "")
	})
	t.Run("UnmarshalXML", func(t *testing.T) {
		var udx struct {
			XMLName xml.Name `xml:"ASD"`
			A       ItemUserDefinedExtension
		}

		txt := []byte(`<ASD>
  <ITEM_UDX xmlns="http://www.opentrans.org/XMLSchema/2.1">
    <UDX.DROPSHIPMENT>true</UDX.DROPSHIPMENT>
    <UDX.TEST>foo</UDX.TEST>
  </ITEM_UDX>
</ASD>`)

		err := xml.Unmarshal(txt, &udx)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(udx.A.Elements))
		val, ok := udx.A.Elements["DROPSHIPMENT"]
		assert.Equal(t, true, ok)
		assert.Equal(t, "true", string(val))
		val, ok = udx.A.Elements["TEST"]
		assert.Equal(t, true, ok)
		assert.Equal(t, "foo", string(val))
	})
}
