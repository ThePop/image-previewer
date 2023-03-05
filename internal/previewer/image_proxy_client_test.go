package previewer

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetImageSuccess(t *testing.T) {
	t.Parallel()

	sourceHeader := http.Header{}

	t.Run("success_get_image", func(t *testing.T) {
		t.Parallel()

		gotImage, err := getImage(sourceHeader, previewImageURL)
		if err != nil {
			t.Errorf("getImage() error = %v", err)
			return
		}

		expectedImage := getExampleImage(sourceImageExampleName)
		if !reflect.DeepEqual(gotImage, expectedImage) {
			t.Errorf("getImage() got = %v, expected %v", gotImage, expectedImage)
		}
	})
}

func TestGetImageIncorrectUrl(t *testing.T) {
	t.Parallel()

	sourceHeader := http.Header{}
	errTextExcepted := "error during getting image"

	t.Run("error_get_image_incorrect_url", func(t *testing.T) {
		t.Parallel()

		_, err := getImage(sourceHeader, "incorrect_url")
		require.Errorf(t, err, errTextExcepted)
	})
}

func TestGetImageIncorrectHeader(t *testing.T) {
	t.Parallel()

	sourceHeader := http.Header{}
	errTextExcepted := "wrong Content-Type for image"

	t.Run("error_get_image_incorrect_header", func(t *testing.T) {
		t.Parallel()

		_, err := getImage(sourceHeader, "example.com")
		require.Errorf(t, err, errTextExcepted)
	})
}
