package events

import (
	"context"

	"github.com/lyft/flytestdlib/config"
	"github.com/lyft/flytestdlib/logger"
)

//go:generate pflags Config

type EventReportingType = string

// The reserved config section key for storage.
const configSectionKey = "Event"

const (
	EventSinkLog   EventReportingType = "log"
	EventSinkFile  EventReportingType = "file"
	EventSinkAdmin EventReportingType = "admin"
)

type Config struct {
	Type     EventReportingType `json:"type" pflag:",Sets the type of EventSink to configure [log/admin/file]."`
	FilePath string             `json:"file-path" pflag:",For file types, specify where the file should be located."`
}

// Retrieve current global config for storage.
func GetConfig(ctx context.Context) *Config {
	if c, ok := config.GetSection(configSectionKey).(*Config); ok {
		return c
	}

	logger.Warnf(ctx, "Failed to retrieve config section [%v].", configSectionKey)
	return nil
}

func init() {
	if err := config.RegisterSection(configSectionKey, &Config{}); err != nil {
		panic(err)
	}
}