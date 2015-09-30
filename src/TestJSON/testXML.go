package main

import (
	//"bytes"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	//"io"
)




func testXML001() {

	a1 := &Address{"Privat", "Schnüffel Allee 43", "Knatterdorf"}
	a2 := &Address{"Arbeit", "Industriestr. 1", "Bad Bommel"}
	a3 := &Address{"HomeOffice", "hinter der Hecke", "Waldrand"}
	a4 := &Address{"Urlaub", "2. Strand rechts", "Malediven"}
	a5 := &Address{"Test1", "Stadt Test1", "Land Test1"}
	a6 := &Address{"Test2", "Stadt Test2", "Land Test2"}
	a7 := &Address{"Test3", "Stadt Test3", "Land Test3"}
	a8 := &Address{"Test4", "Stadt Test4", "Land Test4"}
	a9 := &Address{"Test5", "Stadt Test5", "Land Test5"}
	a10 := &Address{"Test6", "Stadt Test6", "Land Test6"}
	a11 := &Address{"Test7", "Stadt Test7", "Land Test7"}

	vc := VCard{"Peter", "Mustermann", []*Address{a1, a2, a3, a4, a5}, "(vc)", base64.StdEncoding.EncodeToString([]byte{0x14, 0x02, 0xba, 0x52, 0xfe, 0x02, 0x00, 0x00, 0x10, 0xde, 0xad, 0xc0, 0xde})}
	fmt.Println(vc)

	vc2 := VCard{"Klaus", "Klempner", []*Address{a6}, "(vc2)", base64.StdEncoding.EncodeToString([]byte{0x88, 0x32, 0x12, 0xde, 0x17, 0x42, 0x01, 0x02, 0x03, 0xde, 0xad, 0xc0, 0xde})}
	fmt.Println(vc2)

	vc3 := VCard{"A1", "A1", []*Address{a7, a8}, "(vc3)", base64.StdEncoding.EncodeToString([]byte{0x88, 0x32, 0x12, 0xde, 0x17, 0x42, 0x01, 0x02, 0x03, 0xde, 0xad, 0xc0, 0xde})}
	fmt.Println( vc3)

	vc4 := VCard{"A2", "A2", []*Address{a9}, "(vc4)", base64.StdEncoding.EncodeToString([]byte{0x88, 0x32, 0x12, 0xde, 0x17, 0x42, 0x01, 0x02, 0x03, 0xde, 0xad, 0xc0, 0xde})}
	fmt.Println( vc4)
	
	vc5 := VCard{"A3", "A3", []*Address{a10, a11}, "(vc5)", base64.StdEncoding.EncodeToString([]byte{0x88, 0x32, 0x12, 0xde, 0x17, 0x42, 0x01, 0x02, 0x03, 0xde, 0xad, 0xc0, 0xde})}
	fmt.Println( vc5)
	

	var card VCard
		
	bytesXML, err := xml.Marshal(vc)
	if nil == err {
		fmt.Println("XML:", string(bytesXML))

		// und zurück ...
		// wenn ich den Container wiederverwende, muss ich den auch alleine aufräumen, der Unmarshaller hängt nur an, bzw. überschreibt vorhandene Inhalte.
		card.Adressen = card.Adressen[0:0]
		err := xml.Unmarshal( bytesXML, &card)
		if nil == err {
			fmt.Println(card) // oops, hier stehen die Adressen doppelt drin?!? Noch schlimmer, hier stehen die vier Adressen aus dem eigentlichen 
		}
	}

	strTestXML := "<VCard><FirstName>Peter</FirstName><LastName>Lustig</LastName><Adressen><Type>Arbeit</Type><City>Bauarbeiterwagen</City><Country>Hinterm Hof</Country></Adressen><Remark>Moderator</Remark><PrivateKey>iDIS3hdCAQIDLMAOROFL==</PrivateKey></VCard>"
	
	if nil == err {
		fmt.Println("XML:", strTestXML)

		// und zurück ...
		// wenn ich den Container wiederverwende, muss ich den auch alleine aufräumen, der Unmarshaller hängt nur an, bzw. überschreibt vorhandene Inhalte.
		card.Adressen = card.Adressen[0:0]
		err := xml.Unmarshal( []byte(strTestXML), &card)
		if nil == err {
			fmt.Println(card) // oops, hier stehen die Adressen doppelt drin?!? Noch schlimmer, hier stehen die vier Adressen aus dem ersten VCard auch drin
							// ist das ein Referenz-Problem mit dem Adressen slice der Struktur?!? Wo sollen die herkommen?!? 
		}
	}
	
	strTestXML = "<VCard><FirstName>B1</FirstName><LastName>B1</LastName><Adressen><Type>B1</Type><City>B1</City><Country>Country</Country></Adressen><Remark>statisch String</Remark><PrivateKey>iDIS3hdCAQIDLMAOROFL==</PrivateKey></VCard>"
	
	if nil == err {
		fmt.Println("XML:", strTestXML)

		// und zurück ...
		// wenn ich den Container wiederverwende, muss ich den auch alleine aufräumen, der Unmarshaller hängt nur an, bzw. überschreibt vorhandene Inhalte.
		card.Adressen = card.Adressen[0:0]

		err := xml.Unmarshal( []byte(strTestXML), &card)
		if nil == err {
			fmt.Println(card) // oops, hier stehen die Adressen doppelt drin?!? Noch schlimmer, hier stehen die vier Adressen aus dem ersten VCard auch drin
							// ist das ein Referenz-Problem mit dem Adressen slice der Struktur?!? Wo sollen die herkommen?!? 
		}
	}
	
	strTestXML = "<VCard><FirstName>B2</FirstName><LastName>B2</LastName><Adressen><Type>B2</Type><City>B2</City><Country>Country</Country></Adressen><Remark>statisch String</Remark><PrivateKey>iDIS3hdCAQIDLMAOROFL==</PrivateKey></VCard>"
	
	if nil == err {
		fmt.Println("XML:", strTestXML)

		// und zurück ...
		
		// Depp!!!! der Unmarshaller kann auch nicht zaubern, wenn ich ihm ein Objekt übergebe, dann füllt er nur die Felder, die er identifizieren kann, mit den Daten aus dem XML-String.
		// Er initialisiert das Objekt aber nicht! Wenn da also schon was drinsteht, wie ja hier der Fall in dem Adressen-Slice, werden die neu gelesenen Daten einfach nur angehängt.
		// Ich muss schon alleine dafür sorgen, dem Unmarshaller ein leeres(!) Objekt zu geben.
		card.Adressen = card.Adressen[0:0]
		 
		err := xml.Unmarshal( []byte(strTestXML), &card)
		if nil == err {
			fmt.Println(card) // oops, hier stehen die Adressen doppelt drin?!? Noch schlimmer, hier stehen die vier Adressen aus dem ersten VCard auch drin
							// ist das ein Referenz-Problem mit dem Adressen slice der Struktur?!? Wo sollen die herkommen?!? 
		}
	}

}