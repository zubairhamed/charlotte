package td

import (
	"encoding/json"
	"io/ioutil"
)

// Models
type Thing struct {
	Context      interface{}    `json:"@context"`
	Type         string         `json:"@type"`
	Name         string         `json:"name"`
	Base         string         `json:"base"`
	Endpoints    []*Endpoint    `json:"endpoint"`
	Security     *Security      `json:"security"`
	Interactions []*Interaction `json:"interactions"`
}

type Endpoint struct {
	Uri       string `json:"uri"`
	MediaType string `json:"mediaType"`
}

type Security struct {
	Cat       string `json:"cat"`
	Algorithm string `json:"alg"`
	As        string `json:"as"`
}

type Interaction struct {
	Id         string      `json:"@id"`
	Types      []string    `json:"@type"`
	Unit       string      `json:"unit"`
	Reference  string      `json:"reference"`
	Name       string      `json:"name"`
	ValueType  interface{} `json:"valueType"`
	OutputType interface{} `json:"outputType"`
	Writeable  bool        `json:"writeable"`
	Hrefs      []string    `json:"hrefs"`
	Links      []*Link     `json:"links"`
}

type Link struct {
	Href      string `json:"href"`
	MediaType string `json:"mediaType"`
}

func LoadFile(f string) (td *Thing, err error) {

	content, err := ioutil.ReadFile(f)

	err = json.Unmarshal(content, &td)

	return
}
