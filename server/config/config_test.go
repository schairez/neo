package config

import (
	"testing"
)

func TestBuildConfigStructFromYAMLFile(t *testing.T) {
	c, err := ReadYAMLFile()
	if err != nil {
		t.Errorf("Failed to read YAML file into struct %v", err)

	}
	t.Log(c)
}
