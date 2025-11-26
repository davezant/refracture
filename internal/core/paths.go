package core

import (
	"fmt"
	"os"
	"path/filepath"
	"refracture/pkg/refr_strings"
	"refracture/pkg/shutils"
)

func (s *Structure) CreateFolders() error {
	paths := []string{s.Shaders, s.Themes, s.Scripts, s.Assets, s.Scenes, s.Textures, s.Audios, s.Models, s.Media}

	for _, p := range paths {
		if err := os.MkdirAll(p, 0777); err != nil {
			return err
		}
	}
	return nil
}

func (s *Structure) CopyFilesToDesignates(slicesVariants []VariantPath) error {
	for _, i := range slicesVariants {
		err := shutils.Copy(i.OldPath, i.DesignatedPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetRecursivePaths(basePath string) *PathManager {
	pm := NewPathManager()
	var files []string

	basePath = filepath.Clean(basePath)

	ignoreFolderMap := make(map[string]struct{})
	for _, f := range ignoreFolders {
		ignoreFolderMap[filepath.Join(basePath, f)] = struct{}{}
	}

	ignoreFileMap := make(map[string]struct{})
	for _, f := range ignoreFiles {
		ignoreFileMap[f] = struct{}{}
	}

	filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		cleanPath := filepath.Clean(path)

		if info.IsDir() {
			if _, skip := ignoreFolderMap[cleanPath]; skip {
				return filepath.SkipDir
			}
			return nil
		}

		if _, skip := ignoreFileMap[info.Name()]; skip {
			return nil
		}

		files = append(files, cleanPath)
		return nil
	})

	pm.OldPathFiles = files
	return pm
}

func (s *Structure) CopyRawSpecialFolders(copyAddons bool, projectBase string) error {
	var addons = ""
	icon := filepath.Join(projectBase, "icon.svg")
	iconImport := filepath.Join(projectBase, "icon.svg.import")
	if copyAddons {
		addons = filepath.Join(projectBase, "addons")
	}
	output := s.BasePath

	if err := os.MkdirAll(output, 0777); err != nil {
		return fmt.Errorf("failed to create base output folder: %w", err)
	}

	// Função auxiliar para copiar arquivos com log de erro
	copyFileSafe := func(src, dst string) {
		if err := shutils.Copy(src, dst); err != nil {
			fmt.Printf("⚠️ Error copying %s -> %s: %v\n", src, dst, err)
		} else {
			fmt.Printf("✅ Copied %s -> %s\n", src, dst)
		}
	}

	if _, err := os.Stat(icon); err == nil {
		copyFileSafe(icon, filepath.Join(output, "icon.svg"))
	}

	if _, err := os.Stat(iconImport); err == nil {
		copyFileSafe(iconImport, filepath.Join(output, "icon.svg.import"))
	}
	if copyAddons {
		if info, err := os.Stat(addons); err == nil && info.IsDir() {
			filepath.Walk(addons, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					fmt.Printf("⚠️ Error getting %s: %v\n", path, err)
					return nil
				}

				rel, _ := filepath.Rel(projectBase, path)
				dst := filepath.Join(output, rel)

				if info.IsDir() {
					if err := os.MkdirAll(dst, 0777); err != nil {
						fmt.Printf("⚠️ Error creating folder %s: %v\n", dst, err)
					}
					return nil
				}

				copyFileSafe(path, dst)
				return nil
			})
		}
	}
	return nil
}

func (n *PathManager) DesignateFiles(fileStructure *Structure, project_prefix string) []VariantPath {
	var variantsArray = []VariantPath{}

	for _, x := range n.OldPathFiles {
		var name = filepath.Base(x)
		var commonOldFolder = refr_strings.RemoveProjectPrefix(x, project_prefix)
		fixUid := refr_strings.StringFlat(name, UidFormat)
		fixImport := refr_strings.StringFlat(fixUid, ImportFormat)
		extensionFormat := filepath.Ext(fixImport)
		switch extensionFormat {

		// Basics
		case ProjectFormat:
			newPath := fileStructure.BasePath + name
			variantsArray = append(variantsArray, NewPathMod(x, newPath, commonOldFolder, newPath))

		case ScriptsFormat:
			newPath := fileStructure.Scripts + name
			variantsArray = append(variantsArray, NewPathMod(x, newPath, commonOldFolder, newPath))

		case ShaderFormat:
			newPath := fileStructure.Shaders + name
			variantsArray = append(variantsArray, NewPathMod(x, newPath, commonOldFolder, newPath))

		/*case ThemesFormat:
		newPath := fileStructure.Themes + name
		variantsArray = append(variantsArray, NewPathMod(x, newPath, commonOldFolder, newPath))
		*/
		case ScenesFormat:
			newPath := fileStructure.Scenes + name
			variantsArray = append(variantsArray, NewPathMod(x, newPath, commonOldFolder, newPath))

		// Assets
		case refr_strings.ContainsShow(AudioFormats, extensionFormat):
			newPath := fileStructure.Audios + name
			variantsArray = append(variantsArray, NewPathMod(x, newPath, commonOldFolder, newPath))

		case refr_strings.ContainsShow(TexturesFormats, extensionFormat):
			newPath := fileStructure.Textures + name
			variantsArray = append(variantsArray, NewPathMod(x, newPath, commonOldFolder, newPath))

		case refr_strings.ContainsShow(ModelsFormats, extensionFormat):
			newPath := fileStructure.Models + name
			variantsArray = append(variantsArray, NewPathMod(x, newPath, commonOldFolder, newPath))

		default:
			newPath := fileStructure.Media + name
			variantsArray = append(variantsArray, NewPathMod(x, newPath, commonOldFolder, newPath))
		}
	}

	return variantsArray
}
