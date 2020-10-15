package main

import (
	"encoding/xml"
	"encoding/json"
	"time"
	"log"
    "io/ioutil"
	"net/http"
)

type Infogempa struct {
	XMLName xml.Name `xml:"Infogempa"`
	Gempa   struct {
		Tanggal string `xml:"Tanggal"`
		Jam     string `xml:"Jam"`
		Point   struct {
			Coordinates string `xml:"coordinates"`
		} `xml:"point"`
		Lintang   string `xml:"Lintang"`
		Bujur     string `xml:"Bujur"`
		Magnitude string `xml:"Magnitude"`
		Kedalaman string `xml:"Kedalaman"`
		Symbol    string `xml:"_symbol"`
		Wilayah1  string `xml:"Wilayah1"`
		Wilayah2  string `xml:"Wilayah2"`
		Wilayah3  string `xml:"Wilayah3"`
		Wilayah4  string `xml:"Wilayah4"`
		Wilayah5  string `xml:"Wilayah5"`
		Potensi   string `xml:"Potensi"`
	} `xml:"gempa"`
}

func Xmltojson(w http.ResponseWriter, r *http.Request)  {
	
	url := "https://data.bmkg.go.id/autogempa.xml"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var infogempa Infogempa;

	xml.Unmarshal(body, &infogempa)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(infogempa)
}

func main()  {
	http.HandleFunc("/",Xmltojson)
	http.ListenAndServe(":4000", nil)
}