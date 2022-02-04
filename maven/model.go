package maven

import (
	"encoding/xml"
)

type Project struct {
	XMLName           xml.Name `xml:"project"`
	XmlNs             string   `xml:"xmlns,attr,omitempty"`
	XmlNsXsi          string   `xml:"xmlns:xsi,attr,omitempty"`
	XsiSchemaLocation string   `xml:"xsi:schemaLocation,attr,omitempty"`
	ModelVersion      string   `xml:"modelVersion,omitempty"`
	// The Basics
	Parent *Parent `xml:"parent,omitempty"`
	Coordinates
	Packaging     string `xml:"packaging,omitempty"`
	Name          string `xml:"name,omitempty"`
	Description   string `xml:"description,omitempty"`
	Url           string `xml:"url,omitempty"`
	InceptionYear string `xml:"inceptionYear,omitempty"`

	DependencyManagement *DependencyManagement `xml:"dependencyManagement,omitempty"`
	Dependencies         *Dependencies         `xml:"dependencies,omitempty"`
	Properties           *Properties           `xml:"properties,omitempty"`
	Modules              *Modules              `xml:"modules,omitempty"`
}

type Parent struct {
	Coordinates
	RelativePath string `xml:"relativePath,omitempty"`
}

type Coordinates struct {
	GroupID    string `xml:"groupId,omitempty"`
	ArtifactID string `xml:"artifactId,omitempty"`
	Version    string `xml:"version,omitempty"`
}

type DependencyManagement struct {
	Dependencies *Dependencies `xml:"dependencies,omitempty"`
}

type Modules struct {
	Elements []string `xml:"module,omitempty"`
}

type Properties struct {
	Entries map[string]string
}

/*
	<!-- Build Settings -->
	  <build>...</build>
	  <reporting>...</reporting>

	  <!-- More Project Information -->
	  <licenses>...</licenses>
	  <organization>...</organization>
	  <developers>...</developers>
	  <contributors>...</contributors>

	  <!-- Environment Settings -->
	  <issueManagement>...</issueManagement>
	  <ciManagement>...</ciManagement>
	  <mailingLists>...</mailingLists>
	  <scm>...</scm>
	  <prerequisites>...</prerequisites>
	  <repositories>...</repositories>
	  <pluginRepositories>...</pluginRepositories>
	  <distributionManagement>...</distributionManagement>
	  <profiles>...</profiles>
*/

type Dependencies struct {
	Elements []Dependency `xml:"dependency,omitempty"`
}

type Dependency struct {
	Coordinates
	Type     string `xml:"type,omitempty"`
	Scope    string `xml:"scope,omitempty"`
	Optional bool   `xml:"optional,omitempty"`
}
