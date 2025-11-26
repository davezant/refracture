package core_test

import (
	"fmt"
	"refracture/internal/core"
	"testing"
)

func TestPathDesignancyClearResourcesReader(t *testing.T) {
	basePath := "./tests_folder/moba-demo/"

	pm := core.GetRecursivePaths(basePath)
	base := "./tests_folder/"
	s := core.NewStructure(base)
	d := pm.DesignateFiles(s, basePath)

	s.CreateFolders()

	if err := s.CopyFilesToDesignates(d); err != nil {
		fmt.Println(err)
	}

	if err := s.CopyRawSpecialFolders(false, basePath); err != nil {
		fmt.Println(err)
	}

	reader := core.NewReader(d, basePath, "./tests_folder/refractureOut/")
	reader.ReplaceInFiles()
}

func ClearOldFiles(t *testing.T) {
	// TODO
}
