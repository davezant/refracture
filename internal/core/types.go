package core

import (
	"fmt"
	"refracture/pkg/refr_strings"
)

const (
	ShadersPath = "/shaders/"
	ThemesPath  = "/themes/"
	ScriptsPath = "/scripts/"
	AssetsPath  = "/assets/"
	ScenesPath  = "/scenes/"

	ModelsPath   = "models/"
	TexturesPath = "textures/"
	AudiosPath   = "audio/"
	MediaPath    = "media/"
	TextPath     = "text-sources/"
	AnimPath     = "anim/"

	ShaderFormat   = ".gdshader"
	ThemesFormat   = ".tres"
	ScriptsFormat  = ".gd"
	ScenesFormat   = ".tscn"
	TexturesFormat = ".png"
	AudioFormat    = ".ogg"
	MediaFormat    = ".mp4"
	TextFormat     = ".txt"
	UidFormat      = ".uid"
	ImportFormat   = ".import"
	ProjectFormat  = ".godot"
)

var (
	TexturesFormats = []string{
		".png",
		".jpg",
		".jpeg",
		".bmp",
		".tga",
		".svg",
		ThemesFormat,
	}

	AudioFormats = []string{
		".ogg",
		".mp3",
		".wav",
	}

	MediaFormats = []string{
		".mp4",
		".webm",
		".avi",
	}

	ModelsFormats = []string{
		".glb",
		".gltf",
		".fbx",
		".obj",
		".dae",
		".3ds",
		".blend",
		".ply",
		".stl",
		".res",
	}
)
var IterableReadings = []string{
	ShaderFormat, ThemesFormat, TextFormat, ScriptsFormat, ScenesFormat, ProjectFormat, ImportFormat,
}

// Ignore folder can both be '/.godot/' or '.godot'
var ignoreFolders = []string{"/addons/", "/.godot/", "/DoNotMess/", "icon.svg"}
var ignoreFiles = []string{".gitignore", ".editorconfig", ".gitattributes", "LICENSE"}

type Structure struct {
	BasePath string

	//	Main Paths
	Shaders string
	Themes  string
	Scripts string
	Assets  string
	Scenes  string

	// 	Assets Subpaths
	Models   string
	Textures string
	Audios   string
	Media    string
	Anim     string
}

func NewStructure(base string) *Structure {
	if string(base[len(base)-1]) != "/" {
		base = base + "/"
	}
	return &Structure{
		BasePath: base + "refractureOut/",

		Shaders: base + "refractureOut" + ShadersPath,
		Themes:  base + "refractureOut" + ThemesPath,
		Scripts: base + "refractureOut" + ScriptsPath,
		Assets:  base + "refractureOut" + AssetsPath,
		Scenes:  base + "refractureOut" + ScenesPath,

		Models:   base + "refractureOut" + AssetsPath + ModelsPath,
		Textures: base + "refractureOut" + AssetsPath + TexturesPath,
		Audios:   base + "refractureOut" + AssetsPath + AudiosPath,
		Media:    base + "refractureOut" + AssetsPath + MediaPath,
		Anim:     base + "refractureOut" + AssetsPath + AnimPath,
	}
}

type PathManager struct {
	OldPathFiles []string
}

func NewPathManager() *PathManager {
	return &PathManager{}
}

type VariantPath struct {
	// Paths
	OldPath        string
	DesignatedPath string
	// Folders
	OldBaseFolder        string
	DesignatedBaseFolder string
}

func NewPathMod(oldPath, newPath, oldFolder, newFolder string) VariantPath {
	return VariantPath{
		OldPath:              oldPath,
		DesignatedPath:       newPath,
		OldBaseFolder:        oldFolder,
		DesignatedBaseFolder: newFolder,
	}
}

type Reader struct {
	filePathOld []string
	filePathNew []string
	resOld      []string
	resNew      []string
}

func NewReader(newVariantsPaths []VariantPath, projectRoot, outRoot string) Reader {
	var pathOld = []string{}
	var pathNew = []string{}
	var resOld = []string{}
	var resNew = []string{}
	for _, vp := range newVariantsPaths {
		pathOld = append(pathOld, refr_strings.CleanSlashes(vp.OldPath))
		pathNew = append(pathNew, refr_strings.CleanSlashes(vp.DesignatedPath))
		resOld = append(resOld, refr_strings.ToResPath(vp.OldBaseFolder, projectRoot))
		resNew = append(resNew, refr_strings.ToResPath(vp.DesignatedBaseFolder, outRoot))
		fmt.Println("Old Res:'", refr_strings.ToResPath(vp.OldBaseFolder, projectRoot), "'", "New Res:", refr_strings.ToResPath(vp.DesignatedBaseFolder, outRoot), "'")
	}
	return Reader{filePathOld: pathOld, filePathNew: pathNew, resOld: resOld, resNew: resNew}
}
