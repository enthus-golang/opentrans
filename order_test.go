package opentrans

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {
	d := time.Date(2018, time.October, 1, 18, 54, 55, 0, time.UTC)

	t.Run("Test1", func(t *testing.T) {
		o1 := NewOrder(OrderTypeStandard)
		o1.OrderHeader.ControlInfo = &OrderControlInfo{
			GeneratorInfo:  "test",
			GenerationDate: &d,
		}
		o1.OrderHeader.OrderInfo = &OrderInfo{
			OrderID:   "ORDER_123456",
			OrderDate: d,
			Language: []Language{
				{
					Default: true,
					Value:   "deu",
				},
			},
			Parties: Parties{
				Party: []Party{
					{
						PartyID: PartyID{
							Type:  PartyTypeBuyerSpecific,
							Value: "1234569",
						},
						PartyRole: PartyRoleSupplier,
						Address: []Address{
							{
								Name: []Name{
									{
										Value: "XYZ Name",
									},
								},
								Street: []Street{
									{
										Value: "Some sort of street 26",
									},
								},
								Zip: []Zip{
									{
										Value: "ZIP 123",
									},
								},
								City: []City{
									{
										Value: "Awesomo",
									},
								},
								Country: []Country{
									{
										Value: "Deutschland",
									},
								},
								CountryCoded: "DE",
							},
						},
					},
					{
						PartyID: PartyID{
							Type:  PartyTypeSupplierSpecific,
							Value: "25874852",
						},
						PartyRole: PartyRoleBuyer,
						Address: []Address{
							{
								Name: []Name{
									{
										Value: "ASD",
									},
								},
								ContactDetails: []ContactDetails{
									{
										ContactName: []ContactName{
											{
												Value: "Bin",
											},
										},
										FirstName: []FirstName{
											{
												Value: "Foo",
											},
										},
										Phone: []Phone{
											{
												Type:  PhoneOffice,
												Value: "+49 487 1474447448",
											},
										},
										EMails: &EMails{
											EMail: []string{
												"foo@bar.baz",
											},
										},
									},
								},
								Street: []Street{
									{
										Value: "Another street 12",
									},
								},
								Zip: []Zip{
									{
										Value: "45123",
									},
								},
								City: []City{
									{
										Value: "Gotham",
									},
								},
								Country: []Country{
									{
										Value: "Deutschland",
									},
								},
								CountryCoded: "DE",
								Phone: []Phone{
									{
										Type:  PhoneOffice,
										Value: "+49 148/147854788",
									},
								},
								Fax: []Fax{
									{
										Type:  FaxOffice,
										Value: "+49 148/147854789",
									},
								},
								URL: "www.example.tld",
							},
						},
					},
					{
						PartyID: PartyID{
							Type:  PartyTypeSupplierSpecific,
							Value: "25874852",
						},
						PartyRole: PartyRoleDelivery,
						Address: []Address{
							{
								Name: []Name{
									{
										Value: "ASD",
									},
								},
								ContactDetails: []ContactDetails{
									{
										ContactName: []ContactName{
											{
												Value: "Bin",
											},
										},
										FirstName: []FirstName{
											{
												Value: "Foo",
											},
										},
										Phone: []Phone{
											{
												Type:  PhoneOffice,
												Value: "+49 487 1474447448",
											},
										},
										EMails: &EMails{
											EMail: []string{
												"foo@bar.baz",
											},
										},
									},
								},
								Street: []Street{
									{
										Value: "Another street 12",
									},
								},
								Zip: []Zip{
									{
										Value: "45123",
									},
								},
								City: []City{
									{
										Value: "Gotham",
									},
								},
								Country: []Country{
									{
										Value: "Deutschland",
									},
								},
								CountryCoded: "DE",
								Phone: []Phone{
									{
										Type:  PhoneOffice,
										Value: "+49 148/147854788",
									},
								},
								Fax: []Fax{
									{
										Type:  FaxOffice,
										Value: "+49 148/147854789",
									},
								},
								URL: "www.example.tld",
							},
						},
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
            <bmecat:URL>www.example.tld</bmecat:URL>
          </ADDRESS>
          <ACCOUNT></ACCOUNT>
          <MIME_INFO></MIME_INFO>
        </PARTY>
      </PARTIES>
    </ORDER_INFO>
  </ORDER_HEADER>
  <ORDER_ITEM_LIST></ORDER_ITEM_LIST>
  <ORDER_SUMMARY></ORDER_SUMMARY>
</ORDER>`, string(s), "")
	})
}
