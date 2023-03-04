package burp

import (
	b64 "encoding/base64"
	"encoding/xml"
	"os"
)

type Run struct {
	Items []Item `xml:"item" json:"item"`
}

type Item struct {
	Time           string   `xml:"time" json:"time"`
	Url            string   `xml:"url" json:"url"`
	Host           string   `xml:"host" json:"host"`
	Port           int      `xml:"port" json:"port"`
	Protocol       string   `xml:"protocol" json:"protocol"`
	Method         string   `xml:"method" json:"method"`
	Path           string   `xml:"path" json:"path"`
	Ext            string   `xml:"extension" json:"extension"`
	Request        Request  `xml:"request" json:"request"`
	Status         int      `xml:"status" json:"status"`
	ResponseLength int      `xml:"responselength" json:"responselength"`
	MimeType       string   `xml:"mimetype" json:"mimetype"`
	Response       Response `xml:"response" json:"response"`
	Comment        string   `xml:"comment" json:"comment"`
}

type Request struct {
	Base64 bool   `xml:"base64,attr"`
	Data   string `xml:",chardata"`
}

type Response struct {
	Base64 bool   `xml:"base64,attr"`
	Data   string `xml:",chardata"`
}

func (item *Item) GetRequest() string {
	return item.getBase64DecodedIfRequired(item.Request.Data, item.Request.Base64)
}

func (item *Item) GetResponse() string {
	return item.getBase64DecodedIfRequired(item.Response.Data, item.Response.Base64)
}

func (item *Item) getBase64DecodedIfRequired(content string, base64Encoded bool) string {
	if base64Encoded {
		sDec, _ := b64.StdEncoding.DecodeString(content)
		return string(sDec)
	}

	return content
}

func Parse(content []byte) (*Run, error) {
	r := &Run{}
	err := xml.Unmarshal(content, r)
	return r, err
}

func ParseFromFile(fileName string) (*Run, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return Parse(data)
}
