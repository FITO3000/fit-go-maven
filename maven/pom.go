package maven

import (
	"encoding/xml"
	"fmt"
)

const (
	modelVersion      string = "4.0.0"
	xmlns             string = "http://maven.apache.org/POM/4.0.0"
	xmlnsxsi          string = "http://www.w3.org/2001/XMLSchema-instance"
	xsischemalocation string = "http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd"
	prefix            string = ""
	indent            string = "  "
)

type POM struct {
	path  string
	Model *Project
}

func NewPOM() POM {
	return POM{
		Model: &Project{
			ModelVersion:      modelVersion,
			XmlNs:             xmlns,
			XmlNsXsi:          xmlnsxsi,
			XsiSchemaLocation: xsischemalocation,
		},
	}
}

func LoadPOM(path string) (*POM, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (p POM) String() string {
	if data, err := xml.MarshalIndent(p.Model, prefix, indent); err != nil {
		panic(err)
	} else {
		return string(data)
	}
}

func (p *POM) SetCoordinates(coordinates Coordinates) {
	p.Model.GroupID = coordinates.GroupID
	p.Model.ArtifactID = coordinates.ArtifactID
	p.Model.Version = coordinates.Version
}

func (p *POM) SetParent(coordinates Coordinates) {
	p.Model.Parent = &Parent{
		Coordinates: coordinates,
	}
}

func (p *POM) AddDepemdency(coordinates Coordinates) {

}
