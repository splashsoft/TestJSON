package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"crypto/sha1"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName  string
	LastName   string
	//Addresses  []*Address
	Adressen	[]*Address
	Remark     string
	PrivateKey string
}

func (c VCard) String() string {
	key, _ := base64.StdEncoding.DecodeString(c.PrivateKey)
	str := fmt.Sprintf("%s %s - Bemerkung: %s - Key: '%v'", c.FirstName, c.LastName, c.Remark, key)
	str += fmt.Sprintf("\n   Bekannte Adressen ($%x): ", c.Adressen)
	if len(c.Adressen) > 0 {
		for i := range c.Adressen {
			str += fmt.Sprintf("\n   %s: %s %s", c.Adressen[i].Type, c.Adressen[i].Country, c.Adressen[i].City)
		}
	}

	return str
}

func init() {
        // For generic encoding/decoding to work, you need to register
        // types you want to use with gob at init time.

	// ändert das was? - Nö, erstmal nicht

        x := &VCard{}
        gob.Register(x)
}


func testHash( str string) {
	
	hasher := sha1.New()
	io.WriteString( hasher, str )
	
	b := []byte{}
	
	fmt.Printf("Ergebnis: %x\n", hasher.Sum( b ))
	fmt.Printf("Ergebnis: %d\n", hasher.Sum( b ))
	
	hasher.Reset()
	
	data := []byte("Das ist ein Test!")
	
	n,err := hasher.Write( data )
	if n != len(data) || err != nil {
		fmt.Printf("Hash write error; (%d/%d) / %v\n", n, len(data), err )
	}
	
	checksum := hasher.Sum(b)
	fmt.Printf("Ergebnis für '%v': %d\n", data, checksum)

}

func gobDecode( data []byte , t interface{} ) error {

			r := bytes.NewReader( data)
        	dec := gob.NewDecoder(r)
        	return dec.Decode(t)
}

func gobDecodeV2( ptrBuf *bytes.Buffer , t interface{} ) error {

        	dec := gob.NewDecoder(ptrBuf)
        	return dec.Decode(t)
}

