package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/ingrowco/night-watch/configurator"
	"github.com/ingrowco/night-watch/logger"
	"github.com/ingrowco/night-watch/postman"
	"github.com/ingrowco/night-watch/statistics"
)

func main() {
	printBuildInformation()

	ctx, cancel := context.WithCancel(context.Background())
	ctx = configurator.WithContext(ctx)
	ctx = logger.WithContext(ctx)

	logger.FromContext(ctx).Info("Warming up NiW...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go handleInterruptSignal(c, cancel)

	loop(ctx, time.Tick(configurator.FromContext(ctx).GetDuration("main.interval")))
}

func loop(ctx context.Context, tick <-chan time.Time) {
	project := configurator.FromContext(ctx).GetString("ingrow.project")
	checkRequiredValue(ctx, "Project", project)
	stream := configurator.FromContext(ctx).GetString("ingrow.stream")
	checkRequiredValue(ctx, "Stream", stream)
	apiKey := configurator.FromContext(ctx).GetString("ingrow.apikey")
	checkRequiredValue(ctx, "API Key", apiKey)
	baseURL := configurator.FromContext(ctx).GetString("ingrow.url")
	checkRequiredValue(ctx, "Base Url", baseURL)

	log := logger.FromContext(ctx).WithFields(map[string]interface{}{"project": project, "stream": stream, "baseurl": baseURL})

	for {
		select {
		case <-ctx.Done():
			log.Info("shutting down...")
			return
		case <-tick:
			go func() {
				stats, err := statistics.Get(ctx)
				if err != nil {
					log.Errorf("error on getting statistics, %v", err)
					return
				}
				log.WithField("stats", stats).Trace("statistics has been generated")

				err = postman.Send(ctx, baseURL, apiKey, project, stream, stats)
				if err != nil {
					log.Errorf("error on sending the event message, %v", err)
					return
				}
				log.Debug("event has been sent successfully")
			}()
		}
	}
}

func handleInterruptSignal(c chan os.Signal, cancel context.CancelFunc) {
	<-c
	signal.Stop(c)
	cancel()
}

func checkRequiredValue(ctx context.Context, field, value string) {
	log := logger.FromContext(ctx)

	if strings.Trim(value, " ") == "" {
		log.Fatalf("`%s` is required, you should set it before using NiW", field)
	}
}
