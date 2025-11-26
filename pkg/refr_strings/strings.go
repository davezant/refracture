package refr_strings

import (
	"os"
	"path/filepath"
	"strings"
)

func AppendPathString(basePath, overPath string) string {
	if len(basePath) == 0 {
		return overPath
	}
	if len(overPath) == 0 {
		return basePath
	}

	if basePath[len(basePath)-1] == overPath[0] {
		return basePath + overPath[1:]
	}

	return basePath + overPath
}

func ContainsShow(slice []string, item string) string {
	for _, v := range slice {
		if v == item {
			return item
		}
	}
	return slice[0]
}

func StringFlat(baseString, suffix string) string {
	if strings.HasSuffix(baseString, suffix) {
		return strings.TrimSuffix(baseString, suffix)
	}
	return baseString
}

func FileStringReplace(filePath, oldString, newString string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	updated := strings.ReplaceAll(string(data), oldString, newString)

	err = os.WriteFile(filePath, []byte(updated), 0644)
	if err != nil {
		return err
	}

	return nil
}

func ToResPath(fullPath, projectRoot string) string {
	cleaned := filepath.Clean(fullPath)
	cleaned = strings.ReplaceAll(cleaned, `\`, `/`)

	projectRoot = filepath.Clean(projectRoot)
	projectRoot = strings.ReplaceAll(projectRoot, `\`, `/`)

	cleaned = strings.TrimPrefix(cleaned, projectRoot)
	cleaned = strings.TrimPrefix(cleaned, "/")

	return "res://" + cleaned
}

func CleanSlashes(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}
func RemoveProjectPrefix(fullPath, projectRoot string) string {

	cleaned := filepath.Clean(fullPath)

	cleaned = strings.ReplaceAll(cleaned, `\`, `/`)

	projectRoot = strings.ReplaceAll(filepath.Clean(projectRoot), `\`, `/`)
	cleaned = strings.TrimPrefix(cleaned, projectRoot)

	cleaned = strings.TrimPrefix(cleaned, "/")

	return cleaned
}
