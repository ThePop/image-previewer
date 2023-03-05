package previewer

import (
	"bytes"
	"context"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

func getImage(originalHeader http.Header, url string) (image.Image, error) {
	client := http.Client{}
	req, err := http.NewRequestWithContext(context.Background(), "GET", "http://"+url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error during creating request")
	}

	for key, val := range originalHeader {
		req.Header.Set(key, val[0])
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error during getting image")
	}
	defer resp.Body.Close()

	buffer := &bytes.Buffer{}
	if _, err := io.Copy(buffer, resp.Body); err != nil {
		return nil, errors.Wrap(err, "error during reading body message")
	}

	if !strings.HasPrefix(resp.Header.Get("Content-Type"), "image/") {
		return nil, errors.New("wrong Content-Type for image")
	}

	result, err := jpeg.Decode(buffer)
	if err != nil {
		return nil, errors.Wrap(err, "error during decoding jpeg image")
	}

	return result, nil
}
