package maven

import (
	"fmt"
	"testing"
)

func TestX(t *testing.T) {
	pom := CreateProject()
	pom.Parent = &Parent{
		Coordinates: Coordinates{
			GroupID:    "tech.f3s",
			ArtifactID: "parent",
			Version:    "999",
		},
	}
	pom.GroupID = "tech.f3s"
	pom.ArtifactID = "market-model"
	pom.Version = "1.0.0-SNAPSHOT"
	pom.Packaging = "pom"
	pom.Dependencies = &Dependencies{
		Elements: []Dependency{
			{
				Coordinates: Coordinates{
					GroupID:    "tech.f3s",
					ArtifactID: "lib-a",
					Version:    "1.0",
				},
			},
		},
	}
	pom.DependencyManagement = &DependencyManagement{
		Dependencies: &Dependencies{
			Elements: []Dependency{
				{
					Coordinates: Coordinates{
						GroupID:    "tech.f3s",
						ArtifactID: "lib-a",
						Version:    "1.0",
					},
				},
			},
		},
	}

	pom.Modules = &Modules{
		Elements: []string{"a", "b", "c"},
	}

	if data, err := pom.MarshalIndent("", "  "); err != nil {
		t.Fail()
	} else {
		fmt.Println(string(data))
	}
}
