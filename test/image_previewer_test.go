package test

import (
	"bytes"
	"context"
	"fmt"
	"image"
	_ "image/jpeg"
	"io"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const (
	imageURL    = "nginx:8000/_gopher_original_1024x504.jpg"
	imageWidth  = "80"
	imageHeight = "60"
)

type TestSuite struct {
	suite.Suite
	client *http.Client
}

func NewTestSuite() *TestSuite {
	return &TestSuite{client: http.DefaultClient}
}

func TestImagePreview(t *testing.T) {
	s := NewTestSuite()

	width, height := imageWidth, "60"

	// nolint:bodyclose
	res, body, err := s.sendRequest(t, imageURL, width, height)
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, res.StatusCode)
	require.True(t, res.Header.Get("Content-Type") == "image/jpeg")

	config, _, _ := image.DecodeConfig(bytes.NewReader(body))

	require.Equal(t, strconv.Itoa(config.Width), width)
	require.Equal(t, strconv.Itoa(config.Height), height)
}

func TestServerDoesntExist(t *testing.T) {
	s := NewTestSuite()

	url := "http://not_exist.com/gopher.jpg"
	width, height := imageWidth, imageHeight

	// nolint:bodyclose
	res, _, err := s.sendRequest(t, url, width, height)
	require.NoError(t, err)

	require.Equal(t, http.StatusInternalServerError, res.StatusCode)
}

func TestImageNotFound(t *testing.T) {
	s := NewTestSuite()

	url := "http://example.com/not_exists.jpg"
	width, height := imageWidth, imageHeight

	// nolint:bodyclose
	res, _, err := s.sendRequest(t, url, width, height)
	require.NoError(t, err)

	require.Equal(t, http.StatusInternalServerError, res.StatusCode)
}

func TestNotImage(t *testing.T) {
	s := NewTestSuite()

	url := "http://ngingx:80/text_file.txt"
	width, height := imageWidth, imageHeight

	// nolint:bodyclose
	res, _, err := s.sendRequest(t, url, width, height)
	require.NoError(t, err)

	require.Equal(t, http.StatusInternalServerError, res.StatusCode)
}

func TestWrongWidth(t *testing.T) {
	s := NewTestSuite()

	width, height := "test", imageHeight

	// nolint:bodyclose
	res, _, err := s.sendRequest(t, imageURL, width, height)
	require.NoError(t, err)

	require.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestWrongHeight(t *testing.T) {
	s := NewTestSuite()

	width, height := imageWidth, "test"

	// nolint:bodyclose
	res, _, err := s.sendRequest(t, imageURL, width, height)
	require.NoError(t, err)

	require.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func (s *TestSuite) sendRequest(t *testing.T, url string, width, height string) (*http.Response, []byte, error) {
	t.Helper()
	previewerURL := "http://127.0.0.1:8888"
	requestURL := fmt.Sprintf(
		"%s/fill/%s/%s/%s",
		previewerURL,
		width,
		height,
		url,
	)
	req, err := http.NewRequestWithContext(context.Background(), "GET", requestURL, nil)
	require.NoError(t, err)

	res, err := s.client.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	require.NoError(t, err)

	return res, b, err
}
