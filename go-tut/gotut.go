package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Mp3 struct {
	Mp3 string `xml:"guid"`
}

func getXML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Read body: %v", err)
	}

	return string(data), nil
}

func main() {

	xml_string, err := getXML("http://www.rtl.fr/podcast/les-grosses-tetes.xml")

	if err != nil {
		panic(err)
	}

	var mp3s Mp3
	xml.Unmarshal(xml_string, &mp3s)

	//fmt.Println(xml)

}
