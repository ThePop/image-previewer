package previewer

import (
	"net/http"
	"testing"

	"github.com/ThePop/image-previewer/internal/cache"
)

func TestGetPreviewSuccess(t *testing.T) {
	t.Parallel()
	c := cache.NewCache(2)
	testPreviewer := NewPreviewer(c)
	sourceHeader := http.Header{}

	t.Run("success_get_preview", func(t *testing.T) {
		t.Parallel()

		expectedImage := getExampleImage(resizedImageExampleName)

		gotImage, err := testPreviewer.GetPreview(sourceHeader, previewImageURL, 800, 300)
		if err != nil {
			t.Errorf("getPreview() error = %v", err)
			return
		}

		if gotImage.Bounds().Size() != expectedImage.Bounds().Size() {
			t.Errorf(
				"getPreview() got = '%v', expected '%v'",
				gotImage.Bounds().Size(),
				expectedImage.Bounds().Size(),
			)
		}
	})
}
