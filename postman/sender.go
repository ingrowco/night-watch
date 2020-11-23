package postman

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	http "github.com/valyala/fasthttp"

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

	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	resp := http.AcquireResponse()
	defer http.ReleaseResponse(resp)

	req.SetRequestURI(baseUrl)
	req.SetBody(data)
	req.Header.Set("api-key", apiKey)
	req.Header.SetMethod("POST")

	err := getClient(cfg.GetDuration("main.timeout")).Do(req, resp)
	if err != nil {
		return err
	}

	respCode := resp.Header.StatusCode()
	if respCode > 299 {
		return fmt.Errorf("%d - %s", respCode, http.StatusMessage(respCode))
	}

	return nil
}

var httpClient *http.Client

func getClient(timeout time.Duration) *http.Client {
	if httpClient != nil {
		return httpClient
	}
	httpClient = &http.Client{
		Name:        fmt.Sprintf("niw-%s", configurator.AppVersion),
		ReadTimeout: timeout,
	}

	return httpClient
}