func testJSON001() {
	
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

	vc := VCard{"Peter", "Mustermann", []*Address{a1, a2, a3, a4, a5}, "geheim!", base64.StdEncoding.EncodeToString([]byte{0x14, 0x02, 0xba, 0x52, 0xfe, 0x02, 0x00, 0x00, 0x10, 0xde, 0xad, 0xc0, 0xde})}

	vc2 := VCard{"Klaus", "Klempner", []*Address{a6}, "Normal", base64.StdEncoding.EncodeToString([]byte{0x88, 0x32, 0x12, 0xde, 0x17, 0x42, 0x01, 0x02, 0x03, 0xde, 0xad, 0xc0, 0xde})}

	vc3 := VCard{"A1", "A1", []*Address{a7, a8}, "Test", base64.StdEncoding.EncodeToString([]byte{0x88, 0x32, 0x12, 0xde, 0x17, 0x42, 0x01, 0x02, 0x03, 0xde, 0xad, 0xc0, 0xde})}
	fmt.Println( vc3)

	vc4 := VCard{"A2", "A2", []*Address{a9}, "Test", base64.StdEncoding.EncodeToString([]byte{0x88, 0x32, 0x12, 0xde, 0x17, 0x42, 0x01, 0x02, 0x03, 0xde, 0xad, 0xc0, 0xde})}
	fmt.Println( vc4)
	
	vc5 := VCard{"A3", "A3", []*Address{a10, a11}, "Test", base64.StdEncoding.EncodeToString([]byte{0x88, 0x32, 0x12, 0xde, 0x17, 0x42, 0x01, 0x02, 0x03, 0xde, 0xad, 0xc0, 0xde})}
	fmt.Println( vc5)
	
	fmt.Println(vc)

	js, _ := json.Marshal(vc5)

	strJson := string(js)
	fmt.Println("JSON:", strJson)
	// und zurück ...
	var f interface{}
	var card VCard
	fmt.Println("initial card:", card)
	
	//err := json.Unmarshal(js, &card)
	err := json.Unmarshal(js, &f)
	if err == nil {
		// stimmt so nicht ganz { hier kommt mal gleich gar nichts durch, f ist nil (ohne Fehler!), mit card geht es aber }
		// komisch, jetzt gehts, f is map[string]interface{} - nochmal durchiterieren
		//fmt.Printf("card  - %v ...\n", card) 
		fmt.Printf("f ist %T - %v ...\n", f, f) 
	} else {
		fmt.Printf("error unmarshalling JSON via f interace{} - %v ...\n", err)
	}
	

	// GOB encoding geht nur über streams
	// ich nehme mal einen byte buffer als stream ersatz
	var stream bytes.Buffer
	gobEncoder := gob.NewEncoder(&stream)

	// Encoding (Marshalling)
	err = gobEncoder.Encode(vc)
	if err != nil {
		fmt.Printf("error encoding - %v ...\n", err)
	} else {
		gobEncoder.Encode(vc2) // das steht was anderes drin, als in vc1
		//gobEncoder.Encode(card)
		
		// was kann ich denn eigentlich mit dem bytes.Buffer machen?
		// kann ich den zu einem Base64 string konvertieren?

		strGob := base64.StdEncoding.EncodeToString(stream.Bytes())
		fmt.Println("GOB als string: ", strGob)

		testHash( strGob )

		// den string wieder in einen bytes.Buffer
		var data bytes.Buffer
//		var data2 bytes.Buffer
		
		
		arrBytes, err := base64.StdEncoding.DecodeString(strGob)
		if nil == err {
			data.Write(arrBytes)

			gobDecoder := gob.NewDecoder(&data)

			// und wieder zurück
			var c VCard
			err = gobDecoder.Decode(&c)
//			var any interface{}
//			any = &data
			//err = gobDecodeV2( &data, &c)
			//err = gobDecode( data.Bytes(), &c) 
			
			if nil != err {
				fmt.Printf("error decoding - %v ...\n", err)
			} else {
				fmt.Println("zurück von GOB: ", c) // na klar, GOB geht, aber XML bekommt es nicht hin! grrrr ... :[
			
			// zwei  hintereinander geht auch ...
			// geht das auch mit interface? -> Nö, error return  ...
			//var c1 VCard
			//var any interface{}	

			//err = gobDecoder.Decode(&any) // das liefert den untenstehenden compiler error
			//err = gobDecoder.Decode(any) // das liefert nil ohne error

			// hier muss ich schon den konkreten Typ übergeben, allerdings ist innerhalb der Methode der Typ mit interface{} deklariert
			var c1 VCard
			//err = gobDecode( data.Bytes(), &c1) // das liefert die Daten, die am Anfang im Puffer stehen, also wieder die erste VCard. Wenn ich das so machen will, muss ich getrennte Puffer für jeden GOB haben 
			//err = gobDecodeV2( &data, &c1) // gob: unknown type id or corrupted data ..., evtl. ist das keine gute Idee, zwei Decoder auf denselben string anzusetzen 
			err = gobDecoder.Decode( &c1) // wenn ich mit dem einen Decoder arbeite, klappt es
			 
			if( nil == err ){
				fmt.Println("zurück von GOB: ", c1)
			// nächster Test	 
			} else {
				// error decoding 2. gob über any - gob: local interface type *interface {} can only be decoded from remote interface type; received concrete type VCard = struct { FirstName string; LastName string; Addresses [] = struct { Type string; City string; Country string; }; Remark string; PrivateKey string; } ...
				// schade, sonst könnte man diverse unterschiedliche Objekte empfangen und mit type switch unterscheiden ...
				fmt.Printf("error decoding 2. gob über any - %v ...\n", err )
			}
				
			}

		} else {
			fmt.Printf("error decoding base64 - %v ...\n", err)
		}
	}
}


func main() {

	testJSON001();
	//testXML001()
}