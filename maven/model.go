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
	Build                *Build                `xml:"build,omitempty"`
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

type Build struct {
	Directory             string            `xml:"directory,omitempty"`
	OutputDirectory       string            `xml:"outputDirectory,omitempty"`
	FinalName             string            `xml:"finalName,omitempty"`
	TestOutputDirectory   string            `xml:"testOutputDirectory,omitempty"`
	SourceDirectory       string            `xml:"sourceDirectory,omitempty"`
	ScriptSourceDirectory string            `xml:"scriptSourceDirectory,omitempty"`
	TestSourceDirectory   string            `xml:"testSourceDirectory,omitempty"`
	Resources             *Resources        `xml:"resources,omitempty"`
	TestResources         *TestResources    `xml:"testResources,omitempty"`
	PluginManagement      *PluginManagement `xml:"pluginManagement,omitempty"`
}

type Resources struct {
	Elements []Resource `xml:"resource,omitempty"`
}

type Resource struct {
	TargetPath string    `xml:"targetPath,omitempty"`
	Filtering  string    `xml:"filtering,omitempty"`
	Directory  string    `xml:"directory,omitempty"`
	Includes   *Includes `xml:"includes,omitempty"`
	Excludes   *Excludes `xml:"excludes,omitempty"`
}

type Includes struct {
	Elements []string `xml:"include,omitempty"`
}

type Excludes struct {
	Elements []string `xml:"exclude,omitempty"`
}

type TestResources struct {
	Elements []Resource `xml:"testResource,omitempty"`
}

type PluginManagement struct {
	Plugins *Plugins `xml:"plugins,omitempty"`
}

type Plugins struct {
	Elements []Plugin `xml:"plugin,omitempty"`
}

type Plugin struct {
	Coordinates
	Extensions bool `xml:"extensions,omitempty"`
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
