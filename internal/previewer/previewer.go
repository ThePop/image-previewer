package previewer

import (
	"crypto/sha256"
	"encoding/base32"
	"fmt"
	"image"
	"log"
	"net/http"

	"github.com/ThePop/image-previewer/internal/cache"
	"github.com/disintegration/imaging"
	"github.com/pkg/errors"
)

type Previewer interface {
	GetPreview(originalHeader http.Header, url string, width int, height int) (image.Image, error)
}

type previewer struct {
	cache cache.Cache
}

func NewPreviewer(c cache.Cache) Previewer {
	return &previewer{
		cache: c,
	}
}

func (p *previewer) GetPreview(sourceHeader http.Header, url string, width int, height int) (image.Image, error) {
	cacheKey := p.GetCacheKey(url, width, height)

	if cachedImage, isFound := p.cache.Get(cacheKey); isFound {
		log.Println("found cached image")
		return cachedImage.(image.Image), nil
	}

	sourceImage, err := getImage(sourceHeader, url)
	if err != nil {
		return nil, errors.Wrap(err, "error during getting image")
	}

	resizedImage := imaging.Fill(sourceImage, width, height, imaging.Center, imaging.Lanczos)

	p.cache.Set(cacheKey, resizedImage)
	return resizedImage, nil
}

func (p *previewer) GetCacheKey(url string, width int, height int) string {
	keyBytes := sha256.Sum256([]byte(fmt.Sprintf("%v%v%v", url, width, height)))
	return base32.StdEncoding.EncodeToString(keyBytes[:])
}
