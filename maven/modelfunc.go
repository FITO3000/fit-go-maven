package maven

import (
	"encoding/xml"
)

const (
	modelVersion      string = "4.0.0"
	xmlns             string = "http://maven.apache.org/POM/4.0.0"
	xmlnsxsi          string = "http://www.w3.org/2001/XMLSchema-instance"
	xsischemalocation string = "http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd"
)

func CreateProject() *Project {
	return &Project{
		ModelVersion:      modelVersion,
		XmlNs:             xmlns,
		XmlNsXsi:          xmlnsxsi,
		XsiSchemaLocation: xsischemalocation,
	}
}

func (p *Project) MarshalIndent(prefix, indent string) ([]byte, error) {
	return xml.MarshalIndent(p, prefix, indent)
}
