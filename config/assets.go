package config

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

const (
	ASSETS_PATH          = "assets/"
	COMPILED_ASSETS_PATH = "dist/"
	BUILD_FILE           = "build.json"
)

type assetsConfig struct {
	OriginRelativePath   string
	OriginFullPath       string
	CompiledRelativePath string
	CompiledFullPath     string
	BuildFile            string
	AssetsMapping        map[string]string
}

var Assets = assetsConfig{}

func (assets *assetsConfig) Configure() {
	assets.OriginRelativePath = ASSETS_PATH
	assets.CompiledRelativePath = COMPILED_ASSETS_PATH
	assets.OriginFullPath = Application.RootPath + "/" + assets.OriginRelativePath
	assets.CompiledFullPath = Application.RootPath + "/" + assets.CompiledRelativePath
	assets.BuildFile = BUILD_FILE
	assets.AssetsMapping = loadAssetsMapping()
}

func loadAssetsMapping() map[string]string {
	jsonAssets := map[string]string{}
	assetsWithPath := make(map[string]string)

	buildFile, err := os.ReadFile(Assets.CompiledFullPath + Assets.BuildFile)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(buildFile), &jsonAssets)
	for key, value := range jsonAssets {
		key = strings.ReplaceAll(key, Assets.OriginRelativePath, "")
		value = strings.ReplaceAll(value, Assets.CompiledRelativePath, "")
		assetsWithPath[key] = value
	}

	return assetsWithPath
}
