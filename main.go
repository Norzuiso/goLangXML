package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Comprobante struct {
	PayMethod string       `xml:"FormaPago,attr"`
	Emitor    EmisorStruct `xml:"Emisor"`
}

type EmisorStruct struct {
	Rfc string `xml:"Rfc,attr"`
}

type ReceptorStruct struct {
	Rfc string `xml:"Rfc,attr"`
}

type ComplementoStruct struct {
	TimbreFiscalDigital string
}

type Concepto struct {
	Rfc       string      `xml:"ClaveProdServ,attr"`
	Impuestos []Translado `xml:"Impuestos"`
}

type Translado struct {
	Base string `xml:"Base,attr"`
}

func main() {
	data, _ := ioutil.ReadFile("factura.xml")

	content := []byte(data)

	e := new(Comprobante)
	xml.Unmarshal(content, e)

	fmt.Println(1)
	fmt.Println(e)
	fmt.Println(2)

}
