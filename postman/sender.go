package postman

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ingrowco/night-watch/configurator"
)

type lungo struct {
	Project string `json:"project"`
	Stream  string `json:"stream"`
}

type event map[string]interface{}

type message struct {
	Lungo *lungo `json:"lungo"`
	Event event  `json:"event"`
}

func newMessage(project, stream string, events map[string]interface{}) *message {
	return &message{
		Lungo: &lungo{
			Project: project,
			Stream:  stream,
		},
		Event: events,
	}
}

// Send submits an event (a map to strings to interfaces) to the ingrow system
func Send(ctx context.Context, baseUrl, apiKey, project, stream string, stat map[string]interface{}) error {
	event := newMessage(project, stream, stat)
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return sender(ctx, baseUrl, apiKey, data)
}

func sender(ctx context.Context, baseUrl, apiKey string, data []byte) error {
	cfg := configurator.FromContext(ctx)

	r, err := http.NewRequestWithContext(ctx, "POST", baseUrl+"/events", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	r.Header.Set("api-key", apiKey)

	client := http.Client{
		Timeout: cfg.GetDuration("main.timeout"),
	}

	resp, err := client.Do(r)
	if err != nil {
		return err
	}

	if resp.StatusCode > 299 {
		return fmt.Errorf("%d - %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return nil
}
