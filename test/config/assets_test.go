package config_test

import (
	"lucienne/config"
	"maps"
	"os"
	"testing"
)

func TestAssetsConfigure(t *testing.T) {
	config.Application.Configure("test")
	t.Run("set value for OriginRelativePath ", func(t *testing.T) {
		config.Assets.Configure()
		expectedResult := "assets/"
		if expectedResult != config.Assets.OriginRelativePath {
			t.Errorf("Expected: %s, Got: %s", expectedResult, config.Assets.OriginRelativePath)
		}
	})

	t.Run("sets right value for CompiledRelativePath", func(t *testing.T) {
		config.Assets.Configure()
		expectedResult := "dist/"
		if expectedResult != config.Assets.CompiledRelativePath {
			t.Errorf("Expected: %s, Got: %s", expectedResult, config.Assets.CompiledRelativePath)
		}
	})

	t.Run("sets right value for BuildFile", func(t *testing.T) {
		config.Assets.Configure()
		expectedResult := "build.json"
		if expectedResult != config.Assets.BuildFile {
			t.Errorf("Expected: %s, Got: %s", expectedResult, config.Assets.CompiledRelativePath)
		}
	})

	t.Run("sets right OriginFullPath", func(t *testing.T) {
		config.Assets.Configure()
		rootDir, _ := os.Getwd()
		expectedResult := rootDir + "/assets/"
		if expectedResult != config.Assets.OriginFullPath {
			t.Errorf("Expected: %s, Got: %s", expectedResult, config.Assets.OriginFullPath)
		}
	})

	t.Run("sets right CompiledFullPath", func(t *testing.T) {
		config.Assets.Configure()
		rootDir, _ := os.Getwd()
		expectedResult := rootDir + "/dist/"
		if expectedResult != config.Assets.CompiledFullPath {
			t.Errorf("Expected: %s, Got: %s", expectedResult, config.Assets.CompiledFullPath)
		}
	})

	t.Run("loads build json mapping file", func(t *testing.T) {
		config.Assets.Configure("my_origin_assets/", "test/config/test_compiled_path/", "test-build.json")
		expectedResult := map[string]string{
			"some_dir_1/some-file.js":      "first-file-AJFJEO.js",
			"some_dir_2/another-file.scss": "another-css-AJITD2.css",
			"random-file.jpg":              "random-file-GDSJOQR.jpg",
		}
		if !maps.Equal(expectedResult, config.Assets.AssetsMapping) {
			t.Errorf("Expected: %#v, Got: %#v", expectedResult, config.Assets.AssetsMapping)
		}
	})
}
