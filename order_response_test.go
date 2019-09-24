package opentrans

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderResponse(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		responseXML := `<?xml version="1.0" encoding="UTF-8"?>
<ORDERRESPONSE
    xmlns="http://www.opentrans.org/XMLSchema/2.1"
    version="2.1"
    xmlns:bmecat="http://www.bmecat.org/bmecat/2005"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://www.opentrans.org/XMLSchema/2.1 http://www.opentrans.org/XMLSchema/2.1/opentrans_2_1.xsd">
    <ORDERRESPONSE_HEADER>
        <CONTROL_INFO>
            <GENERATOR_INFO>Generator 1</GENERATOR_INFO>
            <GENERATION_DATE>2019-09-24T09:02+01:00</GENERATION_DATE>
        </CONTROL_INFO>
        <ORDERRESPONSE_INFO>
            <ORDER_ID>1552292</ORDER_ID>
            <ORDERRESPONSE_DATE>2019-09-24T09:00:26</ORDERRESPONSE_DATE>
            <ORDER_DATE>2019-09-24T08:58:41</ORDER_DATE>
            <SUPPLIER_ORDER_ID>64527411</SUPPLIER_ORDER_ID>
            <PARTIES>
                <PARTY>
                    <bmecat:PARTY_ID type="iln">4051047000009</bmecat:PARTY_ID>
                    <bmecat:PARTY_ID type="supplier_specific">Customer2</bmecat:PARTY_ID>
                    <PARTY_ROLE>document_creator</PARTY_ROLE>
                    <PARTY_ROLE>supplier</PARTY_ROLE>
                    <ADDRESS>
                        <bmecat:NAME>Some GmbH</bmecat:NAME>
                        <bmecat:STREET>a street 123</bmecat:STREET>
                        <bmecat:ZIP>Z-12345</bmecat:ZIP>
                        <bmecat:CITY>asdasd</bmecat:CITY>
                        <bmecat:COUNTRY>Germany</bmecat:COUNTRY>
                        <bmecat:COUNTRY_CODED>DE</bmecat:COUNTRY_CODED>
                        <bmecat:VAT_ID>DE123456789</bmecat:VAT_ID>
                        <bmecat:PHONE type="office">0</bmecat:PHONE>
                        <bmecat:URL>https://example.com/</bmecat:URL>
                    </ADDRESS>
                </PARTY>
                <PARTY>
                    <bmecat:PARTY_ID type="supplier_specific">DE44457715</bmecat:PARTY_ID>
                    <PARTY_ROLE>buyer</PARTY_ROLE>
                    <ADDRESS>
                        <bmecat:NAME>abc</bmecat:NAME>
                        <bmecat:STREET>def 123</bmecat:STREET>
                        <bmecat:ZIP>12345</bmecat:ZIP>
                        <bmecat:CITY>AAAA</bmecat:CITY>
                        <bmecat:COUNTRY>Germany</bmecat:COUNTRY>
                        <bmecat:COUNTRY_CODED>DE</bmecat:COUNTRY_CODED>
                        <bmecat:VAT_ID>DE123456788</bmecat:VAT_ID>
                    </ADDRESS>
                </PARTY>
                <PARTY>
                    <bmecat:PARTY_ID type="supplier_specific">ShipAddress</bmecat:PARTY_ID>
                    <PARTY_ROLE>delivery</PARTY_ROLE>
                    <PARTY_ROLE>final_delivery</PARTY_ROLE>
                    <ADDRESS>
                        <bmecat:NAME>asd</bmecat:NAME>
                        <bmecat:STREET>dddd</bmecat:STREET>
                        <bmecat:ZIP>55555</bmecat:ZIP>
                        <bmecat:CITY>Z City</bmecat:CITY>
                    </ADDRESS>
                </PARTY>
                <PARTY>
                    <bmecat:PARTY_ID type="supplier_specific">DEM641</bmecat:PARTY_ID>
                    <PARTY_ROLE>manufacturer</PARTY_ROLE>
                    <ADDRESS>
                        <bmecat:NAME>HP INC.</bmecat:NAME>
                        <bmecat:NAME2>HP - COMM MOBILE ACCESSORIES (MP)</bmecat:NAME2>
                    </ADDRESS>
                </PARTY>
            </PARTIES>
            <ORDER_PARTIES_REFERENCE>
                <bmecat:BUYER_IDREF type="supplier_specific">Customer1</bmecat:BUYER_IDREF>
                <bmecat:SUPPLIER_IDREF type="supplier_specific">Customer2</bmecat:SUPPLIER_IDREF>
                <SHIPMENT_PARTIES_REFERENCE>
                    <DELIVERY_IDREF type="supplier_specific">ShipAddress</DELIVERY_IDREF>
                    <FINAL_DELIVERY_IDREF type="supplier_specific">ShipAddress</FINAL_DELIVERY_IDREF>
                </SHIPMENT_PARTIES_REFERENCE>
            </ORDER_PARTIES_REFERENCE>
            <bmecat:CURRENCY>EUR</bmecat:CURRENCY>
            <HEADER_UDX>
                <UDX.A.B.C>aNOrdeR</UDX.A.B.C>
            </HEADER_UDX>
        </ORDERRESPONSE_INFO>
    </ORDERRESPONSE_HEADER>
    <ORDERRESPONSE_ITEM_LIST>
        <ORDERRESPONSE_ITEM>
            <LINE_ITEM_ID>1</LINE_ITEM_ID>
            <PRODUCT_ID>
                <bmecat:SUPPLIER_PID type="supplier_specific">asdasd</bmecat:SUPPLIER_PID>
                <bmecat:SUPPLIER_IDREF type="supplier_specific">Customer2</bmecat:SUPPLIER_IDREF>
                <bmecat:INTERNATIONAL_PID type="EAN">4514953933439</bmecat:INTERNATIONAL_PID>
                <bmecat:DESCRIPTION_SHORT>HP USB 3.0 TO GIGABIT ADAPTER</bmecat:DESCRIPTION_SHORT>
                <bmecat:DESCRIPTION_LONG>F/ DEDICATED HP NOTEBOOKS</bmecat:DESCRIPTION_LONG>
                <MANUFACTURER_INFO>
                    <bmecat:MANUFACTURER_IDREF type="supplier_specific">DEM641</bmecat:MANUFACTURER_IDREF>
                    <bmecat:MANUFACTURER_PID>N7P47AA#AC3</bmecat:MANUFACTURER_PID>
                </MANUFACTURER_INFO>
            </PRODUCT_ID>
            <QUANTITY>4</QUANTITY>
            <bmecat:ORDER_UNIT>C62</bmecat:ORDER_UNIT>
            <PRODUCT_PRICE_FIX>
                <bmecat:PRICE_AMOUNT>20.39</bmecat:PRICE_AMOUNT>
                <bmecat:PRICE_FLAG type="incl_duty">false</bmecat:PRICE_FLAG>
                <bmecat:PRICE_FLAG type="incl_freight">false</bmecat:PRICE_FLAG>
                <bmecat:PRICE_FLAG type="incl_insurance">true</bmecat:PRICE_FLAG>
                <bmecat:PRICE_FLAG type="incl_packing">true</bmecat:PRICE_FLAG>
                <bmecat:PRICE_QUANTITY>1</bmecat:PRICE_QUANTITY>
            </PRODUCT_PRICE_FIX>
            <PRICE_LINE_AMOUNT>99999.91</PRICE_LINE_AMOUNT>
            <ITEM_UDX>
                <UDX.B.C>123</UDX.B.C>
            </ITEM_UDX>
        </ORDERRESPONSE_ITEM>
    </ORDERRESPONSE_ITEM_LIST>
    <ORDERRESPONSE_SUMMARY>
        <TOTAL_ITEM_NUM>1</TOTAL_ITEM_NUM>
        <TOTAL_AMOUNT>199999.91</TOTAL_AMOUNT>
    </ORDERRESPONSE_SUMMARY>
</ORDERRESPONSE>`

		var response OrderResponse
		err := xml.Unmarshal([]byte(responseXML), &response)
		assert.NoError(t, err)
		assert.Equal(t, response.Version, "2.1", "version")
		assert.Equal(t, response.OrderResponseHeader.OrderResponseInfo.OrderID, "1552292", "order id")
		assert.Equal(t, response.OrderResponseHeader.OrderResponseInfo.SupplierOrderID, "64527411", "supplier order id")
		assert.Equal(t, len(response.OrderResponseHeader.OrderResponseInfo.Parties.Party), 4, "parties")
		assert.Equal(t, string(response.OrderResponseHeader.OrderResponseInfo.HeaderUserDefinedExtension.Elements["A.B.C"]), "aNOrdeR", "udx A.B.C")
	})
}
