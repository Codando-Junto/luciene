package renderer_test

import (
	"bytes"
	"lucienne/pkg/renderer"
	"os"
	"strings"
	"testing"
)

func TestRederingHTML(t *testing.T) {
	renderer.HTML.Configure("/assets", "./", map[string]string{"some_path/something.asset": "some_path/other_path/random.asset"})
	teardown := setupHTMLFile(t, "./test.html")

	htmlBuffer := bytes.NewBuffer([]byte(""))
	renderer.HTML.Render(htmlBuffer, map[string]string{"TestContent": "some content"}, "test.html")

	t.Run("render inner variables", func(t *testing.T) {
		if !strings.Contains(htmlBuffer.String(), "some content") {
			t.Error("Expected: contains rendered value \"some content\", got: nothing")
		}
	})

	t.Run("render asset path", func(t *testing.T) {
		if !strings.Contains(htmlBuffer.String(), "<script src=/assets/some_path/other_path/random.asset></script>") {
			t.Error("Expected: contains rendered asset path \"/assets/some_path/other_path/random.asset\", got: nothing")
		}
	})

	teardown(t)
}

func setupHTMLFile(t *testing.T, filePath string) func(t *testing.T) {
	htmlContent := []byte(`
		<html>
			<head>
				<title>Testing</title>
			</head>
			<body>
				{{ .TestContent }}
				<script src={{ assetsPath "some_path/something.asset" }}></script>
			</body>
		</html>
	`)

	os.WriteFile(filePath, htmlContent, 0644)

	return func(t *testing.T) {
		os.Remove(filePath)
	}
}
