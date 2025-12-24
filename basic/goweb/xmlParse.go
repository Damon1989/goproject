package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

type Servers struct {
	XMLName xml.Name         `xml:"servers"`
	Version string           `xml:"version,attr"`
	Svs     []generateServer `xml:"server"`
}

type generateServer struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func loadXmlParse() {
	file, err := os.Open("servers.xml")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	v := Recurlyservers{}

	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(v)
	fmt.Println("+++++++++++++===========================")
	fmt.Println(v.XMLName)
	fmt.Println(v.Version)
	fmt.Println(v.Svs)
	fmt.Println(v.Svs[0])
	fmt.Println(v.Svs[0].ServerName)
	fmt.Println(v.Svs[0].ServerIP)

	fmt.Println("-------------------------")

	for _, s := range v.Svs {
		fmt.Println(s.ServerName)
		fmt.Println(s.ServerIP)
	}
}

func generateXml() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, generateServer{"Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, generateServer{"Beijing_VPN", "127.0.0.2"})
	xml.MarshalIndent(v, "", "  ")
}
