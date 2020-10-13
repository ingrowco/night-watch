package configurator

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Config represent a configurator
type Config interface {
	Get(key string) interface{}
	GetBool(key string) bool
	GetFloat64(key string) float64
	GetInt(key string) int
	GetIntSlice(key string) []int
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	IsSet(key string) bool
}

type configKey struct{}

var v *viper.Viper

func init() {
	v = viper.New()

	setDefaultValues(v)
	addConfigLayers(v)
}

// Get return an instance of global configurator
func Get() *viper.Viper {
	return v
}

// WithContext saves a configurator inside of a context
func WithContext(ctx context.Context) context.Context {
	if c := ctx.Value(configKey{}); c != nil {
		return ctx
	}

	return context.WithValue(ctx, configKey{}, v)
}

// FromContext returns a configurator from a given context
func FromContext(ctx context.Context) *viper.Viper {
	return ctx.Value(configKey{}).(*viper.Viper)
}

func addConfigLayers(c *viper.Viper) {
	c.SetConfigName("config")
	c.SetConfigType("yaml")
	c.AddConfigPath("/etc/niw/")
	c.AddConfigPath("$HOME/.config/niw/")
	c.AddConfigPath("$PWD/")
	err := c.ReadInConfig()
	if err != nil {
		if _, notFound := err.(viper.ConfigFileNotFoundError); !notFound {
			panic(fmt.Sprintf("error on parsing configuration file: %s", err))
		}
	}

	c.SetEnvPrefix("niw")
	c.AutomaticEnv()
	c.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func setDefaultValues(c *viper.Viper) {
	// Logger
	c.SetDefault("log.level", "info")

	// Main
	c.SetDefault("main.interval", "30s")
	c.SetDefault("main.plugins", "os,cpu,disk,uptime,memory,loadavg,network")

	// Ingrow Admin Options
	c.SetDefault("ingrow.project", "")
	c.SetDefault("ingrow.stream", "")
	c.SetDefault("ingrow.apikey", "")
	c.SetDefault("ingrow.url", "")

	// Linux Paths
	c.SetDefault("linux.path.cpu", "/proc/stat")
	c.SetDefault("linux.path.loadavg", "/proc/loadavg")
	c.SetDefault("linux.path.disk", "/proc/diskstats")
	c.SetDefault("linux.path.memory", "/proc/meminfo")
	c.SetDefault("linux.path.network", "/proc/net/dev")
	c.SetDefault("linux.path.uptime", "/proc/uptime")
}
