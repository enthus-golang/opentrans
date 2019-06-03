package opentrans

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanXMLNamespaces(t *testing.T) {
	rawXML := `<ORDER xmlns="http://www.opentrans.org/XMLSchema/2.1" version="2.1" type="standard">
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
          <ACCOUNT xmlns="http://www.opentrans.org/XMLSchema/2.1"></ACCOUNT>
          <MIME_INFO xmlns="http://www.opentrans.org/XMLSchema/2.1"></MIME_INFO>
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
          <ACCOUNT xmlns="http://www.opentrans.org/XMLSchema/2.1"></ACCOUNT>
          <MIME_INFO xmlns="http://www.opentrans.org/XMLSchema/2.1"></MIME_INFO>
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
          <ACCOUNT xmlns="http://www.opentrans.org/XMLSchema/2.1"></ACCOUNT>
          <MIME_INFO xmlns="http://www.opentrans.org/XMLSchema/2.1"></MIME_INFO>
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
      <PRICE_LINE_AMOUNT xmlns="http://www.opentrans.org/XMLSchema/2.1">3.51</PRICE_LINE_AMOUNT>
      <PRODUCT_PRICE_FIX xmlns="http://www.opentrans.org/XMLSchema/2.1">
        <PRICE_AMOUNT xmlns="http://www.bmecat.org/bmecat/2005">3.51</PRICE_AMOUNT>
      </PRODUCT_PRICE_FIX>
    </ORDER_ITEM>
  </ORDER_ITEM_LIST>
  <ORDER_SUMMARY xmlns="http://www.opentrans.org/XMLSchema/2.1"></ORDER_SUMMARY>
</ORDER>`
	formatted, err := CleanXMLNamespaces(rawXML)
	assert.NoError(t, err)
	assert.Equal(t, `<ORDER xmlns:bmecat="http://www.bmecat.org/bmecat/2005" xmlns="http://www.opentrans.org/XMLSchema/2.1" version="2.1" type="standard">
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
      <bmecat:CURRENCY>EUR</bmecat:CURRENCY>
      <PARTIAL_SHIPMENT_ALLOWED>false</PARTIAL_SHIPMENT_ALLOWED>
      <REMARKS type="general">RemarkTheFirst</REMARKS>
      <HEADER_UDX>
        <UDX.DROPSHIPMENT>true</UDX.DROPSHIPMENT>
      </HEADER_UDX>
    </ORDER_INFO>
  </ORDER_HEADER>
  <ORDER_ITEM_LIST>
    <ORDER_ITEM>
      <LINE_ITEM_ID>123456</LINE_ITEM_ID>
      <PRODUCT_ID>
        <bmecat:SUPPLIER_PID type="supplier_specific">5095055</bmecat:SUPPLIER_PID>
        <bmecat:INTERNATIONAL_PID>5702015867511</bmecat:INTERNATIONAL_PID>
      </PRODUCT_ID>
      <QUANTITY>1</QUANTITY>
      <bmecat:ORDER_UNIT>C62</bmecat:ORDER_UNIT>
      <PRICE_LINE_AMOUNT>3.51</PRICE_LINE_AMOUNT>
      <PRODUCT_PRICE_FIX>
        <bmecat:PRICE_AMOUNT>3.51</bmecat:PRICE_AMOUNT>
      </PRODUCT_PRICE_FIX>
    </ORDER_ITEM>
  </ORDER_ITEM_LIST>
  <ORDER_SUMMARY></ORDER_SUMMARY>
</ORDER>`, formatted)
}
