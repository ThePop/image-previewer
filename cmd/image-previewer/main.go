package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ThePop/image-previewer/internal"
	apphttp "github.com/ThePop/image-previewer/internal/app/http"
	"github.com/ThePop/image-previewer/internal/cache"
	"github.com/ThePop/image-previewer/internal/previewer"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func main() {
	logger := log.Logger

	config, err := internal.Configure()
	if err != nil {
		logger.Error().Err(err).Send()
		return
	}

	c := cache.NewCache(config.CacheCapacity)
	p := previewer.NewPreviewer(c)
	previewerHandler := apphttp.NewFillHandler(p, &logger)

	router := chi.NewRouter()
	router.Get("/fill/{width}/{height}/{target}*", previewerHandler.Fill)
	router.NotFoundHandler()

	server := &http.Server{
		Addr:        config.Host + ":" + config.Port,
		Handler:     router,
		ReadTimeout: 30 * time.Second,
	}

	logger.Info().Msg(fmt.Sprintf("start serving http at :%s", config.Port))

	go func() {
		if err = server.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
			logger.Error().Err(err).Send()
		}
	}()

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP)
	for {
		<-signalChannel
		logger.Info().Msg("interrupted")
		return
	}
}
