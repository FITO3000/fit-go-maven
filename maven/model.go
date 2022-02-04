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
	GroupID      string        `xml:"groupId,omitempty"`
	ArtifactID   string        `xml:"artifactId,omitempty"`
	Version      string        `xml:"version,omitempty"`
	Packaging    string        `xml:"packaging,omitempty"`
	Dependencies *Dependencies `xml:"dependencies,omitempty"`
}

//Dependencies
//Parent
//DependencyManagement
//Modules
//Properties

/*
	<!-- Build Settings -->
	  <build>...</build>
	  <reporting>...</reporting>

	  <!-- More Project Information -->
	  <name>...</name>
	  <description>...</description>
	  <url>...</url>
	  <inceptionYear>...</inceptionYear>
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
	GroupID    string `xml:"groupId,omitempty"`
	ArtifactID string `xml:"artifactId,omitempty"`
	Version    string `xml:"version,omitempty"`
	Type       string `xml:"type,omitempty"`
	Scope      string `xml:"scope,omitempty"`
	Optional   bool   `xml:"optional,omitempty"`
}
