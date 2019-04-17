package opentrans

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.com/mclgmbh/gomod/bmecat"
)

func TestNewOrder(t *testing.T) {
	d := Datetime{Time: time.Date(2018, time.October, 1, 18, 54, 55, 0, time.UTC)}

	t.Run("Test1", func(t *testing.T) {
		o1 := NewOrder(OrderTypeStandard)
		o1.OrderHeader.ControlInfo = &ControlInfo{
			GeneratorInfo:  "test",
			GenerationDate: d.Format(datetimeFormat),
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
								URL: "www.example.tld",
								EMail: []bmecat.EMail{
									"foo2@bar.baz",
								},
							},
						},
					},
				},
			},
			OrderPartiesReference: OrderPartiesReference{
				BuyerIDRef: &bmecat.BuyerIDRef{
					PartyID: bmecat.PartyID{
						Type:  bmecat.PartyTypeSupplierSpecific,
						Value: "123456",
					},
				},
				SupplierIDRef: &bmecat.SupplierIDRef{
					PartyID: bmecat.PartyID{
						Type:  bmecat.PartyTypeBuyerSpecific,
						Value: "102315123",
					},
				},
			},
		}

		s, err := xml.MarshalIndent(o1, "", "  ")
		if err != nil {
			t.Error(err.Error())
			return
		}

		assert.Equalf(t, `<ORDER xmlns="http://www.opentrans.org/XMLSchema/2.1" xmlns:bmecat="http://www.bmecat.org/bmecat/2005" version="2.1" type="standard">
  <ORDER_HEADER>
    <CONTROL_INFO>
      <GENERATOR_INFO>test</GENERATOR_INFO>
      <GENERATION_DATE>2018-10-01T18:54:55Z</GENERATION_DATE>
    </CONTROL_INFO>
    <ORDER_INFO>
      <ORDER_ID>ORDER_123456</ORDER_ID>
      <ORDER_DATE>2018-10-01T18:54:55Z</ORDER_DATE>
      <bmecat:LANGUAGE default="true">deu</bmecat:LANGUAGE>
      <PARTIES>
        <PARTY>
          <bmecat:PARTY_ID type="buyer_specific">1234569</bmecat:PARTY_ID>
          <PARTY_ROLE>supplier</PARTY_ROLE>
          <ADDRESS>
            <bmecat:NAME>XYZ Name</bmecat:NAME>
            <bmecat:STREET>Some sort of street 26</bmecat:STREET>
            <bmecat:ZIP>ZIP 123</bmecat:ZIP>
            <bmecat:CITY>Awesomo</bmecat:CITY>
            <bmecat:COUNTRY>Deutschland</bmecat:COUNTRY>
            <bmecat:COUNTRY_CODED>DE</bmecat:COUNTRY_CODED>
          </ADDRESS>
          <ACCOUNT></ACCOUNT>
          <MIME_INFO></MIME_INFO>
        </PARTY>
        <PARTY>
          <bmecat:PARTY_ID type="supplier_specific">25874852</bmecat:PARTY_ID>
          <PARTY_ROLE>buyer</PARTY_ROLE>
          <ADDRESS>
            <bmecat:NAME>ASD</bmecat:NAME>
            <CONTACT_DETAILS>
              <bmecat:CONTACT_NAME>Bin</bmecat:CONTACT_NAME>
              <bmecat:FIRST_NAME>Foo</bmecat:FIRST_NAME>
              <bmecat:PHONE type="office">+49 487 1474447448</bmecat:PHONE>
              <bmecat:EMAILS>
                <bmecat:EMAIL>foo@bar.baz</bmecat:EMAIL>
              </bmecat:EMAILS>
            </CONTACT_DETAILS>
            <bmecat:STREET>Another street 12</bmecat:STREET>
            <bmecat:ZIP>45123</bmecat:ZIP>
            <bmecat:CITY>Gotham</bmecat:CITY>
            <bmecat:COUNTRY>Deutschland</bmecat:COUNTRY>
            <bmecat:COUNTRY_CODED>DE</bmecat:COUNTRY_CODED>
            <bmecat:PHONE type="office">+49 148/147854788</bmecat:PHONE>
            <bmecat:FAX type="office">+49 148/147854789</bmecat:FAX>
            <bmecat:URL>www.example.tld</bmecat:URL>
          </ADDRESS>
          <ACCOUNT></ACCOUNT>
          <MIME_INFO></MIME_INFO>
        </PARTY>
        <PARTY>
          <bmecat:PARTY_ID type="supplier_specific">25874852</bmecat:PARTY_ID>
          <PARTY_ROLE>delivery</PARTY_ROLE>
          <ADDRESS>
            <bmecat:NAME>ASD</bmecat:NAME>
            <CONTACT_DETAILS>
              <bmecat:CONTACT_NAME>Bin</bmecat:CONTACT_NAME>
              <bmecat:FIRST_NAME>Foo</bmecat:FIRST_NAME>
              <bmecat:PHONE type="office">+49 487 1474447448</bmecat:PHONE>
              <bmecat:EMAILS>
                <bmecat:EMAIL>foo@bar.baz</bmecat:EMAIL>
              </bmecat:EMAILS>
            </CONTACT_DETAILS>
            <bmecat:STREET>Another street 12</bmecat:STREET>
            <bmecat:ZIP>45123</bmecat:ZIP>
            <bmecat:CITY>Gotham</bmecat:CITY>
            <bmecat:COUNTRY>Deutschland</bmecat:COUNTRY>
            <bmecat:COUNTRY_CODED>DE</bmecat:COUNTRY_CODED>
            <bmecat:PHONE type="office">+49 148/147854788</bmecat:PHONE>
            <bmecat:FAX type="office">+49 148/147854789</bmecat:FAX>
            <bmecat:EMAIL>foo2@bar.baz</bmecat:EMAIL>
            <bmecat:URL>www.example.tld</bmecat:URL>
          </ADDRESS>
          <ACCOUNT></ACCOUNT>
          <MIME_INFO></MIME_INFO>
        </PARTY>
      </PARTIES>
      <ORDER_PARTIES_REFERENCE>
        <bmecat:BUYER_IDREF type="supplier_specific">123456</bmecat:BUYER_IDREF>
        <bmecat:SUPPLIER_IDREF type="buyer_specific">102315123</bmecat:SUPPLIER_IDREF>
      </ORDER_PARTIES_REFERENCE>
    </ORDER_INFO>
  </ORDER_HEADER>
  <ORDER_ITEM_LIST></ORDER_ITEM_LIST>
  <ORDER_SUMMARY></ORDER_SUMMARY>
</ORDER>`, string(s), "")
	})
}
