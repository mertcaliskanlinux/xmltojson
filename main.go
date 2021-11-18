package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//ÇALIŞANLAR JSON TİPİ
type jsonStaff struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
}

//ÇALIŞANLAR XML TİPİ
type Staff struct {
	XMLName   xml.Name `xml:"staff"`
	ID        int      `xml:"id"`
	FirstName string   `xml:"firstname"`
	LastName  string   `xml:"lastname"`
	UserName  string   `xml:"username"`
}

//DOSYA'DAKİ TÜM KULLANICILARIN DİZİSİ
type Company struct {
	XMLName xml.Name `xml:"company"`
	Staffs  []Staff  `xml:"staff"`
}

func (s Staff) String() string {
	return fmt.Sprintf("\t ID	:	%d - FirstName : %s	- LastName : %s - UserName : %s \n", s.ID, s.FirstName, s.LastName, s.UserName)
}

func main() {
	//XML DOSYASINI AÇIYORUZ
	xmlFile, err := os.Open("Example.xml")

	//XML DOSYAMIZDA BİR HATA VARMI DİYE KONTROL EDİYORUZ
	if err != nil {
		fmt.Println("Dosya Okuma Hatası:", err)
		return
	}
	//DEFER KEYWORD İLE EN SON YAPILACAK İŞLEMİ VERİYORUZ DOSYAYI KAPATIYORUZ.
	defer xmlFile.Close()

	//XML DOSYAMIZI OKUYORUZ BYTE OLARAK GELİYOR
	xmlData, _ := ioutil.ReadAll(xmlFile)

	//YERLEŞTİRME İŞLEMİ İÇİN  DEĞİŞKEN OLUŞTURUYORUZ
	var c Company
	xml.Unmarshal(xmlData, &c)
	fmt.Println(c.Staffs)

	// JSON DÖNÜŞÜMÜ
	var oneStaff jsonStaff
	var allStaffs []jsonStaff

	for _, value := range c.Staffs {
		oneStaff.ID = value.ID
		oneStaff.FirstName = value.FirstName
		oneStaff.LastName = value.LastName
		oneStaff.UserName = value.UserName

		allStaffs = append(allStaffs, oneStaff)
	}

	//JSON VERİYİ DÖNDÜRÜCEK
	//ERR İSE BİR HATA HALİNDE HATA MESAJINI DÖNDÜRCEK os.Exit(-1)ÇIKIŞ YAPICAK
	jsonData, err := json.Marshal(allStaffs)
	if err != nil {
		fmt.Println("HATA HATA HATA ", err)
		os.Exit(-1)

	}
	fmt.Println(string(jsonData))

	// DOSYA OLUŞTURMA VEYA DOSYAYA YAZMA
	jsonFile, err := os.Create("./ExamplaJson.json")

	if err != nil {
		fmt.Println("HATA HATA HATA ", err)

	}

	defer jsonFile.Close()
	jsonFile.Write(jsonData)
	jsonFile.Close()

}
