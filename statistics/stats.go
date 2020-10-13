package statistics

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ingrowco/night-watch/configurator"
)

// Stats is a type for collecting data. All keys are string and values can be any primitive type (in Ingrow system)
type Stats map[string]interface{}

type initializersMap map[string]func(Stats, string) error

var initializers = initializersMap{
	"os":      osPopulate,
	"cpu":     cpuPopulate,
	"disk":    diskPopulate,
	"uptime":  uptimePopulate,
	"memory":  memoryPopulate,
	"loadavg": loadAvgPopulate,
	"network": networkPopulate,
}

// Init collect data based on selected plugins
func (s Stats) Init(ctx context.Context) error {
	c := configurator.FromContext(ctx)
	plugins := getPlugins(c.GetStringSlice("main.plugins"))

	for i := range plugins {
		if f, ok := initializers[plugins[i]]; ok {
			if err := f(s, c.GetString(fmt.Sprintf("linux.path.%s", plugins[i]))); err != nil {
				return err
			}
		}
	}

	return nil
}

// Get returns collected data
func Get(ctx context.Context) (Stats, error) {
	stats := make(Stats)

	if err := stats.Init(ctx); err != nil {
		return nil, err
	}

	if len(stats) == 0 {
		return nil, errors.New("no active plugin has been found")
	}

	return stats, nil
}

func fillKey(placeHolder, key string) string {
	return strings.NewReplacer("-", "_").Replace(fmt.Sprintf(placeHolder, key))
}

func getPlugins(p []string) []string {

	if len(p) == 1 {
		return strings.Split(p[0], ",")
	}

	return p
}
