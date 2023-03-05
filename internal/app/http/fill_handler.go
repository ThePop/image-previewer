package http

import (
	"bytes"
	"image/jpeg"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/ThePop/image-previewer/internal/previewer"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type FillHandler struct {
	previewer previewer.Previewer
	logger    *zerolog.Logger
}

func NewFillHandler(p previewer.Previewer, l *zerolog.Logger) *FillHandler {
	return &FillHandler{
		previewer: p,
		logger:    l,
	}
}

func (h *FillHandler) Fill(w http.ResponseWriter, r *http.Request) {
	width := chi.URLParam(r, "width")
	height := chi.URLParam(r, "height")

	url := strings.Split(r.URL.String(), "/")
	sourceImageURL := strings.Join(url[4:], "/")

	widthInt, err := strconv.Atoi(width)
	if err != nil {
		h.logger.Err(err).Msg("wrong width format")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	heightInt, err := strconv.Atoi(height)
	if err != nil {
		h.logger.Err(err).Msg("wrong height format")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	image, err := h.previewer.GetPreview(r.Header, sourceImageURL, widthInt, heightInt)
	if err != nil {
		h.logger.Err(err).Msg("error during image processing")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	imgBuffer := new(bytes.Buffer)
	err = jpeg.Encode(imgBuffer, image, nil)
	if err != nil {
		h.logger.Err(err).Msg("error during encoding jpg")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(w, imgBuffer)
	if err != nil {
		h.logger.Err(err).Msg("error during forming response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Info().Msg("finished request")
}
