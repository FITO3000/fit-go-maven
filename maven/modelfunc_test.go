package maven

import (
	"fmt"
	"testing"
)

func TestX(t *testing.T) {
	pom := CreateProject()
	pom.GroupID = "tech.f3s"
	pom.ArtifactID = "market-model"
	pom.Version = "1.0.0-SNAPSHOT"
	pom.Packaging = "pom"
	pom.Dependencies = &Dependencies{
		Elements: []Dependency{
			{
				GroupID:    "tech.f3s",
				ArtifactID: "lib-a",
				Version:    "1.0",
			},
		},
	}

	if data, err := pom.MarshalIndent("", "  "); err != nil {
		t.Fail()
	} else {
		fmt.Println(string(data))
	}
}
